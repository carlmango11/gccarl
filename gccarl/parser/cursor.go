package parser

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type Cursor struct {
	grammar map[grammar.RuleName]*grammar.Rule
	Stack   []Index
	Current *Node
	Top     *Node
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

	c.Current.Values = append(c.Current.Values, &Value{
		Token: token,
	})

	c.advance()
	c.stepUp()

	return true
}

func (c *Cursor) stepUp() bool {
	for len(c.Stack) > 0 && c.endOfRule() {
		c.pop()

		if c.finished() {
			return true
		} // todo shit

		switch c.next().Cardinality {
		case grammar.CardSingle, grammar.CardZeroOrOne:
			c.Stack[len(c.Stack)-1].Index++
		}
	}

	return c.finished()
}

func (c *Cursor) Branch() (ready []*Cursor, needBranch []*Cursor) {
	if c.finished() {
		return nil, nil
	}

	nextPart := c.next()
	if nextPart.Cardinality == grammar.CardSingle && !nextPart.IsRule() {
		return []*Cursor{c}, nil
	}

	if !nextPart.IsRule() {
		// optional token. stay here and also jump over
		stay := c.Clone()

		jump := c.Clone()
		jump.advance()
		jump.stepUp()

		return []*Cursor{stay}, []*Cursor{jump}
	}

	// it's a rule
	optionBranches := c.branchOptions()
	if nextPart.Cardinality == grammar.CardSingle {
		return nil, optionBranches
	}

	// need to jump over it also
	jump := c.Clone()
	jump.advance()
	jump.stepUp()

	optionBranches = append(optionBranches, jump)

	return nil, optionBranches
}

func (c *Cursor) Clone() *Cursor {
	stack := make([]Index, len(c.Stack))
	copy(stack, c.Stack)

	node := c.Current.Clone()
	top := node

	for top.parent != nil {
		top = top.parent
	}

	return &Cursor{
		grammar: c.grammar,
		Stack:   stack,
		Current: node,
		Top:     top,
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

		ruleKey := RuleKey{
			Rule:   nextPart.Rule,
			Option: o.Name,
		}

		key := Index{
			RuleKey: ruleKey,
		}

		cloned.Stack = append(cloned.Stack, key)

		newNode := &Node{
			Key:    ruleKey,
			parent: cloned.Current,
		}

		cloned.Current.Values = append(cloned.Current.Values, &Value{
			Node: newNode,
		})

		cloned.Current = newNode

		cursors = append(cursors, cloned)
	}

	return cursors
}

func (c *Cursor) endOfRule() bool {
	top := c.Stack[len(c.Stack)-1]
	return len(c.parts()) == top.Index
}

func (c *Cursor) advance() {
	c.Stack[len(c.Stack)-1].Index++
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

func (c *Cursor) finished() bool {
	return len(c.Stack) == 0
}

func (c *Cursor) terminalState() bool {
	for {
		if c.stepUp() {
			return true
		}

		nextPart := c.next()
		if nextPart.Cardinality == grammar.CardSingle {
			return false
		}

		c.advance()
	}
}

func (c *Cursor) pop() {
	c.Stack = c.Stack[:len(c.Stack)-1]
	c.Current = c.Current.parent
}
