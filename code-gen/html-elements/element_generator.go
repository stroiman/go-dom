package htmlelements

import (
	"errors"
	"fmt"
	"strings"

	g "github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/webref/elements"
	"github.com/stroiman/go-dom/code-gen/webref/idl"
)

// CreateHTMLElementGenerator creates a generator for the element with
func CreateHTMLElementGenerator(name string) (g.Generator, error) {
	html, err1 := idl.LoadIdlParsed("html")
	el, err2 := elements.Load()
	tagName, err3 := el.GetTagNameForInterfaceError(name)
	err := errors.Join(err1, err2, err3)
	if err != nil {
		return nil, err
	}
	return htmlElementGenerator{
		html.Interfaces[name],
		g.NewType(toStructName(name)),
		tagName,
	}.Generator(), nil
}

type htmlElementGenerator struct {
	idlType idl.Interface
	type_   g.Type
	tagName string
}

func (gen htmlElementGenerator) Generator() g.Generator {
	return g.StatementList(
		gen.GenerateInterface(),
		g.Line,
		// TODO: Make this configuratble
		// gen.GenerateStruct(),
		// g.Line,
		// gen.GenerateConstructor(),
		g.Line,
		gen.GenerateAttributes(),
	)
}

func toStructName(name string) string {
	return strings.Replace(name, "HTML", "html", 1)
}

func (gen htmlElementGenerator) GenerateInterface() g.Generator {
	attributes := make([]IdlInterfaceAttribute, 0)

	interfaces := make([]idl.Interface, 1+len(gen.idlType.Includes))
	interfaces[0] = gen.idlType
	copy(interfaces[1:], gen.idlType.Includes)

	for _, i := range interfaces {
		for _, a := range i.Attributes {
			attributes = append(attributes, IdlInterfaceAttribute{
				Name:     a.Name,
				ReadOnly: a.Readonly,
			})
		}
	}
	return IdlInterface{
		Name:       gen.idlType.Name,
		Inherits:   gen.idlType.InternalSpec.Inheritance,
		Attributes: attributes,
	}
}
func (gen htmlElementGenerator) GenerateStruct() g.Generator {
	res := g.Struct{Name: g.NewType(toStructName(gen.idlType.Name))}
	res.Embed(g.Id("HTMLElement"))
	// for a := range gen.idlType.Attributes() {
	// 	res.Field(g.Id(idl.SanitizeName(a.Name)), g.Id("string"))
	// }
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
	for _, a := range gen.idlType.Attributes {
		result.Append(IDLAttribute{
			AttributeName: a.Name,
			ReadOnly:      a.Readonly,
			Receiver: Receiver{
				Name: g.NewValue("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		})
	}
	return result
}

type FileGeneratorSpec struct {
	Name      string
	Generator g.Generator
}

func CreateHTMLElementGenerators() ([]FileGeneratorSpec, error) {
	generator, error := CreateHTMLElementGenerator("HTMLAnchorElement")
	return []FileGeneratorSpec{
		{"html_anchor_element", generator},
	}, error
}
