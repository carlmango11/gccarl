package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/ast"
	"github.com/carlmango11/gccarl/gccarl/compiler"
	"github.com/carlmango11/gccarl/gccarl/parser"
	"github.com/carlmango11/gccarl/gccarl/semantic"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

//go:embed grammar.txt
var grammar string

//go:embed tokens.txt
var tokenDef string

func main() {
	var outputName string
	var debug bool
	flag.StringVar(&outputName, "o", "", "output file name")
	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()

	input, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer input.Close()

	inputBytes, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	tk, err := tokens.New(tokenDef, string(inputBytes))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	parser, err := parser.New(strings.NewReader(grammar), false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	parsed, err := parser.Parse(tk)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	astProg, err := ast.Build(parsed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	program, err := semantic.Build(astProg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	cc := compiler.New()

	c, err := cc.Compile(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	outputFile, err := os.Create(outputName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer outputFile.Close()

	_, err = outputFile.Write(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
