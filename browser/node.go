package browser

import (
	"errors"
	"fmt"
	"slices"

	"golang.org/x/net/html"
)

type Node interface {
	EventTarget
	AppendChild(node Node) Node
	ChildNodes() []Node
	Connected() bool
	InsertBefore(newNode Node, referenceNode Node) (Node, error)
	NodeName() string
	Parent() Node
	RemoveChild(node Node) error
	// unexported
	createHtmlNode() *html.Node
	insertBefore(newNode Node, referenceNode Node) (Node, error)
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
	n.InsertBefore(child, nil)
	return child
}

func (n NodeHelper) InsertBefore(newChild Node, referenceNode Node) (Node, error) {
	if fragment, ok := newChild.(DocumentFragment); ok {
		for len(fragment.ChildNodes()) > 0 {
			n.InsertBefore(fragment.ChildNodes()[0], referenceNode)
		}
		return fragment, nil
	}
	result, err := n.insertBefore(newChild, referenceNode)
	if err == nil {
		newChild.setParent(n.Node)
	}
	return result, err
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

// removeNodeFromParent removes the node from the current parent, _if_ it has
// one. Does nothing for disconnected nodes.
func removeNodeFromParent(node Node) {
	parent := node.Parent()
	if parent != nil {
		parent.RemoveChild(node)
	}
}

func (n *node) RemoveChild(node Node) error {
	idx := slices.Index(n.childNodes, node)
	if idx == -1 {
		return errors.New("Not found")
	}
	n.childNodes = slices.Delete(n.childNodes, idx, idx+1)
	return nil
}

func (n *node) insertBefore(newNode Node, referenceNode Node) (Node, error) {
	// TODO, Don't allow newNode to be inserted in it's own branch (circular tree)
	// TODO, Handle a fragment. Also returns nil
	if referenceNode == nil {
		n.childNodes = append(n.childNodes, newNode)
	} else {
		i := slices.Index(n.childNodes, referenceNode)
		if i == -1 {
			return nil, errors.New("Reference node not found")
		}
		n.childNodes = slices.Insert(n.childNodes, i, newNode)
	}
	removeNodeFromParent(newNode)
	return newNode, nil
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
