package tokens

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRejectMalformedDefinitions(t *testing.T) {
	for _, text := range []string{"NUM:", "NUM: '", "NUM: ''", "NUM: [", "NUM: a*"} {
		t.Run(text, func(t *testing.T) {
			if _, err := ParseDefinitions(strings.NewReader(text)); err == nil {
				t.Fatal("accepted malformed token definition")
			}
		})
	}
}

func TestRegexAlternativesAreAnchored(t *testing.T) {
	reader, err := New(strings.NewReader("WORD: a|b"), strings.NewReader("xb"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := reader.Next(); err == nil {
		t.Fatal("skipped unmatched input before regex alternative")
	}
}

func TestRegexChar(t *testing.T) {
	const def = `CHAR: \'[^']*\'`
	const text = `'\n'`

	reader, err := New(strings.NewReader(def), strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}

	expected := &Token{
		Name: "CHAR",
		Val:  `'\n'`,
	}

	token, err := reader.Next()
	assert.NoError(t, err)
	assert.Equal(t, expected, token)
}

func TestNestedDirectives(t *testing.T) {
	const input = `#ifdef NO
    wrong
    #ifndef FOUR
        nested wrong
    #endif
#endif`

	r, err := New(strings.NewReader(tokenDefs), strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"#", "ifdef", "NO", "\n", "wrong", "\n", "#", "ifndef", "FOUR", "\n", "nested", "wrong", "\n", "#", "endif", "\n", "#", "endif",
	}

	var i int
	for {
		token, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
		}

		assert.Equal(t, expected[i], token.Val)
		i++
	}
}

func TestNewline(t *testing.T) {
	const def = `NEWLINE: (\r\n|\r|\n)`
	const text = "   \n \n   "

	reader, err := New(strings.NewReader(def), strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}

	expected := &Token{
		Name:   "NEWLINE",
		Val:    "\n",
		Offset: 3,
	}

	token, err := reader.Next()
	assert.NoError(t, err)
	assert.Equal(t, expected, token)

	expected.Offset = 5

	token, err = reader.Next()
	assert.NoError(t, err)
	assert.Equal(t, expected, token)
}
