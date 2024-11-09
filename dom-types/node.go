package dom_types

type node struct {
	// Not correct, but currently we have only elements
	childNodes []*Element
	name       string
}

func (parent *node) AppendChild(child *Element) *Element {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *node) ChildNodes() []*Element { return n.childNodes }
