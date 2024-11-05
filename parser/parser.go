package parser

import (
	"fmt"

	dom "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/interfaces"
	"github.com/stroiman/go-dom/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func Parse(tokens []lexer.Token) interfaces.Node {
	p := createParser(tokens)

	e := parseElement(p, nil)
	expect(p, lexer.EOF)

	return e
}

func parseElement(p *parser, stack []string) interfaces.Element {
	expect(p, lexer.TAG_START)
	token := expect(p, lexer.IDENTIFIER)
	expect(p, lexer.TAG_END)
	expect(p, lexer.TAG_CLOSE_START)
	expect(p, lexer.IDENTIFIER)
	expect(p, lexer.TAG_END)
	return dom.NewHTMLElement(token.Data)
}

func expect(p *parser, kind lexer.TokenKind) lexer.Token {
	token := p.advance()
	if token.Kind != kind {
		panic(fmt.Sprintf("Unexpected token. Expected %s, got %s", kind, token))
	}
	return token
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens,
		0,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}
