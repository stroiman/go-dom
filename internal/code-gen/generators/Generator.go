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

func Raw(stmt *jen.Statement) Generator { return RawStatement{stmt} }

func Return(exp Generator) Generator {
	return Raw(jen.Return(exp.Generate()))
}

func Id(id string) Generator {
	return Raw(jen.Id(id))
}

func Assign(ids Generator, expression Generator) Generator {
	return Raw(ids.Generate().Op(":=").Add(expression.Generate()))
}
