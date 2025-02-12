package wrappers

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
)

type Generator = generators.Generator

type TargetGenerators interface {
	CreateInitFunction(ESConstructorData) Generator
	CreateJSConstructorGenerator(ESConstructorData) Generator
}

// PrototypeWrapperGenerator generates code to create a JavaScript prototype
// that wraps an internal Go type.
type PrototypeWrapperGenerator struct {
	Platform TargetGenerators
	Data     ESConstructorData
}

func (g PrototypeWrapperGenerator) Generate() *jen.Statement {
	return generators.StatementList(
		g.Platform.CreateInitFunction(g.Data),
		generators.Line,
		g.Platform.CreateJSConstructorGenerator(g.Data),
	).Generate()
}
