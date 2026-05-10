package main

import (
	"compiler/compiler"
	"compiler/parser"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed c.txt
var grammar string

func main() {
	var outputName string
	flag.StringVar(&outputName, "o", "", "output file name")
	flag.Parse()

	p, err := parser.New(strings.NewReader(grammar), false)
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

	ast, err := p.Parse(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	c, err := compiler.Compile(ast)
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
