package tokens

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
