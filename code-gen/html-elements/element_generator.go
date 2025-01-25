package htmlelements

import (
	g "github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/idl"
)

func GenerateHTMLElement(name string) (g.Generator, error) {
	html, err := idl.LoadIdlParsed("html")
	if err != nil {
		return nil, err
	}
	a := html.IdlNames["HTMLAnchorElement"]
	result := g.StatementList()
	for a := range a.Attributes() {
		result.Append(IDLAttribute{
			AttributeName: idl.SanitizeName(a.Name),
			Receiver: Receiver{
				Name: g.NewValue("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		})
	}
	return result, nil
}
