package parser

import (
	"fmt"
	"io"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type Node struct {
	Key         RuleKey
	Values      []*Value
	Cardinality grammar.Cardinality
}

type RuleKey struct {
	Rule   grammar.RuleName
	Option grammar.OptionName
}

type Key struct {
	RuleKey RuleKey
	Index   int
}

type Cursor struct {
	Current Key
	Path    []*Node
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
	p.cursors = []*Cursor{
		{
			Current: Key{
				RuleKey: RuleKey{
					Rule:   "main",
					Option: "main",
				},
			},
			Path: []*Node{
				&Node{},
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
			fmt.Printf("token: %s\n", token)
		}

		path, ok := p.handleToken(token)
		if ok {
			return path, nil
		}
	}

	return nil, fmt.Errorf("not terminated")
}

func (p *Parser) handleToken(token *tokens.Token) (*Node, bool) {
	cursors := p.advance(p.cursors)

	var nextCursors []*Cursor
	for _, c := range cursors {
		currentNode := c.Path[len(c.Path)-1]

		ok := p.apply(c, token)

		optional := currentNode.Cardinality == grammar.CardOptional || currentNode.Cardinality == grammar.CardMultiple
		if !ok && optional {
			c.Path = c.Path[:len(c.Path)-1]
		}
		if ok {
			nextCursors = append(nextCursors, c)
		}
	}

	cursors = nextCursors

	p.clean(cursors)

	for _, c := range cursors {
		if len(c.Path) == 0 {
			return c.Path[0], true
		}
	}

	p.cursors = cursors

	return nil, false
}

func (p *Parser) advance(cursors []*Cursor) []*Cursor {
	var newCursors []*Cursor

	for _, c := range cursors {
		nextPart := p.getParts(c.Current.RuleKey)[c.Current.Index]
		if !nextPart.IsRule() {
			newCursors = append(newCursors, c)
			break
		}

		var next []*Cursor
		for _, o := range p.grammar[nextPart.Rule].Options {
			next = append(next, newCursor(c, nextPart, o))
		}

		advanced := p.advance(next)
		for _, c := range advanced {
			newCursors = append(newCursors, c)
		}
	}

	return newCursors
}

func newCursor(c *Cursor, part *grammar.Part, o *grammar.Option) *Cursor {
	newNode := &Node{
		Key: RuleKey{
			Rule:   part.Rule,
			Option: o.Name,
		},
		Cardinality: part.Cardinality,
	}

	//node := copyNode(c.Path[len(c.Path)-1])
	//node.Values = append(node.Values, &Value{
	//	Node: newNode,
	//})

	return &Cursor{
		Current: Key{
			RuleKey: RuleKey{
				Rule:   part.Rule,
				Option: o.Name,
			},
		},
		Path: append(copySlice(c.Path), newNode),
	}
}

func copySlice[T any](stack []T) []T {
	c := make([]T, len(stack))
	copy(c, stack)
	return c
}

func copyNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	var vals []*Value
	for _, val := range n.Values {
		vals = append(vals, &Value{
			Node:  copyNode(val.Node),
			Token: val.Token,
		})
	}

	return &Node{
		Key:    n.Key,
		Values: vals,
	}
}

func (p *Parser) apply(c *Cursor, token *tokens.Token) bool {
	parts := p.getParts(c.Current.RuleKey)
	part := parts[c.Current.Index]

	currentNode := c.Path[len(c.Path)-1]

	if part.Token == token.Name {
		if part.Cardinality != grammar.CardMultiple {
			c.Current.Index++
		}

		currentNode.Values = append(currentNode.Values, &Value{
			Token: token,
		})

		return true
	}

	return false
}

func (p *Parser) getParts(key RuleKey) []*grammar.Part {
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
			if c.Current.Index != len(p.getParts(c.Current.RuleKey)) {
				break
			}

			if len(c.Path) == 0 {
				break
			}

			// end of rule
			c.Path = c.Path[:len(c.Path)-1]

			if p.getParts(c.Current.RuleKey)[c.Current.Index].Cardinality != grammar.CardMultiple {
				c.Current.Index++
			}
		}
	}
}
