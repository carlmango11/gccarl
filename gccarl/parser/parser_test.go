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
  main:control* type

control:
	o1:LPAREN
	o2:RPAREN

type:
	int:INT
	char:CHAR
`,
			text: `() int`,
			expected: []RuleKey{
				{"main", "main"},
				{"control", "o1"},
				{"control", "o2"},
				{"type", "int"},
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
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, path)
		})
	}
}
