package grammar

import (
	"strings"
	"testing"
)

func TestMalformedGrammar(t *testing.T) {
	for _, text := range []string{"  a:NUM", "main\n a:NUM", "main:\n a", "main:\n a:", ""} {
		t.Run(text, func(t *testing.T) {
			if _, err := Parse(strings.NewReader(text)); err == nil {
				t.Fatal("accepted malformed grammar")
			}
		})
	}
}

func TestPartCardinality(t *testing.T) {
	rules, err := Parse(strings.NewReader("main:\n  main: NUM+   IDEN?\tPLUS*"))
	if err != nil {
		t.Fatal(err)
	}
	parts := rules["main"].Options[0].Parts
	if len(parts) != 3 || parts[0].Cardinality != CardOneOrMore || parts[1].Cardinality != CardZeroOrOne || parts[2].Cardinality != CardMultiple {
		t.Fatalf("wrong parts: %#v", parts)
	}
}
