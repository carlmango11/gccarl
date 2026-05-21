package parser

import (
	"strings"
	"testing"

	"github.com/carlmango11/gccarl/gccarl/tokens"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const tokenDef = `
IDEN: [a-zA-Z_][a-zA-Z0-9_]*
NUM: [0-9]+

INT: 'int'
CHAR: 'char'

ASTERISKS: '*'
HASH: '#'
LESS_THAN: '<'
GREATER_THAN: '>'
EQUALS: '='
QUOTE: '''
LPAREN: '('
RPAREN: ')'
LBRACE: '{'
RBRACE: '}'
COMMA: ','
SEMI: ';'
LSQUARE: '['
RSQUARE: ']'
PLUS: '+'
MINUS: '-'

INCLUDE: 'include'
RETURN: 'return'


`

func TestParse(t *testing.T) {
	tcs := []struct {
		grammar  string
		text     string
		expected []RuleKey
	}{
		{
			grammar: `
main:
  main:statement*

statement:
	def:type IDEN EQUALS NUM SEMI
	add:NUM PLUS NUM SEMI

type:
	int:INT
	char:CHAR
`,
			text: `int x = 5; 8 + 3;`,
			expected: []RuleKey{
				{"main", "main"},
				{"statement", "def"},
				{"type", "int"},
				{"statement", "add"},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.grammar, func(t *testing.T) {
			p, err := New(strings.NewReader(tc.grammar), true)
			require.NoError(t, err)

			tks, err := tokens.New(tokenDef, tc.text)
			require.NoError(t, err)

			path, err := p.parsePath(tks)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, path)
		})
	}
}
