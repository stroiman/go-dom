package dom

import (
	"io"

	"golang.org/x/net/html"
)

type DocumentEvent = string

type staticNodeList NodeList

const (
	DocumentEventDOMContentLoaded DocumentEvent = "DOMContentLoaded"
	DocumentEventLoad             DocumentEvent = "load"
)

type DocumentParentWindow interface {
	EventTarget
	Location() Location
	Document() Document
	ParseFragment(ownerDocument Document, reader io.Reader) (DocumentFragment, error)
}

type Document interface {
	RootNode
	Body() Element
	Head() Element
	CreateDocumentFragment() DocumentFragment
	CreateAttribute(string) Attr
	CreateElement(string) Element
	CreateElementNS(string, string) Element
	DocumentElement() Element
	Location() Location
	// unexported
	parseFragment(reader io.Reader) (DocumentFragment, error)
}

type elementConstructor func(doc *document) Element

type document struct {
	rootNode
	ownerWindow DocumentParentWindow
}

func NewDocument(window DocumentParentWindow) Document {
	result := &document{newRootNode(), window}
	// Hmmm, can document be replaced; and now the old doc's event goes to a
	// window they shouldn't?
	// What about disconnected documents, e.g. `new Document()` in the browser?
	result.parentTarget = window
	result.SetSelf(result)
	return result
}

func (d *document) CloneNode(deep bool) Node {
	result := NewDocument(d.ownerWindow)
	if deep {
		result.Append(d.cloneChildren()...)
	}
	return result
}

func (d *document) parseFragment(reader io.Reader) (DocumentFragment, error) {
	return d.ownerWindow.ParseFragment(d, reader)
}

func (d *document) ChildElementCount() int {
	return len(d.childElements())
}

func (d *document) Body() Element {
	root := d.DocumentElement()
	if root != nil {
		for _, child := range root.ChildNodes().All() {
			if e, ok := child.(Element); ok {
				if e.TagName() == "BODY" {
					return e
				}
			}
		}
	}
	return nil
}

func (d *document) Head() Element {
	root := d.DocumentElement()
	if root != nil {
		for _, child := range root.ChildNodes().All() {
			if e, ok := child.(Element); ok {
				if e.TagName() == "HEAD" {
					return e
				}
			}
		}
	}
	return nil
}

func (d *document) CreateAttribute(name string) Attr {
	return newAttr(name, "")
}

func (d *document) CreateElement(name string) Element {
	return NewElement(name, d)
}

func (d *document) CreateElementNS(_ string, name string) Element {
	return NewElement(name, d)
}

func (d *document) CreateDocumentFragment() DocumentFragment {
	return NewDocumentFragment(d)
}

func (d *document) Append(nodes ...Node) error {
	return d.append(nodes...)
}

func (d *document) DocumentElement() Element {
	for _, c := range d.ChildNodes().All() {
		if e, ok := c.(Element); ok {
			return e
		}
	}
	return nil
}

func (d *document) NodeName() string { return "#document" }

func (d *document) IsConnected() bool {
	return d.ownerWindow.Document() == d.getSelf()
}

func (d *document) GetElementById(id string) Element {
	return rootNodeHelper{d}.GetElementById(id)
}

func (d *document) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func (d *document) Location() Location {
	return d.ownerWindow.Location()
}

func (d *document) QuerySelector(pattern string) (Element, error) {
	return cssHelper{d}.QuerySelector(pattern)
}

func (d *document) QuerySelectorAll(pattern string) (staticNodeList, error) {
	return cssHelper{d}.QuerySelectorAll(pattern)
}

func (d *document) OwnerDocument() Document { return d }

func (d *document) NodeType() NodeType { return NodeTypeDocument }
