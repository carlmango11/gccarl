package ast

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/parser"
)

func Build(n *parser.Node) (*Program, error) {
	a := &Program{}

	for _, v := range n.Values {
		switch v.Node.Name {
		case "include":
		case "func-def":
			f, err := toFuncDef(v.Node)
			if err != nil {
				return nil, err
			}

			a.FuncDefs = append(a.FuncDefs, f)
		}
	}

	return a, nil
}

func toFuncDef(node *parser.Node) (*FuncDef, error) {
	retType, err := toType(node.Values[0])
	if err != nil {
		return nil, err
	}

	f := &FuncDef{
		Name:       node.Values[1].Identifier,
		ReturnType: retType,
	}

	if node.Values[3].Node != nil && node.Values[3].Node.Name == "params" {
		for i, v := range node.Values[3].Node.Values {
			paramNode := v.Node
			if i > 0 {
				paramNode = v.Node.Values[1].Node
			}

			p, err := toParamDef(paramNode)
			if err != nil {
				return nil, err
			}

			f.Params = append(f.Params, p)
		}
	}

	for i := 4; i < len(node.Values); i++ {
		n := node.Values[i].Node
		if n == nil {
			continue
		}

		switch n.Name {
		case "return":
			expr, err := toExpr(n.Values[1].Node)
			if err != nil {
				return nil, err
			}

			f.ReturnExpr = expr
		case "statement":
			s, err := toStatement(n.Values[0].Node)
			if err != nil {
				return nil, err
			}

			f.Statements = append(f.Statements, s)
		}
	}

	return f, nil
}

func toParamDef(n *parser.Node) (*ParamDef, error) {
	paramType, err := toType(n.Values[0])
	if err != nil {
		return nil, err
	}

	return &ParamDef{
		Type: paramType,
		Name: n.Values[1].Identifier,
	}, nil
}

func toStatement(n *parser.Node) (*Statement, error) {
	switch n.Name {
	case "dec-assign":
		vd, err := toVarDec(n.Values[0].Node.Values)
		if err != nil {
			return nil, err
		}

		return &Statement{
			VarDec: vd,
		}, nil
	case "assign":
		a, err := toAssign(n.Values)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
		}, nil
	case "func-call":
		fc, err := toFuncCall(n.Values)
		if err != nil {
			return nil, err
		}

		return &Statement{
			FuncCall: fc,
		}, nil
	}

	return nil, fmt.Errorf("unknown node: %s", n.Name)
}

func toFuncCall(vs []*parser.Value) (*FuncCall, error) {
	fc := &FuncCall{
		Name: vs[0].Identifier,
	}

	for i, p := range vs[2].Node.Values {
		exprNode := p.Node
		if i > 0 {
			exprNode = p.Node.Values[1].Node
		}

		expr, err := toExpr(exprNode)
		if err != nil {
			return nil, err
		}

		fc.Params = append(fc.Params, expr)
	}

	return fc, nil
}

func toAssign(vs []*parser.Value) (*Assign, error) {
	v, err := toVariable(vs[0].Node)
	if err != nil {
		return nil, err
	}

	expr, err := toExpr(vs[2].Node)
	if err != nil {
		return nil, err
	}

	return &Assign{
		Var:  v,
		Expr: expr,
	}, nil
}

func toValue(n *parser.Node) (*Value, error) {
	if len(n.Values) != 1 {
		return nil, fmt.Errorf("invalid number of values: %d", len(n.Values))
	}

	val := n.Values[0]

	switch n.Name {
	case "int":
		return &Value{
			Int: val.Number,
		}, nil
	case "variable":
		v, err := toVariable(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Value{
			Var: v,
		}, nil
	default:
		return nil, fmt.Errorf("unknown variable: %s", n.Name)
	}
}

func toVariable(n *parser.Node) (*Var, error) {
	v := &Var{
		Name: n.Values[0].Identifier,
	}

	if len(n.Values) == 2 {
		idx := n.Values[1].Node.Values[1].Number
		v.Index = &idx
	}

	return v, nil
}

func toExpr(n *parser.Node) (*Expr, error) {
	switch n.Name {
	case "add":
		add, err := toAddExpr(n.Values)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Add: add,
		}, nil
	case "value":
		v, err := toValue(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Val: v,
		}, nil
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Name)
	}
}

func toAddExpr(vs []*parser.Value) (*AddExpr, error) {
	val, err := toValue(vs[0].Node)
	if err != nil {
		return nil, err
	}

	expr, err := toExpr(vs[2].Node)
	if err != nil {
		return nil, err
	}

	return &AddExpr{
		Val:  val,
		Expr: expr,
	}, nil
}

func toVarDec(vs []*parser.Value) (*VarDec, error) {
	typ, err := toType(vs[0])
	if err != nil {
		return nil, err
	}

	v, err := toVariable(vs[1].Node)
	if err != nil {
		return nil, err
	}

	expr, err := toExpr(vs[3].Node)
	if err != nil {
		return nil, err
	}

	return &VarDec{
		Type: typ,
		Var:  v,
		Expr: expr,
	}, nil
}

func toType(v *parser.Value) (Type, error) {
	if v.Node == nil {
		return TypeVoid, fmt.Errorf("expected type, got %+v", v)
	}

	switch v.Node.Values[0].Literal {
	case "int":
		return TypeInt, nil
	default:
		return TypeVoid, fmt.Errorf("unknown type: %s", v.Node.Values[0].Literal)
	}
}
