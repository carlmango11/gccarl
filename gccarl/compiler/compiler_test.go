package compiler

import (
	"bytes"
	"os"
	"testing"

	"github.com/carlmango11/gccarl/gccarl/ast"
	"github.com/carlmango11/gccarl/gccarl/parser"
	"github.com/carlmango11/gccarl/gccarl/semantic"
	"github.com/carlmango11/gccarl/gccarl/tokens"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	texts := []string{
		`int main() {
    char x[16] = "eat some sun sun";
    print(1, x, 16);
}`,
	}

	tokenDef, err := os.ReadFile("../cmd/gccarl/tokens.txt")
	if err != nil {
		t.Fatal(err)
	}

	grammar, err := os.ReadFile("../cmd/gccarl/grammar.txt")
	if err != nil {
		t.Fatal(err)
	}

	for _, text := range texts {
		tk, err := tokens.New(string(tokenDef), text)
		if err != nil {
			t.Fatal(err)
		}

		p, err := parser.New(bytes.NewReader(grammar), false)
		if err != nil {
			t.Fatal(err)
		}

		parsed, err := p.Parse(tk)
		if err != nil {
			t.Fatal(err)
		}

		astProg, err := ast.Build(parsed)
		require.NoError(t, err)

		prog, err := semantic.Build(astProg)
		require.NoError(t, err)

		c := New()
		_, err = c.Compile(prog)
		require.NoError(t, err)
	}
}
