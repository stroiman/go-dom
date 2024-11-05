package interfaces

type EventTarget interface{}

type Node interface {
	EventTarget
	NodeName() string
}

type Document interface {
	Node
}

type Element interface {
	Node
	TagName() string
}

type HTMLElement interface {
	Element
}

type HTMLHtmlElement interface {
	HTMLElement
}

type HTMLHeadElement interface {
	HTMLElement
}
