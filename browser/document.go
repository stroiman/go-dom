package browser

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type DocumentEvent = string

type StaticNodeList []Node

const (
	DocumentEventDOMContentLoaded DocumentEvent = "DOMContentLoaded"
	DocumentEventLoad             DocumentEvent = "load"
)

type Document interface {
	RootNode
	Body() Element
	CreateElement(string) Element
	DocumentElement() Element
	// unexported
	createElement(*html.Node) Element
}
type elementConstructor func(doc *document) Element

type document struct {
	rootNode
}

func NewDocument() Document {
	return &document{newRootNode()}
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
	return NewHTMLElement(node, d)
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
	for _, c := range d.ChildNodes() {
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

func (d *document) QuerySelector(pattern string) (Node, error) {
	return CSSHelper{d}.QuerySelector(pattern)
}

func (d *document) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	return CSSHelper{d}.QuerySelectorAll(pattern)
}
