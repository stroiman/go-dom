package generators

import "github.com/dave/jennifer/jen"

type Gte struct {
	Lhs Generator
	Rhs Generator
}

func (gte Gte) Generate() *jen.Statement {
	return gte.Lhs.Generate().Op(">=").Add(gte.Rhs.Generate())
}

type Eq struct {
	Lhs Generator
	Rhs Generator
}

func (eq Eq) Generate() *jen.Statement {
	return eq.Lhs.Generate().Op("==").Add(eq.Rhs.Generate())
}

type Neq struct {
	Lhs Generator
	Rhs Generator
}

func (eq Neq) Generate() *jen.Statement {
	return eq.Lhs.Generate().Op("!=").Add(eq.Rhs.Generate())
}
