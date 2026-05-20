package parser

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type Cursor struct {
	grammar map[grammar.RuleName]*grammar.Rule
	Stack   []Key
	Path    []RuleKey
}

func (c *Cursor) String() string {
	return fmt.Sprintf("%v", c.Stack)
}

func (c *Cursor) Apply(token *tokens.Token) bool {
	nextPart := c.next()
	if nextPart.IsRule() {
		panic("unexpected rule")
	}

	if nextPart.Token != token.Name {
		return false
	}

	c.Stack[len(c.Stack)-1].Index++

	c.stepUp()

	return true
}

func (c *Cursor) stepUp() bool {
	for len(c.Stack) > 0 && c.endOfRule() {
		c.Stack = c.Stack[:len(c.Stack)-1]

		if len(c.Stack) == 0 {
			return true
		} // todo shit

		switch c.next().Cardinality {
		case grammar.CardSingle, grammar.CardOptional:
			c.Stack[len(c.Stack)-1].Index++
		}
	}

	return len(c.Stack) == 0
}

func (c *Cursor) Branch() ([]*Cursor, bool) {
	nextPart := c.next()
	if nextPart.Cardinality == grammar.CardSingle && !nextPart.IsRule() {
		return nil, false
	}

	optionBranches := c.branchOptions()
	if nextPart.Cardinality == grammar.CardSingle {
		return optionBranches, true
	}

	var all []*Cursor

	for _, ob := range optionBranches {
		all = append(all, ob.Clone())

		without := ob.Clone()
		without.Stack[len(without.Stack)-1].Index++

		without.stepUp()

		all = append(all, without)
	}

	return all, true
}

func (c *Cursor) Clone() *Cursor {
	stack := make([]Key, len(c.Stack))
	copy(stack, c.Stack)
	path := make([]RuleKey, len(c.Path))
	copy(path, c.Path)

	return &Cursor{
		grammar: c.grammar,
		Stack:   stack,
		Path:    path,
	}
}

func (c *Cursor) next() *grammar.Part {
	top := c.Stack[len(c.Stack)-1]

	for _, o := range c.grammar[top.RuleKey.Rule].Options {
		if o.Name == top.RuleKey.Option {
			return o.Parts[top.Index]
		}
	}

	panic("undefined option")
}

func (c *Cursor) branchOptions() []*Cursor {
	nextPart := c.next()
	if !nextPart.IsRule() {
		return []*Cursor{c}
	}

	var cursors []*Cursor
	for _, o := range c.grammar[nextPart.Rule].Options {
		cloned := c.Clone()

		key := Key{
			RuleKey: RuleKey{
				Rule:   nextPart.Rule,
				Option: o.Name,
			},
		}

		cloned.Stack = append(cloned.Stack, key)
		cloned.Path = append(cloned.Path, key.RuleKey)

		cursors = append(cursors, cloned)
	}

	return cursors
}

func (c *Cursor) endOfRule() bool {
	top := c.Stack[len(c.Stack)-1]
	return len(c.parts()) == top.Index
}

func (c *Cursor) parts() []*grammar.Part {
	top := c.Stack[len(c.Stack)-1]

	for _, o := range c.grammar[top.RuleKey.Rule].Options {
		if o.Name == top.RuleKey.Option {
			return o.Parts
		}
	}

	panic("undefined option")
}
