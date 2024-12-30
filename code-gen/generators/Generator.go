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

func AssignMany(ids []Generator, expression Generator) Generator {
	var gs []jen.Code
	for _, id := range ids {
		gs = append(gs, id.Generate())
	}
	return Raw(jen.List(gs...).Op(":=").Add(expression.Generate()))
}

func ReAssignMany(ids []Generator, expression Generator) Generator {
	var gs []jen.Code
	for _, id := range ids {
		gs = append(gs, id.Generate())
	}
	return Raw(jen.List(gs...).Op("=").Add(expression.Generate()))
}

type Type struct {
	RawStatement
}

func NewType(name string) Type                    { return Type{Raw(jen.Id(name))} }
func NewTypePackage(name string, pkg string) Type { return Type{Raw(jen.Qual(pkg, name))} }
func (t Type) Pointer() Generator                 { return Raw(jen.Op("*").Add(t.Generate())) }

func List(generators ...Generator) []Generator { return generators }

var Noop = GeneratorFunc(func() *jen.Statement { return nil })
var Nil Generator = Raw(jen.Nil())
var Line Generator = Raw(jen.Line())

func Lit(value any) Generator { return Raw(jen.Lit(value)) }

type Value struct{ Generator }

func NewValue(name string) Value {
	return Value{Id(name)}
}

func (v Value) Field(name string) Value {
	return Value{Raw(v.Generate().Dot(name))}
}

func (v Value) Method(name string) Value {
	return Value{Raw(v.Generate().Dot(name))}
}

func (m Value) Call(args ...Generator) Value {
	return Value{Raw(m.Generate().CallFunc(func(g *jen.Group) {
		for _, arg := range args {
			g.Add(arg.Generate())
		}
	}))}
}
