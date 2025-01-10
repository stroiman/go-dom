package wrappers

import (
	"fmt"

	g "github.com/stroiman/go-dom/code-gen/generators"

	"github.com/dave/jennifer/jen"
)

type NewV8FunctionTemplate struct {
	iso JenGenerator
	f   JenGenerator
}

func (t NewV8FunctionTemplate) Generate() *jen.Statement {
	return jen.Qual(v8, "NewFunctionTemplateWithError").Call(t.iso.Generate(), t.f.Generate())
}

func CreateV8Generator(data ESConstructorData) g.Generator {
	generator := StatementList(
		CreateV8Constructor(data),
		CreateV8ConstructorWrapper(data),
		CreateV8WrapperMethods(data),
	)

	if data.Spec.WrapperStruct {
		generator = StatementList(
			CreateV8WrapperTypeGenerator(data),
			generator,
		)
	}

	return generator
}

func CreateV8WrapperTypeGenerator(data ESConstructorData) g.Generator {
	typeName := fmt.Sprintf("%sV8Wrapper", data.Name())
	constructorName := fmt.Sprintf("New%s", typeName)
	innerType := g.Raw(jen.Qual(html, data.Name()))
	wrapperStruct := g.NewStruct(typeName)
	wrapperStruct.Embed(g.Raw(jen.Id("NodeV8WrapperBase").Index(innerType)))

	wrapperConstructor := g.FunctionDefinition{
		Name:     constructorName,
		Args:     g.Arg(g.Id("host"), g.NewType("ScriptHost").Pointer()),
		RtnTypes: g.List(g.NewType(typeName).Pointer()),
		Body: g.Return(g.Raw(
			jen.Op("&").Id(typeName).Values(
				jen.Id("NewNodeV8WrapperBase").Index(innerType.Generate()).Call(jen.Id("host")),
			),
		)),
	}

	return StatementList(wrapperStruct, wrapperConstructor, g.Line)
}

func CreateV8ConstructorWrapper(data ESConstructorData) JenGenerator {
	var body g.Generator
	if IsNodeType(data.InnerTypeName) {
		body = CreateV8IllegalConstructorBody(data)
	} else {
		body = CreateV8ConstructorWrapperBody(data)
	}
	return StatementList(
		g.Line,
		g.FunctionDefinition{
			Name: "Constructor",
			Receiver: g.FunctionArgument{
				Name: g.Id(data.Receiver),
				Type: g.Id(data.WrapperTypeName),
			},
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     body,
		},
	)
}

func CreateV8WrapperMethods(data ESConstructorData) JenGenerator {
	list := StatementList()
	for _, op := range data.Operations {
		if op.MethodCustomization.Ignored {
			continue
		}
		list.Append(CreateV8WrapperMethod(data, op))
	}
	for _, attr := range data.Attributes {
		if attr.Getter != nil {
			list.Append(CreateV8WrapperMethod(data, *attr.Getter))
		}
		if attr.Setter != nil {
			list.Append(CreateV8WrapperMethod(data, *attr.Setter))
		}
	}
	return list
}

func CreateV8WrapperMethod(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	if op.CustomImplementation {
		return g.Noop
	}
	return StatementList(
		g.Line,
		g.FunctionDefinition{
			Receiver: g.FunctionArgument{
				Name: g.Id(data.Receiver),
				Type: g.Id(data.WrapperTypeName),
			},
			Name:     idlNameToGoName(op.Name),
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     CreateV8FunctionTemplateCallbackBody(data, op),
		})
}

func CreateV8FunctionTemplateCallbackBody(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	if op.NotImplemented {
		errMsg := fmt.Sprintf("Not implemented: %s.%s", data.Name(), op.Name)
		return g.Return(g.Nil, g.Raw(jen.Qual("errors", "New").Call(jen.Lit(errMsg))))
	}
	receiver := WrapperInstance{g.NewValue(data.Receiver)}
	instance := g.NewValue("instance")
	readArgsResult := ReadArguments(data, op)
	err := g.Id("err0")
	if len(op.Arguments) == 0 {
		err = g.Id("err")
	}
	requireContext := false
	var CreateCall = func(functionName string, argnames []g.Generator, op ESOperation) g.Generator {
		callInstance := CallInstance{
			Name:     functionName,
			Args:     argnames,
			Op:       op,
			Instance: &instance,
			Receiver: receiver,
		}.GetGenerator()
		requireContext = requireContext || callInstance.RequireContext
		return callInstance.Generator
	}
	errNames := make([]g.Generator, len(readArgsResult.ErrNames)+1)
	errNames[0] = err
	copy(errNames[1:], readArgsResult.ErrNames)
	statements := StatementList(
		AssignArgs(data, op),
		GetInstanceAndError(instance, err, data),
		readArgsResult,
		CreateV8WrapperMethodInstanceInvocations(
			op,
			idlNameToGoName(op.Name),
			readArgsResult.ArgNames,
			errNames,
			CreateCall,
			true,
		),
	)
	if requireContext {
		statements.Prepend(V8RequireContext(receiver))
	}
	return statements
}

func CreateV8Constructor(data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name:     fmt.Sprintf("Create%sPrototype", data.InnerTypeName),
		Args:     g.Arg(g.Id("host"), scriptHostPtr),
		RtnTypes: g.List(v8FunctionTemplatePtr),
		Body:     CreateV8ConstructorBody(data),
	}
}

func CreateV8ConstructorBody(data ESConstructorData) g.Generator {
	builder := NewConstructorBuilder()
	scriptHost := g.NewValue("host")
	constructor := v8FunctionTemplate{g.NewValue("constructor")}

	createWrapperFunction := g.NewValue(fmt.Sprintf("New%s", data.WrapperTypeName))

	statements := StatementList(
		builder.v8Iso.Assign(scriptHost.Field("iso")),
		g.Assign(builder.Wrapper, createWrapperFunction.Call(scriptHost)),
		g.Assign(constructor, builder.NewFunctionTemplateOfWrappedMethod("Constructor")),
		g.Line,
		g.Assign(builder.InstanceTmpl, constructor.GetInstanceTemplate()),
		builder.InstanceTmpl.SetInternalFieldCount(1),
		g.Line,
		g.Assign(builder.Proto, constructor.GetPrototypeTemplate()),
		builder.InstallFunctionHandlers(data),
		builder.InstallAttributeHandlers(data),
		g.Line,
	)
	if data.RunCustomCode {
		statements.Append(
			g.Raw(jen.Id("wrapper").Dot("CustomInitialiser").Call(jen.Id("constructor"))),
		)
	}
	statements.Append(g.Return(constructor))
	return statements
}

func CreateV8ConstructorWrapperBody(data ESConstructorData) g.Generator {
	receiver := WrapperInstance{g.NewValue(data.Receiver)}
	if data.Constructor == nil {
		return CreateV8IllegalConstructorBody(data)
	}
	var readArgsResult V8ReadArguments
	op := *data.Constructor
	readArgsResult = ReadArguments(data, op)
	statements := StatementList(
		AssignArgs(data, op),
		readArgsResult)
	statements.Append(V8RequireContext(receiver))
	baseFunctionName := "CreateInstance"
	var CreateCall = func(functionName string, argnames []g.Generator, op ESOperation) g.Generator {
		return StatementList(
			g.Return(
				g.Raw(jen.Id(data.Receiver).Dot(functionName).CallFunc(func(grp *jen.Group) {
					grp.Add(jen.Id("ctx"))
					grp.Add(jen.Id("info").Dot("This").Call())
					for _, name := range argnames {
						grp.Add(name.Generate())
					}
				})),
			),
		)
	}
	statements.Append(
		CreateV8WrapperMethodInstanceInvocations(
			op,
			baseFunctionName,
			readArgsResult.ArgNames,
			readArgsResult.ErrNames,
			CreateCall,
			false,
		),
	)
	return statements
}

func CreateV8WrapperMethodInstanceInvocations(
	op ESOperation,
	baseFunctionName string,
	argNames []g.Generator,
	errorsNames []g.Generator,
	createCallInstance func(string, []g.Generator, ESOperation) g.Generator,
	extraError bool,
) g.Generator {
	arguments := op.Arguments
	statements := StatementList()
	for i := len(arguments); i >= 0; i-- {
		functionName := baseFunctionName
		for j, arg := range arguments {
			if j < i {
				if arg.Optional {
					functionName += idlNameToGoName(arg.Name)
				}
			}
		}
		argnames := argNames[0:i]
		ei := i
		if extraError {
			ei++
		}
		errNames := errorsNames[0:ei]
		callInstance := createCallInstance(functionName, argnames, op)
		if i > 0 {
			arg := arguments[i-1]
			statements.Append(StatementList(
				IfStmt{
					Condition: g.Raw(jen.Id("args").Dot("noOfReadArguments").Op(">=").Lit(i)),
					Block: StatementList(
						ReturnOnAnyError(errNames),
						callInstance,
					),
				}))
			if !(arg.Optional) {
				statements.Append(
					g.Return(
						g.Nil,
						g.Raw(jen.Qual("errors", "New").Call(jen.Lit("Missing arguments"))),
					),
				)
				break
			}
		} else {
			statements.Append(ReturnOnAnyError(errNames))
			statements.Append(callInstance)
		}
	}
	return statements
}

func V8RequireContext(wrapper WrapperInstance) g.Generator {
	info := v8ArgInfo(g.NewValue("info"))
	return g.Assign(
		g.Id("ctx"),
		wrapper.GetScriptHost().Method("MustGetContext").Call(info.GetV8Context()),
	)
}

type CallInstance struct {
	Name     string
	Args     []g.Generator
	Op       ESOperation
	Instance *g.Value
	Receiver WrapperInstance
}

type GetGeneratorResult struct {
	Generator      g.Generator
	HasValue       bool
	HasError       bool
	RequireContext bool
}

func (c CallInstance) PerformCall() (genRes GetGeneratorResult) {
	args := []g.Generator{}
	genRes.HasError = c.Op.GetHasError()
	genRes.HasValue = c.Op.RetType.IsDefined() // != "undefined"
	var stmt *jen.Statement
	if genRes.HasValue {
		stmt = jen.Id("result")
	}
	if genRes.HasError {
		if stmt != nil {
			stmt = stmt.Op(",").Id("callErr")
		} else {
			stmt = jen.Id("callErr")
		}
	}
	if stmt != nil {
		stmt = stmt.Op(":=")
	}

	for _, a := range c.Args {
		args = append(args, a)
	}
	list := StatementListStmt{}
	var evaluation g.Value
	if c.Instance == nil {
		evaluation = g.NewValue(idlNameToGoName(c.Name)).Call(args...)
	} else {
		evaluation = c.Instance.Method(idlNameToGoName(c.Name)).Call(args...)
	}
	if stmt == nil {
		list.Append(evaluation)
	} else {
		list.Append(g.Raw(stmt.Add(evaluation.Generate())))
	}
	genRes.Generator = list
	return
}

func (c CallInstance) GetGenerator() GetGeneratorResult {
	genRes := c.PerformCall()
	list := StatementListStmt{}
	list.Append(genRes.Generator)
	if !genRes.HasValue {
		if genRes.HasError {
			list.Append(Stmt{jen.Return(jen.Nil(), jen.Id("callErr"))})
		} else {
			list.Append(Stmt{jen.Return(jen.Nil(), jen.Nil())})
		}
	} else {
		retType := c.Op.RetType
		if retType.IsNode() {
			genRes.RequireContext = true
			list.Append(g.Return(g.Raw(jen.Id("ctx").Dot("GetInstanceForNode").Call(jen.Id("result")))))
		} else {
			converter := "To"
			if retType.Nullable {
				converter += "Nullable"
			}
			converter += idlNameToGoName(retType.TypeName)
			genRes.RequireContext = true
			valueReturn := g.Return(c.Receiver.Method(converter).Call(g.Id("ctx"), g.Id("result")))
			if genRes.HasError {
				list.Append(IfStmt{
					Condition: Stmt{jen.Id("callErr").Op("!=").Nil()},
					Block:     Stmt{jen.Return(jen.Nil(), jen.Id("callErr"))},
					Else:      valueReturn,
				})
			} else {
				list.Append(valueReturn)
			}
		}
	}
	genRes.Generator = list
	return genRes
}

func CreateV8IllegalConstructorBody(data ESConstructorData) g.Generator {
	return g.Return(g.Nil,
		g.Raw(jen.Qual(v8, "NewTypeError").Call(
			jen.Id(data.Receiver).Dot("host").Dot("iso"), jen.Lit("Illegal Constructor"),
		)),
	)
}

type V8ReadArguments struct {
	ArgNames  []g.Generator
	ErrNames  []g.Generator
	Generator g.Generator
}

func (r V8ReadArguments) Generate() *jen.Statement {
	if r.Generator != nil {
		return r.Generator.Generate()
	} else {
		return g.Noop.Generate()
	}
}
