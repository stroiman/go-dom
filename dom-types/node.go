package dom_types

type Node interface {
	EventTarget
	NodeName() string
	AppendChild(node Node) Node
	ChildNodes() []Node
}

type node struct {
	// Not correct, but currently we have only elements
	childNodes []Node
	name       string
}

func (parent *node) AppendChild(child Node) Node {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *node) ChildNodes() []Node { return n.childNodes }
