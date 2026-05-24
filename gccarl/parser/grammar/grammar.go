package grammar

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type RuleName string
type OptionName string
type Identifier string

type Cardinality int

const (
	CardSingle Cardinality = iota
	CardZeroOrOne
	CardMultiple
)

func (c Cardinality) Optional() bool {
	return c == CardZeroOrOne || c == CardMultiple
}

func (c Cardinality) String() string {
	switch c {
	case CardZeroOrOne:
		return "?"
	case CardMultiple:
		return "*"
	default:
		return ""
	}
}

type Part struct {
	Cardinality Cardinality
	Token       tokens.Name
	Rule        RuleName
}

func (t Part) IsRule() bool {
	return t.Rule != ""
}

type Option struct {
	Name  OptionName
	Parts []*Part
}

type Rule struct {
	Options []*Option
}

var tabChars = map[string]bool{
	" ":  true,
	"\t": true,
}

func Parse(input io.Reader) (map[RuleName]*Rule, error) {
	sc := bufio.NewScanner(input)

	rules := map[RuleName]*Rule{}
	var ruleName RuleName

	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if tabChars[line[0:1]] {
			line = strings.TrimSpace(line)
			bits := strings.Split(line, ":")

			if len(bits) != 2 {
				return nil, fmt.Errorf("invalid syntax: %s", line)
			}

			label := bits[0]
			tokenStrs := strings.Split(bits[1], " ")

			if len(tokenStrs) == 1 && tokenStrs[0] == "" {
				tokenStrs = nil
			}

			tokens := make([]*Part, len(tokenStrs))
			for i, tokenStr := range tokenStrs {
				tokens[i] = parseToken(tokenStr)
			}

			option := &Option{
				Name:  OptionName(label),
				Parts: tokens,
			}

			rules[ruleName].Options = append(rules[ruleName].Options, option)
		} else {
			ruleName = RuleName(line[:len(line)-1])
			rules[ruleName] = &Rule{}
		}
	}

	err := validate(rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func validate(rules map[RuleName]*Rule) error {
	for ruleName, rule := range rules {
		if len(rule.Options) == 0 {
			return fmt.Errorf("no options for rule %q", ruleName)
		}

		for _, option := range rule.Options {
			if len(option.Parts) == 0 {
				return fmt.Errorf("no tokens for rule %q, option %q", ruleName, option.Name)
			}

			for _, token := range option.Parts {
				if !token.IsRule() {
					continue
				}

				_, ok := rules[token.Rule]
				if !ok {
					return fmt.Errorf("rule %q not found", token.Rule)
				}
			}
		}
	}

	return nil
}

func parseToken(s string) *Part {
	var card Cardinality

	if strings.HasSuffix(s, "*") {
		card = CardMultiple
		s = s[:len(s)-1]
	} else if strings.HasSuffix(s, "?") {
		card = CardZeroOrOne
		s = s[:len(s)-1]
	}

	t := &Part{
		Cardinality: card,
	}

	if strings.ToUpper(s) == s {
		t.Token = tokens.Name(s)
	} else {
		t.Rule = RuleName(s)
	}

	return t
}
