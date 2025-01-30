package dom

import "github.com/gost-dom/browser/internal/entity"

type NodeList interface {
	entity.Entity
	Length() int
	Item(index int) Node
	All() []Node
	setNodes([]Node)
	append(Node)
}

type nodeList struct {
	entity.Entity
	nodes []Node
}

type staticNodeSource []Node

func (s staticNodeSource) ChildNodes() []Node { return s }

func newNodeList(nodes ...Node) NodeList {
	return &nodeList{entity.New(), nodes}
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
