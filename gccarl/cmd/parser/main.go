package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/parser"
)

//go:embed grammar.txt
var grammar string

//go:embed tokens.txt
var tokenDef string

func main() {
	var grammarFile, tokensFile, outDir, packageName string
	var debug bool
	flag.StringVar(&grammarFile, "grammar", "", "grammar file (defaults to bundled grammar)")
	flag.StringVar(&tokensFile, "tokens", "", "token definitions (defaults to bundled definitions)")
	flag.StringVar(&outDir, "out", "generated", "output parent directory")
	flag.StringVar(&packageName, "package", "ast", "generated package name")
	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()
	if err := generate(grammarFile, tokensFile, outDir, packageName, debug); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generate(grammarFile, tokensFile, outDir, packageName string, debug bool) error {
	if flag.NArg() != 0 {
		return fmt.Errorf("generation takes no program argument; use -grammar and -tokens")
	}
	grammarText, err := readDefinition(grammarFile, grammar)
	if err != nil {
		return err
	}
	tokenText, err := readDefinition(tokensFile, tokenDef)
	if err != nil {
		return err
	}
	p, err := parser.New(strings.NewReader(grammarText), debug)
	if err != nil {
		return err
	}
	return p.Generate(strings.NewReader(tokenText), outDir, packageName)
}

func readDefinition(path, fallback string) (string, error) {
	if path == "" {
		return fallback, nil
	}
	data, err := os.ReadFile(path)
	return string(data), err
}
