package compiler

import (
	"compiler/parser"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const grammar = `
main:
  statements:statement*

value:
  int:NUM
  variable:IDEN
  
expr:
  add:value "+" expr
  sub:value "-" expr
  parens:"(" expr ")"
  value:value

statement:
  dec-assign:type IDEN "=" NUM ";"
  assign:IDEN "=" expr ";"
  
type:
  int:"int"
`

func TestGenerator(t *testing.T) {
	tcs := []struct {
		text     string
		expected []Instr
	}{
		{
			text: `
				int x = 5;
				int y = 3;
				y = x + y;`,
			expected: []Instr{
				"mov dword [rbp-4], 5",
				"mov dword [rbp-8], 3",
				"mov rax, [rbp-8]",
				"mov rbx, [rbp-4]",
				"add rax, rbx",
				"mov dword [rbp-8], rax",
			},
		},
	}

	p := parser.New(strings.NewReader(grammar), true)
	g := &Compiler{}

	for _, tc := range tcs {
		t.Run(tc.text, func(t *testing.T) {
			ast, err := p.Parse(tc.text)
			assert.NoError(t, err)

			output, err := g.compile(ast)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, output)
		})
	}
}
