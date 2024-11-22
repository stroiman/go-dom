package browser

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// TODO: In the DOM, this is a `NamedNodeMap`. Is that useful in Go?
type Attributes []html.Attribute

func (attrs Attributes) Length() int {
	return len(attrs)
}

type Element interface {
	ElementContainer
	Append(Element) Element
	GetAttribute(name string) string
	SetAttribute(name string, value string)
	GetAttributes() Attributes
	OuterHTML() string
	TagName() string
}

type element struct {
	node
	tagName    string
	namespace  string
	attributes Attributes
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string) Element {
	// TODO: handle namespace
	return &element{newNode(), tagName, "", Attributes(nil)}
}

func newElementFromNode(node *html.Node) Element {
	return &element{newNode(), node.Data, node.Namespace, node.Attr}
}

func (e *element) NodeName() string {
	return e.TagName()
}

func (e *element) TagName() string {
	return strings.ToUpper(e.tagName)
}

func (parent *element) Append(child Element) Element {
	NodeHelper{parent}.AppendChild(child)
	return child
}

func (parent *element) AppendChild(child Node) Node {
	return NodeHelper{parent}.AppendChild(child)
}

func (e *element) InsertBefore(newChild Node, reference Node) (Node, error) {
	return NodeHelper{e}.InsertBefore(newChild, reference)
}

func (e *element) OuterHTML() string {
	writer := &strings.Builder{}
	html.Render(writer, toHtmlNode(e))
	return string(writer.String())
}

func (e *element) GetAttribute(name string) string {
	for _, a := range e.attributes {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}

func (e *element) GetAttributes() Attributes { return e.attributes }

func (e *element) SetAttribute(name string, value string) {
	idx := slices.IndexFunc(e.attributes, func(a html.Attribute) bool {
		return a.Key == name && a.Namespace == e.namespace
	})
	if idx == -1 {
		e.attributes = append(e.attributes, html.Attribute{
			Key:       name,
			Val:       value,
			Namespace: e.namespace})
	} else {
		e.attributes[idx].Val = value
	}
}
func (e *element) createHtmlNode() *html.Node {
	tag := strings.ToLower(e.tagName)
	return &html.Node{
		Type:      html.ElementNode,
		Data:      tag,
		DataAtom:  atom.Lookup([]byte(tag)),
		Namespace: e.namespace,
		Attr:      e.attributes,
	}
}

func (e *element) QuerySelector(pattern string) (Node, error) {
	return CSSHelper{e}.QuerySelector(pattern)
}

func (e *element) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	return CSSHelper{e}.QuerySelectorAll(pattern)
}
