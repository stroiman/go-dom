package parser

import (
	"fmt"

	dom "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/interfaces"
	"github.com/stroiman/go-dom/lexer"
)

// Provides tokens in a pure lazy-loaded list such that we can start multiple
// branches of parsing that don't affect each other.
type tokenWrapper struct {
	lexer.Token
	next   *tokenWrapper
	stream <-chan lexer.Token
}

func createWrapperFromStream(stream <-chan lexer.Token) (*tokenWrapper, bool) {
	if nextToken, ok := <-stream; ok {
		fmt.Printf("Consume token: %s\n", nextToken)
		return &tokenWrapper{nextToken, nil, stream}, true
	} else {
		return nil, false
	}
}

func (w *tokenWrapper) nextToken() (*tokenWrapper, bool) {
	ok := true
	if w.next == nil {
		w.next, ok = createWrapperFromStream(w.stream)
	}
	return w.next, ok
}

type parser struct {
	w   *tokenWrapper
	eof bool
}

func Parse(tokens <-chan lexer.Token) interfaces.Node {
	p := createParser(tokens)
	e := parseElement(p, nil)
	if !p.eof {
		panic("Didn't parse to EOF")
	}
	return e
}

func parseElement(p *parser, stack []string) interfaces.Element {
	token := expect(p, lexer.TAG_OPEN_BEGIN)
	expect(p, lexer.TAG_END)
	expect(p, lexer.TAG_CLOSE_BEGIN)
	expect(p, lexer.TAG_END)
	return dom.NewHTMLElement(token.Data)
}

func expect(p *parser, kind lexer.TokenKind) lexer.Token {
	token := p.currentToken()
	if token.Kind != kind {
		panic(fmt.Sprintf("Unexpected token. Expected %s, got %s", kind, token))
	}
	p.advance()
	return token
}

func createParser(tokens <-chan lexer.Token) *parser {
	if first, ok :=
		createWrapperFromStream(tokens); ok {
		return &parser{
			first,
			false,
		}
	} else {
		panic("Stream has no data")
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.w.Token
}

func (p *parser) advance() lexer.Token {
	var ok bool
	result := p.w.Token
	p.w, ok = p.w.nextToken()
	if !ok {
		p.eof = true
	}
	return result
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) hasTokens() bool {
	return !p.eof
}
