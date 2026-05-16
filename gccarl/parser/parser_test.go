package lexer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tcs := []struct {
		grammar  string
		text     string
		expected *Node
	}{
		{
			grammar: `
main:
  dec:type IDEN "=" NUM

type:
  int:"int"
  float:"float"
`,
			text: `int hello = 4`,
			expected: &Node{
				Name: "dec",
				Values: []*Value{
					{
						Node: &Node{
							Name: "int",
							Values: []*Value{
								{
									Literal: "int",
								},
							},
						},
					},
					{
						Identifier: "hello",
					},
					{
						Literal: "=",
					},
					{
						Number: 4,
					},
				},
			},
		},
		{
			grammar: `
main:
  l:"var" "wrong"
  a:"var" func*

func:
  b:"f"
  c:"g"
`,
			text: `var g f`,
			expected: &Node{
				Name: "a",
				Values: []*Value{
					{
						Literal: "var",
					},
					{
						Node: &Node{
							Name: "c",
							Values: []*Value{
								{
									Literal: "g",
								},
							},
						},
					},
					{
						Node: &Node{
							Name: "b",
							Values: []*Value{
								{
									Literal: "f",
								},
							},
						},
					},
				},
			},
		},
		{
			grammar: `
main:
  funcs:func*

func:
  f:"f"
`,
			text: `f f`,
			expected: &Node{
				Name: "funcs",
				Values: []*Value{
					{
						Node: &Node{
							Name: "f",
							Values: []*Value{
								{
									Literal: "f",
								},
							},
						},
					},
					{
						Node: &Node{
							Name: "f",
							Values: []*Value{
								{
									Literal: "f",
								},
							},
						},
					},
				},
			},
		},
		{
			grammar: `
main:
  a:"x"* "y" "z"?
`,
			text: `xxyz`,
			expected: &Node{
				Name: "a",
				Values: []*Value{
					{
						Literal: "x",
					},
					{
						Literal: "x",
					},
					{
						Literal: "y",
					},
					{
						Literal: "z",
					},
				},
			},
		},
		{
			grammar: `
main:
  a:"x"* "z"? "p"
`,
			text: `xxp`,
			expected: &Node{
				Name: "a",
				Values: []*Value{
					{
						Literal: "x",
					},
					{
						Literal: "x",
					},
					{
						Literal: "p",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.grammar, func(t *testing.T) {
			p, err := New(strings.NewReader(tc.grammar), true)
			require.NoError(t, err)

			node, err := p.Parse(strings.NewReader(tc.text))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, node)
		})
	}
}
