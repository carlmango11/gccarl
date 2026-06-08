package semantic

import (
	"fmt"
	"strconv"

	"github.com/carlmango11/gccarl/gccarl/generated/ast"
)

type builder struct {
	vars  map[ast.IDEN]PrimitiveType
	funcs map[ast.IDEN]Type
	strs  []string
}

var opType = map[ast.OperatorType]CompareOp{
	ast.OperatorTypeLess:  OpLessThan,
	ast.OperatorTypeEqual: OpEquals,
}

func Build(program *ast.Main) (*Program, error) {
	b := &builder{
		vars: make(map[ast.IDEN]PrimitiveType),
		funcs: map[ast.IDEN]Type{
			"do_syscall": {
				Kind: KindVoid,
			},
		},
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
				err := b.declareVar(locals, decAssign.Type, decAssign.Variable, false)
				if err != nil {
					return nil, err
				}
			} else if s.VarDec != nil {
				err := b.declareVar(locals, s.VarDec.Type, s.VarDec.Variable, false)
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

			err = b.declareVar(locals, astParam.Param.Type, astParam.Param.Variable, true)
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

func (b *builder) declareVar(vars map[ast.IDEN]Type, astType *ast.Type, v *ast.Variable, isParam bool) error {
	_, ok := vars[v.Variable.IDEN]
	if ok {
		return fmt.Errorf("variable %s already declared", v.Variable.IDEN)
	}

	typ, err := b.toType(astType, v, isParam)
	if err != nil {
		return err
	}

	vars[v.Variable.IDEN] = typ
	return nil
}

func (b *builder) toType(typ *ast.Type, v *ast.Variable, isParam bool) (Type, error) {
	if len(v.Variable.ArrayIndex) > 0 {
		if isParam {
			// arrays decay to pointers in params
			return Type{
				Kind: KindPointer,
				Prim: astTypeToPrim(typ),
			}, nil
		}

		return b.toArrayType(typ, v)
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

func (b *builder) toArrayType(typ *ast.Type, v *ast.Variable) (Type, error) {
	var size int

	sizeStr := v.Variable.ArrayIndex[0].ArrayIndex.NUM

	if sizeStr != "" {
		var err error
		size, err = strconv.Atoi(string(sizeStr))
		if err != nil {
			panic("invalid size: " + sizeStr)
		}
	}

	sub, err := b.subType(typ, v)
	if err != nil {
		return Type{}, err
	}

	return Type{
		Kind:      KindArray,
		SubType:   &sub,
		ArraySize: size,
	}, nil
}

func (b *builder) subType(typ *ast.Type, v *ast.Variable) (Type, error) {
	copyVar := *v
	copyVarOption := *v.Variable
	copyVar.Variable = &copyVarOption
	copyVar.Variable.ArrayIndex = copyVar.Variable.ArrayIndex[1:]

	return b.toType(typ, &copyVar, false)
}

func (b *builder) toReturnType(i *ast.Type) (Type, error) {
	kind := KindPrimitive

	var prim PrimitiveType
	switch i.Type {
	case ast.TypeTypeInt:
		prim = PrimInt32
	case ast.TypeTypeChar:
		prim = PrimChar
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
	typ, err := b.toType(p.Param.Type, p.Param.Variable, true)
	if err != nil {
		return nil, err
	}

	return &ParamDef{
		Type: typ,
		Name: VarName(p.Param.Variable.Variable.IDEN),
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
	v, ok := vars[a.Variable.Variable.IDEN]
	if !ok {
		v, ok = vars[a.Variable.Variable.IDEN]
		if !ok {
			return nil, fmt.Errorf("variable %s not declared", a.Variable.Variable.IDEN)
		}
	}

	indexable := v.Kind == KindPointer || v.Kind == KindArray
	if len(a.Variable.Variable.ArrayIndex) > 0 && !indexable {
		return nil, fmt.Errorf("variable %s is not indexable", a.Variable.Variable.IDEN)
	}

	expr, err := b.toExpr(a.Expr, vars)
	if err != nil {
		return nil, err
	}

	if !v.Equals(expr.Type) {
		if compatibleTypes(v, expr.Type) {
			expr = &Expr{
				Cast: &Cast{
					To:   v,
					Expr: expr,
				},
			}
		}
	}

	return &Assign{
		Name: VarName(a.Variable.Variable.IDEN),
		Expr: expr,
	}, nil
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

		return &Expr{
			Type: boolType(),
			Compare: &CompareOpExpr{
				Left:  e1,
				Op:    opType[compExpr.Operator.Type],
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
			b.strs = append(b.strs, s)

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
			vr := v.Value.Variable.Variable.Variable
			name := vr.IDEN

			typ, ok := locals[name]
			if !ok {
				return nil, fmt.Errorf("variable %s not declared", name)
			}

			if len(vr.ArrayIndex) > 0 {
				// TODO: handle nested
				i, err := strconv.Atoi(string(vr.ArrayIndex[0].ArrayIndex.NUM))
				if err != nil {
					return nil, err
				}

				return &Expr{
					Type: *typ.SubType,
					IndexedVar: &IndexedVar{
						Index: i,
						Name:  VarName(name),
					},
				}, nil
			}

			if typ.Kind == KindArray {
				// decay to pointer
				return &Expr{
					Type: Type{
						Kind:    KindPointer,
						SubType: typ.SubType,
					},
					AddressOf: VarName(name),
				}, nil
			}

			return &Expr{
				Type: typ,
				Var:  VarName(name),
			}, nil
		case ast.ValueTypeCompLit:
			av := v.Value.CompLit.CompositeLiteral.ArrayVal

			exprs, err := b.toCompLit(av, locals)
			if err != nil {
				return nil, err
			}

			return &Expr{
				Type: Type{
					Kind:    KindArray,
					SubType: &exprs[0].Type,
				},
				CompLiteral: exprs,
			}, nil
		}
	}

	panic("invalid sub expression: " + string(sub.Type))
}

func (b *builder) toCompLit(av *ast.CompositeLiteral_ArrayValOption, locals map[ast.IDEN]Type) ([]*Expr, error) {
	exprsNodes := []*ast.Expr{av.CompEntries.Entries.Expr}
	for _, e := range av.CompEntries.Entries.CommaExpr {
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

	var ifLines []*Line
	for _, n := range i.Line {
		l, err := b.toLine(locals, n)
		if err != nil {
			return nil, err
		}

		ifLines = append(ifLines, l)
	}

	var elseLines []*Line
	var elseIfExpr *Expr

	if i.Else != nil {
		if i.Else.Else.ElseIf != nil {
			elseIfN := i.Else.Else.ElseIf

			var err error
			elseIfExpr, err = b.toExpr(elseIfN.ElseIf.Expr, locals)
			if err != nil {
				return nil, err
			}

			if elseIfExpr.Type.Prim != PrimBool {
				return nil, fmt.Errorf("if condition must be a boolean")
			}
		}

		for _, n := range i.Else.Else.Line {
			l, err := b.toLine(locals, n)
			if err != nil {
				return nil, err
			}

			elseLines = append(elseLines, l)
		}
	}

	return &If{
		Condition: expr,
		Lines:     ifLines,
		ElseIf:    elseIfExpr,
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
	panic("implement me")
}
