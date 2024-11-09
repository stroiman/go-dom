package dom_types

import (
	"strings"

	"github.com/stroiman/go-dom/interfaces"
)

type elementConstructor func(doc *Document) *Element

var defaultElements map[string]elementConstructor = map[string]elementConstructor{
	"html": func(doc *Document) *Element { return NewHTMLHtmlElement(doc) },
	"body": func(doc *Document) *Element { return NewHTMLElement("body") },
	"head": func(doc *Document) *Element { return NewHTMLElement("head") },
}

type Document struct {
	// For HTML, it's an HTML element, for XML, it's an XML document.
	// While HTMLDocument doesn't exist as a separate type; it's an alias for
	// Document, XMLDocument inherits from Document; whish is why we can't be more
	// explicit in the type
	//
	// ... unless internally there are two implementations of the interface.
	documentElement *Element
	body            interfaces.Element
}

func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Body() interfaces.Element {
	return d.body
}

func (d *Document) CreateElement(name string) *Element {
	constructor, ok := defaultElements[strings.ToLower(name)]
	if ok {
		return constructor(d)
	}
	return NewHTMLUnknownElement(name)
}

func (d *Document) SetBody(body interfaces.Element) {
	d.body = body
}

func (d *Document) Append(element *Element) *Element {
	d.documentElement = element
	return element
}

func (d *Document) DocumentElement() *Element {
	return d.documentElement
}

func (d Document) NodeName() string { return "#document" }
