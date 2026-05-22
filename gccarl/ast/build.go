package ast

import (
	"fmt"
	"strconv"

	"github.com/carlmango11/gccarl/gccarl/parser"
)

func Build(n *parser.Node) (*Program, error) {
	a := &Program{}

	for _, v := range n.Values {
		switch v.Node.Key.Option {
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

	funcName := node.Values[1].Token.Val

	f := &FuncDef{
		Name: Identifier(funcName),
		ReturnType: &TypeDef{
			Type: retType,
		},
	}

	if node.Values[3].Node != nil && node.Values[3].Node.Key.Option == "params" {
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

		switch n.Key.Option {
		case "return":
			expr, err := toExpr(n.Values[1].Node)
			if err != nil {
				return nil, err
			}

			f.ReturnExpr = expr
		case "line":
			for _, v := range n.Values[0].Node.Values {
				switch v.Node.Key.Option {
				case "statement":
					s, err := toStatement(n.Values[0].Node)
					if err != nil {
						return nil, err
					}

					f.Lines = append(f.Lines, &Line{
						Statement: s,
					})
				case "control":
					c, err := toControl(n.Values[0].Node)
					if err != nil {
						return nil, err
					}

					f.Lines = append(f.Lines, &Line{
						Control: c,
					})
				}
			}
		case "statement":
		}
	}

	return f, nil
}

func toControl(n *parser.Node) (*Control, error) {
	switch n.Key.Option {
	case "if":
		ifStatement, err := toIf(n.Values[1].Node)
		if err != nil {
			return nil, err
		}

		return &Control{
			If: ifStatement,
		}, nil
	}

	panic("invalid control")
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
		Type:   typ,
		Arrays: v.Arrays,
	}

	return &ParamDef{
		Type: typeDef,
		Name: v.Name,
	}, nil
}

func toStatement(n *parser.Node) (*Statement, error) {
	switch n.Key.Option {
	case "var-dec":
		panic("impl")
	case "dec-assign":
		vd, err := toDecAssign(n.Values[0].Node)
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
	case "expr":
		expr, err := toExpr(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Statement{
			Expr: expr,
		}, nil
	}

	return nil, fmt.Errorf("unknown node: %s", n.Key.Option)
}

func toIf(n *parser.Node) (*If, error) {
	expr, err := toExpr(n.Values[2].Node)
	if err != nil {
		return nil, err
	}

	var statements []*Statement
	for _, sn := range n.Values[4 : len(n.Values)-1] {
		s, err := toStatement(sn.Node)
		if err != nil {
			return nil, err
		}

		statements = append(statements, s)
	}

	return &If{
		Condition:  expr,
		Statements: statements,
	}, nil
}

func toFuncCall(vs []*parser.Value) (*FuncCall, error) {
	fc := &FuncCall{
		Name: Identifier(vs[0].Token.Val),
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

		fc.Args = append(fc.Args, expr)
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

	switch n.Key.Option {
	case "int":
		num, err := strconv.Atoi(val.Token.Val)
		if err != nil {
			return nil, err
		}

		return &Value{
			Int: &num,
		}, nil
	case "variable":
		v, err := toVariable(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Value{
			Var: v,
		}, nil
	case "str":
		str := n.Values[0].Token.Val
		str = str[1 : len(str)-1]

		return &Value{
			Str: &str,
		}, nil
	case "char":
		ch := n.Values[0].Token.Val[1]

		return &Value{
			Char: &ch,
		}, nil
	case "comp-lit":
		exprs, err := toCompositeLiteral(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		return &Value{
			CompLit: exprs,
		}, nil
	default:
		panic(fmt.Sprintf("unknown node: %s", n.Key.Option))
	}
}

func toVariable(n *parser.Node) (*Var, error) {
	var arrays []*ArrayDef

	for i := 1; i < len(n.Values); i++ {
		var hasSize bool
		var size int

		arrayIndexNode := n.Values[i].Node
		if len(arrayIndexNode.Values) == 3 {
			num, err := strconv.Atoi(arrayIndexNode.Values[1].Token.Val)
			if err != nil {
				return nil, err
			}

			size = num
			hasSize = true
		}

		arrays = append(arrays, &ArrayDef{
			HasSize: hasSize,
			Size:    size,
		})
	}

	return &Var{
		Name:   Identifier(n.Values[0].Token.Val),
		Arrays: arrays,
	}, nil
}

func toExpr(n *parser.Node) (*Expr, error) {
	switch n.Key.Option {
	case "is-equal":
		left, err := toExpr(n.Values[0].Node)
		if err != nil {
			return nil, err
		}

		right, err := toExpr(n.Values[2].Node)
		if err != nil {
			return nil, err
		}

		return &Expr{
			IfEqual: &IfEqual{
				Left:  left,
				Right: right,
			},
		}, nil
	case "comp":
		return handleCompExpr(n.Values[0].Node)
	case "sub-expr":
		return handleSubExpr(n.Values[0].Node)
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Key.Option)
	}
}

func handleCompExpr(n *parser.Node) (*Expr, error) {
	switch n.Key.Option {
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Key.Option)
	}
}

func handleSubExpr(n *parser.Node) (*Expr, error) {
	switch n.Key.Option {
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
		return nil, fmt.Errorf("unknown expr type: %s", n.Key.Option)
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

// dec-assign node
func toDecAssign(n *parser.Node) (*DecAssign, error) {
	switch n.Key.Option {
	//case "array":
	//	typ, err := toType(n.Values[0])
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	v, err := toVariable(n.Values[1].Node)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return &DecAssign{
	//		Array: &ArrayDecAssign{
	//			Dec: &Dec{
	//				Type: &TypeDef{
	//					Type:   typ,
	//					Arrays: v.Arrays,
	//				},
	//				Name: v.Name,
	//			},
	//			Exprs: compLit,
	//		},
	//	}, nil
	case "standard":
		typ, err := toType(n.Values[0])
		if err != nil {
			return nil, err
		}

		v, err := toVariable(n.Values[1].Node)
		if err != nil {
			return nil, err
		}

		expr, err := toExpr(n.Values[3].Node)
		if err != nil {
			return nil, err
		}

		return &DecAssign{
			Dec: &Dec{
				Type: &TypeDef{
					Type:   typ,
					Arrays: v.Arrays,
				},
				Name: v.Name,
			},
			Assign: &Assign{
				Var: &Var{
					Name: v.Name,
				},
				Expr: expr,
			},
		}, nil
	}

	return nil, fmt.Errorf("unknown expr type: %s", n.Key.Option)
}

func toCompositeLiteral(n *parser.Node) ([]*Expr, error) {
	entriesNode := n.Values[1].Node

	exprNodes := []*parser.Node{
		entriesNode.Values[0].Node,
	}

	for i, n := range entriesNode.Values[1].Node.Values {
		if i%2 == 0 {
			continue // skip commas
		}

		exprNodes = append(exprNodes, n.Node)
	}

	var exprs []*Expr
	for _, exprNode := range exprNodes {
		expr, err := toExpr(exprNode)
		if err != nil {
			return nil, err
		}

		exprs = append(exprs, expr)
	}

	return exprs, nil
}

// "type" node
func toType(v *parser.Value) (*RawType, error) {
	if v.Node == nil {
		return nil, fmt.Errorf("expected type, got %+v", v)
	}

	if v.Node.Key.Option == "custom" {
		return &RawType{
			Type: Identifier(v.Node.Values[0].Token.Val),
		}, nil
	}

	// tood pointer

	return &RawType{
		Type: Identifier(v.Node.Values[0].Token.Val),
	}, nil
}
