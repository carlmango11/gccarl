package parser

import (
	"strconv"
	"testing"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
	"github.com/stretchr/testify/assert"
)

func TestNode_Clone(t *testing.T) {
	n := &Node{
		Key: rk(1),
		Values: []*Value{
			token("1_1"),
			token("1_2"),
			node(2,
				token("2_1"),
				token("2_2"),
			),
			token("1_3"),
			node(3,
				token("3_1"),
				node(4,
					token("4_1"),
					node(5,
						token("5_1"),
					),
				),
			),
		},
	}

	attachParents(n)

	toClone := n.Values[4].Node.Values[1].Node.Values[1].Node

	cloned := toClone.Clone()

	assert.Equal(t, toClone, cloned)
}

func attachParents(n *Node) {
	for _, v := range n.Values {
		if v.Node != nil {
			v.Node.parent = n
			attachParents(v.Node)
		}
	}
}

func rk(n int) RuleKey {
	return RuleKey{
		Rule:   grammar.RuleName(strconv.Itoa(n)),
		Option: grammar.OptionName(strconv.Itoa(n)),
	}
}

func node(n int, vals ...*Value) *Value {
	return &Value{
		Node: &Node{
			Key:    rk(n),
			Values: vals,
		},
	}
}

func token(name tokens.Name) *Value {
	return &Value{
		Token: &tokens.Token{
			Name: name,
		},
	}
}
