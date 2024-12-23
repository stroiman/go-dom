package browser

import "golang.org/x/net/html"

type TextNode interface {
	Node
	Text() string
}

type textNode struct {
	node
	text string
}

func NewTextNode(node *html.Node, text string) Node {
	return &textNode{newNode(), text}
}

func (n *textNode) AppendChild(newChild Node) Node {
	return NodeHelper{n}.AppendChild(newChild)
}

func (n *textNode) InsertBefore(newChild Node, reference Node) (Node, error) {
	return NodeHelper{n}.InsertBefore(newChild, reference)
}

func (n *textNode) Text() string { return n.text }

func (n *textNode) NodeType() NodeType { return NodeTypeText }

func (n *textNode) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: n.text,
	}
}
