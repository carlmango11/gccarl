package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
)

//go:embed c.txt
var grammar string

//go:embed tokens.txt
var tokens string

func main() {
	var outputName string
	var debug bool
	flag.StringVar(&outputName, "o", "", "output file name")
	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()

	//_, err := parser.New(strings.NewReader(grammar), debug)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}

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

	tk, err := tokens.New(tokens, string(inputBytes))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for {
		t, err := tk.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Printf("%s\n", t)
	}
	return

	//parsed, err := p.Parse(input)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//
	//astProg, err := ast.Build(parsed)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//
	//program, err := semantic.Build(astProg)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//
	//cc := compiler.New()
	//
	//c, err := cc.Compile(program)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//
	//outputFile, err := os.Create(outputName)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//
	//defer outputFile.Close()
	//
	//_, err = outputFile.Write(c)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	return
	//}
}
