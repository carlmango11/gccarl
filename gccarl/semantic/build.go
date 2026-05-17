package semantic

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/ast"
)

type builder struct {
	vars map[ast.Identifier]PrimitiveType
}

func Build(program *ast.Program) (*Program, error) {
	b := &builder{
		vars: make(map[ast.Identifier]PrimitiveType),
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
	}, nil
}

func (b *builder) toFuncDef(f *ast.FuncDef) (*FuncDef, error) {
	// TODO func scoped types

	returnType, err := b.toType(f.ReturnType)
	if err != nil {
		return nil, err
	}

	locals := map[ast.Identifier]Type{}

	for _, s := range f.Statements {
		dec := s.Dec
		if s.DecAssign != nil {
			dec = s.DecAssign.Dec
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
		returnExpr, err = b.toExpr(f.ReturnExpr)
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
	kind := KindPrimitive
	if i.Array {
		kind = KindArray
	}

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
		a, err := b.toAssign(vars, s.DecAssign.Assign)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
		}, nil
	case s.Assign != nil:
		a, err := b.toAssign(vars, s.Assign)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
		}, nil
	case s.Expr != nil:
		expr, err := b.toExpr(s.Expr)
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

	indexable := v.Kind == KindArray || v.Kind == KindPointer
	if a.Var.Indexed && !indexable {
		return nil, fmt.Errorf("variable %s is not indexable", a.Var.Name)
	}

	expr, err := b.toExpr(a.Expr)
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
		Name:  VarName(a.Var.Name),
		Index: a.Var.Index,
		Expr:  expr,
	}, nil
}

func (b *builder) toExpr(expr *ast.Expr) (*Expr, error) {
	switch {
	case expr.Val != nil:
		v := expr.Val
		switch {
		case v.Int != nil:
			// TODO other sizes
			return &Expr{
				Type: int32Type(),
				Literal: Literal{
					Int32: int32(*v.Int),
				},
			}, nil
		case v.Char != nil:
			return &Expr{
				Type: charType(),
				Literal: Literal{
					Char: *v.Char,
				},
			}, nil
		case v.Array != nil:
			return b.toArrayExpr(v.Array)
		}
	}

	panic("invalid expression")
}

func (b *builder) toArrayExpr(array *ast.Array) (*Expr, error) {
	var typ Type

	var vals []*Expr
	for _, v := range array.Entries {
		val, err := b.toExpr(v)
		if err != nil {
			return nil, err
		}

		typ = val.Type
		// TODO: check

		vals = append(vals, val)
	}

	return &Expr{
		Type: Type{Kind: KindArray, Prim: typ.Prim},
		Array: &Array{
			Vals: vals,
		},
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
