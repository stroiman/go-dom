package dom

import (
	"strings"

	"golang.org/x/net/html"
)

type CommentNode interface {
	Node
	// Text() string
}

type commentNode struct {
	node
	text string
}

func NewCommentNode(text string) Node {
	result := &commentNode{newNode(), text}
	result.SetSelf(result)
	return result
}

func (n *commentNode) Render(builder *strings.Builder) {
	builder.WriteString("<!--")
	builder.WriteString(n.text)
	builder.WriteString("-->")
}

func (n *commentNode) NodeType() NodeType {
	return NodeTypeComment
}

func (n *commentNode) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.CommentNode,
		Data: n.text,
	}
}
