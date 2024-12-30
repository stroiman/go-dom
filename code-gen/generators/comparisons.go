package generators

import "github.com/dave/jennifer/jen"

type Gte struct {
	Lhs Generator
	Rhs Generator
}

func (gte Gte) Generate() *jen.Statement {
	return gte.Lhs.Generate().Op(">=").Add(gte.Rhs.Generate())
}
