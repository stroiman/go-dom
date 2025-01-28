package htmlelements

import (
	"errors"
	"fmt"
	"strings"

	g "github.com/gost-dom/browser/code-gen/generators"
	"github.com/gost-dom/browser/code-gen/webref/elements"
	"github.com/gost-dom/browser/code-gen/webref/idl"
)

// HTMLGeneratorReq specifies what to generate for a specific Web IDL spec. The
// name of the spec is in the SpecName field, and the interface to generate is in
// the Interface field. The Generate... fields specify what to generate.
//
// Note: As more needs for customisation arises, so will the Generate... fields
// also more likely become more complex.
//
// E.g., the URL type is loaded from url.json, and has the interface name URL in
// the idl, so this is specified by:
//
//	HTMLGeneratorReq {
//		InterfaceName: "URL",
//		SpecName:      "url",
//	}
type HTMLGeneratorReq struct {
	InterfaceName       string
	SpecName            string
	GenerateStruct      bool
	GenerateConstructor bool
	GenerateInterface   bool
	GenerateAttributes  bool
}

/* -------- baseGenerator -------- */

type baseGenerator struct {
	req     HTMLGeneratorReq
	idlType idl.Interface
	type_   g.Type
}

func CreateGenerator(req HTMLGeneratorReq) (baseGenerator, error) {
	html, err := idl.LoadIdlParsed(req.SpecName)
	return baseGenerator{
		req,
		html.Interfaces[req.InterfaceName],
		g.NewType(toStructName(req.InterfaceName)),
	}, err
}

func (gen baseGenerator) GenerateInterface() g.Generator {
	attributes := make([]IdlInterfaceAttribute, 0)
	operations := make([]IdlInterfaceOperation, 0)

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
		for _, o := range i.Operations {
			operations = append(operations, IdlInterfaceOperation{o})
		}
	}
	return IdlInterface{
		Name:       gen.idlType.Name,
		Inherits:   gen.idlType.InternalSpec.Inheritance,
		Attributes: attributes,
		Operations: operations,
	}
}

/* -------- htmlElementGenerator -------- */

// CreateHTMLElementGenerator creates a generator for the element with
func CreateHTMLElementGenerator(req HTMLGeneratorReq) (htmlElementGenerator, error) {
	base, err1 := CreateGenerator(req)
	el, err2 := elements.Load()
	tagName, err3 := el.GetTagNameForInterfaceError(req.InterfaceName)
	err := errors.Join(err1, err2, err3)
	if err != nil {
		return htmlElementGenerator{}, err
	}
	return htmlElementGenerator{
		base,
		tagName,
	}, nil
}

type htmlElementGenerator struct {
	baseGenerator
	tagName string
}

func (gen htmlElementGenerator) Generator() g.Generator {
	result := g.StatementList()
	if gen.req.GenerateInterface {
		result.Append(
			gen.GenerateInterface(),
			g.Line,
		)
	}
	if gen.req.GenerateStruct {
		result.Append(gen.GenerateStruct(),
			g.Line,
		)
	}
	if gen.req.GenerateConstructor {
		result.Append(
			gen.GenerateConstructor(),
			g.Line,
		)
	}
	if gen.req.GenerateAttributes {
		result.Append(gen.GenerateAttributes())
	}
	return result
}

func toStructName(name string) string {
	return strings.Replace(name, "HTML", "html", 1)
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
				Type: gen.type_.Pointer(),
			},
		})
	}
	return result
}

type FileGeneratorSpec struct {
	Name      string
	Package   string
	Generator g.Generator
}

var HTMLAnchorElementSpecs = HTMLGeneratorReq{
	InterfaceName:      "HTMLAnchorElement",
	SpecName:           "html",
	GenerateInterface:  true,
	GenerateAttributes: true,
}

func CreateHTMLElementGenerators() ([]FileGeneratorSpec, error) {
	generator, error := CreateHTMLElementGenerator(HTMLAnchorElementSpecs)
	return []FileGeneratorSpec{
		{"html_anchor_element",
			"github.com/gost-dom/browser/browser/html",
			generator.Generator(),
		},
	}, errors.Join(error)
}

var URLSpec = HTMLGeneratorReq{
	InterfaceName:     "URL",
	SpecName:          "url",
	GenerateInterface: true,
}

func CreateDOMGenerators() ([]FileGeneratorSpec, error) {
	// return []FileGeneratorSpec{}, nil
	generator, error := CreateGenerator(HTMLGeneratorReq{
		InterfaceName:      "URL",
		SpecName:           "url",
		GenerateInterface:  true,
		GenerateAttributes: true,
	})
	return []FileGeneratorSpec{{
		"url",
		"github.com/stroiman/go-dom/browser/dom",
		generator.GenerateInterface(),
	}}, errors.Join(error)
}
