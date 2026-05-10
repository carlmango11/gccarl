package parser

import (
	"compiler/parser/grammar"
	"fmt"
	"io"
	"strings"
)

type Node struct {
	Name   string
	Values []*Value
}

type Identifier string

type Value struct {
	Node       *Node
	Literal    string
	Identifier Identifier
	Number     float64
}

type Parser struct {
	grammar   map[grammar.RuleName]*grammar.Rule
	err       error
	debug     bool
	nodeDepth int
}

func New(r io.Reader, debug bool) (*Parser, error) {
	gr, err := grammar.Parse(r)
	if err != nil {
		return nil, err
	}

	return &Parser{
		grammar: gr,
		debug:   debug,
	}, nil
}

func (p *Parser) Parse(r io.Reader) (*Node, error) {
	text, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	sc := &Scanner{
		text: strings.TrimSpace(string(text)),
	}

	n, err := p.parseNode(sc, "main")
	if err != nil {
		return nil, err
	}

	if !sc.Finish() && p.err != nil && p.err != io.EOF {
		return nil, fmt.Errorf("parser failed to finish: %v", p.err)
	}

	return n, nil
}

func (p *Parser) parseNode(sc *Scanner, ruleName grammar.RuleName) (*Node, error) {
	p.nodeDepth++
	defer func() {
		p.nodeDepth--
	}()

	p.fullDebugf(sc, "parsing rule: %s", ruleName)

	rule, ok := p.grammar[ruleName]
	if !ok {
		return nil, fmt.Errorf("rule %q not found", ruleName)
	}

	originalIndex := sc.Index()

	var err error
	for _, o := range rule.Options {
		p.fullDebugf(sc, "parsing option: %s", o.Name)

		var vals []*Value

		vals, err = p.parseTokens(sc, o.Tokens)
		if err != nil {
			sc.Reset(originalIndex)
			continue
		}

		p.err = nil
		return &Node{
			Name:   string(o.Name),
			Values: vals,
		}, nil
	}

	return nil, err
}

func (p *Parser) parseTokens(sc *Scanner, tokens []*grammar.Token) ([]*Value, error) {
	p.fullDebugf(sc, "parsing tokens: %v", tokens)

	var vals []*Value

	for _, token := range tokens {
		v, err := p.parseToken(sc, token)
		if err != nil {
			return nil, err
		}

		vals = append(vals, v...)
	}

	return vals, nil
}

func (p *Parser) parseToken(sc *Scanner, token *grammar.Token) ([]*Value, error) {
	if !token.Multi {
		val, err := p.parseSingleToken(sc, token)
		if err != nil {
			return nil, err
		}

		return []*Value{val}, nil
	}

	p.fullDebugf(sc, "parsing multi token: %v", token)

	var vals []*Value
	for {
		val, err := p.parseSingleToken(sc, token)
		if err != nil {
			p.err = err
			return vals, nil
		}

		vals = append(vals, val)
	}
}

func (p *Parser) parseSingleToken(sc *Scanner, token *grammar.Token) (*Value, error) {
	p.fullDebugf(sc, "parsing single token: %v", token)

	switch token.Type {
	case grammar.TTRule:
		return p.parseNodeVal(sc, token.Rule)
	case grammar.TTLiteral:
		return p.parseLiteralVal(sc, token.Literal)
	case grammar.TTIdentifier:
		return p.parseIdentifier(sc)
	case grammar.TTNumber:
		return p.parseNumber(sc)
	}

	panic("invalid token type")
}

func (p *Parser) parseNodeVal(sc *Scanner, ruleName grammar.RuleName) (*Value, error) {
	node, err := p.parseNode(sc, ruleName)
	if err != nil {
		return nil, err
	}

	return &Value{
		Node: node,
	}, nil
}

func (p *Parser) parseLiteralVal(sc *Scanner, literal string) (*Value, error) {
	p.fullDebugf(sc, "parsing lit: %q", literal)

	ok, err := sc.ReadLiteral(literal)
	if err != nil {
		return nil, err
	}

	if !ok {
		p.debugf("literal %q doesn't exist", literal)
		return nil, fmt.Errorf("literal %s not found", literal)
	}

	p.fullDebugf(sc, "parsed lit: %q", literal)

	return &Value{
		Literal: literal,
	}, nil
}

func (p *Parser) parseIdentifier(sc *Scanner) (*Value, error) {
	p.fullDebugf(sc, "parsing id")

	iden, err := sc.ParseIdentifier()
	if err != nil {
		return nil, err
	}

	p.fullDebugf(sc, "parsed id: %v", iden)

	return &Value{
		Identifier: iden,
	}, nil
}

func (p *Parser) parseNumber(sc *Scanner) (*Value, error) {
	num, err := sc.ParseNumber()
	if err != nil {
		return nil, err
	}

	return &Value{
		Number: num,
	}, nil
}

func (p *Parser) debugf(format string, args ...any) {
	if p.debug {
		fmt.Printf(strings.Repeat(" ", p.nodeDepth*2)+format+"\n", args...)
	}
}

func (p *Parser) fullDebugf(sc *Scanner, format string, args ...any) {
	p.debugf(format, args...)
	p.debugf("Sc: %s", sc)
}
