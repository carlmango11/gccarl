package grammar

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	const grammar = `
main:
  a:imports* func*

imports:
  b:"#include" IDEN NUM

func:
  c:"t1" "t3"?
  d:"t2"`

	rules, err := Parse(strings.NewReader(grammar))
	require.NoError(t, err)

	expected := map[RuleName]*Rule{
		"main": {
			Options: []*Option{
				{
					Name: "a",
					Parts: []*Part{
						{
							Type:        TTRule,
							Rule:        "imports",
							Cardinality: CardMultiple,
						},
						{
							Type:        TTRule,
							Rule:        "func",
							Cardinality: CardMultiple,
						},
					},
				},
			},
		},
		"imports": {
			Options: []*Option{
				{
					Name: "b",
					Parts: []*Part{
						{
							Type:  TTLiteral,
							Token: "#include",
						},
						{
							Type: TTIdentifier,
						},
						{
							Type: TTNumber,
						},
					},
				},
			},
		},
		"func": {
			Options: []*Option{
				{
					Name: "c",
					Parts: []*Part{
						{
							Type:  TTLiteral,
							Token: "t1",
						},
						{
							Type:        TTLiteral,
							Token:       "t3",
							Cardinality: CardZeroOrOne,
						},
					},
				},
				{
					Name: "d",
					Parts: []*Part{
						{
							Type:  TTLiteral,
							Token: "t2",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, rules)
}
