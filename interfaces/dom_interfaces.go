package interfaces

type EventTarget interface{}

type Node interface {
	EventTarget
	NodeName() string
}

type Element interface {
	Node
	// Append(Element) Element
	Children() []Element
	IsConnected() bool
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
