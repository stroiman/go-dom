package dom

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/gost-dom/browser/internal/constants"
	. "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/log"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Attribute = html.Attribute

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
	GetAttributeNode(string) *Attribute
	SetAttributeNode(*Attribute) *Attribute
	RemoveAttributeNode(*Attribute) (*Attribute, error)
	Attributes() NamedNodeMap
	InsertAdjacentHTML(position string, text string) error
	OuterHTML() string
	InnerHTML() string
	TagName() string
	Matches(string) (bool, error)
	// unexported
	getAttributes() Attributes
	getSelfElement() Element
}

type element struct {
	node
	tagName          string
	namespace        string
	attributes       Attributes
	ownerDocument    Document
	selfElement      Element
	selfRenderer     Renderer
	childrenRenderer ChildrenRenderer
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string, ownerDocument Document) Element {
	// return newElement(tagName, ownerDocument)
	// // TODO: handle namespace
	result := &element{newNode(), tagName, "", Attributes(nil), ownerDocument, nil, nil, nil}
	result.SetSelf(result)
	return result
}

func newElement(tagName string, ownerDocument Document) *element {
	// TODO: handle namespace
	result := &element{newNode(), tagName, "", Attributes(nil), ownerDocument, nil, nil, nil}
	result.SetSelf(result)
	return result
}

func (e *element) ChildElementCount() int {
	return len(e.childElements())
}

func (e *element) SetSelf(n Node) {
	if self, ok := n.(Element); ok {
		e.selfElement = self
	} else {
		panic("Setting a non-element as element self")
	}
	if self, ok := n.(Renderer); ok {
		e.selfRenderer = self
	} else {
		panic("Setting a non-renderer as element self")
	}
	if self, ok := n.(ChildrenRenderer); ok {
		e.childrenRenderer = self
	} else {
		panic("Setting a non-child-renderer as element self")
	}
	e.node.SetSelf(n)
}

func (e *element) NodeName() string {
	return e.selfElement.TagName()
}

func (e *element) TagName() string {
	return strings.ToLower(e.tagName)
}

func (e *element) Append(nodes ...Node) error {
	return e.append(nodes...)
}

func (e *element) ClassList() DOMTokenList {
	return NewClassList(e)
}

func (e *element) OuterHTML() string {
	writer := &strings.Builder{}
	e.selfRenderer.Render(writer)
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
	if a := e.GetAttributeNode(name); a != nil {
		return a.Val, true
	} else {
		return "", false
	}
}

func (e *element) GetAttributeNode(name string) *Attribute {
	for _, a := range e.attributes {
		if a.Key == name && a.Namespace == e.namespace {
			return a
		}
	}
	return nil
}

func (e *element) SetAttributeNode(node *Attribute) *Attribute {
	for i, a := range e.attributes {
		if a.Key == node.Key && a.Namespace == node.Namespace {
			e.attributes[i] = node
			return a
		}
	}
	e.attributes = append(e.attributes, node)
	return nil
}

func (e *element) RemoveAttributeNode(node *Attribute) (*Attribute, error) {
	for i, a := range e.attributes {
		if a == node {
			e.attributes = slices.Delete(e.attributes, i, i+1)
			return node, nil
		}
	}
	return nil, newDomErrorCode("Node was not found", domErrorNotFound)
}

func (e *element) getAttributes() Attributes {
	return e.attributes
}

func (e *element) getSelfElement() Element {
	if r := e.selfElement; r != nil {
		return r
	}
	panic(
		"Calling method on an element which isn't an element. Did a custom type forget to call 'setSelf()'?",
	)
}

func (e *element) Attributes() NamedNodeMap {
	return newNamedNodeMapForElement(e)
}

func (e *element) SetAttribute(name string, value string) {
	if a := e.GetAttributeNode(name); a != nil {
		a.Val = value
	} else {
		e.attributes = append(e.attributes, &html.Attribute{
			Key:       name,
			Val:       value,
			Namespace: e.namespace})
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

func renderElement(e *element, writer *strings.Builder) {
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
	e.childrenRenderer.RenderChildren(writer)
	writer.WriteString("</")
	writer.WriteString(tagName)
	writer.WriteRune('>')
}

var tagNameRegExp = regexp.MustCompile("(?m:^[a-zA-Z]+$)")
var attributeRegExp = regexp.MustCompile("(?m:^[[]([a-zA-Z-]+)[]]$)")
var tagNameAndAttribute = regexp.MustCompile(`(?m:^([a-zA-Z]+)+[[]([a-zA-Z-]+)="([a-zA-Z-]+)"[]]$)`)

// Element.Matches returns true if the current element matches the specified CSS
// selectors; accepting a comma-separated list of selectors with any leading and
// trailing whitespace trimmed. Returns an error if the patterns is not
// supported (or invalid)
//
// Note: This only supports a subset of CSS selectors; primarily those used by
// HTMX.
func (e *element) Matches(pattern string) (bool, error) {
	// Implementation note:
	// This might have been implemented in terms of QuerySelector == self, but
	// QuerySelector find child elements, not the element itself it is queried
	// upon.

	patterns := strings.Split(pattern, ",")
	log.Debug("Element.Matches", "pattern", pattern, "element", e.getSelfElement().OuterHTML())

	for _, p := range patterns {
		p = strings.TrimSpace(p)

		knownPattern := false

		if tagNameRegExp.MatchString(p) {
			knownPattern = true
			if strings.ToLower(e.getSelfElement().TagName()) == strings.ToLower(p) {
				return true, nil
			}
		}
		if m := attributeRegExp.FindStringSubmatch(p); m != nil {
			knownPattern = true
			if len(m) != 2 {
				panic(
					fmt.Sprintf(
						"Element.Matches: Unexpected no of matches\n%s\n%s\n - Pattern: %s\n - Element: %s",
						constants.BUG_USSUE_URL,
						constants.BUG_USSUE_DETAILS,
						p,
						e.getSelfElement().OuterHTML(),
					),
				)
			}
			_, hasAttribute := e.GetAttribute(m[1])
			if hasAttribute {
				return true, nil
			}
		}
		if m := tagNameAndAttribute.FindStringSubmatch(p); m != nil {
			knownPattern = true
			if len(m) != 4 {
				panic(
					fmt.Sprintf(
						"Element.Matches: Unexpected no of matches\n%s\n%s\n - Pattern: %s\n - Element: %s",
						constants.BUG_USSUE_URL,
						constants.BUG_USSUE_DETAILS,
						p,
						e.getSelfElement().OuterHTML(),
					),
				)
			}
			tag := m[1]
			key := m[2]
			val := m[3]
			v, found := e.GetAttribute(key)
			if strings.ToLower(e.getSelfElement().TagName()) == strings.ToLower(tag) && found &&
				val == v {
				return true, nil
			}
		}
		if !knownPattern {
			log.Error("Element.Matches: unsupported pattern", "pattern", p)
			return false, fmt.Errorf(
				"Element.matches: Unsupported pattern - patterns: %s\n%s\n",
				p,
				constants.MISSING_FEATURE_ISSUE_URL,
			)
		}
	}
	log.Debug("Element.Matches: no match", "pattern", pattern, "e", e.getSelfElement().OuterHTML())
	return false, nil
}

func (e *element) String() string {
	childLen := e.ChildNodes().Length()

	id, found := e.GetAttribute("id")
	if found {
		id = "id='" + id + "'"
	}
	return fmt.Sprintf("<%s %s(child count=%d) />", e.tagName, id, childLen)
}
