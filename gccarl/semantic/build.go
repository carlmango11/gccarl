package semantic

import (
	"fmt"
	"strconv"

	"github.com/carlmango11/gccarl/gccarl/generated/ast"
)

type builder struct {
	vars    map[ast.IDEN]PrimitiveType
	funcs   map[ast.IDEN]Type
	structs map[TypeName]Type // todo local structs
	strs    []string
}

var compareOp = map[ast.OperatorType]CompareOp{
	ast.OperatorTypeLess:  OpLessThan,
	ast.OperatorTypeEqual: OpEquals,
}

var numericalOps = map[ast.OperatorType]NumericOp{
	ast.OperatorTypePlus: NumOpAdd,
}

func Build(program *ast.Main) (*Program, error) {
	b := &builder{
		vars: make(map[ast.IDEN]PrimitiveType),
		funcs: map[ast.IDEN]Type{
			"do_syscall": {
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

func (b *builder) build(p *ast.Main) (*Program, error) {
	var funcDecs []*FuncDef

	for _, dd := range p.Main.DecDef {
		switch dd.Type {
		case ast.DecDefTypeFuncDef:
			f, err := b.toFuncDef(dd.FuncDef)
			if err != nil {
				return nil, err
			}

			funcDecs = append(funcDecs, f)
		case ast.DecDefTypeTypeDef:
			b.defineType(dd.TypeDef)
		}
	}

	return &Program{
		FuncDefs: funcDecs,
		Strings:  b.strs,
	}, nil
}

func (b *builder) toFuncDef(f *ast.DecDef_FuncDefOption) (*FuncDef, error) {
	// TODO func scoped types

	returnType, err := b.toReturnType(f.Type)
	if err != nil {
		return nil, err
	}

	b.funcs[f.IDEN] = returnType

	locals := map[ast.IDEN]Type{}

	for _, l := range f.Line {
		switch l.Type {
		case ast.LineTypeControl: // TODO
		case ast.LineTypeStatement:
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
		ps := []*ast.ParamDef{f.ParamsDef.Params.ParamDef}
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

func (b *builder) toLine(locals map[ast.IDEN]Type, l *ast.Line) (*Line, error) {
	switch l.Type {
	case ast.LineTypeControl:
		c, err := b.toControl(locals, l.Control.Control)
		if err != nil {
			return nil, err
		}

		return &Line{
			Control: c,
		}, nil
	case ast.LineTypeStatement:
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

func (b *builder) declareVar(vars map[ast.IDEN]Type, astType *ast.Type, v *ast.VariableDef, isParam bool) error {
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

func varDefName(v *ast.VariableDef) ast.IDEN {
	switch v.Type {
	case ast.VariableDefTypePointer:
		p := v.Pointer.VariableDef
		for {
			if p.Type == ast.VariableDefTypeVariable {
				return p.Variable.IDEN
			}

			p = p.Pointer.VariableDef
		}
	case ast.VariableDefTypeVariable:
		return v.Variable.IDEN
	default:
		panic(fmt.Sprintf("unhandled variable type %s", v.Type))
	}
}

func (b *builder) toType(typ *ast.Type, v *ast.VariableDef, isParam bool) (Type, error) {
	switch v.Type {
	case ast.VariableDefTypePointer:
		sub, err := b.toType(typ, v.Pointer.VariableDef, false) // no need for array decay check anymore
		if err != nil {
			return Type{}, err
		}

		return Type{
			Kind:    KindPointer,
			SubType: &sub,
		}, nil
	case ast.VariableDefTypeVariable:
		return b.toVarType(typ, v.Variable.ArrayIndexDef, isParam)
	}

	panic("invalid type")
}

func (b *builder) toVarType(typ *ast.Type, arrs []*ast.ArrayIndexDef, isParam bool) (Type, error) {
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

	if typ.Type == ast.TypeTypeStruct {
		st, ok := b.structs[TypeName(typ.Struct.IDEN)]
		if !ok {
			panic("missing struct")
		}

		return st, nil
	}

	return Type{
		Kind: KindPrimitive,
		Prim: astTypeToPrim(typ),
	}, nil
}

func astTypeToPrim(typ *ast.Type) PrimitiveType {
	switch typ.Type {
	case ast.TypeTypeInt:
		return PrimInt32
	case ast.TypeTypeChar:
		return PrimChar
	default:
		// custom
		panic("impl")
	}
}

func (b *builder) toReturnType(i *ast.Type) (Type, error) {
	kind := KindPrimitive

	var prim PrimitiveType
	switch i.Type {
	case ast.TypeTypeInt:
		prim = PrimInt32
	case ast.TypeTypeChar:
		prim = PrimChar
	case ast.TypeTypeVoid:
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

func (b *builder) toParamDec(p *ast.ParamDef) (*ParamDef, error) {
	typ, err := b.toType(p.Param.Type, p.Param.VariableDef, true)
	if err != nil {
		return nil, err
	}

	return &ParamDef{
		Type: typ,
		Name: VarName(p.Param.VariableDef.Variable.IDEN),
	}, nil
}

func (b *builder) toStatement(vars map[ast.IDEN]Type, s *ast.Statement) (*Statement, error) {
	switch s.Type {
	case ast.StatementTypeDecAssign:
		a, err := b.toDecAssign(vars, s.DecAssign.DecAssign.Standard)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
		}, nil
	case ast.StatementTypeAssign:
		a, err := b.toAssign(vars, s.Assign.VariableAccess, s.Assign.Expr)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
		}, nil
	case ast.StatementTypeVarDec:
		// handled in the normal local func vars
		return nil, nil
	case ast.StatementTypeReturn:
		expr, err := b.toExpr(s.Return.Expr, vars)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Return: expr,
		}, nil
	case ast.StatementTypeExpr:
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

func (b *builder) toDecAssign(vars map[ast.IDEN]Type, a *ast.DecAssign_StandardOption) (*Assign, error) {
	varName := varDefName(a.VariableDef)

	varType, ok := vars[varName]
	if !ok {
		varType, ok = vars[varName] // todo global
		if !ok {
			return nil, fmt.Errorf("variable %s not declared", varName)
		}
	}

	expr, err := b.toExpr(a.Expr, vars)
	if err != nil {
		return nil, err
	}

	if varType.Kind == KindStruct {
		// TODO this is shit
		panic("impl")
		//expr = b.rewriteAsStructExpr(varType, a.Type.Struct.IDEN, expr)
	}

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

	return &Assign{
		Var: VarWrite{
			Direct: []VarRead{
				{Name: VarName(varName)},
			},
		},
		Expr: expr,
	}, nil
}

func (b *builder) toVarWrite(v *ast.VariableAccess) VarWrite {
	switch v.Type {
	case ast.VariableAccessTypeDeref:
		a := b.toVarWrite(v.Deref.VariableAccess)

		return VarWrite{
			Deref: &a,
		}
	case ast.VariableAccessTypeVariable:
		return VarWrite{
			Direct: b.toVarDirects(v),
		}
	}

	panic("invalid var access: " + v.Type)
}

func (b *builder) toVarRead(v *ast.VariableAccess_VariableOption) VarRead {
	directs := []VarRead{b.toVarDirect(v.SubVariableAccess)}

	for _, x := range v.InnerSubVariableAccess {
		if x.Type == ast.InnerSubVariableAccessTypeArrow {
			return VarRead{
				Deref: VarRead{},
			}
		}
	}
}

func (b *builder) toVarDirects(v *ast.VariableAccess_VariableOption) []VarRead {
	all := []*ast.SubVariableAccess_VOption{v.SubVariableAccess.V}

	for _, x := range v.InnerSubVariableAccess {
		all = append(all, x.)
	}

	dws := []VarRead{b.toVarDirect(v.Variable.SubVariableAccess, false)}

	for _, x := range v.Variable.InnerSubVariableAccess {
		switch x.Type {
		case ast.InnerSubVariableAccessTypeDot:
			dws = append(dws, b.toVarDirect(x.Dot.SubVariableAccess, false))
		case ast.InnerSubVariableAccessTypeArrow:
			dws = append(dws, b.toVarDirect(x.Arrow.SubVariableAccess, true))
		}
	}

	return dws
}

func (b *builder) toVarDirect(v *ast.SubVariableAccess) VarRead {
	var index []int
	for _, ia := range v.V.ArrayIndexAccess {
		i, err := strconv.Atoi(string(ia.ArrayIndex.NUM))
		if err != nil {
			panic(fmt.Sprintf("invalid array index %v: %v", ia.ArrayIndex.NUM, err))
		}

		index = append(index, i)
	}

	return VarRead{
		Name:    VarName(v.V.IDEN),
		Index:   index,
	}
}

func (b *builder) getVarType(vars map[ast.IDEN]Type, v *ast.VariableAccess_VariableOption) (Type, error) {
	typ, ok := vars[v.SubVariableAccess.V.IDEN]
	if !ok {
		typ, ok = vars[v.SubVariableAccess.V.IDEN] // todo global
		if !ok {
			return Type{}, fmt.Errorf("variable %s not declared", v.SubVariableAccess.V.IDEN)
		}
	}

	fieldName := v.SubVariableAccess.V.IDEN

	for _, iv := range v.InnerSubVariableAccess {
		var thisStruct StructType

		switch iv.Type {
		case ast.InnerSubVariableAccessTypeDot:
			if typ.Kind != KindStruct {
				return Type{}, fmt.Errorf("cannot use %v.%v for non-struct", fieldName, iv.Dot.SubVariableAccess.V.IDEN)
			}

			thisStruct = typ.Struct
			fieldName = iv.Dot.SubVariableAccess.V.IDEN
		case ast.InnerSubVariableAccessTypeArrow:
			if typ.Kind != KindPointer {
				return Type{}, fmt.Errorf("cannot use %v->%v for non-pointer", fieldName, iv.Arrow.SubVariableAccess.V.IDEN)
			}

			if typ.SubType.Kind != KindStruct {
				return Type{}, fmt.Errorf("%v-%v is not a struct", fieldName, iv.Arrow.SubVariableAccess.V.IDEN)
			}

			thisStruct = typ.SubType.Struct
			fieldName = iv.Arrow.SubVariableAccess.V.IDEN
		}

		for _, f := range thisStruct.Fields {
			if f.Name == VarName(fieldName) {
				typ = f.Type
				break
			}
		}
	}

	return typ, nil
}

func (b *builder) toAssign(vars map[ast.IDEN]Type, v *ast.VariableAccess, e *ast.Expr) (*Assign, error) {
	if v.Type == ast.VariableAccessTypeAddressOf {
		return nil, fmt.Errorf("cannot write to &")
	}

	varType, err := b.getVarType(vars, v.Variable)
	if err != nil {
		return nil, err
	}

	// TODO re-add
	//indexable := varType.Kind == KindPointer || varType.Kind == KindArray
	//if len(v.Variable.ArrayIndexAccess) > 0 && !indexable {
	//	return nil, fmt.Errorf("variable %s is not indexable", v.Variable.IDEN)
	//}

	expr, err := b.toExpr(e, vars)
	if err != nil {
		return nil, err
	}

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

	return &Assign{
		Var:  b.toVarWrite(v),
		Expr: expr,
	}, nil
}

func checkCompatibility(varType Type, expr *Expr) *Expr {
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

	return expr
}

func (b *builder) toExpr(expr *ast.Expr, locals map[ast.IDEN]Type) (*Expr, error) {
	switch expr.Type {
	case ast.ExprTypeComp:
		compExpr := expr.Comp.CompExpr.CompExpr

		e1, err := b.toExpr(compExpr.Expr, locals)
		if err != nil {
			return nil, err
		}

		e2, err := b.fromSubExpr(compExpr.SubExpr, locals)
		if err != nil {
			return nil, err
		}

		op, ok := compareOp[compExpr.Operator.Type]
		if ok {
			return &Expr{
				Type: boolType(),
				Compare: &CompareOpExpr{
					Left:  e1,
					Op:    op,
					Right: e2,
				},
			}, nil
		}

		// todo: check for casts

		return &Expr{
			Type: e1.Type,
			Numeric: &NumericOpExpr{
				Left:  e1,
				Op:    numericalOps[compExpr.Operator.Type],
				Right: e2,
			},
		}, nil
	case ast.ExprTypeSubExpr:
		return b.fromSubExpr(expr.SubExpr.SubExpr, locals)
	}

	panic("invalid expression: " + string(expr.Type))
}

func (b *builder) fromSubExpr(sub *ast.SubExpr, locals map[ast.IDEN]Type) (*Expr, error) {
	switch sub.Type {
	case ast.SubExprTypeFuncCall:
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
	case ast.SubExprTypeValue:
		v := sub.Value

		switch v.Value.Type {
		case ast.ValueTypeStr:
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
		case ast.ValueTypeInt:
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
		case ast.ValueTypeChar:
			return &Expr{
				Type: charType(),
				Literal: &Literal{
					Char: v.Value.Char.CHAR[1],
				},
			}, nil
		case ast.ValueTypeVariable:
			typ, vr, err := b.fromValueVar(locals, v.Value.Variable.VariableAccess)
			if err != nil {
				return nil, err
			}

			return &Expr{
				Type: typ,
				Var:  &vr,
			}, nil
		case ast.ValueTypeArrayLit:
			av := v.Value.ArrayLit.ArrayEntries.Entries

			exprs, err := b.toArrayLit(av, locals)
			if err != nil {
				return nil, err
			}

			return &Expr{
				Type: Type{
					Kind:    KindArray,
					SubType: &exprs[0].Type,
				},
				ArrayLiteral: exprs,
			}, nil
		case ast.ValueTypeCompLit:
			av := v.Value.CompLit.CompEntries

			typ, err := b.toVarType(v.Value.CompLit.Type, v.Value.CompLit.ArrayIndexDef, false)
			if err != nil {
				return nil, err
			}

			compLit, err := b.toCompLit(av, locals)
			if err != nil {
				return nil, err
			}

			return &Expr{
				Type:        typ,
				CompLiteral: compLit,
			}, nil
		}
	}

	panic("invalid sub expression: " + string(sub.Type))
}

func (b *builder) fromValueVar(vars map[ast.IDEN]Type, v *ast.VariableAccess) (Type, VarRead, error) {
	switch v.Type {
	case ast.VariableAccessTypeDeref:
		typ, subRead, err := b.fromValueVar(vars, v.Deref.VariableAccess)
		if err != nil {
			return Type{}, VarRead{}, err
		}

		vr := VarRead{
			Deref: &subRead,
		}

		// type becomes whatever we derefed
		typ = *typ.SubType

		return typ, vr, nil
	case ast.VariableAccessTypeAddressOf:
		typ, subRead, err := b.fromValueVar(vars, v.AddressOf.VariableAccess)
		if err != nil {
			return Type{}, VarRead{}, err
		}

		if subRead.Direct == nil {
			return Type{}, VarRead{}, fmt.Errorf("cannot take address of address")
		}

		returnType := Type{
			Kind:    KindPointer,
			SubType: &typ,
		}

		vr := VarRead{
			AddressOf: subRead.Direct,
		}

		return returnType, vr, nil
	case ast.VariableAccessTypeVariable:
		typ, err := b.getVarType(vars, v.Variable)
		if err != nil {
			return Type{}, VarRead{}, err
		}

		if typ.Kind == KindArray {
			// decay to pointer
			returnType := Type{
				Kind:    KindPointer,
				SubType: typ.SubType,
			}

			vr := VarRead{
				AddressOf: b.toVarDirects(v),
			}

			return returnType, vr, nil
		}

		vr := b.toVarRead(v.Variable)

		return typ, vr, nil
	default:
		panic("invalid value type: " + string(v.Type))
	}
}

func (b *builder) toArrayLit(av *ast.ArrayEntries_EntriesOption, locals map[ast.IDEN]Type) ([]*Expr, error) {
	exprsNodes := []*ast.Expr{av.Expr}
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

func (b *builder) toCompLit(e *ast.CompEntries, locals map[ast.IDEN]Type) (*CompLiteral, error) {
	exprsNodes := []*ast.CompEntry{e.Entries.CompEntry}
	for _, e := range e.Entries.CommaCompEntry {
		exprsNodes = append(exprsNodes, e.E.CompEntry)
	}

	var entries []*CompLiteralEntry

	for _, node := range exprsNodes {
		var name VarName
		var exprNode *ast.Expr

		switch node.Type {
		case ast.CompEntryTypeAnon:
			exprNode = node.Anon.Expr
		case ast.CompEntryTypeLabelled:
			name = VarName(node.Labelled.IDEN)
			exprNode = node.Labelled.Expr
		}

		expr, err := b.toExpr(exprNode, locals)
		if err != nil {
			return nil, err
		}

		// type check TODO

		entries = append(entries, &CompLiteralEntry{
			Name: name,
			Expr: expr,
		})
	}

	return &CompLiteral{
		Entries: entries,
	}, nil
}

func (b *builder) toFuncCall(call *ast.SubExpr_FuncCallOption, locals map[ast.IDEN]Type) (*FuncCall, error) {
	var params []*ast.Expr
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

func (b *builder) toIf(locals map[ast.IDEN]Type, i *ast.Control_IfOption) (*If, error) {
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

	elseLines, err := b.toLines(locals, i.Else.Else.BlockOrLine)
	if err != nil {
		return nil, err
	}

	return &If{
		Condition: expr,
		Lines:     ifLines,
		ElseLines: elseLines,
	}, nil
}

func (b *builder) toWhile(locals map[ast.IDEN]Type, w *ast.Control_WhileOption) (*While, error) {
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

func (b *builder) toControl(locals map[ast.IDEN]Type, c *ast.Control) (*Control, error) {
	switch c.Type {
	case ast.ControlTypeIf:
		ifC, err := b.toIf(locals, c.If)
		if err != nil {
			return nil, err
		}

		return &Control{
			If: ifC,
		}, nil
	case ast.ControlTypeWhile:
		w, err := b.toWhile(locals, c.While)
		if err != nil {
			return nil, err
		}

		return &Control{
			While: w,
		}, nil
	}

	panic("invalid control")
}

func (b *builder) toLines(locals map[ast.IDEN]Type, e *ast.BlockOrLine) ([]*Line, error) {
	switch e.Type {
	case ast.BlockOrLineTypeLine:
		l, err := b.toLine(locals, e.Line.Line)
		if err != nil {
			return nil, err
		}

		return []*Line{l}, nil
	case ast.BlockOrLineTypeBlock:
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

func (b *builder) defineType(d *ast.DecDef_TypeDefOption) error {
	switch d.TypeDef.Type {
	case ast.TypeDefTypeStructDef:
		return b.toStructDef(d.TypeDef.StructDef)
	}

	panic("invalid type")
}

func (b *builder) toStructDef(td *ast.TypeDef_StructDefOption) error {
	var all []*ast.VarDecColon

	if td.StructBlock.Block.VarDecColon != nil {
		all = append(all, td.StructBlock.Block.VarDecColon)
	}

	for _, x := range td.StructBlock.Block.VarDecComma {
		all = append(all, x.DecComma.VarDecColon)
	}

	var vars []StructField
	structVars := map[ast.IDEN]Type{}

	for _, a := range all {
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
