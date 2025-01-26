package htmlelements

import (
	"fmt"
	"strings"

	g "github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/webref/idl"
)

func GenerateHTMLElement(name string) (g.Generator, error) {
	html, err := idl.LoadIdlParsed("html")
	if err != nil {
		return nil, err
	}
	return htmlElementGenerator{
		html.IdlNames[name],
		g.NewType(toStructName(name)),
	}.Generator(), nil
}

type htmlElementGenerator struct {
	idlType idl.Name
	type_   g.Type
}

func (gen htmlElementGenerator) Generator() g.Generator {
	return g.StatementList(
		gen.GenerateStruct(),
		g.Line,
		gen.GenerateConstructor(),
		g.Line,
		gen.GenerateAttributes(),
	)
}

func toStructName(name string) string {
	return strings.Replace(name, "HTML", "html", 1)
}

func (gen htmlElementGenerator) GenerateStruct() g.Generator {
	res := g.Struct{Name: g.NewType(toStructName(gen.idlType.Name))}
	res.Embed(g.Id("HTMLElement"))
	for a := range gen.idlType.Attributes() {
		res.Field(g.Id(idl.SanitizeName(a.Name)), g.Id("string"))
	}
	return res
}

func (gen htmlElementGenerator) GenerateConstructor() g.Generator {
	res := g.NewValue("result")
	i := g.NewType(gen.idlType.Name)
	t := g.NewType(toStructName(gen.idlType.Name))
	owner := g.Id("ownerDoc")
	return g.FunctionDefinition{
		Name:     fmt.Sprintf("New%s", gen.idlType.Name),
		RtnTypes: g.List(i),
		Args:     g.Arg(owner, g.Id("HTMLDocument")),
		Body: g.StatementList(
			g.Assign(
				res,
				t.CreateInstance(g.NewValue("NewHTMLElement").Call(g.Lit("a"), owner)).Reference(),
			),
			res.Field("SetSelf").Call(res),
			g.Return(res),
		),
	}
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
