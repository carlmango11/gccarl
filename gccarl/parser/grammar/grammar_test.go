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
	a:imports* states=(statement COMMA)*

statement:
	b:ADD NUM
`

	rules, err := Parse(strings.NewReader(grammar))
	require.NoError(t, err)

	expected := map[RuleName][]*Option{
		"main": {
			{
				Name: "a",
				Parts: []*Part{
					{
						Rule:        "imports",
						Cardinality: CardMultiple,
					},
					{
						Rule:        "states",
						Cardinality: CardMultiple,
					},
				},
			},
		},
		"states": {
			{
				Name: "states",
				Parts: []*Part{
					{
						Rule:        "statement",
						Cardinality: CardSingle,
					},
					{
						Token:       "COMMA",
						Cardinality: CardSingle,
					},
				},
			},
		},
		"statement": {
			{
				Name: "b",
				Parts: []*Part{
					{
						Token: "ADD",
					},
					{
						Token: "NUM",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, rules)
}
