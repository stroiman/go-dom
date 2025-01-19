package dom

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	. "github.com/stroiman/go-dom/browser/internal/dom"

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
	ClassList() DOMTokenList
	HasAttribute(name string) bool
	GetAttribute(name string) (string, bool)
	SetAttribute(name string, value string)
	Attributes() NamedNodeMap
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
	selfElement   Element
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string, ownerDocument Document) Element {
	// return newElement(tagName, ownerDocument)
	// // TODO: handle namespace
	result := &element{newNode(), tagName, "", Attributes(nil), ownerDocument, nil}
	result.SetSelf(result)
	return result
}

func newElement(tagName string, ownerDocument Document) *element {
	// TODO: handle namespace
	result := &element{newNode(), tagName, "", Attributes(nil), ownerDocument, nil}
	result.SetSelf(result)
	return result
}

func (e *element) SetSelf(n Node) {
	self, ok := n.(Element)
	if !ok {
		panic("Setting a non-element as element self")
	}
	e.selfElement = self
	e.node.SetSelf(n)
}

func (e *element) NodeName() string {
	return e.selfElement.TagName()
}

func (e *element) TagName() string {
	return strings.ToLower(e.tagName)
}

func (parent *element) Append(child Element) (Element, error) {
	_, err := parent.AppendChild(child)
	return child, err
}

func (e *element) ClassList() DOMTokenList {
	return NewClassList(e)
}

func (e *element) OuterHTML() string {
	writer := &strings.Builder{}
	if renderer, ok := e.self.(Renderer); ok {
		renderer.Render(writer)
	}
	return writer.String()
}

func (e *element) InnerHTML() string {
	writer := &strings.Builder{}
	e.renderChildren(writer)
	return writer.String()
}

func (e *element) HasAttribute(name string) bool {
	for _, a := range e.attributes {
		if a.Key == name {
			return true
		}
	}
	return false
}

func (e *element) GetAttribute(name string) (string, bool) {
	for _, a := range e.attributes {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}

func (e *element) getAttributes() Attributes {
	return e.attributes
}

func (e *element) Attributes() NamedNodeMap {
	return newNamedNodeMapForElement(e)
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

func (e *element) QuerySelector(pattern string) (Element, error) {
	return cssHelper{e}.QuerySelector(pattern)
}

func (e *element) QuerySelectorAll(pattern string) (staticNodeList, error) {
	return cssHelper{e}.QuerySelectorAll(pattern)
}

func (n *element) InsertAdjacentHTML(position string, text string) error {
	var (
		parent    Node
		reference Node
	)
	switch position {
	case "beforebegin":
		parent = n.Parent()
		reference = n.getSelf()
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
	fragment, err := n.ownerDocument.parseFragment(strings.NewReader(text))
	if err == nil {
		_, err = parent.InsertBefore(fragment, reference)
	}
	return err
}

func (n *element) NodeType() NodeType { return NodeTypeElement }

func (n *element) Click() bool {
	return n.DispatchEvent(NewCustomEvent("click", EventCancelable(true), EventBubbles(true)))
}

func (e *element) Render(writer *strings.Builder) {
	renderElement(e, writer)
}

func renderElement(e Element, writer *strings.Builder) {
	tagName := strings.ToLower(e.TagName())
	writer.WriteRune('<')
	writer.WriteString(tagName)
	for a := range e.Attributes().All() {
		writer.WriteRune(' ')
		writer.WriteString(a.Name())
		writer.WriteString("=\"")
		writer.WriteString(a.Value())
		writer.WriteString("\"")
	}
	writer.WriteRune('>')
	if childRenderer, ok := e.getSelf().(ChildrenRenderer); ok {
		childRenderer.RenderChildren(writer)
	}
	writer.WriteString("</")
	writer.WriteString(tagName)
	writer.WriteRune('>')
}

func (e *element) String() string {
	childLen := e.ChildNodes().Length()

	id, found := e.GetAttribute("id")
	if found {
		id = "id='" + id + "'"
	}
	return fmt.Sprintf("<%s %s(child count=%d) />", e.tagName, id, childLen)
}
