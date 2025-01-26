package htmlelements

import (
	"errors"
	"fmt"
	"strings"

	g "github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/webref/elements"
	"github.com/stroiman/go-dom/code-gen/webref/idl"
)

func GenerateHTMLElement(name string) (g.Generator, error) {
	html, err1 := idl.LoadIdlParsed("html")
	el, err2 := elements.Load()
	tagName, err3 := el.GetTagNameForInterfaceError(name)
	err := errors.Join(err1, err2, err3)
	if err != nil {
		return nil, err
	}
	return htmlElementGenerator{
		html.IdlNames[name],
		g.NewType(toStructName(name)),
		tagName,
	}.Generator(), nil
}

type htmlElementGenerator struct {
	idlType idl.Name
	type_   g.Type
	tagName string
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
				t.CreateInstance(g.NewValue("NewHTMLElement").Call(g.Lit(gen.tagName), owner)).
					Reference(),
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
