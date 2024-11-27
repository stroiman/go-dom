package browser

import (
	"errors"
	"slices"

	"golang.org/x/net/html"
)

type NodeType int

const (
	NodeTypeElement               NodeType = 1
	NodeTypeAttribute             NodeType = 2
	NodeTypeText                  NodeType = 3
	NodeTypeCDataSection          NodeType = 4
	NodeTypeProcessingInstruction NodeType = 7
	NodeTypeComment               NodeType = 8
	NodeTypeDocument              NodeType = 9
	NodeTypeDocumentType          NodeType = 10
	NodeTypeDocumentFragment      NodeType = 11
)

type Node interface {
	EventTarget
	AppendChild(node Node) Node
	ChildNodes() NodeList
	Connected() bool
	InsertBefore(newNode Node, referenceNode Node) (Node, error)
	NodeName() string
	NodeType() NodeType
	OwnerDocument() Document
	Parent() Node
	RemoveChild(node Node) error
	NextSibling() Node
	// unexported
	createHtmlNode() *html.Node
	insertBefore(newNode Node, referenceNode Node) (Node, error)
	setParent(node Node)
	nodes() []Node
}

type node struct {
	eventTarget
	childNodes NodeList
	parent     Node
}

func newNode() node {
	return node{newEventTarget(), NewNodeList(), nil}
}

type NodeHelper struct{ Node }

func (n NodeHelper) AppendChild(child Node) Node {
	n.InsertBefore(child, nil)
	return child
}

func (n NodeHelper) InsertBefore(newChild Node, referenceNode Node) (Node, error) {
	if fragment, ok := newChild.(DocumentFragment); ok {
		for fragment.ChildNodes().Length() > 0 {
			n.InsertBefore(fragment.ChildNodes().Item(0), referenceNode)
		}
		return fragment, nil
	}
	result, err := n.insertBefore(newChild, referenceNode)
	if err == nil {
		newChild.setParent(n.Node)
	}
	return result, err
}

func (n *node) ChildNodes() NodeList { return n.childNodes }

func (n *node) Parent() Node { return n.parent }

func (n *node) setParent(parent Node) {
	n.parent = parent
	n.parentTarget = parent
}

func (n *node) Connected() (result bool) {
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
	idx := slices.Index(n.childNodes.All(), node)
	if idx == -1 {
		return errors.New("Not found")
	}
	n.childNodes.setNodes(slices.Delete(n.childNodes.All(), idx, idx+1))
	return nil
}

func (n *node) insertBefore(newNode Node, referenceNode Node) (Node, error) {
	// TODO, Don't allow newNode to be inserted in it's own branch (circular tree)
	// TODO, Handle a fragment. Also returns nil
	if referenceNode == nil {
		n.childNodes.append(newNode)
	} else {
		i := slices.Index(n.childNodes.All(), referenceNode)
		if i == -1 {
			return nil, errors.New("Reference node not found")
		}
		n.childNodes.setNodes(slices.Insert(n.childNodes.All(), i, newNode))
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
	for _, child := range n.nodes() {
		htmlNode.AppendChild(NodeIterator{child}.toHtmlNode(m))
	}
	return htmlNode
}

func (n *node) OwnerDocument() Document {
	parent := n.Parent()
	if parent != nil {
		return parent.OwnerDocument()
	}
	return nil
}

func (n *node) NextSibling() Node {
	children := n.Parent().nodes()
	idx := slices.IndexFunc(
		children,
		func(child Node) bool { return n.ObjectId() == child.ObjectId() },
	) + 1
	if idx == 0 {
		panic("We should exist in our parent's collection")
	}
	if idx >= len(children) {
		return nil
	}
	return children[idx]
}

func (n *node) nodes() []Node {
	return n.childNodes.All()
}
