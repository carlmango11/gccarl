package semantic

import (
	"fmt"
	"strconv"

	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
)

type builder struct {
	funcs     map[cparser.IDEN]Type
	strs      []string
	scopes    []*Scope
	nextVarID VarID
}

type Scope struct {
	vars    map[cparser.IDEN]Var
	structs map[cparser.IDEN]Type
}

var compareOp = map[cparser.OperatorType]CompareOp{
	cparser.OperatorTypeLess:  OpLessThan,
	cparser.OperatorTypeEqual: OpEquals,
}

var numericalOps = map[cparser.OperatorType]NumericOp{
	cparser.OperatorTypePlus: NumOpAdd,
}

func Build(program *cparser.Main) (*Program, error) {
	b := &builder{
		funcs: map[cparser.IDEN]Type{
			"do_syscall": {
				Kind: KindVoid,
			},
			"assert": {
				Kind: KindVoid,
			},
		},
		scopes: []*Scope{newScope()},
	}

	b.funcs["print"] = Type{
		Kind: KindVoid,
	}

	return b.build(program)
}

func (b *builder) build(p *cparser.Main) (*Program, error) {
	var funcDecs []*FuncDef

	for _, dd := range p.Main.DecDef {
		switch dd.Type {
		case cparser.DecDefTypeFuncDef:
			f, err := b.toFuncDef(dd.FuncDef)
			if err != nil {
				return nil, err
			}

			funcDecs = append(funcDecs, f)
		case cparser.DecDefTypeTypeDef:
			err := b.defineType(dd.TypeDef)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Program{
		FuncDefs: funcDecs,
		Strings:  b.strs,
	}, nil
}

func (b *builder) newScope() {
	b.scopes = append(b.scopes, newScope())
}

func (b *builder) dropScope() {
	b.scopes = b.scopes[:len(b.scopes)-1]
}

func (b *builder) toFuncDef(f *cparser.DecDef_FuncDefOption) (*FuncDef, error) {
	b.newScope()
	defer b.dropScope()

	returnType, err := b.toReturnType(f.Type)
	if err != nil {
		return nil, err
	}

	b.funcs[f.IDEN] = returnType

	var paramVars []Var
	if f.ParamsDef != nil {
		ps := []*cparser.ParamDef{f.ParamsDef.Params.ParamDef}
		for _, x := range f.ParamsDef.Params.CommaParamDef {
			ps = append(ps, x.Param.ParamDef)
		}

		for _, astParam := range ps {
			v, err := b.declareVar(astParam.Param.Type, astParam.Param.VariableDef, true)
			if err != nil {
				return nil, err
			}

			paramVars = append(paramVars, v)
		}
	}

	var statements []*Statement
	for _, s := range f.Block.Block.Statement {
		statement, err := b.toStatement(s)
		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)
	}

	return &FuncDef{
		ReturnType: returnType,
		Name:       FuncName(f.IDEN),
		Params:     paramVars,
		Statements: statements,
	}, nil
}

func (b *builder) declareVar(astType *cparser.Type, v *cparser.VariableDef, isParam bool) (Var, error) {
	typ, err := b.toType(astType, v, isParam)
	if err != nil {
		return Var{}, err
	}

	varName := varDefName(v)

	_, ok := b.scope().vars[varName]
	if ok {
		return Var{}, fmt.Errorf("variable %s already declared", varName)
	}

	newVar := b.newVar(varName, typ)

	b.scope().vars[varName] = newVar

	return newVar, nil
}

func varDefName(v *cparser.VariableDef) cparser.IDEN {
	switch v.Type {
	case cparser.VariableDefTypePointer:
		p := v.Pointer.VariableDef
		for {
			if p.Type == cparser.VariableDefTypeVariable {
				return p.Variable.IDEN
			}

			p = p.Pointer.VariableDef
		}
	case cparser.VariableDefTypeVariable:
		return v.Variable.IDEN
	default:
		panic(fmt.Sprintf("unhandled variable type %s", v.Type))
	}
}

func (b *builder) toType(typ *cparser.Type, v *cparser.VariableDef, isParam bool) (Type, error) {
	switch v.Type {
	case cparser.VariableDefTypePointer:
		sub, err := b.toType(typ, v.Pointer.VariableDef, false) // no need for array decay check anymore
		if err != nil {
			return Type{}, err
		}

		return Type{
			Kind:    KindPointer,
			SubType: &sub,
		}, nil
	case cparser.VariableDefTypeVariable:
		return b.toVarType(typ, v.Variable.ArrayIndexDef, isParam)
	}

	panic("invalid type")
}

func (b *builder) toVarType(typ *cparser.Type, arrs []*cparser.ArrayIndexDef, isParam bool) (Type, error) {
	if len(arrs) > 0 {
		if isParam {
			// arrays decay to pointers in params
			return Type{
				Kind: KindPointer,
				Prim: astTypeToPrim(typ),
			}, nil
		}

		var size int
		sizeStr := arrs[0].ArrayIndex.NUM

		if sizeStr != "" {
			var err error
			size, err = strconv.Atoi(string(sizeStr))
			if err != nil {
				panic("invalid size: " + sizeStr)
			}
		}

		sub, err := b.toVarType(typ, arrs[1:], false)
		if err != nil {
			return Type{}, err
		}

		return Type{
			Kind:      KindArray,
			SubType:   &sub,
			ArraySize: size,
		}, nil
	}

	if typ.Type == cparser.TypeTypeStruct {
		st, ok := b.structType(typ.Struct.IDEN)
		if !ok {
			panicf("%v is not defined", typ.Struct.IDEN)
		}

		return st, nil
	}

	return Type{
		Kind: KindPrimitive,
		Prim: astTypeToPrim(typ),
	}, nil
}

func (b *builder) structType(n cparser.IDEN) (Type, bool) {
	for i := len(b.scopes) - 1; i >= 0; i-- {
		s, ok := b.scopes[i].structs[n]
		if ok {
			return s, true
		}
	}

	return Type{}, false
}

func (b *builder) getVar(n cparser.IDEN) (Var, error) {
	for i := len(b.scopes) - 1; i >= 0; i-- {
		s, ok := b.scopes[i].vars[n]
		if ok {
			return s, nil
		}
	}

	return Var{}, fmt.Errorf("variable %s not declared", n)
}

func (b *builder) scope() *Scope {
	return b.scopes[len(b.scopes)-1]
}

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func astTypeToPrim(typ *cparser.Type) PrimitiveType {
	switch typ.Type {
	case cparser.TypeTypeInt:
		return PrimInt32
	case cparser.TypeTypeChar:
		return PrimChar
	case cparser.TypeTypeUnsignedInt:
		return PrimUChar
	default:
		// custom
		panic("impl")
	}
}

func (b *builder) toReturnType(i *cparser.Type) (Type, error) {
	kind := KindPrimitive

	var prim PrimitiveType
	switch i.Type {
	case cparser.TypeTypeInt:
		prim = PrimInt32
	case cparser.TypeTypeChar:
		prim = PrimChar
	case cparser.TypeTypeVoid:
		return Type{
			Kind: KindVoid,
		}, nil
	default:
		// custom
		panic("impl")
	}

	return Type{
		Kind: kind,
		Prim: prim,
	}, nil
}

func (b *builder) toStatement(s *cparser.Statement) (*Statement, error) {
	switch s.Type {
	case cparser.StatementTypeIf:
		ifC, err := b.toIf(s.If)
		if err != nil {
			return nil, err
		}

		return &Statement{
			If: ifC,
		}, nil
	case cparser.StatementTypeWhile:
		w, err := b.toWhile(s.While)
		if err != nil {
			return nil, err
		}

		return &Statement{
			While: w,
		}, nil
	case cparser.StatementTypeFor:
		f, err := b.toFor(s.For)
		if err != nil {
			return nil, err
		}

		return &Statement{
			For: f,
		}, nil
	case cparser.StatementTypeDecAssign:
		a, err := b.toDecAssign(s.DecAssign.DecAssign.Standard)
		if err != nil {
			return nil, err
		}

		return &Statement{
			DeclareInit: a,
		}, nil
	case cparser.StatementTypeVarDec:
		_, err := b.declareVar(s.VarDec.VarDec.VarDec.Type, s.VarDec.VarDec.VarDec.VariableDef, false)
		if err != nil {
			return nil, err
		}

		return nil, nil
	case cparser.StatementTypeReturn:
		expr, err := b.toExpr(s.Return.Expr)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Return: expr,
		}, nil
	case cparser.StatementTypeExpr:
		expr, err := b.toExpr(s.Expr.Expr)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Expr: expr,
		}, nil
	case cparser.StatementTypeCompound:
		var ss []*Statement

		for _, cs := range s.Compound.Block.Block.Statement {
			compoundS, err := b.toStatement(cs)
			if err != nil {
				return nil, err
			}

			ss = append(ss, compoundS)
		}

		return &Statement{
			Compound: &Compound{
				Statements: ss,
			},
		}, nil
	}

	panic("invalid statement: " + s.Type)
}

func (b *builder) toDecAssign(a *cparser.DecAssign_StandardOption) (*InitVar, error) {
	v, err := b.declareVar(a.Type, a.VariableDef, false)
	if err != nil {
		return nil, err
	}

	init, err := b.toInitialiser(v.Type, a.Initialiser)
	if err != nil {
		return nil, err
	}

	return &InitVar{
		Var:         v,
		Initialiser: init,
	}, nil
}

func (b *builder) toExpr(expr *cparser.Expr) (*Expr, error) {
	switch expr.Type {
	case cparser.ExprTypeComp:
		compExpr := expr.Comp.CompExpr.CompExpr

		rightExpr, err := b.toExpr(compExpr.Expr)
		if err != nil {
			return nil, err
		}

		leftExpr, err := b.fromSubExpr(compExpr.SubExpr)
		if err != nil {
			return nil, err
		}

		if compExpr.Operator.Type == cparser.OperatorTypeAssign {
			if !leftExpr.Writeable() {
				return nil, fmt.Errorf("cannot assign to expression %v", rightExpr)
			}

			return &Expr{
				Type: leftExpr.Type,
				Assign: &Assign{
					To:   leftExpr,
					Expr: rightExpr,
				},
			}, nil
		}

		op, ok := compareOp[compExpr.Operator.Type]
		if ok {
			return &Expr{
				Type: boolType(),
				Compare: &CompareOpExpr{
					Left:  rightExpr,
					Op:    op,
					Right: leftExpr,
				},
			}, nil
		}

		// todo: check for casts

		return &Expr{
			Type: rightExpr.Type,
			Numeric: &NumericOpExpr{
				Left:  rightExpr,
				Op:    numericalOps[compExpr.Operator.Type],
				Right: leftExpr,
			},
		}, nil
	case cparser.ExprTypeSubExpr:
		return b.fromSubExpr(expr.SubExpr.SubExpr)
	}

	panic("invalid expression: " + string(expr.Type))
}

func (b *builder) fromSubExpr(sub *cparser.SubExpr) (*Expr, error) {
	switch sub.Type {
	case cparser.SubExprTypeFuncCall:
		fc, err := b.toFuncCall(sub.FuncCall)
		if err != nil {
			return nil, err
		}

		returnType, ok := b.funcs[sub.FuncCall.IDEN]
		if !ok {
			return nil, fmt.Errorf("function %s not declared", sub.FuncCall.IDEN)
		}

		return &Expr{
			Type:     returnType,
			FuncCall: fc,
		}, nil
	case cparser.SubExprTypeAddressOf:
		expr, err := b.toExpr(sub.AddressOf.Expr)
		if err != nil {
			return nil, err
		}

		if !expr.Writeable() {
			return nil, fmt.Errorf("cannot assign to expression %v", expr)
		}

		return &Expr{
			Type: Type{
				Kind:    KindPointer,
				SubType: &expr.Type,
			},
			AddressOf: expr,
		}, nil
	case cparser.SubExprTypeDeref:
		expr, err := b.toExpr(sub.Deref.Expr)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Type:  expr.Type,
			Deref: expr,
		}, nil
	case cparser.SubExprTypeVariable:
		return b.fromVarOption(sub.Variable)
	case cparser.SubExprTypeValue:
		v := sub.Value

		switch v.Value.Type {
		case cparser.ValueTypeStr:
			s := string(v.Value.Str.STR)
			b.strs = append(b.strs, s[1:len(s)-1])

			return &Expr{
				Type: Type{
					Kind: KindArray,
					SubType: &Type{
						Kind: KindPrimitive,
						Prim: PrimChar,
					},
					ArraySize: len(s),
				},
				StringID: StringID(len(b.strs)),
			}, nil
		case cparser.ValueTypeInt:
			// TODO other sizes
			i, err := strconv.Atoi(string(v.Value.Int.NUM))
			if err != nil {
				return nil, err
			}

			return &Expr{
				Type: int32Type(),
				Literal: &Literal{
					Int32: int32(i),
				},
			}, nil
		case cparser.ValueTypeChar:
			return &Expr{
				Type: charType(),
				Literal: &Literal{
					Char: v.Value.Char.CHAR[1],
				},
			}, nil
		}
	}

	panic("invalid sub expression: " + string(sub.Type))
}

func (b *builder) fromVarOption(vo *cparser.SubExpr_VariableOption) (*Expr, error) {
	v, err := b.getVar(vo.SubVariableAccess.V.IDEN)
	if err != nil {
		return nil, err
	}

	e := &Expr{
		Type: v.Type,
		Var:  &v,
	}

	// array indexing
	e, err = wrapArrayIndexing(e, string(vo.SubVariableAccess.V.IDEN), vo.SubVariableAccess.V.ArrayIndexAccess)
	if err != nil {
		return nil, err
	}

	return b.fromVarOptionRecursive(e, string(vo.SubVariableAccess.V.IDEN), vo.InnerSubVariableAccess)
}

func (b *builder) fromVarOptionRecursive(inner *Expr, parentName string, sub []*cparser.InnerSubVariableAccess) (*Expr, error) {
	if len(sub) == 0 {
		return inner, nil
	}

	x := sub[0]
	rest := sub[1:]

	var structType Type
	var subVar *cparser.SubVariableAccess

	switch x.Type {
	case cparser.InnerSubVariableAccessTypeDot:
		if inner.Type.Kind != KindStruct {
			return nil, fmt.Errorf("%v is not a struct", parentName)
		}

		structType = inner.Type
		subVar = x.Dot.SubVariableAccess

	case cparser.InnerSubVariableAccessTypeArrow:
		if inner.Type.Kind != KindPointer {
			return nil, fmt.Errorf("cannot use %v->%v for non-pointer", parentName, x.Arrow.SubVariableAccess.V.IDEN)
		}

		structType = *inner.Type.SubType
		subVar = x.Arrow.SubVariableAccess

		inner = &Expr{
			Type:  *inner.Type.SubType,
			Deref: inner,
		}

	default:
		panic("invalid sub expression: " + string(x.Type))
	}

	if structType.Kind != KindStruct {
		return nil, fmt.Errorf("%v is not a struct", parentName)
	}

	f, ok := structType.Struct.Field(FieldName(subVar.V.IDEN))
	if !ok {
		return nil, fmt.Errorf("%v does not have field %v", parentName, subVar.V.IDEN)
	}

	inner = &Expr{
		Type: f.Type,
		Field: &FieldExpr{
			Expr:  inner,
			Field: f.Name,
		},
	}

	// array indexing
	var err error
	inner, err = wrapArrayIndexing(inner, string(f.Name), subVar.V.ArrayIndexAccess)
	if err != nil {
		return nil, err
	}

	return b.fromVarOptionRecursive(inner, string(f.Name), rest)
}

func wrapArrayIndexing(inner *Expr, parentName string, arr []*cparser.ArrayIndexAccess) (*Expr, error) {
	if len(arr) == 0 {
		return inner, nil
	}

	x := arr[0]
	rest := arr[1:]

	if !inner.Type.Indexable() {
		return nil, fmt.Errorf("%v is not indexable", parentName)
	}

	idx, err := strconv.Atoi(string(x.ArrayIndex.NUM))
	if err != nil {
		return nil, fmt.Errorf("invalid array index %v: %v", x.ArrayIndex.NUM, err)
	}

	inner = &Expr{
		Type: *inner.Type.SubType,
		Index: &IndexExpr{
			Expr:  inner,
			Index: idx,
		},
	}

	return wrapArrayIndexing(inner, "fix me", rest) // todo fix parent
}

func (b *builder) toArrayLit(av *cparser.ArrayEntries_EntriesOption) ([]*Expr, error) {
	exprsNodes := []*cparser.Expr{av.Expr}
	for _, e := range av.CommaExpr {
		exprsNodes = append(exprsNodes, e.CommaExpr.Expr)
	}

	var exprs []*Expr

	for _, astExpr := range exprsNodes {
		expr, err := b.toExpr(astExpr)
		if err != nil {
			return nil, err
		}

		// type check TODO

		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func (b *builder) toInitList(typ Type, e *cparser.CompEntries) (*InitList, error) {
	if !typ.TakesInitList() {
		return nil, fmt.Errorf("%v cannot take an initialiser list", typ)
	}

	exprsNodes := []*cparser.CompEntry{e.Entries.CompEntry}
	for _, e := range e.Entries.CommaCompEntry {
		exprsNodes = append(exprsNodes, e.E.CompEntry)
	}

	var entries []*ListEntry

	var i int

	for _, node := range exprsNodes {
		var names []FieldName
		var subType Type
		var initNode *cparser.Initialiser

		switch node.Type {
		case cparser.CompEntryTypeAnon:
			fieldType, err := typ.Field(i)
			if err != nil {
				return nil, err
			}

			subType = fieldType
			initNode = node.Anon.Initialiser

		case cparser.CompEntryTypeLabelled:
			if len(node.Labelled.EntryLabelField) == 0 {
				return nil, fmt.Errorf("labelled expression cannot have no label field")
			}

			for _, n := range node.Labelled.EntryLabelField {
				names = append(names, FieldName(n.C.IDEN))
			}

			if typ.Kind != KindStruct {
				return nil, fmt.Errorf("%v is not a struct so called take label %v", typ, names)
			}

			fieldType, ok := fieldType(typ, names)
			if !ok {
				return nil, fmt.Errorf("%v does not have field %v", typ, names)
			}

			subType = fieldType
			initNode = node.Labelled.Initialiser

			// jump to labelled field
			i = typ.Struct.FieldIndex(names[0])
		}

		init, err := b.toInitialiser(subType, initNode)
		if err != nil {
			return nil, err
		}

		i++

		// type check TODO

		entries = append(entries, &ListEntry{
			Name: names,
			Init: init,
		})
	}

	return &InitList{
		Entries: entries,
	}, nil
}

func fieldType(t Type, names []FieldName) (Type, bool) {
	for _, f := range t.Struct.Fields {
		if f.Name == names[0] {
			if len(names) == 1 {
				return f.Type, true
			}

			return fieldType(f.Type, names[1:])
		}
	}

	return Type{}, false
}

func (b *builder) toFuncCall(call *cparser.SubExpr_FuncCallOption) (*FuncCall, error) {
	var params []*cparser.Expr
	if call.Params != nil {
		params = append(params, call.Params.Params.Expr)

		for _, x := range call.Params.Params.CommaExpr {
			params = append(params, x.CommaExpr.Expr)
		}
	}

	var args []*Expr
	for _, e := range params {
		arg, err := b.toExpr(e)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	return &FuncCall{
		Func: FuncName(call.IDEN),
		Args: args,
	}, nil
}

func (b *builder) toIf(i *cparser.Statement_IfOption) (*If, error) {
	expr, err := b.toExpr(i.Expr)
	if err != nil {
		return nil, err
	}

	if expr.Type.Prim != PrimBool {
		return nil, fmt.Errorf("if condition must be a boolean")
	}

	ifStatement, err := b.toStatement(i.Statement)
	if err != nil {
		return nil, err
	}

	var elseStatement *Statement

	if i.Else != nil {
		elseStatement, err = b.toStatement(i.Else.Else.Statement)
		if err != nil {
			return nil, err
		}
	}

	return &If{
		Condition:     expr,
		Statement:     ifStatement,
		ElseStatement: elseStatement,
	}, nil
}

func (b *builder) toWhile(w *cparser.Statement_WhileOption) (*While, error) {
	expr, err := b.toExpr(w.Expr)
	if err != nil {
		return nil, err
	}

	if expr.Type.Prim != PrimBool {
		return nil, fmt.Errorf("while condition must be a boolean")
	}

	s, err := b.toStatement(w.Statement)
	if err != nil {
		return nil, err
	}

	return &While{
		Condition: expr,
		Statement: s,
	}, nil
}

func (b *builder) toFor(f *cparser.Statement_ForOption) (*For, error) {
	init, err := b.toStatement(f.Statement0)
	if err != nil {
		return nil, err
	}

	cond, err := b.toStatement(f.Statement1)
	if err != nil {
		return nil, err
	}

	action, err := b.toStatement(f.Statement2)
	if err != nil {
		return nil, err
	}

	s, err := b.toStatement(f.Statement3)
	if err != nil {
		return nil, err
	}

	return &For{
		Init:      init,
		Condition: cond,
		Action:    action,
		Statement: s,
	}, nil
}

func (b *builder) defineType(d *cparser.DecDef_TypeDefOption) error {
	switch d.TypeDef.Type {
	case cparser.TypeDefTypeStructDef:
		return b.toStructDef(d.TypeDef.StructDef)
	}

	panic("invalid type")
}

func (b *builder) toStructDef(td *cparser.TypeDef_StructDefOption) error {
	var vars []StructField
	structVars := map[cparser.IDEN]bool{}

	for _, a := range td.VarDecColon {
		t := a.C.VarDec.VarDec.Type
		def := a.C.VarDec.VarDec.VariableDef

		fieldName := varDefName(def)

		_, ok := structVars[fieldName]
		if ok {
			return fmt.Errorf("duplicate field name %s on %s", fieldName, td.IDEN)
		}

		structVars[fieldName] = true

		typ, err := b.toType(t, def, false)
		if err != nil {
			return err
		}

		vars = append(vars, StructField{
			Name: FieldName(fieldName),
			Type: typ,
		})
	}

	// todo check exists

	st := StructType{
		Name:   TypeName(td.IDEN),
		Fields: vars,
	}

	b.scope().structs[td.IDEN] = Type{
		Kind:   KindStruct,
		Struct: st,
	}

	return nil
}

func (b *builder) toInitialiser(typ Type, init *cparser.Initialiser) (*Initialiser, error) {
	switch init.Type {
	case cparser.InitialiserTypeExpr:
		expr, err := b.toExpr(init.Expr.Expr)
		if err != nil {
			return nil, err
		}

		expr, err = b.handleExprType(typ, expr)
		if err != nil {
			return nil, err
		}

		return &Initialiser{
			Expr: expr,
		}, nil

	case cparser.InitialiserTypeList:
		cl, err := b.toInitList(typ, init.List.CompEntries)
		if err != nil {
			return nil, err
		}

		// TODO: check compatible

		return &Initialiser{
			List: cl,
		}, nil

	default:
		panic("invalid initialiser")
	}
}

func (b *builder) handleExprType(varType Type, expr *Expr) (*Expr, error) {
	if !varType.Equals(expr.Type) {
		if compatibleTypes(varType, expr.Type) {
			expr = &Expr{
				Cast: &Cast{
					To:   varType,
					Expr: expr,
				},
			}
		} else {
			panic("not compatible")
		}
	}

	return expr, nil
}

func (b *builder) newVar(iden cparser.IDEN, typ Type) Var {
	b.nextVarID++

	return Var{
		ID:   b.nextVarID,
		Name: VarName(iden),
		Type: typ,
	}
}

func int32Type() Type {
	return Type{Kind: KindPrimitive, Prim: PrimInt32}
}

func boolType() Type {
	return Type{Kind: KindPrimitive, Prim: PrimBool}
}

func charType() Type {
	return Type{Kind: KindPrimitive, Prim: PrimChar}
}

func compatibleTypes(t1, t2 Type) bool {
	if t1.Kind == KindStruct {
		return t2.Kind == KindArray
	}

	//switch t1.Kind {
	//case KindPointer:
	//	switch t2.Kind {
	//	case KindPointer:
	//		return compatibleTypes(*t1.SubType, *t2.SubType)
	//	case
	//	}
	//}
	//
	//return false
	panic("imlp")
}

func newScope() *Scope {
	return &Scope{
		vars:    make(map[cparser.IDEN]Var),
		structs: make(map[cparser.IDEN]Type),
	}
}
