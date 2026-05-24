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

var debug bool = true

type Generator struct {
	packageName string
	grammar     map[grammar.RuleName]*grammar.Rule
	nodeC       int
	ast         strings.Builder
	types       strings.Builder
	astNodes    strings.Builder
}

func (g *Generator) generate(n *Node, outDir string) error {
	g.generateTypes()
	g.generateAST(n)

	outDir = outDir + "/" + g.packageName + "/"

	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	typesFile := outDir + "/types.go"
	astFile := outDir + "/ast.go"

	err = g.writeFile(g.types.String(), typesFile)
	if err != nil {
		return err
	}

	err = g.writeFile(g.astNodes.String()+"\n"+g.ast.String(), astFile)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) writeFile(output, fileName string) error {
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

	g.generateNode(n)
}

func (g *Generator) generateNode(n *Node) {
	n.ID = g.nodeC
	g.nodeC++

	for _, v := range n.Values {
		if v.Node != nil {
			g.generateNode(v.Node)
		}
	}

	g.ast.WriteString(fmt.Sprintf("\nvar n%v = &%v {\n", n.ID, kebabToCamel(string(n.Key.Rule))))
	g.ast.WriteString(fmt.Sprintf("\tType: %s,\n", typeEnumField(n.Key.Rule, n.Key.Option)))
	g.ast.WriteString(fmt.Sprintf("\t%s: &%v{\n", optionField(n.Key.Option), optionStruct(n.Key.Rule, n.Key.Option)))

	for _, v := range n.Values {
		if v.Node != nil {
			g.ast.WriteString(fmt.Sprintf("\t\t%v: n%v,\n", optionRuleFieldName(v.Node.Key.Rule), v.Node.ID))
		} else {
			g.ast.WriteString(fmt.Sprintf("\t\t%v: %q,\n", v.Token.Name, v.Token.Val))
		}
	}

	g.ast.WriteString("\t},\n}\n")
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

	for name, rule := range g.grammar {
		s := g.generateRuleType(name, rule.Options)
		g.types.WriteString("\n" + s)
	}

	g.types.WriteString("\n" + generateTokens(tks))
	g.types.WriteString("\n" + generateOptionTypes(optionTypes))
}

func generateOptionTypes(types map[RuleKey][]*grammar.Part) string {
	var sb strings.Builder

	for rk, parts := range types {
		output := generateOptionType(rk, parts)
		sb.WriteString(output)
	}

	sb.WriteString("\n")
	return sb.String()
}

func optionStruct(r grammar.RuleName, o grammar.OptionName) string {
	ruleName := kebabToCamel(string(r))
	optionName := kebabToCamel(string(o))

	return fmt.Sprintf("%s_%sOption", ruleName, optionName)
}

func generateOptionType(rk RuleKey, parts []*grammar.Part) string {
	var sb strings.Builder

	structName := optionStruct(rk.Rule, rk.Option)

	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	for _, part := range parts {
		card := ""
		if part.Cardinality == grammar.CardMultiple {
			card = "[]"
		}

		if part.Token != "" {
			sb.WriteString(fmt.Sprintf("\t%s %s\n", part.Token, part.Token))
		} else {
			ruleCodeName := optionRuleFieldName(part.Rule)
			sb.WriteString(fmt.Sprintf("\t%s %s*%s\n", ruleCodeName, card, ruleCodeName))
		}
	}

	sb.WriteString("}\n\n")
	out := sb.String()

	debugf("generated option %q:\n%s", structName, out)

	return out
}

func optionRuleFieldName(r grammar.RuleName) string {
	return kebabToCamel(string(r))
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

	var enums, fields strings.Builder

	for _, o := range options {
		optionName := kebabToCamel(string(o.Name))

		enums.WriteString(fmt.Sprintf("\n\t%s %s = \"%s\"", typeEnumField(rule, o.Name), typeEnumType(rule), o.Name))
		fields.WriteString(fmt.Sprintf("\n\t%s *%s", optionName, optionStruct(rule, o.Name)))
	}

	text = strings.ReplaceAll(text, "{{type-enums}}", enums.String())
	text = strings.ReplaceAll(text, "{{fields}}", fields.String())

	return text
}

func (g *Generator) isMultiple(key RuleKey) bool {
	for _, o := range g.grammar[key.Rule].Options {
		if o.Name == key.Option {
			return o.
		}
	}
}

func typeEnumField(rule grammar.RuleName, option grammar.OptionName) string {
	return fmt.Sprintf("%sType%s", kebabToCamel(string(rule)), kebabToCamel(string(option)))
}

func typeEnumType(rule grammar.RuleName) string {
	return fmt.Sprintf("%sType", kebabToCamel(string(rule)))
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

func optionField(o grammar.OptionName) string {
	return kebabToCamel(string(o))
}
