package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/carlmango11/gccarl/gccarl/parser/grammar"
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
	Number     int
}

type Parser struct {
	grammar map[grammar.RuleName]*grammar.Rule
	debug   bool
	nodes   []string
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

	if !sc.Finish() {
		return nil, fmt.Errorf("parser failed to finish")
	}

	return n, nil
}

func (p *Parser) parseNode(sc *Scanner, ruleName grammar.RuleName) (*Node, error) {
	p.nodes = append(p.nodes, string(ruleName))
	defer func() {
		p.nodes = p.nodes[:len(p.nodes)-1]
	}()

	p.fullDebugf(sc, "parsing rule: %s", ruleName)

	rule, ok := p.grammar[ruleName]
	if !ok {
		return nil, fmt.Errorf("rule %q not found", ruleName)
	}

	originalIndex := sc.Index()

	var err error
	var node *Node
	for _, o := range rule.Options {
		node, err = p.parseOption(sc, o)
		if err != nil {
			sc.Reset(originalIndex)
			continue
		}

		return node, nil
	}

	return nil, fmt.Errorf("cannot parse %s: %w", ruleName, err)
}

func (p *Parser) parseOption(sc *Scanner, o *grammar.Option) (*Node, error) {
	p.nodes = append(p.nodes, string(o.Name))
	defer func() {
		p.nodes = p.nodes[:len(p.nodes)-1]
	}()

	p.fullDebugf(sc, "parsing option: %s", o.Name)

	// TODO: there's a bug here. "x"* "z"? "x" wont' work because the first x* eats everything

	vals, err := p.parseTokens(sc, o.Tokens)
	if err != nil {
		return nil, err
	}

	return &Node{
		Name:   string(o.Name),
		Values: vals,
	}, nil
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
	switch token.Cardinality {
	case grammar.CardSingle:
		val, err := p.parseSingleToken(sc, token)
		if err != nil {
			return nil, err
		}

		return []*Value{val}, nil
	case grammar.CardOptional:
		p.fullDebugf(sc, "parsing optional token: %v", token)

		index := sc.Index()

		val, err := p.parseSingleToken(sc, token)
		if err != nil {
			sc.Reset(index)
			p.fullDebugf(sc, "did not find single token, reverted: %v", token)
			return nil, nil
		}

		p.fullDebugf(sc, "found single token: %v", token)
		return []*Value{val}, nil
	case grammar.CardMultiple:
		p.fullDebugf(sc, "parsing multi token: %v", token)

		index := sc.Index()

		var vals []*Value
		for {
			val, err := p.parseSingleToken(sc, token)
			if err != nil {
				sc.Reset(index)
				p.fullDebugf(sc, "did not find single token in multi, reverted: %v", token)
				return vals, nil
			}

			index = sc.Index()
			p.fullDebugf(sc, "found single token in multi: %v", token)
			vals = append(vals, val)
		}
	default:
		return nil, fmt.Errorf("unexpected token cardinality: %d", token.Cardinality)
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
		Number: int(num),
	}, nil
}

func (p *Parser) debugf(format string, args ...any) {
	if p.debug {
		fmt.Printf(strings.Join(p.nodes, "/")+": "+format+"\n", args...)
	}
}

func (p *Parser) fullDebugf(sc *Scanner, format string, args ...any) {
	p.debugf(format, args...)
	p.debugf("Sc: %s", sc)
}
