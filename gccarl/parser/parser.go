package parser

import (
	"fmt"
	"io"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type RuleKey struct {
	Rule   grammar.RuleName
	Option grammar.OptionName
}

func (k RuleKey) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%v:%v"`, k.Rule, k.Option)), nil
}

type Index struct {
	RuleKey RuleKey
	Index   int
}

type Value struct {
	Node  *Node         `json:"node,omitempty"`
	Token *tokens.Token `json:"token,omitempty"`
}

type Parser struct {
	grammar map[grammar.RuleName]*grammar.Rule
	debug   bool

	cursors []*Cursor
}

func New(r io.Reader, debugParam bool) (*Parser, error) {
	debug = debugParam

	gr, err := grammar.Parse(r)
	if err != nil {
		return nil, err
	}

	return &Parser{
		grammar: gr,
		debug:   debug,
	}, nil
}

func (p *Parser) Parse(r *tokens.Reader, outDir, packageName string) error {
	top := &Node{
		Key: RuleKey{
			Rule:   "main",
			Option: "main",
		},
	}

	p.cursors = []*Cursor{
		{
			grammar: p.grammar,
			Stack: []Index{
				{
					RuleKey: RuleKey{
						Rule:   "main",
						Option: "main",
					},
				},
			},
			Current: top,
			Top:     top,
		},
	}

	for {
		token, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		if debug {
			fmt.Printf("token: %v\n", token)
		}

		if len(p.cursors) == 0 {
			return fmt.Errorf("expected at least one cursor")
		}

		p.handleToken(token)
	}

	if len(p.cursors) != 1 {
		return fmt.Errorf("expected 1 cursor but got %d", len(p.cursors))
	}

	if !p.cursors[0].terminalState() {
		return fmt.Errorf("last cursor did not terminate: %v", p.cursors[0])
	}

	g := &Generator{
		grammar:     p.grammar,
		packageName: packageName,
	}

	return g.generate(p.cursors[0].Top, outDir)
}

func debugf(format string, args ...any) {
	if debug {
		fmt.Printf(format+"\n", args...)
	}
}

func (p *Parser) handleToken(token *tokens.Token) {
	cursors := p.advance(p.cursors)

	for _, c := range cursors {
		debugf("handling %v with %v", token.Name, c)
	}

	var nextCursors []*Cursor
	for _, c := range cursors {
		ok := c.Apply(token)
		if ok {
			debugf("applied %v to %v", token.Name, c)
			nextCursors = append(nextCursors, c)
		}
	}

	p.cursors = nextCursors
}

func (p *Parser) advance(cursors []*Cursor) []*Cursor {
	needBranch := cursors

	var ready []*Cursor
	var nextNeedBranch []*Cursor

	var i int
	for {
		for _, c := range needBranch {
			innerReady, innerNeedBrach := c.Branch()

			ready = append(ready, innerReady...)
			nextNeedBranch = append(nextNeedBranch, innerNeedBrach...)
		}

		needBranch = nextNeedBranch
		nextNeedBranch = nil

		if len(needBranch) == 0 {
			break
		}

		const maxBranches = 100

		i++
		if i > maxBranches {
			panic("too many branches")
		}
	}

	return ready
}

func (p *Parser) getParts(k RuleKey) []*grammar.Part {
	for _, o := range p.grammar[k.Rule].Options {
		if o.Name == k.Option {
			return o.Parts
		}
	}

	panic("invalid rule")
}
