package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/carlmango11/gccarl/gccarl/parser"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

func main() {
	var outputDir, tokenDefFile, grammarFile, textFile, packageName string
	var debug bool

	flag.StringVar(&packageName, "p", "", "generated package name")
	flag.StringVar(&outputDir, "o", "", "output directory")
	flag.StringVar(&tokenDefFile, "tk", "", "token definition file")
	flag.StringVar(&grammarFile, "g", "", "grammar file")
	flag.StringVar(&textFile, "t", "", "program text file")
	flag.BoolVar(&debug, "d", false, "enable debug logging")
	flag.Parse()

	tokenFile, err := os.Open(tokenDefFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer tokenFile.Close()

	grammarF, err := os.Open(grammarFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer grammarF.Close()

	textF, err := os.Open(textFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer textF.Close()

	tk, err := tokens.New(tokenFile, textF)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	parser, err := parser.New(grammarF, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = parser.Parse(tk, outputDir, packageName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	astOutput, err := os.Create(outputDir + "/ast.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer astOutput.Close()
}
