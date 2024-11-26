package browser

type NodeList interface {
	Entity
	Length() int
	Item(index int) Node

	// Instance properties
	// length
	// Instance methods
	// entries()
	// forEach()
	// item()
	// keys()
	// values()}

	// unexported
	All() []Node
	setNodes([]Node)
	append(Node)
}

type nodeList struct {
	base
	nodes []Node
}

type StaticNodeSource []Node

func (s StaticNodeSource) ChildNodes() []Node { return s }

func NewNodeList() NodeList {
	return &nodeList{newBase(), nil}
}

func (l *nodeList) Length() int { return len(l.nodes) }

func (l *nodeList) Item(index int) Node {
	if index >= len(l.nodes) {
		return nil
	}
	return l.nodes[index]
}

func (l *nodeList) All() []Node           { return l.nodes }
func (l *nodeList) setNodes(nodes []Node) { l.nodes = nodes }
func (l *nodeList) append(node Node)      { l.nodes = append(l.nodes, node) }
