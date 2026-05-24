package parser

import (
	"testing"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
)

func TestGen(t *testing.T) {
	g := &Generator{
		grammar: map[grammar.RuleName]*grammar.Rule{
			"main": {
				Options: []*grammar.Option{
					{
						Name: "a",
						Parts: []*grammar.Part{
							{
								Cardinality: grammar.CardSingle,
								Token:       "PLUS",
							},
							{
								Cardinality: grammar.CardSingle,
								Token:       "PLUS",
							},
						},
					},
				},
			},
		},
	}
}
