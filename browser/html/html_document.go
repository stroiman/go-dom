package html

import (
	. "github.com/stroiman/go-dom/browser/dom"
)

type HTMLDocument interface {
	Document
	// unexported
	getWindow() Window
}

type htmlDocument struct {
	Document
	window Window
}

func NewHTMLDocument(window Window) HTMLDocument {
	result := &htmlDocument{NewDocument(window), window}
	result.SetSelf(result)
	return result
}

func (d *htmlDocument) CreateElement(name string) Element {
	switch name {
	case "template":
		return NewHTMLTemplateElement(d)
	case "a":
		return NewHTMLAnchorElement(d)
	}
	return NewHTMLElement(name, d)
}

func (d *htmlDocument) getWindow() Window { return d.window }
