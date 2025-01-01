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
	result := &documentFragment{newRootNode(), ownerDocument}
	result.SetSelf(result)
	return result
}

func (d *documentFragment) Append(element Element) Element {
	d.AppendChild(element)
	return element
}

func (d *documentFragment) GetElementById(id string) Element {
	return RootNodeHelper{d}.GetElementById(id)
}

func (d *documentFragment) QuerySelector(pattern string) (Element, error) {
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

func (d *documentFragment) NodeName() string { return "#document-fragment" }
