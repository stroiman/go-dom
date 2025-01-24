package generators

import (
	"slices"

	"github.com/dave/jennifer/jen"
)

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

func Reassign(id Generator, expression Generator) Generator {
	return Raw(id.Generate().Op("=").Add(expression.Generate()))
}

func ReassignMany(ids []Generator, expression Generator) Generator {
	var gs []jen.Code
	for _, id := range ids {
		gs = append(gs, id.Generate())
	}
	return Raw(jen.List(gs...).Op("=").Add(expression.Generate()))
}

func List(generators ...Generator) []Generator { return generators }

var Noop = GeneratorFunc(func() *jen.Statement { return nil })
var Nil Generator = Raw(jen.Nil())
var Line Generator = Raw(jen.Line())

func Lit(value any) Generator { return Raw(jen.Lit(value)) }

// WrapLine creates a generator that places a new line in front of the generated
// code. Useful for function function calls or struct initialisation using many
// members.
func WrapLine(g Generator) Generator {
	return GeneratorFunc(func() *jen.Statement {
		return jen.Line().Add(g.Generate())
	})
}

type StatementListStmt struct {
	Statements []Generator
}

func StatementList(statements ...Generator) StatementListStmt {
	return StatementListStmt{statements}
}

func (s *StatementListStmt) Prepend(stmt Generator) {
	s.Statements = slices.Insert(s.Statements, 0, stmt)
}

func (s *StatementListStmt) Append(stmt ...Generator) {
	s.Statements = append(s.Statements, stmt...)
}

func (s StatementListStmt) Generate() *jen.Statement {
	result := []jen.Code{}
	for _, s := range s.Statements {
		jenStatement := s.Generate()
		if jenStatement != nil && len(*jenStatement) != 0 {
			if len(result) != 0 {
				result = append(result, jen.Line())
			}
			result = append(result, jenStatement)
		}
	}
	jenStmt := jen.Statement(result)
	return &jenStmt
}

type IfStmt struct {
	Condition Generator
	Block     Generator
	Else      Generator
}

func (s IfStmt) Generate() *jen.Statement {
	result := jen.If(s.Condition.Generate()).Block(s.Block.Generate())
	if s.Else != nil {
		result.Else().Block(s.Else.Generate())
	}
	return result
}

func ToJenCodes(gg []Generator) (res []jen.Code) {
	for _, g := range gg {
		res = append(res, g.Generate())
	}
	return
}
