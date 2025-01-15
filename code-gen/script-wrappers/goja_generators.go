package wrappers

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	g "github.com/stroiman/go-dom/code-gen/generators"
)

type GojaTargetGenerators struct{}

func (gen GojaTargetGenerators) CreateJSConstructorGenerator(data ESConstructorData) g.Generator {
	generator := g.StatementList()

	generator = g.StatementList(
		gen.CreateWrapperStruct(data),
		generator,
	)
	return generator
}

func (gen GojaTargetGenerators) CreateWrapperStruct(data ESConstructorData) g.Generator {
	typeNameBase := fmt.Sprintf("%sWrapper", data.Name())
	typeName := lowerCaseFirstLetter(typeNameBase)
	constructorName := fmt.Sprintf("new%s", typeNameBase)
	innerType := g.Raw(jen.Qual(html, data.Name()))
	wrapperStruct := g.NewStruct(typeName)
	wrapperStruct.Embed(g.Raw(jen.Id("BaseInstanceWrapper").Index(innerType)))

	wrapperConstructor := g.FunctionDefinition{
		Name:     constructorName,
		Args:     g.Arg(g.Id("instance"), g.NewType("gojaInstance").Pointer()),
		RtnTypes: g.List(g.NewType(typeName).Pointer()),
		Body: g.Return(g.Raw(
			jen.Id(typeName).Values(
				jen.Id("newBaseInstanceWrapper").
					Index(innerType.Generate()).
					Call(jen.Id("instance")),
			),
		)),
	}

	return g.StatementList(wrapperStruct, wrapperConstructor)
}
