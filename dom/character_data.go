package dom

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

/* -------- CharacterData -------- */

type CharacterData interface {
	Node
	Data() string
	Length() int
}

type characterData struct {
	node
	data string
}

func (n *characterData) Data() string {
	return n.data
}

func (n *characterData) Length() int {
	return utf8.RuneCountInString(n.data)
}

func (d *characterData) cloneChildren() []Node { return nil }

/* -------- Comment -------- */

type Comment interface {
	CharacterData
}

type comment struct {
	characterData
}

func NewComment(text string) Comment {
	result := &comment{characterData{newNode(), text}}
	result.SetSelf(result)
	return result
}

func (n *comment) Render(builder *strings.Builder) {
	builder.WriteString("<!--")
	builder.WriteString(n.Data())
	builder.WriteString("-->")
}

func (n *comment) NodeType() NodeType {
	return NodeTypeComment
}

func (n *comment) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.CommentNode,
		Data: n.Data(),
	}
}

func (c *comment) CloneNode(bool) Node {
	return NewComment(c.TextContent())
}

/* -------- Text -------- */

type Text interface {
	CharacterData
}

type textNode struct {
	characterData
}

func NewText(text string) Text {
	result := &textNode{characterData{newNode(), text}}
	result.SetSelf(result)
	return result
}

func (n *textNode) CloneNode(bool) Node {
	return NewText(n.characterData.data)
}

func (n *textNode) Render(builder *strings.Builder) {
	builder.WriteString(n.Data())
}

func (n *textNode) NodeType() NodeType { return NodeTypeText }

func (n *textNode) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: n.Data(),
	}
}

func (n *textNode) TextContent() string {
	return n.Data()
}
