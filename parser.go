package go_dom

import (
	// . "github.com/stroiman/go-dom/dom-types"
	"io"

	dom "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/lexer"
	"github.com/stroiman/go-dom/parser"
)

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

func Parse(s io.Reader) dom.Document {
	return parser.Parse(lexer.TokenizeStream(s))
}
