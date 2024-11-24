package browser

import "golang.org/x/net/html"

type NamedNodeMap interface {
	Entity
	Length() int
	Item(index int) Attr
}

type Attr interface {
	Node
	LocalName() string
	Name() string
	NamespaceURI() string
	OwnerElement() Element
	Prefix() string
	GetValue() string
	SetValue(val string)
}

type namedNodeMap struct {
	base
	ownerElement Element
}

type attr struct {
	node
	ownerElement Element
	attr         *html.Attribute
}

func NewNamedNodeMapForElement(ownerElement Element) NamedNodeMap {
	return &namedNodeMap{newBase(), ownerElement}
}

func (m *namedNodeMap) Length() int {
	attributes := m.ownerElement.getAttributes()
	return len(attributes)
}

func (m *namedNodeMap) Item(index int) Attr {
	attributes := m.ownerElement.getAttributes()
	if index >= len(attributes) {
		return nil
	}
	if index < 0 {
		return nil
	}

	return &attr{newNode(), m.ownerElement, attributes[index]}
}

func (a *attr) LocalName() string     { panic("TODO") }
func (a *attr) Name() string          { return a.attr.Key }
func (a *attr) NamespaceURI() string  { panic("TODO") }
func (a *attr) OwnerElement() Element { return a.ownerElement }
func (a *attr) Prefix() string        { panic("TODO") }
func (a *attr) GetValue() string      { return a.attr.Val }
func (a *attr) SetValue(val string)   { a.attr.Val = val }
func (a *attr) NodeType() NodeType    { return NodeTypeAttribute }

func (a *attr) AppendChild(newChild Node) Node {
	return nil //, newDomError("Atrribute cannot have a child")
}

func (a *attr) InsertBefore(newChild Node, ref Node) (Node, error) {
	return nil, newDomError("Atrribute cannot have a child")
}

func (a *attr) createHtmlNode() *html.Node {
	panic(
		"N/A - createHtmlNode() is intended to be called when traversing child nodes; and attributes shouldn't appear as a child node.",
	)
}
