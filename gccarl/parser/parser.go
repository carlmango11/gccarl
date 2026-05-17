package parser

import (
	"fmt"
	"io"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type Node struct {
	Rule   grammar.RuleName
	Option grammar.OptionName
	Values []*Value
}

type RuleKey struct {
	Rule   grammar.RuleName
	Option grammar.OptionName
}

type Key struct {
	Rule   grammar.RuleName
	Option grammar.OptionName
	Index  int
}

type Cursor struct {
	Current Key
	Stack   []Key
	Path    []RuleKey
}

type Identifier string

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

func (p *Parser) Parse(r *tokens.Reader) ([]RuleKey, error) {
	p.cursors = []*Cursor{
		{
			Current: Key{"main", "main", 0},
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

		path, ok := p.handleToken(token)
		if ok {
			return path, nil
		}
	}

	return nil, fmt.Errorf("not terminated")
}

func (p *Parser) handleToken(token *tokens.Token) ([]RuleKey, bool) {
	cursors := p.advance(p.cursors)

	var nextCursors []*Cursor
	for _, c := range cursors {
		ok := p.apply(c, token)
		if ok {
			nextCursors = append(nextCursors, c)
		}
	}

	cursors = nextCursors

	p.clean(cursors)

	for _, c := range cursors {
		if len(c.Stack) == 0 {
			return c.Path, true
		}
	}

	p.cursors = cursors

	return nil, false
}

func (p *Parser) advance(cursors []*Cursor) []*Cursor {
	var newCursors []*Cursor

	for _, c := range cursors {
		nextPart := p.getParts(c.Current)[c.Current.Index]
		if !nextPart.IsRule() {
			newCursors = append(newCursors, c)
			break
		}

		var next []*Cursor
		for _, o := range p.grammar[nextPart.Rule].Options {
			newC := &Cursor{
				Current: Key{
					Rule:   nextPart.Rule,
					Option: o.Name,
				},
				Stack: append(c.Stack, c.Current),
				Path: append(c.Path, RuleKey{
					Rule:   nextPart.Rule,
					Option: o.Name,
				}),
			}

			next = append(next, newC)
		}

		advanced := p.advance(next)
		for _, c := range advanced {
			newCursors = append(newCursors, c)
		}
	}

	return newCursors
}

func (p *Parser) apply(c *Cursor, token *tokens.Token) bool {
	part := p.getParts(c.Current)[c.Current.Index]

	if part.Token == token.Name {
		if part.Cardinality != grammar.CardMultiple {
			c.Current.Index++
		}

		return true
	}

	if part.Cardinality == grammar.CardOptional {
		return true
	}

	return false
}

func (p *Parser) getParts(key Key) []*grammar.Part {
	for _, o := range p.grammar[key.Rule].Options {
		if o.Name == key.Option {
			return o.Parts
		}
	}

	panic("undefined option")
}

func (p *Parser) clean(cs []*Cursor) {
	for _, c := range cs {
		for {
			if c.Current.Index != len(p.getParts(c.Current)) {
				break
			}

			if len(c.Stack) == 0 {
				break
			}

			// end of rule
			c.Current = c.Stack[len(c.Stack)-1]
			c.Stack = c.Stack[:len(c.Stack)-1]

			if p.getParts(c.Current)[c.Current.Index].Cardinality != grammar.CardMultiple {
				c.Current.Index++
			}
		}
	}
}
