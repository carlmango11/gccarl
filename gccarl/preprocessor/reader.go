package preprocessor

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type reader struct {
	includePaths []string
	tokens       tokens.ReaderPeeker

	macros      map[string]*Macro
	prevNewline bool
	directive   *directiveState
	macro       *macroState
	ifs         []*ifState
}

type macroState struct {
	def      *Macro
	tokenIdx int
	params   map[string][]*tokens.Token

	swap    []*tokens.Token
	swapIdx int
}

type ifState struct {
	executed bool
	skip     bool
}

type directiveState struct {
	dirType  DirectiveType
	finished bool
	subFile  *reader
}

type Macro struct {
	Args   []string
	Tokens []*tokens.Token
}

func newReader(filename string, includePaths []string, system bool, macros map[string]*Macro) (*reader, error) {
	textBytes, err := readFile(filename, includePaths, system)
	if err != nil {
		return nil, err
	}

	r := tokens.NewReader(cparser.TokenDefinitions, string(textBytes))

	return &reader{
		tokens:       r,
		includePaths: includePaths,
		macros:       macros,
		prevNewline:  true,
	}, nil
}

func (r *reader) Next() (*tokens.Token, error) {
	for {
		if r.directive != nil {
			tk, err := r.advanceDirective()
			if err != nil {
				return nil, err
			}

			if tk != nil {
				return tk, nil
			}

			continue
		}

		if r.macro != nil {
			tk, err := r.advanceMacro()
			if err != nil {
				return nil, err
			}

			if tk != nil {
				return tk, nil
			}

			continue
		}

		tk, err := r.nextToken()
		if err != nil {
			return nil, err
		}

		switch tk.Name {
		case "NEWLINE":
			r.prevNewline = true
			continue
		case "HASH":
			if r.prevNewline {
				// directive started
				r.prevNewline = false
				r.directive = &directiveState{}
				continue
			}
		}

		if len(r.ifs) > 0 && r.ifs[len(r.ifs)-1].skip {
			continue
		}

		// C code
		cTk, err := r.processCToken(tk)
		if err != nil {
			return nil, err
		}

		if cTk != nil {
			return cTk, nil
		}
	}
}

func (r *reader) nextToken() (*tokens.Token, error) {
	tk, err := r.tokens.Next()
	if err != nil {
		if err == io.EOF {
			// TODO do EOF checks
		}
		return nil, err
	}
	log.Printf(">> %v", tk.Val)

	return tk, nil
}

func (r *reader) processCToken(tk *tokens.Token) (*tokens.Token, error) {
	macro, ok := r.macros[tk.Val]
	if !ok {
		return tk, nil
	}

	var params map[string][]*tokens.Token

	if len(macro.Args) > 0 {
		next, err := r.tokens.Peek()
		if err == nil {
			if next.Name == "LPAREN" {
				// need to read the params
				params, err = r.readParams(macro.Args)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	r.macro = &macroState{
		def:    macro,
		params: params,
	}

	return nil, nil
}

func (r *reader) advanceDirective() (*tokens.Token, error) {
	if r.directive.dirType == DirNone {
		tk, err := r.nextToken()
		if err != nil {
			return nil, err
		}

		dirType, err := toDirType(tk)
		if err != nil {
			return nil, err
		}

		r.directive.dirType = dirType
	}

	switch r.directive.dirType {
	case DirInclude:
		tk, err := r.advanceInclude()
		if err == nil {
			return tk, nil
		}

		if err != io.EOF {
			return nil, err
		}

		r.directive.finished = true
	case DirDefine:
		err := r.processDefine()
		if err != nil {
			return nil, err
		}
	case DirUndefine:
		err := r.processUndefine()
		if err != nil {
			return nil, err
		}
	case DirIf:
		err := r.processIf()
		if err != nil {
			return nil, err
		}
	case DirEndIf:
		err := r.processEndIf()
		if err != nil {
			return nil, err
		}
	case DirIfNotDef:
		err := r.processIfDef(false)
		if err != nil {
			return nil, err
		}
	case DirIfDef:
		err := r.processIfDef(true)
		if err != nil {
			return nil, err
		}
	default:
		panic(fmt.Sprintf("unexpected directive type: %v", r.directive.dirType))
	}

	// all directives should complete processing by this point
	if !r.directive.finished {
		panic("impl")
	}

	tk, err := r.nextToken()
	if err != nil {
		return nil, err
	}

	if tk.Name == "BSLASH" {
		panic("impl")
	}

	if tk.Name == "NEWLINE" {
		r.directive = nil
		r.prevNewline = true
		return nil, nil
	}

	return nil, fmt.Errorf("expected newline after end of directive")
}

func (r *reader) advanceInclude() (*tokens.Token, error) {
	if r.directive.subFile != nil {
		subTk, err := r.directive.subFile.Next()
		if err != nil {
			return nil, err
		}

		return subTk, nil
	}

	tk, err := r.nextToken()
	if err != nil {
		return nil, err
	}

	var system bool

	switch tk.Name {
	case "SYS_INCLUDE_STR":
		system = true
	case "STR":
	default:
		return nil, fmt.Errorf("invalid #include: %s", tk.Val)
	}

	filename := tk.Val[1 : len(tk.Val)-1]

	sub, err := newReader(filename, r.includePaths, system, r.macros)
	if err != nil {
		return nil, err
	}

	r.directive.subFile = sub
	return nil, nil
}

func readFile(filename string, includePaths []string, system bool) ([]byte, error) {
	if !system {
		return os.ReadFile(filename)
	}

	for _, path := range includePaths {
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

func (r *reader) processDefine() error {
	name, err := r.nextToken()
	if err != nil {
		return err
	}

	next, err := r.nextToken()
	if err != nil {
		return err
	}

	var args []string

	if next.Name == "LPAREN" {
		var finished bool
		for {
			tk, err := r.nextToken()
			if err != nil {
				return err
			}

			switch tk.Name {
			case "RPAREN":
				finished = true
				break
			case "COMMA":
				// do nothing
			case "IDEN":
				args = append(args, tk.Val)
			default:
				return fmt.Errorf("%v is not a valid macro arg", tk.Val)
			}

			if finished {
				break
			}
		}
	}

	macroTokens, err := r.parseRest()
	if err != nil && err != io.EOF {
		return err
	}

	if len(args) == 0 {
		// need to add that first one back in
		macroTokens = append([]*tokens.Token{next}, macroTokens...)
	}

	r.macros[name.Val] = &Macro{
		Args:   args,
		Tokens: macroTokens,
	}

	r.directive.finished = true
	return nil
}

func (r *reader) processUndefine() error {
	name, err := r.nextToken()
	if err != nil {
		return err
	}

	delete(r.macros, name.Val)

	r.directive.finished = true
	return nil
}

func (r *reader) parseRest() ([]*tokens.Token, error) {
	var tks []*tokens.Token

	for {
		tk, err := r.tokens.Peek()
		if err != nil {
			return nil, err
		}

		if tk.Name == "BSLASH" {
			panic("impl")
		}

		if tk.Name == "NEWLINE" {
			break
		}

		tk, _ = r.tokens.Next()
		tks = append(tks, tk)
	}

	return tks, nil
}

func (r *reader) advanceMacro() (*tokens.Token, error) {
	if r.macro.tokenIdx == len(r.macro.def.Tokens) {
		r.macro = nil
		return nil, nil
	}

	if len(r.macro.swap) > 0 {
		if r.macro.swapIdx == len(r.macro.swap) {
			// finished swap
			r.macro.swap = nil
			r.macro.swapIdx = 0
		} else {
			// in a swap phase
			tk := r.macro.swap[r.macro.swapIdx]
			r.macro.swapIdx++

			return tk, nil
		}
	}

	tk := r.macro.def.Tokens[r.macro.tokenIdx]

	for _, arg := range r.macro.def.Args {
		if arg == tk.Val {
			// need to do a swap
			r.macro.swap = r.macro.params[arg]
			r.macro.swapIdx = 1 // we instantly return the first one
			r.macro.tokenIdx++

			return r.macro.swap[0], nil
		}
	}

	r.macro.tokenIdx++
	return tk, nil
}

func (r *reader) readParams(args []string) (map[string][]*tokens.Token, error) {
	var depth int
	var paramIdx int
	params := make(map[string][]*tokens.Token)

	add := func(t *tokens.Token) {
		params[args[paramIdx]] = append(params[args[paramIdx]], t)
	}

	for {
		tk, err := r.nextToken()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("unexpected EOF during macro expansion")
			}

			return nil, err
		}

		switch tk.Name {
		case "LPAREN":
			depth++
			if depth != 1 {
				add(tk)
			}
		case "RPAREN":
			depth--
			if depth != 0 {
				add(tk)
			}
		case "COMMA":
			if depth == 1 {
				paramIdx++
			} else {
				add(tk)
			}
		default:
			add(tk)
		}

		if depth == 0 {
			break
		}
	}

	if paramIdx != len(args)-1 {
		return nil, fmt.Errorf("expected %d arguments and got %d", len(args), paramIdx-1)
	}

	return params, nil
}
