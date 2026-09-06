package tokens

import (
	"strings"
	"testing"
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
