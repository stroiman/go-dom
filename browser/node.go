package browser

import "golang.org/x/net/html"

type Node interface {
	EventTarget
	NodeName() string
	AppendChild(node Node) Node
	ChildNodes() []Node
	Parent() Node
	Connected() bool
	setParent(node Node)
	wrappedNode() *html.Node
	// TODO: Remove
	setWrappedNode(*html.Node)
}

type node struct {
	childNodes []Node
	name       string
	htmlNode   *html.Node
	parent     Node
}

func newNode(htmlNode *html.Node) node {
	return node{[]Node{}, htmlNode.Data, htmlNode, nil}
}

func (parent *node) AppendChild(child Node) Node {
	parent.htmlNode.AppendChild(child.wrappedNode())
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

func (n *node) wrappedNode() *html.Node {
	return n.htmlNode
}
func (n *node) setWrappedNode(node *html.Node) {
	n.htmlNode = node
}

func (n *node) NodeName() string {
	return "#node"
}
