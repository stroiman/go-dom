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
		// fmt.Printf("Consume token: %s\n", nextToken)
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
	doc *dom.Document
	eof bool
}

func Parse(tokens <-chan lexer.Token) interfaces.Document {
	p := createParser(tokens)
	parseElement(p, nil)
	if !p.eof {
		panic("Didn't parse to EOF")
	}
	return p.doc
}

func parseElement(p *parser, parent *dom.Element) {
	token := expect(p, lexer.TAG_OPEN_BEGIN)
	expect(p, lexer.TAG_END)
	e := p.doc.CreateElement(token.Data)
	if parent == nil {
		p.doc.Append(e)
	} else {
		parent.AppendChild(e)
	}
	for p.w.Kind == lexer.TAG_OPEN_BEGIN {
		parseElement(p, e)
	}
	expect(p, lexer.TAG_CLOSE_BEGIN)
	expect(p, lexer.TAG_END)
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
			dom.NewDocument(),
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
