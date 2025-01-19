package dom

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

type CommentNode interface {
	Node
	Data() string
	Length() int
}

type commentNode struct {
	node
	text string
}

func NewCommentNode(text string) CommentNode {
	result := &commentNode{newNode(), text}
	result.SetSelf(result)
	return result
}

func (n *commentNode) Data() string {
	return n.text
}

func (n *commentNode) Length() int {
	return utf8.RuneCountInString(n.text)
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
