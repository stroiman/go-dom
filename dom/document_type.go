package dom

import "golang.org/x/net/html"

type DocumentType interface {
	Node
	Name() string
}

type documentType struct {
	node
	name string
}

func NewDocumentType(name string) DocumentType {
	result := &documentType{newNode(), name}
	result.SetSelf(result)
	return result
}

func (t *documentType) Name() string       { return t.name }
func (t *documentType) NodeType() NodeType { return NodeTypeDocumentType }

func (t *documentType) CloneNode(deep bool) Node {
	return NewDocumentType(t.name)
}

func (t *documentType) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DoctypeNode,
		Data: t.name,
	}
}
