package go_dom

import (
	"io"

	. "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/interfaces"
)

func Parse(r io.Reader) interfaces.Node {
	return NewHTMLHtmlElement()
}
