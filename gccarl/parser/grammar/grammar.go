package grammar

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
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

var tokenRegex = regexp.MustCompile(`(?:[^\s()]*\((?:[^()]|\([^()]*\))*\)\??|[^\s]+)`)

func Parse(input io.Reader) (map[RuleName][]*Option, error) {
	ruleText, err := parse(input)
	if err != nil {
		return nil, err
	}

	rules := make(map[RuleName][]*Option)

	for ruleName, options := range ruleText {
		for _, line := range options {
			bits := strings.Split(line, ":")

			if len(bits) != 2 {
				return nil, fmt.Errorf("invalid syntax: %s", line)
			}

			label := bits[0]
			def := bits[1]

			matches := tokenRegex.FindAllString(def, -1)

			var parts []*Part
			for _, m := range matches {
				var card Cardinality

				if strings.HasSuffix(m, "*") {
					card = CardMultiple
					m = m[:len(m)-1]
				} else if strings.HasSuffix(m, "?") {
					card = CardZeroOrOne
					m = m[:len(m)-1]
				}

				part := &Part{
					Cardinality: card,
				}

				bits = strings.Split(m, "=")

				if len(bits) == 2 {
					var inlineParts []*Part
					for _, b := range strings.Split(bits[1], " ") {
						inlinePart := &Part{
							Cardinality: CardSingle,
						}

						if strings.ToUpper(b) == b {
							inlinePart.Token = tokens.Name(b)
						} else {
							inlinePart.Rule = RuleName(b)
						}
					}

					rules[RuleName(bits[0])] = []*Option{
						{
							Name:  OptionName(bits[0]),
							Parts: inlineParts,
						},
					}

					part.Rule = RuleName(bits[0])
				} else {
					if strings.ToUpper(bits[0]) == bits[0] {
						part.Token = tokens.Name(bits[0])
					} else {
						part.Rule = RuleName(bits[0])
					}

					parts = append(parts, part)
				}
			}

			option := &Option{
				Name:  OptionName(label),
				Parts: parts,
			}

			rules[RuleName(ruleName)] = append(rules[RuleName(ruleName)], option)
		}
	}

	err = validate(rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func parse(input io.Reader) (map[string][]string, error) {
	sc := bufio.NewScanner(input)

	rules := map[string][]string{}
	var ruleName string

	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if tabChars[line[0:1]] {
			line = strings.TrimSpace(line)
			rules[ruleName] = append(rules[ruleName], line)
		} else {
			ruleName = line[:len(line)-1]
		}
	}

	return rules, nil
}

func validate(rules map[RuleName][]*Option) error {
	for ruleName, options := range rules {
		if len(options) == 0 {
			return fmt.Errorf("no options for rule %q", ruleName)
		}

		for _, option := range options {
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
