package wrappers

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	g "github.com/stroiman/go-dom/code-gen/generators"
)

type GojaNamingStrategy struct {
	ESConstructorData
}

func (s GojaNamingStrategy) PrototypeWrapperBaseName() string {
	return fmt.Sprintf("%sWrapper", s.Name())
}

func (s GojaNamingStrategy) PrototypeWrapperTypeName() string {
	return lowerCaseFirstLetter(s.PrototypeWrapperBaseName())
}

func (s GojaNamingStrategy) PrototypeWrapperConstructorName() string {
	return fmt.Sprintf("new%s", s.PrototypeWrapperBaseName())
}

func (s GojaNamingStrategy) ReceiverName() string {
	return "w" // data.Receiver
}

type GojaTargetGenerators struct{}

func (gen GojaTargetGenerators) CreateJSConstructorGenerator(data ESConstructorData) g.Generator {
	generator := g.StatementList()

	generator = g.StatementList(
		gen.CreateWrapperStruct(data),
		gen.CreateWrapperMethods(data),
		generator,
	)
	return generator
}

func (gen GojaTargetGenerators) CreateWrapperStruct(data ESConstructorData) g.Generator {
	naming := GojaNamingStrategy{data}
	typeName := naming.PrototypeWrapperTypeName()
	constructorName := naming.PrototypeWrapperConstructorName()
	innerType := g.Raw(jen.Qual(dom, data.Name()))

	wrapperStruct := g.NewStruct(typeName)
	wrapperStruct.Embed(g.Raw(jen.Id("baseInstanceWrapper").Index(innerType)))

	wrapperConstructor := g.FunctionDefinition{
		Name:     constructorName,
		Args:     g.Arg(g.Id("instance"), g.NewType("GojaContext").Pointer()),
		RtnTypes: g.List(g.NewType("wrapper")),
		Body: g.Return(g.InstantiateStruct(g.Id(typeName),
			g.NewValue("newBaseInstanceWrapper").TypeParam(innerType).Call(g.Id("instance")),
		)),
	}

	return g.StatementList(wrapperStruct, wrapperConstructor)
}

func (gen GojaTargetGenerators) CreateWrapperMethods(data ESConstructorData) g.Generator {
	list := g.StatementList()
	for op := range data.WrapperFunctionsToGenerate() {
		list.Append(gen.CreateWrapperMethod(data, op))
	}
	return list
}

var (
	gojaFc    = g.Raw(jen.Qual(gojaSrc, "FunctionCall"))
	gojaValue = g.Raw(jen.Qual(gojaSrc, "Value"))
)

func (gen GojaTargetGenerators) CreateWrapperMethod(
	data ESConstructorData,
	op ESOperation,
) g.Generator {
	naming := GojaNamingStrategy{data}
	callArgument := g.Id("c")
	return g.StatementList(
		g.Line,
		g.FunctionDefinition{
			Receiver: g.FunctionArgument{
				Name: g.Id(naming.ReceiverName()),
				Type: g.Id(naming.PrototypeWrapperTypeName()),
			},
			Name:     idlNameToGoName(op.Name),
			Args:     g.Arg(callArgument, gojaFc),
			RtnTypes: g.List(gojaValue),
			Body:     gen.CreateWrapperMethodBody(data, op, callArgument),
		})
}

func (gen GojaTargetGenerators) CreateWrapperMethodBody(
	data ESConstructorData,
	op ESOperation,
	callArgument g.Generator,
) g.Generator {
	if op.NotImplemented {
		msg := fmt.Sprintf(
			"%s.%s: Not implemented. Create an issue: %s", data.Name(), op.Name, ISSUE_URL,
		)
		return g.Raw(jen.Panic(jen.Lit(fmt.Sprintf(msg))))
	}
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	// result := g.NewValue("result")
	instance := g.NewValue("instance")
	// converter := g.NewValue(fmt.Sprintf("to%s", op.RetType.TypeName))
	readArgs := g.StatementList()
	argNames := make([]g.Generator, len(op.Arguments))
	for i, a := range op.Arguments {
		argNames[i] = g.Id(a.Name)
		value := g.Raw(callArgument.Generate().Dot("Arguments").Index(jen.Lit(i)))
		converter := fmt.Sprintf("decode%s", a.Type)
		readArgs.Append(g.Assign(argNames[i], receiver.Field(converter).Call(value)))
	}
	list := g.StatementList(
		g.Assign(instance, receiver.Field("getInstance").Call(callArgument)),
		readArgs,
	)
	if op.HasResult() {
		converter := fmt.Sprintf("to%s", upperCaseFirstLetter(op.RetType.TypeName))
		if op.GetHasError() {
			list.Append(
				g.AssignMany(g.List(
					g.Id("result"), g.Id("err")),
					instance.Field(upperCaseFirstLetter(op.Name)).Call(argNames...),
				),
				panicOnNotNil(g.Id("err")),
			)
		} else {
			list.Append(
				g.Assign(
					g.Id("result"),
					instance.Field(upperCaseFirstLetter(op.Name)).Call(argNames...),
				),
			)
		}
		list.Append(g.Return(receiver.Field(converter).Call(g.Id("result"))))
	} else {
		if op.GetHasError() {
			list.Append(
				g.Assign(g.Id("err"), instance.Field(upperCaseFirstLetter(op.Name)).Call(argNames...)),
				panicOnNotNil(g.Id("err")),
			)

		} else {
			list.Append(instance.Field(upperCaseFirstLetter(op.Name)).Call(argNames...))
		}
	}
	return list
}

func panicOnNotNil(lhs g.Generator) g.Generator {
	return g.IfStmt{
		Condition: g.Neq{Lhs: lhs, Rhs: g.Nil},
		Block:     g.Raw(jen.Panic(jen.Id("err"))),
	}
}
