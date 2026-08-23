package gccarl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrammar(t *testing.T) {
	f, err := os.Open("./input/grammar_test.c")
	require.NoError(t, err)

	defer f.Close()

}
