package browser

import (
	"golang.org/x/net/html"
)

type ObjectId = uintptr

var idSeq <-chan ObjectId

func init() {
	c := make(chan ObjectId)
	idSeq = c
	go func() {
		var val ObjectId = 1
		for {
			c <- val
			val = val + 1
		}
	}()
}

func NewObjectId() ObjectId {
	return <-idSeq
}

type Node interface {
	EventTarget
	// ObjectId is used internally to use nodes as keys in a map without keeping
	// the objects reachable.
	ObjectId() ObjectId
	NodeName() string
	AppendChild(node Node) Node
	ChildNodes() []Node
	Parent() Node
	Connected() bool
	setParent(node Node)
	wrappedNode() *html.Node
}

type node struct {
	objectId   ObjectId
	childNodes []Node
	name       string
	htmlNode   *html.Node
	parent     Node
}

func newNode(htmlNode *html.Node) node {
	id := NewObjectId()
	return node{id, []Node{}, htmlNode.Data, htmlNode, nil}
}

func (n *node) ObjectId() ObjectId { return n.objectId }

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

func (n *node) NodeName() string {
	return "#node"
}
