package tokeniser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

var whitespace = regexp.MustCompile(`^[ \t\n]+`)

type TokenDef struct {
	Name    string
	Literal string
	Regex   *regexp.Regexp
}

type Token struct {
	Name    string
	Val     string
	Literal bool
}

type Tokeniser struct {
	tokens []*TokenDef
	text   string
	i      int
}

func New(tokenDefText, text string) (*Tokeniser, error) {
	tokenDefs, err := parseTokenDefs(tokenDefText)
	if err != nil {
		return nil, err
	}

	return &Tokeniser{tokens: tokenDefs, text: text}, nil
}

func parseTokenDefs(r string) ([]*TokenDef, error) {
	sc := bufio.NewScanner(strings.NewReader(r))

	var defs []*TokenDef

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		bits := strings.Split(line, ":")
		if len(bits) != 2 {
			return nil, fmt.Errorf("tokeniser: invalid token definition: %s", line)
		}

		name := strings.TrimSpace(bits[0])
		def := strings.TrimSpace(bits[1])

		if def[0] == '\'' && def[len(def)-1] == '\'' {
			defs = append(defs, &TokenDef{Name: name, Literal: def[1 : len(def)-1]})
		} else {
			reg, err := regexp.Compile("^" + def)
			if err != nil {
				return nil, fmt.Errorf("tokeniser: invalid token definition: %s", def)
			}

			defs = append(defs, &TokenDef{Name: name, Regex: reg})
		}
	}

	return defs, nil
}

func (s *Tokeniser) Next() (*Token, error) {
	s.skipWhitespace()

	if s.i == len(s.text) {
		return nil, io.EOF
	}

	var matches []*Token

	for _, t := range s.tokens {
		match, ok := s.match(t)
		if !ok {
			continue
		}

		matches = append(matches, match)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("expected at least one match")
	}

	sort.Slice(matches, func(i, j int) bool {
		m1 := matches[i]
		m2 := matches[j]

		if m1.Literal && m2.Literal {
			return len(matches[i].Val) > len(matches[j].Val)
		}

		return m1.Literal
	})

	longest := matches[0]
	s.i += len(longest.Val)

	return matches[0], nil
}

func (s *Tokeniser) match(t *TokenDef) (*Token, bool) {
	if t.Literal != "" {
		end := s.i + len(t.Literal)
		if end > len(s.text) {
			return nil, false
		}

		if s.text[s.i:end] == t.Literal {
			return &Token{
				Name:    t.Name,
				Val:     t.Literal,
				Literal: true,
			}, true
		}

		return nil, false
	}

	v := t.Regex.FindString(s.text[s.i:])
	if v != "" {
		return &Token{
			Name: t.Name,
			Val:  v,
		}, true
	}

	return nil, false
}

func (s *Tokeniser) skipWhitespace() {
	v := whitespace.FindString(s.text[s.i:])
	s.i += len(v)
}
