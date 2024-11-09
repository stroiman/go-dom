package dom_types

type Element interface {
	Node
	// Append(Element) Element
	Children() []Element
	IsConnected() bool
	TagName() string
}

type element struct {
	node
	tagName     string
	isConnected bool
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the root
	// node.
}

func NewElement(tagName string) Element { return &element{node{}, tagName, false} }

func (e *element) Children() []Element {
	// TODO: Is encapsulated in an HTMLCollection
	panic("TODO")
}

func (e *element) NodeName() string {
	return e.TagName()
}

func (e *element) TagName() string {
	return e.tagName
}

func (e *element) IsConnected() bool { return e.isConnected }

// func (e *Element) Append(child interfaces.Element) interfaces.Element { return child }
