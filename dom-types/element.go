package dom_types

import "github.com/stroiman/go-dom/interfaces"

type Element struct {
	tagName     string
	isConnected bool
}

func NewElement(tagName string) *Element { return &Element{tagName, false} }

func (e *Element) Children() []interfaces.Element {
	// TODO: Is encapsulated in an HTMLCollection
	return nil
}

func (e *Element) NodeName() string {
	return e.TagName()
}

func (e *Element) TagName() string {
	return e.tagName
}

func (e *Element) IsConnected() bool { return e.isConnected }

func (e *Element) Append(child interfaces.Element) interfaces.Element { return child }
