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

// NewHTMLDocument creates an HTML document with no content. The response is
// equivalent to parsing an empty string using a DOMParser in client script.
//
// The resulting document has the outer HTML
//
//	<html><head></head><body></body></html>
func NewHTMLDocument(window Window) HTMLDocument {
	doc := newHTMLDocument(window)
	docEl := doc.CreateElement("html")
	docEl.AppendChild(doc.CreateElement("head"))
	docEl.AppendChild(doc.CreateElement("body"))
	doc.AppendChild(docEl)
	return doc
}

// newHTMLDocument is used internally to create an empty HTML when parsing an
// HTML input.
func newHTMLDocument(window Window) HTMLDocument {
	result := &htmlDocument{NewDocument(window), window}
	result.SetSelf(result)
	return result
}

func (d *htmlDocument) CreateElementNS(namespace string, name string) Element {
	if namespace == "http://www.w3.org/1999/xhtml" {
		return d.CreateElement(name)
	}
	return d.Document.CreateElementNS(namespace, name)
}

func (d *htmlDocument) CreateElement(name string) Element {
	switch name {
	case "template":
		return NewHTMLTemplateElement(d)
	case "form":
		return NewHtmlFormElement(d)
	case "input":
		return NewHTMLInputElement(d)
	case "button":
		return NewHTMLButtonElement(d)
	case "a":
		return NewHTMLAnchorElement(d)
	}
	return NewHTMLElement(name, d)
}

func (d *htmlDocument) getWindow() Window { return d.window }
