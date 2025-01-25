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
	return htmlElementGenerator{
		html.IdlNames[name],
	}.Generator(), nil
}

type htmlElementGenerator struct {
	idlType idl.IdlName
}

func (gen htmlElementGenerator) Generator() g.Generator {
	return g.StatementList(
		gen.GenerateAttributes(),
	)
}

func (gen htmlElementGenerator) GenerateAttributes() g.Generator {
	result := g.StatementList()
	for a := range gen.idlType.Attributes() {
		result.Append(IDLAttribute{
			AttributeName: idl.SanitizeName(a.Name),
			Receiver: Receiver{
				Name: g.NewValue("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		})
	}
	return result
}
