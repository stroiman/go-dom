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

func (f *documentFragment) OwnerDocument() Document {
	return f.ownerDocument
}

func (f *documentFragment) CloneNode(deep bool) Node {
	result := NewDocumentFragment(f.ownerDocument)
	if deep {
		result.Append(f.cloneChildren()...)
	}
	return result
}

func (f *documentFragment) ChildElementCount() int {
	return len(f.childElements())
}

func (f *documentFragment) Append(nodes ...Node) error {
	return f.append(nodes...)
}

func (d *documentFragment) GetElementById(id string) Element {
	return rootNodeHelper{d}.GetElementById(id)
}

func (d *documentFragment) QuerySelector(pattern string) (Element, error) {
	return cssHelper{d}.QuerySelector(pattern)
}

func (d *documentFragment) QuerySelectorAll(pattern string) (staticNodeList, error) {
	return cssHelper{d}.QuerySelectorAll(pattern)
}

func (d *documentFragment) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func (d *documentFragment) NodeType() NodeType { return NodeTypeDocumentFragment }

func (d *documentFragment) NodeName() string { return "#document-fragment" }
