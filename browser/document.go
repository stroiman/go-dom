package browser

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type EventTarget interface{}

type Document interface {
	Node
	Body() Element
	CreateElement(string) Element
	wrapElement(*html.Node) Element
	DocumentElement() Element
	Append(Element) Element
	SetBody(e Element)
}
type elementConstructor func(doc *document) Element

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
	node := &html.Node{
		Type:      html.ElementNode,
		DataAtom:  atom.Lookup([]byte(name)),
		Data:      name,
		Namespace: "",
	}
	return d.wrapElement(node)
}

func (d *document) wrapElement(node *html.Node) Element {
	return NewHTMLElement(node)
}

func (d *document) SetBody(body Element) {
	d.body = body
}

func (d *document) Append(element Element) Element {
	fmt.Println("Set document element", element)
	d.documentElement = element
	return element
}

func (d *document) AppendChild(node Node) Node {
	if elm, ok := node.(Element); ok {
		return d.Append(elm)
	}
	return node
}

func (d *document) DocumentElement() Element {
	return d.documentElement
}

func (d *document) NodeName() string { return "#document" }

func (d *document) Connected() bool { return true }
