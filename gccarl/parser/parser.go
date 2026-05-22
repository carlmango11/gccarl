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

func (p *Parser) Parse(r *tokens.Reader) (*Node, error) {
	path, err := p.parsePath(r)
	if err != nil {
		return nil, err
	}

	p.debugf("path: %s", path)

	r.Reset() // this is all shit, I should just build it as I parse
	node := p.buildNode(r, RuleKey{"main", "main"}, &PathReader{path: path})

	return node, nil
}

type PathReader struct {
	path []RuleKey
	i    int
}

func (p *PathReader) Next() RuleKey {
	if p.i >= len(p.path) {
		panic("path too short")
	}

	r := p.path[p.i]
	p.i++

	return r
}

func (p *PathReader) Peek() RuleKey {
	if p.i >= len(p.path) {
		panic("peek called on empty path")
	}

	return p.path[p.i]
}

func (p *Parser) buildNode(r *tokens.Reader, k RuleKey, path *PathReader) *Node {
	parts := p.getParts(k)

	var vals []*Value
	for i, part := range parts {
		if part.IsRule() {
			if path.Peek().Rule != part.Rule {
				continue
			}

			n := path.Next()

			p.debugf("[%v] new rule %v", i, n)

			vals = append(vals, &Value{
				Node: p.buildNode(r, n, path),
			})

			continue
		}

		tok, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		p.debugf("[%d] applying %v to token %v", i, tok.Name, k)

		vals = append(vals, &Value{
			Token: tok,
		})
	}

	return &Node{
		Key:    k,
		Values: vals,
	}
}

func (p *Parser) parsePath(r *tokens.Reader) ([]RuleKey, error) {
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

		if len(p.cursors) == 0 {
			return nil, fmt.Errorf("expected at least one cursor")
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
			p.debugf("applied %v to %v", token.Name, c)
			nextCursors = append(nextCursors, c)
		}
	}

	p.cursors = nextCursors
}

func (p *Parser) advance(cursors []*Cursor) []*Cursor {
	needBranch := cursors

	var ready []*Cursor
	var nextNeedBranch []*Cursor

	var i int
	for {
		for _, c := range needBranch {
			innerReady, innerNeedBrach := c.Branch()

			ready = append(ready, innerReady...)
			nextNeedBranch = append(nextNeedBranch, innerNeedBrach...)
		}

		needBranch = nextNeedBranch
		nextNeedBranch = nil

		if len(needBranch) == 0 {
			break
		}

		const maxBranches = 100

		i++
		if i > maxBranches {
			panic("too many branches")
		}
	}

	return ready
}

func (p *Parser) getParts(k RuleKey) []*grammar.Part {
	for _, o := range p.grammar[k.Rule].Options {
		if o.Name == k.Option {
			return o.Parts
		}
	}

	panic("invalid rule")
}
