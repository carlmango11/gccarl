package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/parser"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

//go:embed grammar.txt
var grammar string

//go:embed tokens.txt
var tokenDef string

func main() {
	var debug bool

	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()

	textF, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer textF.Close()

	tk, err := tokens.New(strings.NewReader(tokenDef), textF)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	parser, err := parser.New(strings.NewReader(grammar), true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = parser.Parse(tk, "generated", "ast")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
