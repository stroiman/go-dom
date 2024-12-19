package generators

import "github.com/dave/jennifer/jen"

type st = *jen.Statement

type Generator interface {
	Generate() *jen.Statement
}

type GeneratorFunc func() *jen.Statement

func (f GeneratorFunc) Generate() *jen.Statement { return f() }

type RawStatement struct{ *jen.Statement }

func (s RawStatement) Generate() *jen.Statement { return s.Statement.Clone() }

func Raw(stmt *jen.Statement) RawStatement { return RawStatement{stmt} }

func Return(exp ...Generator) Generator {
	var code []jen.Code
	for _, s := range exp {
		code = append(code, s.Generate())
	}
	return Raw(jen.Return(code...))
}

func Id(id string) Generator {
	return Raw(jen.Id(id))
}

func Assign(ids Generator, expression Generator) Generator {
	return Raw(ids.Generate().Op(":=").Add(expression.Generate()))
}

type Type struct {
	RawStatement
}

func NewType(name string) Type                    { return Type{Raw(jen.Id(name))} }
func NewTypePackage(name string, pkg string) Type { return Type{Raw(jen.Qual(pkg, name))} }
func (t Type) Pointer() Generator                 { return Raw(jen.Op("*").Add(t.Generate())) }

func List(generators ...Generator) []Generator { return generators }

var Line Generator = Raw(jen.Line())
