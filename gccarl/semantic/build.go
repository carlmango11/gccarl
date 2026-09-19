package semantic

import (
	"fmt"
	"strconv"

	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
)

type builder struct {
	vars    map[cparser.IDEN]Type
	funcs   map[cparser.IDEN]Type
	structs map[TypeName]Type // todo local structs
	strs    []string
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
		vars: make(map[cparser.IDEN]Type),
		funcs: map[cparser.IDEN]Type{
			"do_syscall": {
				Kind: KindVoid,
			},
			"assert": {
				Kind: KindVoid,
			},
		},
		structs: make(map[TypeName]Type),
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

func (b *builder) toFuncDef(f *cparser.DecDef_FuncDefOption) (*FuncDef, error) {
	// TODO func scoped types

	returnType, err := b.toReturnType(f.Type)
	if err != nil {
		return nil, err
	}

	b.funcs[f.IDEN] = returnType

	locals := map[cparser.IDEN]Type{}

	for _, l := range f.Line {
		switch l.Type {
		case cparser.LineTypeControl: // TODO
		case cparser.LineTypeStatement:
			s := l.Statement.StatementComma.Statement.Statement

			if s.DecAssign != nil {
				decAssign := s.DecAssign.DecAssign.Standard
				err := b.declareVar(locals, decAssign.Type, decAssign.VariableDef, false)
				if err != nil {
					return nil, err
				}
			} else if s.VarDec != nil {
				err := b.declareVar(locals, s.VarDec.VarDec.VarDec.Type, s.VarDec.VarDec.VarDec.VariableDef, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	var paramDefs []*ParamDef
	if f.ParamsDef != nil {
		ps := []*cparser.ParamDef{f.ParamsDef.Params.ParamDef}
		for _, x := range f.ParamsDef.Params.CommaParamDef {
			ps = append(ps, x.Param.ParamDef)
		}

		for _, astParam := range ps {
			pd, err := b.toParamDec(astParam)
			if err != nil {
				return nil, err
			}

			err = b.declareVar(locals, astParam.Param.Type, astParam.Param.VariableDef, true)
			if err != nil {
				return nil, err
			}

			paramDefs = append(paramDefs, pd)
		}
	}

	var lines []*Line
	for _, l := range f.Line {
		line, err := b.toLine(locals, l)
		if err != nil {
			return nil, err
		}

		if line != nil {
			lines = append(lines, line)
		}
	}

	localsCast := map[VarName]Type{}
	for k, v := range locals {
		localsCast[VarName(k)] = v
	}

	return &FuncDef{
		ReturnType: returnType,
		Name:       FuncName(f.IDEN),
		Locals:     localsCast,
		Params:     paramDefs,
		Lines:      lines,
	}, nil
}

func (b *builder) toLine(locals map[cparser.IDEN]Type, l *cparser.Line) (*Line, error) {
	switch l.Type {
	case cparser.LineTypeControl:
		c, err := b.toControl(locals, l.Control.Control)
		if err != nil {
			return nil, err
		}

		return &Line{
			Control: c,
		}, nil
	case cparser.LineTypeStatement:
		s, err := b.toStatement(locals, l.Statement.StatementComma.Statement.Statement)
		if err != nil {
			return nil, err
		}

		if s == nil {
			return nil, nil
		}

		return &Line{
			Statement: s,
		}, nil
	}

	panic("unreachable")
}

func (b *builder) declareVar(vars map[cparser.IDEN]Type, astType *cparser.Type, v *cparser.VariableDef, isParam bool) error {
	typ, err := b.toType(astType, v, isParam)
	if err != nil {
		return err
	}

	varName := varDefName(v)

	_, ok := vars[varName]
	if ok {
		return fmt.Errorf("variable %s already declared", varName)
	}

	vars[varName] = typ
	return nil
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
		st, ok := b.structs[TypeName(typ.Struct.IDEN)]
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

func (b *builder) toParamDec(p *cparser.ParamDef) (*ParamDef, error) {
	typ, err := b.toType(p.Param.Type, p.Param.VariableDef, true)
	if err != nil {
		return nil, err
	}

	return &ParamDef{
		Type: typ,
		Name: VarName(p.Param.VariableDef.Variable.IDEN),
	}, nil
}

func (b *builder) toStatement(vars map[cparser.IDEN]Type, s *cparser.Statement) (*Statement, error) {
	switch s.Type {
	case cparser.StatementTypeDecAssign:
		a, err := b.toDecAssign(vars, s.DecAssign.DecAssign.Standard)
		if err != nil {
			return nil, err
		}

		return &Statement{
			DeclareInit: a,
		}, nil
	case cparser.StatementTypeVarDec:
		// handled in the normal local func vars
		return nil, nil
	case cparser.StatementTypeReturn:
		expr, err := b.toExpr(s.Return.Expr, vars)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Return: expr,
		}, nil
	case cparser.StatementTypeExpr:
		expr, err := b.toExpr(s.Expr.Expr, vars)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Expr: expr,
		}, nil
	}

	panic("invalid statement: " + s.Type)
}

func (b *builder) toDecAssign(vars map[cparser.IDEN]Type, a *cparser.DecAssign_StandardOption) (*DeclareInit, error) {
	varName := varDefName(a.VariableDef)

	varType, ok := vars[varName]
	if !ok {
		varType, ok = vars[varName] // todo global
		if !ok {
			return nil, fmt.Errorf("variable %s not declared", varName)
		}
	}

	init, err := b.toInitialiser(varType, a.Initialiser, vars)
	if err != nil {
		return nil, err
	}

	return &DeclareInit{
		Type:        varType,
		Name:        VarName(varName),
		Initialiser: init,
	}, nil
}

func (b *builder) toVarDirect(v *cparser.SubVariableAccess) VarRead {
	var index []int
	for _, ia := range v.V.ArrayIndexAccess {
		i, err := strconv.Atoi(string(ia.ArrayIndex.NUM))
		if err != nil {
			panic(fmt.Sprintf("invalid array index %v: %v", ia.ArrayIndex.NUM, err))
		}

		index = append(index, i)
	}

	return VarRead{
		Name:  VarName(v.V.IDEN),
		Index: index,
	}
}

func (b *builder) getVarType(vars map[cparser.IDEN]Type, name cparser.IDEN) (Type, error) {
	typ, ok := vars[name]
	if !ok {
		typ, ok = b.vars[name]
		if !ok {
			return Type{}, fmt.Errorf("variable %s not declared", name)
		}
	}

	return typ, nil
}

func (b *builder) toExpr(expr *cparser.Expr, locals map[cparser.IDEN]Type) (*Expr, error) {
	switch expr.Type {
	case cparser.ExprTypeComp:
		compExpr := expr.Comp.CompExpr.CompExpr

		rightExpr, err := b.toExpr(compExpr.Expr, locals)
		if err != nil {
			return nil, err
		}

		leftExpr, err := b.fromSubExpr(compExpr.SubExpr, locals)
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
		return b.fromSubExpr(expr.SubExpr.SubExpr, locals)
	}

	panic("invalid expression: " + string(expr.Type))
}

func (b *builder) fromSubExpr(sub *cparser.SubExpr, locals map[cparser.IDEN]Type) (*Expr, error) {
	switch sub.Type {
	case cparser.SubExprTypeFuncCall:
		fc, err := b.toFuncCall(sub.FuncCall, locals)
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
		expr, err := b.toExpr(sub.AddressOf.Expr, locals)
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
		expr, err := b.toExpr(sub.Deref.Expr, locals)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Type:  expr.Type,
			Deref: expr,
		}, nil
	case cparser.SubExprTypeVariable:
		return b.toVarExpr(locals, sub.Variable)
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
			//case ast.ValueTypeCompLit:
			//	av := v.Value.CompLit.CompEntries
			//
			//	typ, err := b.toVarType(v.Value.CompLit.Type, v.Value.CompLit.ArrayIndexDef, false)
			//	if err != nil {
			//		return nil, err
			//	}
			//
			//	compLit, err := b.toInitList(av, locals)
			//	if err != nil {
			//		return nil, err
			//	}
			//
			//	return &Expr{
			//		Type:        typ,
			//		CompLiteral: compLit,
			//	}, nil
		}
	}

	panic("invalid sub expression: " + string(sub.Type))
}

func (b *builder) toVarExpr(vars map[cparser.IDEN]Type, v *cparser.SubExpr_VariableOption) (*Expr, error) {
	e := &Expr{}

	typ, err := b.getVarType(vars, v.SubVariableAccess.V.IDEN)
	if err != nil {
		return nil, err
	}

	e.Type = typ
	e.Var = &VarExpr{
		Type:   typ,
		Fields: []VarRead{b.toVarDirect(v.SubVariableAccess)},
	}

	name := v.SubVariableAccess.V.IDEN

	for _, x := range v.InnerSubVariableAccess {
		switch x.Type {
		case cparser.InnerSubVariableAccessTypeArrow:
			if typ.Kind != KindPointer {
				return nil, fmt.Errorf("cannot use %v->%v for non-pointer", name, x.Arrow.SubVariableAccess.V.IDEN)
			}

			if typ.SubType.Kind != KindStruct {
				return nil, fmt.Errorf("%v is not a struct", name)
			}

			f, ok := typ.SubType.Struct.Field(VarName(x.Arrow.SubVariableAccess.V.IDEN))
			if !ok {
				return nil, fmt.Errorf("%v does not have field %v", name, x.Arrow.SubVariableAccess.V.IDEN)
			}

			name = x.Arrow.SubVariableAccess.V.IDEN

			e = &Expr{
				Type: f.Type,
				Var: &VarExpr{
					Expr: &Expr{
						Type:  *typ.SubType,
						Deref: e,
					},
					Fields: []VarRead{
						b.toVarDirect(x.Arrow.SubVariableAccess),
					},
					Type: f.Type,
				},
			}

		case cparser.InnerSubVariableAccessTypeDot:
			if typ.Kind != KindStruct {
				return nil, fmt.Errorf("%v is not a struct", name)
			}

			f, ok := typ.Struct.Field(VarName(x.Dot.SubVariableAccess.V.IDEN))
			if !ok {
				return nil, fmt.Errorf("%v does not have field %v", name, x.Dot.SubVariableAccess.V.IDEN)
			}

			typ = f.Type
			name = x.Dot.SubVariableAccess.V.IDEN

			e.Type = f.Type
			e.Var.Fields = append(e.Var.Fields, b.toVarDirect(x.Dot.SubVariableAccess))
		}
	}

	return e, nil
}

func (b *builder) toArrayLit(av *cparser.ArrayEntries_EntriesOption, locals map[cparser.IDEN]Type) ([]*Expr, error) {
	exprsNodes := []*cparser.Expr{av.Expr}
	for _, e := range av.CommaExpr {
		exprsNodes = append(exprsNodes, e.CommaExpr.Expr)
	}

	var exprs []*Expr

	for _, astExpr := range exprsNodes {
		expr, err := b.toExpr(astExpr, locals)
		if err != nil {
			return nil, err
		}

		// type check TODO

		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func (b *builder) toInitList(typ Type, e *cparser.CompEntries, locals map[cparser.IDEN]Type) (*InitList, error) {
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
		var names []VarName
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
				names = append(names, VarName(n.C.IDEN))
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

		init, err := b.toInitialiser(subType, initNode, locals)
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

func fieldType(t Type, names []VarName) (Type, bool) {
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

func (b *builder) toFuncCall(call *cparser.SubExpr_FuncCallOption, locals map[cparser.IDEN]Type) (*FuncCall, error) {
	var params []*cparser.Expr
	if call.Params != nil {
		params = append(params, call.Params.Params.Expr)

		for _, x := range call.Params.Params.CommaExpr {
			params = append(params, x.CommaExpr.Expr)
		}
	}

	var args []*Expr
	for _, e := range params {
		arg, err := b.toExpr(e, locals)
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

func (b *builder) toIf(locals map[cparser.IDEN]Type, i *cparser.Control_IfOption) (*If, error) {
	expr, err := b.toExpr(i.Expr, locals)
	if err != nil {
		return nil, err
	}

	if expr.Type.Prim != PrimBool {
		return nil, fmt.Errorf("if condition must be a boolean")
	}

	ifLines, err := b.toLines(locals, i.BlockOrLine)
	if err != nil {
		return nil, err
	}

	var elseLines []*Line

	if i.Else != nil {
		elseLines, err = b.toLines(locals, i.Else.Else.BlockOrLine)
		if err != nil {
			return nil, err
		}
	}

	return &If{
		Condition: expr,
		Lines:     ifLines,
		ElseLines: elseLines,
	}, nil
}

func (b *builder) toWhile(locals map[cparser.IDEN]Type, w *cparser.Control_WhileOption) (*While, error) {
	expr, err := b.toExpr(w.Expr, locals)
	if err != nil {
		return nil, err
	}

	if expr.Type.Prim != PrimBool {
		return nil, fmt.Errorf("if condition must be a boolean")
	}

	var lines []*Line
	for _, n := range w.Line {
		l, err := b.toLine(locals, n)
		if err != nil {
			return nil, err
		}

		lines = append(lines, l)
	}

	return &While{
		Condition: expr,
		Lines:     lines,
	}, nil
}

func (b *builder) toFor(locals map[cparser.IDEN]Type, f *cparser.Control_ForOption) (*For, error) {
	init, err := b.toStatement(locals, f.Statement0)
	if err != nil {
		return nil, err
	}

	cond, err := b.toStatement(locals, f.Statement1)
	if err != nil {
		return nil, err
	}

	action, err := b.toStatement(locals, f.Statement2)
	if err != nil {
		return nil, err
	}

	lines, err := b.toLines(locals, f.BlockOrLine)
	if err != nil {
		return nil, err
	}

	return &For{
		Init:      init,
		Condition: cond,
		Action:    action,
		Lines:     lines,
	}, nil
}

func (b *builder) toControl(locals map[cparser.IDEN]Type, c *cparser.Control) (*Control, error) {
	switch c.Type {
	case cparser.ControlTypeIf:
		ifC, err := b.toIf(locals, c.If)
		if err != nil {
			return nil, err
		}

		return &Control{
			If: ifC,
		}, nil
	case cparser.ControlTypeWhile:
		w, err := b.toWhile(locals, c.While)
		if err != nil {
			return nil, err
		}

		return &Control{
			While: w,
		}, nil
	case cparser.ControlTypeFor:
		f, err := b.toFor(locals, c.For)
		if err != nil {
			return nil, err
		}

		return &Control{
			For: f,
		}, nil
	}

	panic("invalid control")
}

func (b *builder) toLines(locals map[cparser.IDEN]Type, e *cparser.BlockOrLine) ([]*Line, error) {
	switch e.Type {
	case cparser.BlockOrLineTypeLine:
		l, err := b.toLine(locals, e.Line.Line)
		if err != nil {
			return nil, err
		}

		return []*Line{l}, nil
	case cparser.BlockOrLineTypeBlock:
		var all []*Line
		for _, n := range e.Block.Line {
			l, err := b.toLine(locals, n)
			if err != nil {
				return nil, err
			}

			all = append(all, l)
		}

		return all, nil
	}

	panic("invalid block or lines")
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
	structVars := map[cparser.IDEN]Type{}

	for _, a := range td.VarDecColon {
		t := a.C.VarDec.VarDec.Type
		def := a.C.VarDec.VarDec.VariableDef

		err := b.declareVar(structVars, t, def, false) // checks for dups
		if err != nil {
			return err
		}

		typ, err := b.toType(t, def, false)
		if err != nil {
			return err
		}

		vars = append(vars, StructField{
			Name: VarName(varDefName(a.C.VarDec.VarDec.VariableDef)),
			Type: typ,
		})
	}

	// todo check exists

	st := StructType{
		Name:   TypeName(td.IDEN),
		Fields: vars,
	}

	b.structs[TypeName(td.IDEN)] = Type{
		Kind:   KindStruct,
		Struct: st,
	}

	return nil
}

func (b *builder) toInitialiser(typ Type, init *cparser.Initialiser, vars map[cparser.IDEN]Type) (*Initialiser, error) {
	switch init.Type {
	case cparser.InitialiserTypeExpr:
		expr, err := b.toExpr(init.Expr.Expr, vars)
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
		cl, err := b.toInitList(typ, init.List.CompEntries, vars)
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

//func (b *builder) rewriteAsStructExpr(structType Type, name ast.IDEN, expr *Expr) *Expr {
//	sd := b.structs[name]
//
//	newExpr := &Expr{
//		Type: structType,
//	}
//
//	for i := range expr.CompLiteral {
//		newFieldExpr := checkCompatibility(sd.Fields[i].Type, expr.CompLiteral[i])
//		newExpr.CompLiteral = append(newExpr.CompLiteral, newFieldExpr)
//	}
//
//	return newExpr
//}

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
