package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/ast"
	"github.com/carlmango11/gccarl/gccarl/compiler"
	"github.com/carlmango11/gccarl/gccarl/parser"
	"github.com/carlmango11/gccarl/gccarl/semantic"
)

//go:embed c.txt
var grammar string

func main() {
	var outputName string
	var debug bool
	flag.StringVar(&outputName, "o", "", "output file name")
	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()

	p, err := parser.New(strings.NewReader(grammar), debug)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	input, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer input.Close()

	parsed, err := p.Parse(input)
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
