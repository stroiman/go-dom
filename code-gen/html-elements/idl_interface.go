package htmlelements

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/webref/idl"
)

/* -------- IdlInterface -------- */

type IdlInterface struct {
	Name       string
	Inherits   string
	Attributes []IdlInterfaceAttribute
	Operations []IdlInterfaceOperation
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
		fields = append(fields, generators.Raw(
			jen.Id(getterName).Params().Params(jen.Id("string")),
		))
		if !a.ReadOnly {
			setterName := fmt.Sprintf("Set%s", getterName)
			fields = append(fields, generators.Raw(
				jen.Id(setterName).Params(jen.Id("string")),
			))
		}
	}
	for _, o := range i.Operations {
		if !o.Static {
			// Todo: Parameters
			// Todo: Return type
			fields = append(fields, generators.Raw(
				jen.Id(upperCaseFirstLetter(o.Name)).
					Params().
					Params(jen.Id("string"), jen.Id("error")),
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

/* -------- IdlInterfaceOperation -------- */

type IdlInterfaceOperation struct {
	idl.Operation
}
