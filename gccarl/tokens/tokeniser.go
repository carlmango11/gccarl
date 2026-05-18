package tokens

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

type Name string

type Token struct {
	Name    Name
	Val     string
	Literal bool
}

type Reader struct {
	tokens []*TokenDef
	text   string
	i      int
}

func New(tokenDefText, text string) (*Reader, error) {
	tokenDefs, err := parseTokenDefs(tokenDefText)
	if err != nil {
		return nil, err
	}

	return &Reader{tokens: tokenDefs, text: text}, nil
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

func (tk *Reader) Next() (*Token, error) {
	tk.skipWhitespace()

	if tk.i == len(tk.text) {
		return nil, io.EOF
	}

	var matches []*Token

	for _, t := range tk.tokens {
		match, ok := tk.match(t)
		if !ok {
			continue
		}

		matches = append(matches, match)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("tokeniser: no matches found for token: %s", tk.text[tk.i])
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
	tk.i += len(longest.Val)

	return matches[0], nil
}

func (tk *Reader) match(t *TokenDef) (*Token, bool) {
	if t.Literal != "" {
		end := tk.i + len(t.Literal)
		if end > len(tk.text) {
			return nil, false
		}

		if tk.text[tk.i:end] == t.Literal {
			return &Token{
				Name:    Name(t.Name),
				Val:     t.Literal,
				Literal: true,
			}, true
		}

		return nil, false
	}

	v := t.Regex.FindString(tk.text[tk.i:])
	if v != "" {
		return &Token{
			Name: Name(t.Name),
			Val:  v,
		}, true
	}

	return nil, false
}

func (tk *Reader) skipWhitespace() {
	v := whitespace.FindString(tk.text[tk.i:])
	tk.i += len(v)
}

func (tk *Reader) Index() int {
	return tk.i
}

func (tk *Reader) Reset(i int) {
	tk.i = i
}
