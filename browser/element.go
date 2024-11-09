package browser

type Element interface {
	Node
	// Children() []Element
	IsConnected() bool
	TagName() string
	Append(Element) Element
}

type element struct {
	node
	tagName     string
	isConnected bool
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string) Element { return &element{node{}, tagName, false} }

// func (e *element) Children() []Element {
// 	panic("TODO")
// }

func (e *element) NodeName() string {
	return e.TagName()
}

func (e *element) TagName() string {
	return e.tagName
}

func (e *element) IsConnected() bool { return e.isConnected }

func (parent *element) Append(child Element) Element {
	parent.AppendChild(child)
	return child
}
