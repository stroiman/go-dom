package wrappers

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
)

type Generator = generators.Generator

type TargetGenerators interface {
	// CreateInitFunction generates an init function intended to register that a
	// class should be created. This doesn't _create_ the class, as that
	// requires a host created at runtime. So this is a registration that _when_
	// a host is created, this class must be added to global scope, optionally
	// with a subclass.
	CreateInitFunction(ESConstructorData) Generator
	// Create a struct definition, and it's constructor, that must contain the
	// methods acting as callback for prototype functions, including the
	// constructor.
	CreateWrapperStruct(ESConstructorData) Generator
	// CreateHostInitialiser creates the function that will register the class
	// in the host's global scope.
	CreateHostInitialiser(ESConstructorData) Generator
	// CreatePrototypeInitializer creates the "initializePrototype" method, which
	// sets all the properties on the prototypes on this class.
	CreatePrototypeInitializer(ESConstructorData) Generator
	// CreateConstructorCallback generates the function to be called whan
	// JavaScript code constructs an instance.
	CreateConstructorCallback(ESConstructorData) Generator
	CreateMethodCallback(ESConstructorData, ESOperation) Generator
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
		g.Platform.CreateHostInitialiser(g.Data),
		g.Platform.CreatePrototypeInitializer(g.Data),
		g.Platform.CreateConstructorCallback(g.Data),
		g.MethodCallbacks(g.Data),
	)

	return list.Generate()
}

func (g PrototypeWrapperGenerator) MethodCallbacks(data ESConstructorData) Generator {
	list := generators.StatementList()
	for op := range data.WrapperFunctionsToGenerate() {
		list.Append(g.Platform.CreateMethodCallback(data, op))
	}
	return list
}
