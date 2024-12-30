package html

import (
	. "github.com/stroiman/go-dom/browser/dom"
)

type HTMLDocument interface {
	Document
}

type htmlDocument struct {
	Document
}

func NewHTMLDocument(window Window) HTMLDocument {
	result := &htmlDocument{NewDocument(window)}
	result.SetSelf(result)
	return result
}
