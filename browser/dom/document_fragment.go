package dom

import "golang.org/x/net/html"

type DocumentFragment interface {
	RootNode
}

type documentFragment struct {
	rootNode
	ownerDocument Document
}

func NewDocumentFragment(ownerDocument Document) DocumentFragment {
	return &documentFragment{newRootNode(), ownerDocument}
}

func (d *documentFragment) Append(element Element) Element {
	NodeHelper{d}.AppendChild(element)
	return element
}

func (d *documentFragment) AppendChild(newChild Node) Node {
	return NodeHelper{d}.AppendChild(newChild)
}

func (d *documentFragment) InsertBefore(newChild Node, reference Node) (Node, error) {
	return NodeHelper{d}.InsertBefore(newChild, reference)
}

func (d *documentFragment) GetElementById(id string) Element {
	return RootNodeHelper{d}.GetElementById(id)
}

func (d *documentFragment) QuerySelector(pattern string) (Node, error) {
	return CSSHelper{d}.QuerySelector(pattern)
}

func (d *documentFragment) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	return CSSHelper{d}.QuerySelectorAll(pattern)
}

func (d *documentFragment) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func (d *documentFragment) NodeType() NodeType { return NodeTypeDocumentFragment }
