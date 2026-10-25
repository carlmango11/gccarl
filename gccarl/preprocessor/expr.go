package preprocessor

import "github.com/carlmango11/gccarl/gccarl/generated/dparser"

type Expr struct {
	ID      dparser.IDEN
	Sub     *Expr
	Defined *dparser.IDEN
	OpExpr  *OpExpr
}

type OpExpr struct {
	Expr1 *Expr
	Op    dparser.OperatorType
	Expr2 *Expr
}

func parseExpr(s string) (*Expr, error) {
	raw, err := dparser.Parse(s)
	if err != nil {
		return nil, err
	}

	return parseRaw(raw)
}

func parseRaw(raw *dparser.Main) (*Expr, error) {
	switch raw.Type {
	case dparser.MainTypeComp:
		e1, err := parseRaw(raw.Comp.Main)
		if err != nil {
			return nil, err
		}

		e2, err := parseSub(raw.Comp.SubExpr)
		if err != nil {
			return nil, err
		}

		return &Expr{
			OpExpr: &OpExpr{
				Expr1: e1,
				Op:    raw.Comp.Operator.Type,
				Expr2: e2,
			},
		}, nil
	}

	panic("unreachable")
}

func parseSub(s *dparser.SubExpr) (*Expr, error) {
	switch s.Type {
	case dparser.SubExprTypeIden:
		return &Expr{
			ID: s.Iden.IDEN,
		}, nil
	case dparser.SubExprTypeDefined:
		return &Expr{
			Defined: &s.Defined.IDEN,
		}, nil
	case dparser.SubExprTypeParens:
		e, err := parseRaw(s.Parens.Main)
		if err != nil {
			return nil, err
		}

		return &Expr{
			Sub: e,
		}, nil
	}

	panic("unreachable")
}

func (p *Processor) evalExpr(e *Expr) bool {
	switch {
	case e.Defined != nil:
		_, ok := p.macros[string(*e.Defined)]
		return ok
	}

	panic("unreachable")
}
