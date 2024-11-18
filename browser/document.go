package browser

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type DocumentEvent = string

const (
	DocumentEventDOMContentLoaded DocumentEvent = "DOMContentLoaded"
	DocumentEventLoad             DocumentEvent = "load"
)

type Document interface {
	Node
	Body() Element
	CreateElement(string) Element
	createElement(*html.Node) Element
	DocumentElement() Element
	Append(Element) Element
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
}

func newDocument(node *html.Node) Document {
	return &document{
		node: newNode(node),
	}
}

func NewDocument() Document {
	return &document{
		node: newNode(&html.Node{
			Type: html.DocumentNode,
		}),
	}
}

func (d *document) Body() Element {
	root := d.DocumentElement()
	if root != nil {
		for _, child := range root.ChildNodes() {
			if e, ok := child.(Element); ok {
				if e.TagName() == "BODY" {
					return e
				}
			}
		}
	}
	return nil
}

func (d *document) CreateElement(name string) Element {
	node := &html.Node{
		Type:      html.ElementNode,
		DataAtom:  atom.Lookup([]byte(name)),
		Data:      name,
		Namespace: "",
	}
	return d.createElement(node)
}

func (d *document) createElement(node *html.Node) Element {
	return NewHTMLElement(node)
}

func (d *document) Append(element Element) Element {
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
