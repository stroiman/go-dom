package browser

// ElementContainer defines common functionality in [Document],
// [DocumentFragment], and [Element]. While they all have [Node] as the direct
// base class in the DOM spec; they share a common set of functions operating on
// elements
type ElementContainer interface {
	Node
	Append(Element) Element
	QuerySelector(string) (Node, error)
	QuerySelectorAll(string) (StaticNodeList, error)
}

// RootNode implements defines common behaviour between [Document] and
// [DocumentFragment]. While they both have [Node] as the direct base class in
// the DOM spec; they share a common set of functions operating on elements.
type RootNode interface {
	ElementContainer
	GetElementById(string) Element
}

type rootNode struct {
	node
}

func newRootNode() rootNode {
	return rootNode{newNode()}
}
