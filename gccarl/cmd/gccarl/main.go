package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/carlmango11/gccarl/gccarl/compiler"
	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
	"github.com/carlmango11/gccarl/gccarl/semantic"
)

//go:embed lib.asm
var libASM string

func main() {
	var outputName string
	var debug bool

	flag.StringVar(&outputName, "o", "", "output file")
	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: gccarl -o output.asm program.c")
		os.Exit(1)
	}

	text, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	tree, err := cparser.Parse(string(text))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	program, err := semantic.Build(tree)
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
