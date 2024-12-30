package html

import "github.com/stroiman/go-dom/browser/dom"

type HTMLDocument interface {
	dom.Document
}

type htmlDocument struct {
	dom.Document
}

func NewHTMLDocument(window dom.Window) HTMLDocument {
	result := &htmlDocument{dom.NewDocument(window)}
	result.SetSelf(result)
	return result
}
