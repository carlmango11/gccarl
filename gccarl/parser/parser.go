package parser

import (
	"fmt"
	"io"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type Node struct {
	Key    RuleKey
	Values []*Value
}

type RuleKey struct {
	Rule   grammar.RuleName
	Option grammar.OptionName
}

type Key struct {
	RuleKey RuleKey
	Index   int
}

type Value struct {
	Node  *Node
	Token *tokens.Token
}

type Parser struct {
	grammar map[grammar.RuleName]*grammar.Rule
	debug   bool
	nodes   []string

	cursors []*Cursor
}

func New(r io.Reader, debug bool) (*Parser, error) {
	gr, err := grammar.Parse(r)
	if err != nil {
		return nil, err
	}

	return &Parser{
		grammar: gr,
		debug:   debug,
	}, nil
}

func (p *Parser) Parse(r *tokens.Reader) ([]*Node, error) {

}

func (p *Parser) Parse(r *tokens.Reader) ([]RuleKey, error) {
	p.cursors = []*Cursor{
		{
			grammar: p.grammar,
			Stack: []Key{
				{
					RuleKey: RuleKey{
						Rule:   "main",
						Option: "main",
					},
				},
			},
		},
	}

	for {
		token, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		if p.debug {
			fmt.Printf("token: %v\n", token)
		}

		p.handleToken(token)
	}

	if len(p.cursors) != 1 {
		return nil, fmt.Errorf("expected 1 cursor but got %d", len(p.cursors))
	}

	if !p.cursors[0].terminalState() {
		return nil, fmt.Errorf("last cursor did not terminate: %v", p.cursors[0])
	}

	return p.cursors[0].Path, nil
}

func (p *Parser) debugf(format string, args ...any) {
	if p.debug {
		fmt.Printf(format+"\n", args...)
	}
}

func (p *Parser) handleToken(token *tokens.Token) {
	cursors := p.advance(p.cursors)

	for _, c := range cursors {
		p.debugf("handling %v with %v", token.Name, c)
	}

	var nextCursors []*Cursor
	for _, c := range cursors {
		ok := c.Apply(token)
		if ok {
			nextCursors = append(nextCursors, c)
		}
	}

	p.cursors = nextCursors
}

func (p *Parser) advance(cursors []*Cursor) []*Cursor {
	needBranch := cursors

	var ready []*Cursor
	var nextNeedBranch []*Cursor

	for {
		for _, c := range needBranch {
			inner, branched := c.Branch()
			if branched {
				nextNeedBranch = append(nextNeedBranch, inner...)
			} else {
				ready = append(ready, c)
			}
		}

		needBranch = nextNeedBranch
		nextNeedBranch = nil

		if len(needBranch) == 0 {
			break
		}
	}

	return ready
}
