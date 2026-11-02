package preprocessor

import (
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	expected := []string{
		"4",
		"ifndef",
		"FOUR",
		"4",
		"(",
		"*",
		"(",
		"int",
		"*",
		")",
		"(",
		"(",
		"x",
		"+",
		"=",
		"sizeof",
		"(",
		"int",
		")",
		")",
		"-",
		"sizeof",
		"(",
		"int",
		")",
		")",
		")",
		";",
	}

	r, err := NewReader([]string{"testdata/testlib"}, "testdata/test1.h")
	require.NoError(t, err)

	var i int
	for {
		fmt.Println(i)
		tk, err := r.Next()
		if err == io.EOF {
			break
		}

		require.NoError(t, err)
		require.NotNil(t, tk, "idx %d", i)
		require.Equal(t, expected[i], tk.Val, "idx %d", i)

		i++
	}
}

func TestPrintf(t *testing.T) {
	r, err := NewReader(nil, "testdata/printf.c")
	require.NoError(t, err)

	out := &strings.Builder{}
	for {
		tk, err := r.Next()
		if err == io.EOF {
			break
		}

		require.NoError(t, err)
		out.WriteString(tk.Val + " ")
	}

	log.Printf("out: %v", out.String())
}
