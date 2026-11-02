package preprocessor

import (
	"io"

	"github.com/carlmango11/gccarl/gccarl/generated/dparser"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

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

type exprTokenReader struct {
	i   int
	tks []*tokens.Token
}

func (r *exprTokenReader) Next() (*tokens.Token, error) {
	if r.i == len(r.tks) {
		return nil, io.EOF
	}

	tk := r.tks[r.i]
	r.i++

	return tk, nil
}

func parseExpr(tks []*tokens.Token) (*Expr, error) {
	r := &exprTokenReader{
		tks: tks,
	}

	raw, err := dparser.Parse(r)
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

func (r *reader) evalExpr(exprTokens []*tokens.Token) (bool, error) {
	e, err := parseExpr(exprTokens)
	if err != nil {
		return false, err
	}

	switch {
	case e.Defined != nil:
		_, ok := r.macros[string(*e.Defined)]
		return ok, nil
	}

	panic("unreachable")
}
