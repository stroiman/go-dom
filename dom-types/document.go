package dom_types

import (
	"strings"

	"github.com/stroiman/go-dom/interfaces"
)

type elementConstructor func(doc *Document) interfaces.Element

var defaultElements map[string]elementConstructor = map[string]elementConstructor{
	"html": func(doc *Document) interfaces.Element { return NewHTMLHtmlElement(doc) },
	"body": func(doc *Document) interfaces.Element { return NewHTMLElement("body") },
	"head": func(doc *Document) interfaces.Element { return NewHTMLElement("head") },
}

type Document struct {
	// For HTML, it's an HTML element, for XML, it's an XML document.
	// While HTMLDocument doesn't exist as a separate type; it's an alias for
	// Document, XMLDocument inherits from Document; whish is why we can't be more
	// explicit in the type
	//
	// ... unless internally there are two implementations of the interface.
	documentElement interfaces.Element
	body            interfaces.Element
}

// I don't know if it's better to return the 'interface', or the type, but
// returning the interface, the compiler helps identify bad implementations
// early.
func NewDocument() interfaces.Document {
	return &Document{}
}

func (d *Document) Body() interfaces.Element {
	return d.body
}

func (d *Document) CreateElement(name string) interfaces.Element {
	constructor, ok := defaultElements[strings.ToLower(name)]
	if ok {
		return constructor(d)
	}
	return NewHTMLUnknownElement(name)
}

func (d *Document) SetBody(body interfaces.Element) {
	d.body = body
}

func (d *Document) Append(element interfaces.Element) interfaces.Element {
	d.documentElement = element
	return element
}

func (d *Document) DocumentElement() interfaces.Element {
	return d.documentElement
}

func (d Document) NodeName() string { return "#document" }
