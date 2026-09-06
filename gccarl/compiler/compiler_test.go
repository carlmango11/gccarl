package compiler

import (
	"testing"

	"github.com/carlmango11/gccarl/gccarl/semantic"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	texts := []string{
		`int main() {
    char x[16] = "eat some sun sun";
    print(1, x, 16);
}`,
	}

	for _, text := range texts {
		astProg, err := cparser.Parse(text)
		require.NoError(t, err)

		prog, err := semantic.Build(astProg)
		require.NoError(t, err)

		c := New()
		_, err = c.Compile(prog)
		require.NoError(t, err)
	}
}

func TestRuntimePrograms(t *testing.T) {
	for _, text := range []string{"int main() { return 7; }", "int main() { return 42; }"} {
		tree, err := cparser.Parse(text)
		require.NoError(t, err)
		program, err := semantic.Build(tree)
		require.NoError(t, err)
		output, err := New().Compile(program)
		require.NoError(t, err)
		require.NotEmpty(t, output)
	}
}
