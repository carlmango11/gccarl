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
	typeNode := node.Values[0]

	retType, err := toType(typeNode)
	if err != nil {
		return nil, err
	}

	funcName := node.Values[1].Identifier

	f := &FuncDef{
		Name: funcName,
		ReturnType: &TypeDef{
			Type: retType,
		},
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
	typ, err := toType(n.Values[0])
	if err != nil {
		return nil, err
	}

	v, err := toVariable(n.Values[1].Node)
	if err != nil {
		return nil, err
	}

	typeDef := &TypeDef{
		Type:  typ,
		Array: v.Indexed,
	}

	return &ParamDef{
		Type: typeDef,
		Name: v.Name,
	}, nil
}

func toStatement(n *parser.Node) (*Statement, error) {
	switch n.Name {
	case "var-dec":
		panic("impl")
	case "dec-assign":
		vd, err := toDecAssign(n.Values[0].Node.Values)
		if err != nil {
			return nil, err
		}

		return &Statement{
			DecAssign: vd,
		}, nil
	case "assign":
		a, err := toAssign(n.Values)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Assign: a,
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
	val := n.Values[0]

	switch n.Name {
	case "int":
		return &Value{
			Int: &val.Number,
		}, nil
	case "variable":
		v, err := toVariable(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Value{
			Var: v,
		}, nil
	//case "str":
	//	s, err := toStr(n.Values[0].Node)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return &Value{
	//		Str: s,
	//	}, nil
	case "char":
		return &Value{
			Char: &n.Values[1].Char,
		}, nil
	case "array":
		a, err := toArray(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Value{
			Array: a,
		}, nil
	default:
		panic(fmt.Sprintf("unknown node: %s", n.Name))
	}
}

func toVariable(n *parser.Node) (*Var, error) {
	var array bool
	var index int
	if len(n.Values) == 2 {
		array = true

		arrayIndexNode := n.Values[1].Node
		if len(arrayIndexNode.Values) == 3 {
			index = arrayIndexNode.Values[2].Number
		}
	}

	return &Var{
		Name:    n.Values[0].Identifier,
		Indexed: array,
		Index:   index,
	}, nil
}

func toArray(n *parser.Node) (*Array, error) {
	a := &Array{}

	if len(n.Values) == 2 {
		return a, nil
	}

	for i, v := range n.Values[1].Node.Values {
		exprNode := v.Node
		if i > 0 {
			exprNode = v.Node.Values[1].Node
		}

		expr, err := toExpr(exprNode)
		if err != nil {
			return nil, err
		}

		a.Entries = append(a.Entries, expr)
	}

	return a, nil
}

func toExpr(n *parser.Node) (*Expr, error) {
	switch n.Name {
	case "comp":
		return handleCompExpr(n.Values[0].Node)
	case "sub-expr":
		return handleSubExpr(n.Values[0].Node)
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Name)
	}
}

func handleCompExpr(n *parser.Node) (*Expr, error) {
	switch n.Name {
	case "value":
		v, err := toValue(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Val: v,
		}, nil
	case "func-call":
		fc, err := toFuncCall(n.Values)
		if err != nil {
			return nil, err
		}

		return &Expr{
			FuncCall: fc,
		}, nil
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Name)
	}
}

func handleSubExpr(n *parser.Node) (*Expr, error) {
	switch n.Name {
	case "value":
		v, err := toValue(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Val: v,
		}, nil
	case "func-call":
		fc, err := toFuncCall(n.Values)
		if err != nil {
			return nil, err
		}

		return &Expr{
			FuncCall: fc,
		}, nil
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Name)
	}
}

func toAddExpr(vs []*parser.Value) (*AddExpr, error) {
	e1, err := toExpr(vs[0].Node)
	if err != nil {
		return nil, err
	}

	e2, err := toExpr(vs[2].Node)
	if err != nil {
		return nil, err
	}

	return &AddExpr{
		Expr1: e1,
		Expr2: e2,
	}, nil
}

func toDecAssign(vs []*parser.Value) (*DecAssign, error) {
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

	return &DecAssign{
		Dec: &Dec{
			Type: &TypeDef{
				Type:  typ,
				Array: v.Indexed,
			},
			Name: v.Name,
			Size: v.Index,
		},
		Assign: &Assign{
			Var: &Var{
				Name: v.Name,
			},
			Expr: expr,
		},
	}, nil
}

// "type" node
func toType(v *parser.Value) (*RawType, error) {
	if v.Node == nil {
		return nil, fmt.Errorf("expected type, got %+v", v)
	}

	if v.Node.Name == "custom" {
		return &RawType{
			Type: v.Node.Values[0].Identifier,
		}, nil
	}

	// tood pointer

	return &RawType{
		Type: parser.Identifier(v.Node.Values[0].Literal),
	}, nil
}
