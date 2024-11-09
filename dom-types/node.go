package dom_types

type Node struct {
	// Not correct, but currently we have only elements
	childNodes []*Element
	name       string
}

func (parent *Node) AppendChild(child *Element) *Element {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *Node) ChildNodes() []*Element { return n.childNodes }
