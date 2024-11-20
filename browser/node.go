package browser

import (
	"errors"
	"fmt"
	"slices"

	"golang.org/x/net/html"
)

type Node interface {
	EventTarget
	appendChild(node Node) Node
	ChildNodes() []Node
	Connected() bool
	InsertBefore(newNode Node, referenceNode Node) error
	NodeName() string
	Parent() Node
	// unexported
	createHtmlNode() *html.Node
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

type NodeHelper struct{ Node }

func (n NodeHelper) AppendChild(child Node) Node {
	n.appendChild(child)
	child.setParent(n.Node)
	return child
}

func (parent *node) appendChild(child Node) Node {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *node) ChildNodes() []Node { return n.childNodes }

func (n *node) Parent() Node { return n.parent }

func (n *node) setParent(parent Node) { n.parent = parent }

func (n *node) Connected() (result bool) {
	fmt.Printf("Check parent on node, %v\n", n)
	if n.parent != nil {
		result = n.parent.Connected()
	}
	return
}

func (n *node) NodeName() string {
	return "#node"
}

func (n *node) InsertBefore(newNode Node, referenceNode Node) error {
	i := slices.Index(n.childNodes, referenceNode)
	if i == -1 {
		return errors.New("Reference node not found")
	}

	n.childNodes = slices.Insert(n.childNodes, i, newNode)
	newNode.setParent(referenceNode.Parent())
	return nil
}

type NodeIterator struct{ Node }

func toHtmlNode(node Node) *html.Node {
	return NodeIterator{node}.toHtmlNode(nil)
}
func toHtmlNodeAndMap(node Node) (*html.Node, map[*html.Node]Node) {
	m := make(map[*html.Node]Node)
	result := NodeIterator{node}.toHtmlNode(m)
	return result, m
}

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
