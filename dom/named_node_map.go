package dom

import (
	"iter"

	"github.com/gost-dom/browser/internal/entity"
	"golang.org/x/net/html"
)

type NamedNodeMap interface {
	entity.Entity
	All() iter.Seq[Attr]
	Length() int
	Item(index int) Attr
}

// Attr is the interface corresponding to the [Attr DOM node]
//
// [Attr DOM node]: https://developer.mozilla.org/en-US/docs/Web/API/Attr
type Attr interface {
	Node
	LocalName() string
	Name() string
	NamespaceURI() string
	OwnerElement() Element
	Prefix() string
	Value() string
	SetValue(val string)

	htmlAttr() html.Attribute
}

type namedNodeMap struct {
	entity.Entity
	ownerElement Element
}

type attr struct {
	node
	ownerElement Element
	attr         *html.Attribute
}

func newAttr(n, v string) Attr {
	res := &attr{
		node: newNode(),
		attr: &html.Attribute{
			Key: n,
			Val: v,
		},
	}
	res.setSelf(res)
	return res
}

func (a *attr) setParent(parent Node) {
	var ok bool
	a.ownerElement = nil
	if a.ownerElement, ok = parent.(Element); !ok && parent != nil {
		panic("Setting non-element owner on an attribute")
	}
	a.node.setParent(parent)
}

func (a *attr) htmlAttr() html.Attribute {
	return *a.attr
}

func (a *attr) CloneNode(deep bool) Node { return newAttr(a.attr.Key, a.attr.Val) }

func newNamedNodeMapForElement(ownerElement Element) NamedNodeMap {
	return &namedNodeMap{entity.New(), ownerElement}
}

func (m *namedNodeMap) All() iter.Seq[Attr] {
	return func(yield func(Attr) bool) {
		for i := 0; i < m.Length(); i++ {
			if !yield(m.Item(i)) {
				return
			}
		}
	}
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

	result := attributes[index]
	return result
}

func (a *attr) LocalName() string     { return a.attr.Key }
func (a *attr) Name() string          { return a.attr.Key }
func (a *attr) NamespaceURI() string  { return "" }
func (a *attr) OwnerElement() Element { return a.ownerElement }
func (a *attr) Prefix() string        { return "" }
func (a *attr) Value() string         { return a.attr.Val }
func (a *attr) SetValue(val string)   { a.attr.Val = val }
func (a *attr) NodeType() NodeType    { return NodeTypeAttribute }

func (a *attr) AppendChild(newChild Node) (Node, error) {
	return nil, newDomError("Atrribute cannot have a child")
}

func (a *attr) InsertBefore(newChild Node, ref Node) (Node, error) {
	return nil, newDomError("Atrribute cannot have a child")
}

func (a *attr) createHtmlNode() *html.Node {
	panic(
		"N/A - createHtmlNode() is intended to be called when traversing child nodes; and attributes shouldn't appear as a child node.",
	)
}
