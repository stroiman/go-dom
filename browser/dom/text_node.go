package dom

import (
	"strings"

	"golang.org/x/net/html"
)

type TextNode interface {
	Node
	Text() string
}

type textNode struct {
	node
	text string
}

func NewTextNode(node *html.Node, text string) Node {
	result := &textNode{newNode(), text}
	result.SetSelf(result)
	return result
}

func (n *textNode) Render(builder *strings.Builder) {
	builder.WriteString(n.text)
}

func (n *textNode) Text() string { return n.text }

func (n *textNode) NodeType() NodeType { return NodeTypeText }

func (n *textNode) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: n.text,
	}
}

func (n *textNode) GetTextContent() string {
	return n.text
}
