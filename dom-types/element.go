package dom_types

import "github.com/stroiman/go-dom/interfaces"

type Element struct {
	node
	tagName     string
	isConnected bool
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the root
	// node.
}

func NewElement(tagName string) *Element { return &Element{node{}, tagName, false} }

func (e *Element) Children() []interfaces.Element {
	// TODO: Is encapsulated in an HTMLCollection
	panic("TODO")
}

func (e *Element) NodeName() string {
	return e.TagName()
}

func (e *Element) TagName() string {
	return e.tagName
}

func (e *Element) IsConnected() bool { return e.isConnected }

// func (e *Element) Append(child interfaces.Element) interfaces.Element { return child }
