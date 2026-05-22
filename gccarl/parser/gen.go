package parser

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
)

//go:embed template_go.txt
var typesTemplate string

type Generator struct {
	packageName string
	grammar     map[grammar.RuleName]*grammar.Rule
}

func (g *Generator) generate(n *Node, outDir string) error {
	outDir += "/" + g.packageName

	err := g.writeTypes(outDir)
	if err != nil {
		return err
	}

	err = g.writeAST(n, outDir)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) writeAST(n *Node, outDir string) error {
	ast, err := g.generateAST(n)
	if err != nil {
		return err
	}

	err = os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	astF, err := os.Create(outDir + "/ast.go")
	if err != nil {
		return err
	}

	defer astF.Close()

	_, err = astF.WriteString(ast)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateAST(n *Node) (string, error) {
	var sb strings.Builder
	sb.WriteString("package " + g.packageName)

	var i int
	for _, v := range n.Values {
		if v.Node == nil {
			sb.WriteString(fmt.Sprintf("\t%s %s"))
		} else {
			generateNode(i, v.Node)
		}
	}

	return sb.String(), nil
}

func generateNode(i int, n *Node) string {
	text := fmt.Sprintf("var n%d = &%v {", i, kebabToCamel(string(n.Key.Rule)))

	for _, v := range n.Values {
		if v.Node == nil {
			generateNode(i, v.Node)

		} else {

		}
	}

	return text
}

func (g *Generator) writeTypes(outDir string) error {
	var sb strings.Builder
	sb.WriteString("package " + g.packageName)

	for name, rule := range g.grammar {
		s := g.generateType(name, rule.Options)
		sb.WriteString("\n" + s)
	}

	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	typesF, err := os.Create(outDir + "/types.go")
	if err != nil {
		return err
	}

	defer typesF.Close()

	_, err = typesF.WriteString(sb.String())
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateType(rule grammar.RuleName, options []*grammar.Option) string {
	ruleName := kebabToCamel(string(rule))

	text := strings.ReplaceAll(typesTemplate, "{{name}}", ruleName)

	var enums, fields strings.Builder

	for _, o := range options {
		optionName := kebabToCamel(string(o.Name))

		enums.WriteString(fmt.Sprintf("\n\t%sType%s %sType = \"%s\"", ruleName, optionName, ruleName, o.Name))
		fields.WriteString(fmt.Sprintf("\n\t%s *%s", optionName, ruleName))
	}

	text = strings.ReplaceAll(text, "{{type-enums}}", enums.String())
	text = strings.ReplaceAll(text, "{{fields}}", fields.String())

	for _, o := range options {
		ruleOption := ruleName + kebabToCamel(string(o.Name))

		text += fmt.Sprintf("\ntype %s struct {", ruleOption)

		for _, p := range o.Parts {
			card := ""
			if p.Cardinality == grammar.CardMultiple {
				card = "[]"
			}

			text += fmt.Sprintf("\t%v %s*%s", p.Token, card, p.Token)
		}
	}

	return text
}

func kebabToCamel(s string) string {
	parts := strings.Split(s, "-")

	for i, part := range parts {
		if part == "" {
			continue
		}

		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}

	return strings.Join(parts, "")
}
