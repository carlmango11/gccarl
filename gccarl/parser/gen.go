package parser

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

//go:embed template_go.txt
var typesTemplate string

type Generator struct {
	packageName string
	grammar     map[grammar.RuleName]*grammar.Rule
	nodeC       int
	ast         strings.Builder
	types       strings.Builder
}

func (g *Generator) generate(n *Node, outDir string) error {
	g.generateTypes()
	g.generateAST(n)

	outDir += outDir + "/" + g.packageName + "/"
	typesFile := outDir + "/types.go"
	astFile := outDir + "/ast.go"

	err := g.writeFile(g.types.String(), typesFile)
	if err != nil {
		return err
	}

	err = g.writeFile(g.ast.String(), astFile)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) writeFile(output, fileName string) error {
	err := os.MkdirAll(fileName, 0755)
	if err != nil {
		return err
	}

	astF, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer astF.Close()

	_, err = astF.WriteString(output)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateAST(n *Node) {
	g.ast.WriteString("package " + g.packageName)

	for _, v := range n.Values {
		if v.Node == nil {
			g.ast.WriteString(fmt.Sprintf("\t%s %s", v.Token.Name, v.Token.Name))
		} else {
			g.generateNode(v.Node)
		}
	}
}

func (g *Generator) generateNode(n *Node) string {
	g.ast.WriteString(fmt.Sprintf("var n%d = &%v {", g.nodeC, kebabToCamel(string(n.Key.Rule))))
	g.nodeC++

	return ""
}

func (g *Generator) generateTypes() {
	g.types.WriteString("package " + g.packageName)

	tks := map[tokens.Name]bool{}
	optionTypes := map[RuleKey][]*grammar.Part{}

	for ruleName, rule := range g.grammar {
		for _, o := range rule.Options {
			optionTypes[RuleKey{ruleName, o.Name}] = o.Parts

			for _, p := range o.Parts {
				if p.Token != "" {
					tks[p.Token] = true
				}
			}
		}
	}

	g.types.WriteString("\n" + generateTokens(tks))
	g.types.WriteString("\n" + generateOptionTypes(optionTypes))

	for name, rule := range g.grammar {
		s := g.generateRuleType(name, rule.Options)
		g.types.WriteString("\n" + s)
	}
}

func generateOptionTypes(types map[RuleKey][]*grammar.Part) string {
	var sb strings.Builder

	for rk, parts := range types {
		ruleName := kebabToCamel(string(rk.Rule))
		optionName := kebabToCamel(string(rk.Option))

		sb.WriteString(fmt.Sprintf("type %s%s struct {\n", ruleName, optionName))

		for _, part := range parts {
			card := ""
			if part.Cardinality == grammar.CardMultiple {
				card = "[]"
			}

			if part.Token != "" {
				sb.WriteString(fmt.Sprintf("\t%s %s\n", part.Token, part.Token))
			} else {
				sb.WriteString(fmt.Sprintf("\t%s %s*%s%s\n", ruleName, card, ruleName, optionName))
			}
		}

		sb.WriteString("}\n\n")
	}

	sb.WriteString("\n}")
	return sb.String()
}

func generateTokens(tks map[tokens.Name]bool) string {
	var sb strings.Builder

	for tk := range tks {
		sb.WriteString(fmt.Sprintf("\ntype %v string", tk))
	}

	return sb.String()
}

func (g *Generator) generateRuleType(rule grammar.RuleName, options []*grammar.Option) string {
	ruleName := kebabToCamel(string(rule))

	text := strings.ReplaceAll(typesTemplate, "{{name}}", ruleName)

	var enums, fields, optionTypes strings.Builder

	for _, o := range options {
		optionName := kebabToCamel(string(o.Name))

		enums.WriteString(fmt.Sprintf("\n\t%sType%s %sType = \"%s\"", ruleName, optionName, ruleName, o.Name))
		fields.WriteString(fmt.Sprintf("\n\t%s *%s%s", optionName, ruleName, optionName))

		optionTypes.WriteString("\n" + optionType(o, ruleName, optionName))
	}

	text = strings.ReplaceAll(text, "{{type-enums}}", enums.String())
	text = strings.ReplaceAll(text, "{{fields}}", fields.String())
	text += optionTypes.String()

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

func optionType(option *grammar.Option, rule, optionName string) string {
	text := fmt.Sprintf("type %s%s struct {", rule, optionName)

	for _, p := range option.Parts {
		brackets := ""
		if p.Cardinality == grammar.CardMultiple {
			brackets = "[]"
		}

		if p.Token != "" {
			text += fmt.Sprintf("\n\t%s %s%s", p.Token, brackets, p.Token)
		} else {
			text += fmt.Sprintf("\n\t%s %s*%s", rule, brackets, rule)
		}
	}

	text += "\n}"
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
