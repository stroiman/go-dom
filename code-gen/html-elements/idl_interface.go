package htmlelements

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stroiman/go-dom/code-gen/generators"
)

/* -------- IdlInterface -------- */

type IdlInterface struct {
	Name       string
	Inherits   string
	Attributes []IdlInterfaceAttribute
}

func (i IdlInterface) Generate() *jen.Statement {
	fields := make(
		[]generators.Generator,
		0,
		2*len(i.Attributes)+1,
	) // Make room for getters and setters
	if i.Inherits != "" {
		fields = append(fields, generators.Id(i.Inherits))
	}

	for _, a := range i.Attributes {
		getterName := upperCaseFirstLetter(a.Name)
		setterName := fmt.Sprintf("Set%s", getterName)
		fields = append(fields, generators.Raw(
			jen.Id(getterName).Params().Params(jen.Id("string")),
		))
		if !a.ReadOnly {
			fields = append(fields, generators.Raw(
				jen.Id(setterName).Params(jen.Id("string")),
			))
		}
	}
	return jen.Type().Add(jen.Id(i.Name)).Interface(generators.ToJenCodes(fields)...)
}

/* -------- IdlInterfaceAttribute -------- */

type IdlInterfaceAttribute struct {
	Name     string
	ReadOnly bool
}

func NewStringAttribute(name string) IdlInterfaceAttribute {
	return IdlInterfaceAttribute{Name: name}
}
