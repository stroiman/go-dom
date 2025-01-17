package dom

// ElementContainer defines common functionality in [Document],
// [DocumentFragment], and [Element]. While they all have [Node] as the direct
// base class in the DOM spec; they share a common set of functions operating on
// elements
type ElementContainer interface {
	Node
	Append(Element) (Element, error)
	QuerySelector(string) (Element, error)
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

type RootNodeHelper struct{ RootNode }

func (d RootNodeHelper) GetElementById(id string) Element {
	var search func(node Node) Element
	search = func(node Node) Element {
		if elm, ok := node.(Element); ok {
			if a, _ := elm.GetAttribute("id"); a == id {
				return elm
			}
		}
		for _, child := range node.ChildNodes().All() {
			if found := search(child); found != nil {
				return found
			}
		}
		return nil
	}
	return search(d)
}
