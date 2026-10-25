package preprocessor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type Processor struct {
	includePaths []string

	macros map[string]*Define
}

func New(includePaths []string) *Processor {
	return &Processor{
		includePaths: includePaths,
		macros:       make(map[string]*Define),
	}
}

type Define struct {
	Name   string
	Params []string
	Tokens []*tokens.Token
}

type DirectiveType int

const (
	DirIf DirectiveType = iota
	DirElif
	DirElse
	DirEndIf
	DirDefine
	DirUndefine
	DirIfDef
	DirIfNotDef
	DirInclude
)

func (d DirectiveType) IsIf() bool {
	switch d {
	case DirIf:
		return true
	default:
		return false
	}
}

type Directive struct {
	Type      DirectiveType
	MacroName string
	Expr      *Expr
	Define    *Define
	Include   *Include
}

type Include struct {
	Name   string
	System bool
}

type section struct {
	dir  *Directive
	code *bytes.Buffer
}

func (p *Processor) Process(filename string) (string, error) {
	return p.processFile(filename, false)
}

func (p *Processor) processFile(filename string, system bool) (string, error) {
	textBytes, err := p.readFile(filename, system)
	if err != nil {
		return "", err
	}

	sc := bufio.NewScanner(bytes.NewReader(textBytes))

	current := bytes.NewBuffer(nil)

	var sections []section

	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)

		// handle comment
		i := strings.Index(sc.Text(), "//")
		if i != -1 {
			line = sc.Text()[:i]
		}

		if len(line) > 0 && line[0] == '#' {
			sections = append(sections, section{
				code: current,
			})
			current = bytes.NewBuffer(nil)

			// handle directive
			if line[len(line)-1] == '\\' {
				panic("impl")
			}

			d, err := p.parseDirective(line)
			if err != nil {
				return "", fmt.Errorf("error parsing directive: %v", err)
			}

			sections = append(sections, section{
				dir: d,
			})
		} else {
			// process code
			current.WriteString(line + "\n")
		}
	}

	sections = append(sections, section{
		code: current,
	})

	return p.processSections(sections)
}

func cTokenReader(text string) *tokens.Reader {
	return tokens.NewReader(cparser.TokenDefinitions, text)
}

func (p *Processor) parseDirective(line string) (*Directive, error) {
	line = line[1:] // remove #

	bits := strings.Split(line, " ")

	dirType := toDirType(bits[0])
	rest := line[len(bits[0]):]

	dir := &Directive{
		Type: dirType,
	}

	switch dirType {
	case DirIf, DirElse:
		expr, err := parseExpr(rest)
		if err != nil {
			return nil, err
		}

		dir.Expr = expr
	case DirIfNotDef, DirUndefine:
		dir.MacroName = strings.TrimSpace(rest)
	case DirInclude:
		rest = strings.TrimSpace(rest)

		filename := rest[1 : len(rest)-1]
		var system bool

		// TODO validate
		if rest[0] == '<' {
			system = true
		}

		dir.Include = &Include{
			Name:   filename,
			System: system,
		}
	case DirDefine:
		define, err := parseDefine(rest)
		if err != nil {
			return nil, err
		}

		dir.Define = &define
	}

	return dir, nil
}

func parseDefine(rest string) (Define, error) {
	rest = strings.TrimSpace(rest)

	var params []string
	var name string

	lp := strings.Index(rest, "(")
	if lp != -1 {
		// func macro
		name = strings.TrimSpace(rest[:lp])

		rp := strings.Index(rest, ")")
		if rp == -1 {
			return Define{}, fmt.Errorf("error parsing define")
		}

		paramStr := rest[lp+1 : rp]
		for _, bit := range strings.Split(paramStr, ",") {
			bit = strings.TrimSpace(bit)
			params = append(params, bit)
		}

		rest = rest[rp+1:]
	} else {
		space := strings.Index(rest, " ")
		if space == -1 {
			return Define{}, fmt.Errorf("error parsing define")
		}

		name = strings.TrimSpace(rest[:space])
		rest = rest[space+1:]
	}

	// rest is just the tokens
	reader := cTokenReader(rest)
	var tokens []*tokens.Token

	for {
		token, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			return Define{}, err
		}

		tokens = append(tokens, token)
	}

	return Define{
		Name:   name,
		Params: params,
		Tokens: tokens,
	}, nil
}

func toDirType(s string) DirectiveType {
	switch s {
	case "if":
		return DirIf
	case "elif":
		return DirElif
	case "ifdef":
		return DirIfDef
	case "ifndef":
		return DirIfNotDef
	case "else":
		return DirElse
	case "define":
		return DirDefine
	case "undef":
		return DirUndefine
	case "endif":
		return DirEndIf
	case "include":
		return DirInclude
	}

	panic(s + " not valid directive")
}

type ifState struct {
	executed bool
	skip     bool
}

func (p *Processor) processSections(sections []section) (string, error) {
	out := bytes.NewBuffer(nil)

	var ifs []*ifState

	for _, s := range sections {
		if s.dir != nil {
			switch s.dir.Type {
			case DirIf, DirIfNotDef, DirIfDef:
				skip := len(ifs) > 0 && ifs[len(ifs)-1].skip
				executed := true // pretend we executed if skipping

				if !skip {
					executed = p.eval(s.dir)
					skip = !executed
				}

				ifs = append(ifs, &ifState{
					executed: executed,
					skip:     skip,
				})

				continue
			case DirElif:
				curr := ifs[len(ifs)-1]

				if curr.executed {
					curr.skip = true
					continue
				}

				curr.executed = p.eval(s.dir)
				curr.skip = !curr.executed

				continue

			case DirElse:
				curr := ifs[len(ifs)-1]

				if curr.executed {
					curr.skip = true
					continue
				}

				curr.executed = true
				curr.skip = false

				continue

			case DirEndIf:
				ifs = ifs[:len(ifs)-1]
				continue
			}
		}

		if len(ifs) > 0 && ifs[len(ifs)-1].skip {
			continue
		}

		if s.code != nil {
			outputCode, err := p.processCode(s.code)
			if err != nil {
				return "", err
			}

			out.WriteString(outputCode)
			continue
		}

		switch s.dir.Type {
		case DirUndefine:
			delete(p.macros, s.dir.MacroName)
		case DirDefine:
			if len(s.dir.Define.Params) > 0 {
				panic("impl")
			}

			p.macros[s.dir.Define.Name] = s.dir.Define
		case DirInclude:
			includeCode, err := p.processFile(s.dir.Include.Name, s.dir.Include.System)
			if err != nil {
				return "", err
			}

			out.WriteString(includeCode)
		}
	}

	return out.String(), nil
}

func (p *Processor) eval(d *Directive) bool {
	switch d.Type {
	case DirIf:
		return p.evalExpr(d.Expr)
	case DirIfNotDef:
		_, ok := p.macros[d.MacroName]
		return !ok
	case DirIfDef:
		_, ok := p.macros[d.MacroName]
		return ok
	}

	panic("error")
}

func (p *Processor) processCode(code *bytes.Buffer) (string, error) {
	out := bytes.NewBuffer(nil)
	reader := cTokenReader(code.String())

	for {
		tk, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		macro, ok := p.macros[tk.Val]
		if !ok {
			out.WriteString(tk.Val + " ")
			continue
		}

		if len(macro.Params) > 0 {
			panic("impl")
		}

		for _, r := range macro.Tokens {
			out.WriteString(r.Val + " ") // TODO this is shit but I've lost the whitespace
		}
	}

	return out.String(), nil
}

func (p *Processor) readFile(filename string, system bool) ([]byte, error) {
	if !system {
		return os.ReadFile(filename)
	}

	for _, path := range p.includePaths {
		textBytes, err := os.ReadFile(filepath.Join(path, filename))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return nil, err
		}

		return textBytes, nil
	}

	return nil, fmt.Errorf("could not find file %s", filename)
}
