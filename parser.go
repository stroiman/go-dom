package go_dom

import (
	// . "github.com/stroiman/go-dom/dom-types"
	"io"

	"github.com/stroiman/go-dom/interfaces"
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

func Parse(s io.Reader) interfaces.Node {
	return parser.Parse(lexer.TokenizeStream(s))
}
