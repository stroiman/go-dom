package generators

import (
	"github.com/dave/jennifer/jen"
)

type Type struct{ Generator }

func NewType(name string) Type                    { return Type{Raw(jen.Id(name))} }
func NewTypePackage(name string, pkg string) Type { return Type{Raw(jen.Qual(pkg, name))} }
func (t Type) Pointer() Generator                 { return Raw(jen.Op("*").Add(t.Generate())) }

func (t Type) TypeParam(g Generator) Value {
	return Value{Raw(t.Generate().Index(g.Generate()))}
}

func (t Type) CreateStruct(values ...Generator) Value {
	return Value{Raw(t.Generate().Values(ToJenCodes(values)...))}
}
