package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/compiler"
	"github.com/carlmango11/gccarl/gccarl/generated/ast"
	"github.com/carlmango11/gccarl/gccarl/parser"
	"github.com/carlmango11/gccarl/gccarl/semantic"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

//go:embed grammar.txt
var grammar string

//go:embed tokens.txt
var tokenDef string

//go:embed lib.asm
var libASM string

func main() {
	var outputName string
	var debug bool

	flag.StringVar(&outputName, "o", "", "output file")
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

	err = parser.Parse(tk, "../../generated", "ast")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	program, err := semantic.Build(ast.MainNode)
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

	_, err = outputFile.WriteString("\n" + libASM)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
