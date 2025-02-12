package wrappers

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
)

type Generator = generators.Generator

type TargetGenerators interface {
	CreateInitFunction(ESConstructorData) Generator
	CreateWrapperStruct(ESConstructorData) Generator
	CreateConstructor(ESConstructorData) Generator
	CreatePrototypeInitializer(ESConstructorData) Generator
	CreateConstructorWrapper(ESConstructorData) Generator
	CreateWrapperMethods(ESConstructorData) Generator
}

// PrototypeWrapperGenerator generates code to create a JavaScript prototype
// that wraps an internal Go type.
type PrototypeWrapperGenerator struct {
	Platform TargetGenerators
	Data     ESConstructorData
}

func (g PrototypeWrapperGenerator) Generate() *jen.Statement {
	list := generators.StatementList(
		g.Platform.CreateInitFunction(g.Data),
		generators.Line,
	)
	if !g.Data.Spec.SkipWrapper {
		list.Append(g.Platform.CreateWrapperStruct(g.Data))

	}
	list.Append(
		g.Platform.CreateConstructor(g.Data),
		g.Platform.CreatePrototypeInitializer(g.Data),
		g.Platform.CreateConstructorWrapper(g.Data),
		g.Platform.CreateWrapperMethods(g.Data),
	)

	return list.Generate()
}
