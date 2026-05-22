package semantic

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/ast"
)

type builder struct {
	vars  map[ast.Identifier]PrimitiveType
	funcs map[ast.Identifier]Type
	strs  []string
}

func Build(program *ast.Program) (*Program, error) {
	b := &builder{
		vars:  make(map[ast.Identifier]PrimitiveType),
		funcs: make(map[ast.Identifier]Type),
	}

	b.funcs["print"] = Type{
		Kind: KindVoid,
	}

	return b.build(program)
}

func (b *builder) build(p *ast.Program) (*Program, error) {
	var funcDecs []*FuncDef
	for _, astF := range p.FuncDefs {
		f, err := b.toFuncDef(astF)
		if err != nil {
			return nil, err
		}

		funcDecs = append(funcDecs, f)
	}

	return &Program{
		Imports:  p.Imports,
		FuncDefs: funcDecs,
		Strings:  b.strs,
	}, nil
}

func (b *builder) toFuncDef(f *ast.FuncDef) (*FuncDef, error) {
	// TODO func scoped types

	returnType, err := b.toType(f.ReturnType)
	if err != nil {
		return nil, err
	}

	b.funcs[f.Name] = returnType

	locals := map[ast.Identifier]Type{}

	for _, s := range f.Statements {
		dec := s.Dec
		if s.DecAssign != nil {
			dec = s.DecAssign.Dec

			//if s.DecAssign.Array != nil {
			//	dec = s.DecAssign.Array.Dec
			//}
		}

		if dec == nil {
			continue
		}

		err := b.declareVar(locals, dec)
		if err != nil {
			return nil, err
		}
	}

	var paramDefs []*ParamDef
	for _, astParam := range f.Params {
		pd, err := b.toParamDec(astParam)
		if err != nil {
			return nil, err
		}

		paramDefs = append(paramDefs, pd)
	}

	var statements []*Statement
	for _, astStatement := range f.Statements {
		s, err := b.toStatement(locals, astStatement)
		if err != nil {
			return nil, err
		}

		statements = append(statements, s)
	}

	var returnExpr *Expr
	if f.ReturnExpr != nil {
		returnExpr, err = b.toExpr(f.ReturnExpr, locals)
		if err != nil {
			return nil, err
		}
	}

	localsCast := map[VarName]Type{}
	for k, v := range locals {
		localsCast[VarName(k)] = v
	}

	return &FuncDef{
		ReturnType: returnType,
		Name:       FuncName(f.Name),
		Locals:     localsCast,
		Params:     paramDefs,
		Statements: statements,
		ReturnExpr: returnExpr,
	}, nil
}

func (b *builder) declareVar(vars map[ast.Identifier]Type, dec *ast.Dec) error {
	_, ok := vars[dec.Name]
	if ok {
		return fmt.Errorf("variable %s already declared", dec.Name)
	}

	typ, err := b.toType(dec.Type)
	if err != nil {
		return err
	}

	vars[dec.Name] = typ
	return nil
}

func (b *builder) toType(i *ast.TypeDef) (Type, error) {
	if len(i.Arrays) > 0 {
		size := i.Arrays[0].Size

		i.Arrays = i.Arrays[1:]
		sub, err := b.toType(i)
		if err != nil {
			return Type{}, err
		}

		return Type{
			Kind:      KindArray,
			SubType:   &sub,
			ArraySize: size,
		}, nil
	}

	kind := KindPrimitive

	var prim PrimitiveType
	switch i.Type.Type {
	case "int":
		prim = PrimInt32
	case "char":
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
	typ, err := b.toType(p.Type)
	if err != nil {
		return nil, err
	}

	return &ParamDef{
		Type: typ,
		Name: VarName(p.Name),
	}, nil
}

func (b *builder) toStatement(vars map[ast.Identifier]Type, s *ast.Statement) (*Statement, error) {
	switch {
	case s.DecAssign != nil:
		switch {
		case s.DecAssign.Assign != nil:
			a, err := b.toAssign(vars, s.DecAssign.Assign)
			if err != nil {
				return nil, err
			}

			return &Statement{
				Assign: a,
			}, nil
			//case s.DecAssign.Array != nil:
			//	a, err := b.toArrayDecAssign(vars, s.DecAssign.Array)
			//	if err != nil {
			//		return nil, err
			//	}
			//
			//	return &Statement{
			//		ArrayAssign: a,
			//	}, nil
		}

	case s.Assign != nil:
		a, err := b.toAssign(vars, s.Assign)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
		}, nil
	case s.Expr != nil:
		expr, err := b.toExpr(s.Expr, vars)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Expr: expr,
		}, nil
	}

	panic("invalid statement")
}

func (b *builder) toAssign(vars map[ast.Identifier]Type, a *ast.Assign) (*Assign, error) {
	v, ok := vars[a.Var.Name]
	if !ok {
		v, ok = vars[a.Var.Name]
		if !ok {
			return nil, fmt.Errorf("variable %s not declared", a.Var.Name)
		}
	}

	indexable := v.Kind == KindPointer
	if len(a.Var.Arrays) > 0 && !indexable {
		return nil, fmt.Errorf("variable %s is not indexable", a.Var.Name)
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
		Name: VarName(a.Var.Name),
		Expr: expr,
	}, nil
}

//func (b *builder) toArrayDecAssign(vars map[ast.Identifier]Type, a *ast.ArrayDecAssign) (*ArrayAssign, error) {
//	v := vars[a.Dec.Name]
//
//	var exprs []*Expr
//
//	for _, entry := range a.Exprs {
//		expr, err := b.toExpr(entry, vars)
//		if err != nil {
//			return nil, err
//		}
//
//		if !v.SubType.Equals(expr.Type) {
//			if compatibleTypes(v, expr.Type) {
//				expr = &Expr{
//					Cast: &Cast{
//						To:   v,
//						Expr: expr,
//					},
//				}
//			}
//		}
//
//		exprs = append(exprs, expr)
//	}
//
//	return &ArrayAssign{
//		Type: v,
//		Name: VarName(a.Dec.Name),
//		Vals: exprs,
//	}, nil
//}

func (b *builder) toExpr(expr *ast.Expr, locals map[ast.Identifier]Type) (*Expr, error) {
	switch {
	case expr.Val != nil:
		v := expr.Val

		switch {
		case v.Str != nil:
			b.strs = append(b.strs, *v.Str)

			return &Expr{
				Type: Type{
					Kind: KindArray,
					SubType: &Type{
						Kind: KindPrimitive,
						Prim: PrimChar,
					},
					ArraySize: len(*v.Str),
				},
				StringID: StringID(len(b.strs)),
			}, nil
		case v.Int != nil:
			// TODO other sizes
			return &Expr{
				Type: int32Type(),
				Literal: &Literal{
					Int32: int32(*v.Int),
				},
			}, nil
		case v.Char != nil:
			return &Expr{
				Type: charType(),
				Literal: &Literal{
					Char: *v.Char,
				},
			}, nil
		case v.Var != nil:
			typ, ok := locals[v.Var.Name]
			if !ok {
				return nil, fmt.Errorf("variable %s not declared", v.Var.Name)
			}

			return &Expr{
				Type: typ,
				Var:  VarName(v.Var.Name),
			}, nil
		case v.CompLit != nil:
			exprs, err := b.toCompLit(v.CompLit, locals)
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
	case expr.FuncCall != nil:
		fc, err := b.toFuncCall(expr.FuncCall, locals)
		if err != nil {
			return nil, err
		}

		returnType, ok := b.funcs[expr.FuncCall.Name]
		if !ok {
			return nil, fmt.Errorf("function %s not declared", expr.FuncCall.Name)
		}

		return &Expr{
			Type:     returnType,
			FuncCall: fc,
		}, nil
	}

	panic("invalid expression")
}

func (b *builder) toCompLit(lit []*ast.Expr, locals map[ast.Identifier]Type) ([]*Expr, error) {
	var exprs []*Expr

	for _, astExpr := range lit {
		expr, err := b.toExpr(astExpr, locals)
		if err != nil {
			return nil, err
		}

		// type check TODO

		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func (b *builder) toFuncCall(call *ast.FuncCall, locals map[ast.Identifier]Type) (*FuncCall, error) {
	var args []*Expr
	for _, e := range call.Args {
		if e.Val != nil && e.Val.Var != nil {
			typ, ok := locals[e.Val.Var.Name]
			if !ok {
				return nil, fmt.Errorf("variable %s not declared", e.Val.Var.Name)
			}

			if typ.Kind == KindArray {
				// need to convert to a pointer to the array for func calls
				arg := &Expr{
					Type: Type{
						Kind:    KindPointer,
						SubType: &typ,
					},
					AddressOf: VarName(e.Val.Var.Name),
				}

				args = append(args, arg)
				continue
			}
		}

		arg, err := b.toExpr(e, locals)
		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	return &FuncCall{
		Func: FuncName(call.Name),
		Args: args,
	}, nil
}

func int32Type() Type {
	return Type{Kind: KindPrimitive, Prim: PrimInt32}
}

func charType() Type {
	return Type{Kind: KindPrimitive, Prim: PrimChar}
}

func compatibleTypes(t1, t2 Type) bool {
	panic("implement me")
}
