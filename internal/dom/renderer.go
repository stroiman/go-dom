package dom

import "strings"

// Renderer is the interface for DOM types that can render themselves.
//
// Renderer is used for generating HTML and TextContent from DOM nodes.
type Renderer interface {
	Render(*strings.Builder)
}

// ChildrenRenderer is the interface for DOM types that have children that can
// be rendered.
//
// ChildrenRenderer is used for generating HTML and TextContent from DOM nodes.
type ChildrenRenderer interface {
	RenderChildren(*strings.Builder)
}
