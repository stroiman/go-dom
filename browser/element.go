package browser

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Element interface {
	Node
	// Children() []Element
	Append(Element) Element
	GetAttribute(name string) string
	IsConnected() bool
	OuterHTML() string
	TagName() string
}

type element struct {
	node
	tagName     string
	isConnected bool
	namespace   string
	attributes  []html.Attribute
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string, node *html.Node) Element {
	return &element{newNode(node), tagName, false, node.Namespace, node.Attr}
}

func (e *element) NodeName() string {
	return e.TagName()
}

func (e *element) TagName() string {
	return strings.ToUpper(e.tagName)
}

func (e *element) IsConnected() bool { return e.isConnected }

func (parent *element) Append(child Element) Element {
	parent.AppendChild(child)
	return child
}

func (e *element) OuterHTML() string {
	writer := &strings.Builder{}
	html.Render(writer, NodeIterator{e}.toHtmlNode(nil))
	return string(writer.String())
}

func (n *element) populateNodeMap(m map[*html.Node]Node) {
	m[n.htmlNode] = n
	for _, c := range n.childNodes {
		c.populateNodeMap(m)
	}
}
func (n *node) GetAttribute(name string) string {
	for _, a := range n.htmlNode.Attr {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
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
