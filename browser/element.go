package browser

import (
	"errors"
	"slices"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// TODO: In the DOM, this is a `NamedNodeMap`. Is that useful in Go?
type Attributes []*html.Attribute

func (attrs Attributes) Length() int {
	return len(attrs)
}

type MouseEvents interface {
	Click() bool
}

// An Element in the document. Can be either an [HTMLElement] or an [XMLElement]
type Element interface {
	ElementContainer
	MouseEvents
	Append(Element) Element
	ClassList() DOMTokenList
	GetAttribute(name string) string
	SetAttribute(name string, value string)
	GetAttributes() NamedNodeMap
	InsertAdjacentHTML(position string, text string) error
	OuterHTML() string
	InnerHTML() string
	TagName() string
	// unexported
	getAttributes() Attributes
}

type element struct {
	node
	tagName       string
	namespace     string
	attributes    Attributes
	ownerDocument Document
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string, ownerDocument Document) Element {
	// TODO: handle namespace
	return &element{newNode(), tagName, "", Attributes(nil), ownerDocument}
}

func newElementFromNode(node *html.Node, ownerDocument Document) *element {
	attributes := make([]*html.Attribute, len(node.Attr))
	for i, a := range node.Attr {
		attributes[i] = new(html.Attribute)
		*attributes[i] = a
	}
	return &element{newNode(), node.Data, node.Namespace, attributes, ownerDocument}
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

func (e *element) ClassList() DOMTokenList {
	return NewClassList(e)
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

func (e *element) InnerHTML() string {
	writer := &strings.Builder{}
	for _, child := range e.ChildNodes().All() {
		html.Render(writer, toHtmlNode(child))
	}
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

func (e *element) getAttributes() Attributes {
	return e.attributes
}

func (e *element) GetAttributes() NamedNodeMap {
	return NewNamedNodeMapForElement(e)
}

func (e *element) SetAttribute(name string, value string) {
	idx := slices.IndexFunc(e.attributes, func(a *html.Attribute) bool {
		return a.Key == name && a.Namespace == e.namespace
	})
	if idx == -1 {
		e.attributes = append(e.attributes, &html.Attribute{
			Key:       name,
			Val:       value,
			Namespace: e.namespace})
	} else {
		e.attributes[idx].Val = value
	}
}
func (e *element) createHtmlNode() *html.Node {
	tag := strings.ToLower(e.tagName)
	attrs := make([]html.Attribute, len(e.attributes))
	for i, a := range e.attributes {
		attrs[i] = *a
	}
	return &html.Node{
		Type:      html.ElementNode,
		Data:      tag,
		DataAtom:  atom.Lookup([]byte(tag)),
		Namespace: e.namespace,
		Attr:      attrs,
	}
}

func (e *element) QuerySelector(pattern string) (Node, error) {
	return CSSHelper{e}.QuerySelector(pattern)
}

func (e *element) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	return CSSHelper{e}.QuerySelectorAll(pattern)
}

func (n *element) InsertAdjacentHTML(position string, text string) error {
	var (
		parent    Node
		reference Node
	)
	switch position {
	case "beforebegin":
		parent = n.Parent()
		reference = n // NOTE This will not work for subclasses
	case "afterbegin":
		parent = n
		reference = n.ChildNodes().Item(0)
	case "beforeend":
		parent = n
		reference = nil
	case "afterend":
		parent = n.Parent()
		reference = n.NextSibling()
	default:
		return errors.New("Invalid position")
	}
	nodes, err := html.ParseFragment(strings.NewReader(text), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err == nil {
		for _, child := range nodes {
			element := createElementFromNode(nil, n.OwnerDocument(), nil, child)
			parent.InsertBefore(element, reference)
		}
	}
	return err
}

func (n *element) NodeType() NodeType { return NodeTypeElement }

func (n *element) Click() bool {
	return n.DispatchEvent(NewCustomEvent("click", EventCancelable(true), EventBubbles(true)))
}
