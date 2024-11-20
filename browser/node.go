package browser

import (
	"golang.org/x/net/html"
)

type Node interface {
	EventTarget
	AppendChild(node Node) Node
	ChildNodes() []Node
	Connected() bool
	NodeName() string
	Parent() Node
	// unexported
	createHtmlNode() *html.Node
	// toHtmlNode(Node, map[*html.Node]Node) *html.Node
	setParent(node Node)
}

type node struct {
	eventTarget
	childNodes []Node
	parent     Node
}

func newNode() node {
	return node{newEventTarget(), []Node{}, nil}
}

func (parent *node) AppendChild(child Node) Node {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *node) ChildNodes() []Node { return n.childNodes }

func (n *node) Parent() Node { return n.parent }

func (n *node) setParent(parent Node) { n.parent = parent }

func (n *node) Connected() (result bool) {
	if n.parent != nil {
		result = n.parent.Connected()
	}
	return
}

func (n *node) NodeName() string {
	return "#node"
}

type NodeIterator struct{ Node }

func (n NodeIterator) toHtmlNode(m map[*html.Node]Node) *html.Node {
	htmlNode := n.Node.createHtmlNode()
	if m != nil {
		m[htmlNode] = n.Node
	}
	for _, child := range n.ChildNodes() {
		htmlNode.AppendChild(NodeIterator{child}.toHtmlNode(m))
	}
	return htmlNode
}
