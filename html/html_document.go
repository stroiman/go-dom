package html

import (
	"strings"

	. "github.com/gost-dom/browser/dom"
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

func mustAppendChild(p, c Node) Node {
	_, err := p.AppendChild(c)
	if err != nil {
		panic(err)
	}
	return c
}

// NewHTMLDocument creates an HTML document for an about:blank page.
//
// The resulting document has an outer HTML similar to this, but there are no
// guarantees about the actual content, so do not depend on this value.
//
//	<html><head></head><body><h1>Gost-DOM</h1></body></html>
func NewHTMLDocument(window Window) HTMLDocument {
	doc := newHTMLDocument(window)
	body := doc.CreateElement("body")
	docEl := doc.CreateElement("html")
	h1 := mustAppendChild(body, doc.CreateElement("h1"))
	h1.SetTextContent("Gost-DOM")
	docEl.Append(
		doc.CreateElement("head"),
		body,
	)
	doc.AppendChild(docEl)
	return doc
}

// newHTMLDocument is used internally to create an empty HTML when parsing an
// HTML input.
func newHTMLDocument(window Window) HTMLDocument {
	var result HTMLDocument = &htmlDocument{NewDocument(window), window}
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
	switch strings.ToLower(name) {
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

func (d *htmlDocument) OwnerDocument() Document { return d }
