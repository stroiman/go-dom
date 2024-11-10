package browser

import (
	"fmt"
	"io"
	"strings"

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
	doc Document
	eof bool
}

func Parse(tokens <-chan lexer.Token) Document {
	p := createParser(tokens)
	parseDocument(p, nil)
	if !p.eof {
		panic("Didn't parse to EOF")
	}
	return p.doc
}

func parseDocument(p *parser, parent Element) {
	// TODO: Handle processing instructions
	if p.w.Kind == lexer.TAG_OPEN_BEGIN && p.w.Data == "html" {
		parseElementCallback(p, parent, parseHtmlChildren)
	} else {
		html := p.doc.Append(p.doc.CreateElement("html"))
		parseHtmlChildren(p, html)
	}
}

func parseHtmlChildren(p *parser, parent Element) {
	if p.w.Kind != lexer.TAG_OPEN_BEGIN || p.w.Data != "head" {
		parent.AppendChild(p.doc.CreateElement("head"))
	} else {
		parseElement(p, parent)
	}
	if p.w.Kind != lexer.TAG_OPEN_BEGIN || p.w.Data != "body" {
		parent = parent.Append(p.doc.CreateElement("body"))
	}
	parseElementChildren(p, parent)
}

func parseElementCallback(
	p *parser,
	parent Element,
	callback func(p *parser, parent Element),
) {
	token := expect(p, lexer.TAG_OPEN_BEGIN)
	expect(p, lexer.TAG_END)
	e := p.doc.CreateElement(token.Data)
	if parent == nil {
		p.doc.Append(e)
	} else {
		parent.AppendChild(e)
	}
	callback(p, e)
	expect(p, lexer.TAG_CLOSE_BEGIN)
	expect(p, lexer.TAG_END)
}

func parseElement(p *parser, parent Element) {
	parseElementCallback(p, parent, parseElementChildren)
}

func parseElementChildren(p *parser, parent Element) {
	for (!p.eof) && p.w.Kind == lexer.TAG_OPEN_BEGIN {
		parseElement(p, parent)
	}
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
			NewDocument(),
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

func streamOfTokens(input []lexer.Token) <-chan lexer.Token {
	resp := make(chan lexer.Token)
	go func() {
		defer close(resp)
		for _, elm := range input {
			resp <- elm
		}
	}()
	return resp
}

func ParseHtmlStream(s io.Reader) Document {
	return Parse(lexer.TokenizeStream(s))
}

func ParseHtmlString(s string) Document {
	return ParseHtmlStream(strings.NewReader(s))
}