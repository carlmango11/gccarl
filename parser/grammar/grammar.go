package grammar

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type RuleName string
type OptionName string
type Identifier string

type TokenType int

const (
	TTLiteral TokenType = iota + 1
	TTRule
	TTIdentifier
	TTNumber
)

func (t TokenType) String() string {
	switch t {
	case TTLiteral:
		return "lit"
	case TTRule:
		return "rule"
	case TTIdentifier:
		return "id"
	case TTNumber:
		return "num"
	default:
		return "?"
	}
}

type Token struct {
	Type  TokenType
	Multi bool

	Literal string
	Rule    RuleName
}

func (t Token) String() string {
	s := t.Type.String()
	if t.Multi {
		s += "*"
	}

	if t.Literal != "" {
		s += ":" + t.Literal
	}

	if t.Rule != "" {
		s += ":" + string(t.Rule)
	}

	return s
}

type Option struct {
	Name   OptionName
	Tokens []*Token
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

			tokens := make([]*Token, len(tokenStrs))
			for i, tokenStr := range tokenStrs {
				tokens[i] = parseToken(tokenStr)
			}

			option := &Option{
				Name:   OptionName(label),
				Tokens: tokens,
			}

			rules[ruleName].Options = append(rules[ruleName].Options, option)
		} else {
			ruleName = RuleName(line[:len(line)-1])
			rules[ruleName] = &Rule{}
		}
	}

	return rules, nil
}

func parseToken(s string) *Token {
	var multi bool
	if strings.HasSuffix(s, "*") {
		multi = true
		s = s[:len(s)-1]
	}

	t := &Token{
		Multi: multi,
	}

	switch {
	case s == "IDEN":
		t.Type = TTIdentifier
	case s == "NUM":
		t.Type = TTNumber
	case isLiteral(s):
		t.Type = TTLiteral
		t.Literal = s[1 : len(s)-1]
	default:
		t.Type = TTRule
		t.Rule = RuleName(s)
	}

	return t
}

func isLiteral(s string) bool {
	return strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")
}
