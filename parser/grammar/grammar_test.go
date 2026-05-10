package grammar

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const grammar = `
main:
  a:imports* func*

imports:
  b:"#include" IDEN NUM

func:
  c:"t1" "t3"
  d:"t2"`

	rules := Parse(strings.NewReader(grammar))

	expected := map[RuleName]*Rule{
		"main": {
			Options: []*Option{
				{
					Name: "a",
					Tokens: []*Token{
						{
							Type:  TTRule,
							Rule:  "imports",
							Multi: true,
						},
						{
							Type:  TTRule,
							Rule:  "func",
							Multi: true,
						},
					},
				},
			},
		},
		"imports": {
			Options: []*Option{
				{
					Name: "b",
					Tokens: []*Token{
						{
							Type:    TTLiteral,
							Literal: "#include",
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
					Tokens: []*Token{
						{
							Type:    TTLiteral,
							Literal: "t1",
						},
						{
							Type:    TTLiteral,
							Literal: "t3",
						},
					},
				},
				{
					Name: "d",
					Tokens: []*Token{
						{
							Type:    TTLiteral,
							Literal: "t2",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, rules)
}
