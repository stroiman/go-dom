package dom

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type DocumentEvent = string

type StaticNodeList NodeList

const (
	DocumentEventDOMContentLoaded DocumentEvent = "DOMContentLoaded"
	DocumentEventLoad             DocumentEvent = "load"
)

type Document interface {
	RootNode
	Body() Element
	Head() Element
	CreateDocumentFragment() DocumentFragment
	CreateElement(string) Element
	DocumentElement() Element
	Location() Location
	// unexported
	createElement(*html.Node) Element
}
type elementConstructor func(doc *document) Element

type document struct {
	rootNode
	ownerWindow Window
}

func NewDocument(window Window) Document {
	result := &document{newRootNode(), window}
	// Hmmm, can document be replaced; and now the old doc's event goes to a
	// window they shouldn't?
	// What about disconnected documents, e.g. `new Document()` in the browser?
	result.parentTarget = window
	return result
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
	if node.Data == "template" {
		return NewHTMLTemplateElement(node, d)
	}
	return NewHTMLElement(node, d)
}

func (d *document) CreateDocumentFragment() DocumentFragment {
	return NewDocumentFragment(d)
}

func (d *document) Append(element Element) Element {
	NodeHelper{d}.AppendChild(element)
	return element
}

func (d *document) AppendChild(newChild Node) Node {
	return NodeHelper{d}.AppendChild(newChild)
}

func (d *document) InsertBefore(newChild Node, reference Node) (Node, error) {
	return NodeHelper{d}.InsertBefore(newChild, reference)
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

func (d *document) Connected() bool {
	return true
}

func (d *document) GetElementById(id string) Element {
	return RootNodeHelper{d}.GetElementById(id)
}

func (d *document) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func (d *document) Location() Location {
	return d.ownerWindow.Location()
}

func (d *document) QuerySelector(pattern string) (Node, error) {
	return CSSHelper{d}.QuerySelector(pattern)
}

func (d *document) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	return CSSHelper{d}.QuerySelectorAll(pattern)
}

func (d *document) OwnerDocument() Document { return d }

func (d *document) NodeType() NodeType { return NodeTypeDocument }
