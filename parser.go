package go_dom

import (
	// . "github.com/stroiman/go-dom/dom-types"
	"io"

	"github.com/stroiman/go-dom/interfaces"
	"github.com/stroiman/go-dom/lexer"
	"github.com/stroiman/go-dom/parser"
)

func Parse(s io.Reader) interfaces.Node {
	tokens := lexer.Tokenize(s)
	return parser.Parse(tokens)
}
