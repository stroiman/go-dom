package browser

import "golang.org/x/net/html"

type Node interface {
	EventTarget
	NodeName() string
	AppendChild(node Node) Node
	ChildNodes() []Node
}

type node struct {
	childNodes []Node
	name       string
	htmlNode   *html.Node
}

func newNode(htmlNode *html.Node) node {
	return node{[]Node{}, htmlNode.Data, htmlNode}
}

func (parent *node) AppendChild(child Node) Node {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *node) ChildNodes() []Node { return n.childNodes }
