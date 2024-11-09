package dom_types

import (
	"strings"
)

type EventTarget interface{}

type Document interface {
	Node
	Body() Element
	CreateElement(string) Element
	DocumentElement() Element
	Append(Element) Element
	SetBody(e Element)
}
type elementConstructor func(doc *document) Element

var defaultElements map[string]elementConstructor = map[string]elementConstructor{
	"html": func(doc *document) Element { return NewHTMLHtmlElement(doc) },
	"body": func(doc *document) Element { return NewHTMLElement("body") },
	"head": func(doc *document) Element { return NewHTMLElement("head") },
}

type document struct {
	node
	// For HTML, it's an HTML element, for XML, it's an XML document.
	// While HTMLDocument doesn't exist as a separate type; it's an alias for
	// Document, XMLDocument inherits from Document; whish is why we can't be more
	// explicit in the type
	//
	// ... unless internally there are two implementations of the interface.
	documentElement Element
	body            Element
}

func NewDocument() Document {
	return &document{}
}

func (d *document) Body() Element {
	return d.body
}

func (d *document) CreateElement(name string) Element {
	constructor, ok := defaultElements[strings.ToLower(name)]
	if ok {
		return constructor(d)
	}
	return NewHTMLUnknownElement(name)
}

func (d *document) SetBody(body Element) {
	d.body = body
}

func (d *document) Append(element Element) Element {
	d.documentElement = element
	return element
}

func (d *document) DocumentElement() Element {
	return d.documentElement
}

func (d *document) NodeName() string { return "#document" }
