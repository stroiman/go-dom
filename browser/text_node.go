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
	return &textNode{newNode(node), text}
}

func (n *textNode) Text() string { return n.text }
