package htmlelements

import (
	g "github.com/stroiman/go-dom/code-gen/generators"
)

func GenerateHtmlAnchor() g.Generator {
	return IDLAttribute{
		AttributeName: "target",
		Receiver: Receiver{
			Name: g.NewValue("e"),
			Type: g.NewType("htmlAnchorElement").Pointer(),
		},
	}
}
