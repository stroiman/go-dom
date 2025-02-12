package wrappers

import (
	"fmt"

	g "github.com/gost-dom/generators"

	"github.com/dave/jennifer/jen"
)

var scriptHost = g.NewValue("scriptHost")

type NewV8FunctionTemplate struct {
	iso JenGenerator
	f   JenGenerator
}

type V8NamingStrategy struct {
	ESConstructorData
}

func (s V8NamingStrategy) Receiver() string { return "w" }
func (s V8NamingStrategy) PrototypeWrapperBaseName() string {
	return fmt.Sprintf("%sV8Wrapper", s.Name())
}

func (s V8NamingStrategy) PrototypeWrapperName() string {
	return lowerCaseFirstLetter(s.PrototypeWrapperBaseName())
}

func (t NewV8FunctionTemplate) Generate() *jen.Statement {
	return jen.Qual(v8, "NewFunctionTemplateWithError").Call(t.iso.Generate(), t.f.Generate())
}

type V8TargetGenerators struct{}

func (_ V8TargetGenerators) CreateJSConstructorGenerator(data ESConstructorData) g.Generator {
	generator := g.StatementList()

	if data.Spec.WrapperStruct {
		generator.Append(CreateV8WrapperTypeGenerator(data))
	}

	generator.Append(
		CreateInitFunction(data),
		g.Line,
		CreateV8Constructor(data),
		CreateV8PrototypeInitialiser(data),
		CreateV8ConstructorWrapper(data),
		CreateV8WrapperMethods(data),
	)

	return generator
}

func CreateInitFunction(data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name: "init",
		Body: g.NewValue("registerJSClass").Call(
			g.Lit(data.Spec.TypeName),
			g.Lit(data.Inheritance),
			g.Id(prototypeFactoryFunctionName(data))),
	}
}

func CreateV8WrapperTypeGenerator(data ESConstructorData) g.Generator {
	typeNameBase := fmt.Sprintf("%sV8Wrapper", data.Name())
	typeName := g.NewType(lowerCaseFirstLetter(typeNameBase))
	constructorName := fmt.Sprintf("new%s", typeNameBase)
	innerType := g.NewTypePackage(data.Name(), data.GetInternalPackage())
	wrapperStruct := g.NewStruct(typeName)
	wrapperStruct.Embed(g.NewType("nodeV8WrapperBase").TypeParam(innerType))

	wrapperConstructor := g.FunctionDefinition{
		Name:     constructorName,
		Args:     g.Arg(scriptHost, scriptHostPtr),
		RtnTypes: g.List(typeName.Pointer()),
		Body: g.Return(
			typeName.CreateInstance(
				g.NewValue("newNodeV8WrapperBase").
					TypeParam(innerType).
					Call(scriptHost),
			).Reference(),
		),
	}

	return g.StatementList(wrapperStruct, wrapperConstructor, g.Line)
}

func CreateV8PrototypeInitialiser(data ESConstructorData) JenGenerator {
	naming := V8NamingStrategy{data}
	builder := NewConstructorBuilder()
	receiver := g.NewValue(naming.Receiver())
	installer := PrototypeInstaller{
		builder.v8Iso,
		builder.Proto,
		WrapperInstance{g.Value{Generator: receiver}},
	}
	return g.FunctionDefinition{
		Name: "installPrototype",
		Receiver: g.FunctionArgument{
			Name: receiver,
			Type: g.Id(naming.PrototypeWrapperName()),
		},
		Args: g.Arg(builder.Proto, v8ObjectTemplatePtr),
		Body: g.StatementList(
			g.Assign(g.NewValue("iso"), receiver.Field("scriptHost").Field("iso")),
			installer.InstallFunctionHandlers(data),
			installer.InstallAttributeHandlers(data),
		),
	}
}

func CreateV8ConstructorWrapper(data ESConstructorData) JenGenerator {
	naming := V8NamingStrategy{data}
	var body g.Generator
	if IsNodeType(data.IdlInterfaceName) {
		body = CreateV8IllegalConstructorBody(data)
	} else {
		body = CreateV8ConstructorWrapperBody(data)
	}
	return g.StatementList(
		g.Line,
		g.FunctionDefinition{
			Name: "Constructor",
			Receiver: g.FunctionArgument{
				Name: g.Id(naming.Receiver()),
				Type: g.Id(naming.PrototypeWrapperName()),
			},
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     body,
		},
	)
}

func CreateV8WrapperMethods(data ESConstructorData) JenGenerator {
	list := g.StatementList()
	for op := range data.WrapperFunctionsToGenerate() {
		list.Append(CreateV8WrapperMethod(data, op))
	}
	return list
}

func CreateV8WrapperMethod(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	naming := V8NamingStrategy{data}
	return g.StatementList(
		g.Line,
		g.FunctionDefinition{
			Receiver: g.FunctionArgument{
				Name: g.Id(naming.Receiver()),
				Type: g.Id(naming.PrototypeWrapperName()),
			},
			Name:     op.WrapperMethodName(),
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     CreateV8FunctionTemplateCallbackBody(data, op),
		})
}

func CreateV8FunctionTemplateCallbackBody(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	naming := V8NamingStrategy{data}
	debug := g.NewValuePackage("Debug", log).Call(
		g.Lit(fmt.Sprintf("V8 Function call: %s.%s", data.Name(), op.Name)))
	if op.NotImplemented {
		errMsg := fmt.Sprintf(
			"%s.%s: Not implemented. Create an issue: %s", data.Name(), op.Name, ISSUE_URL,
		)
		return g.StatementList(
			debug,
			g.Return(g.Nil, g.Raw(jen.Qual("errors", "New").Call(jen.Lit(errMsg)))))
	}
	receiver := WrapperInstance{g.NewValue(naming.Receiver())}
	instance := g.NewValue("instance")
	readArgsResult := ReadArguments(data, op)
	err := g.Id("err0")
	if len(op.Arguments) == 0 {
		err = g.Id("err")
	}
	requireContext := false
	var CreateCall = func(functionName string, argnames []g.Generator, op ESOperation) g.Generator {
		callInstance := V8InstanceInvocation{
			Name:     functionName,
			Args:     argnames,
			Op:       op,
			Instance: &instance,
			Receiver: receiver,
		}.GetGenerator()
		requireContext = requireContext || callInstance.RequireContext
		return callInstance.Generator
	}
	statements := g.StatementList(
		debug,
		AssignArgs(data, op),
		GetInstanceAndError(instance, err, data),
		readArgsResult,
		CreateV8WrapperMethodInstanceInvocations(
			data,
			op,
			idlNameToGoName(op.Name),
			readArgsResult.Args,
			err,
			CreateCall,
			true,
		),
	)
	if requireContext {
		statements.Prepend(V8RequireContext(receiver))
	}
	return statements
}

func prototypeFactoryFunctionName(data ESConstructorData) string {
	return fmt.Sprintf("create%sPrototype", data.IdlInterfaceName)
}

func CreateV8Constructor(data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name:     prototypeFactoryFunctionName(data),
		Args:     g.Arg(scriptHost, scriptHostPtr),
		RtnTypes: g.List(v8FunctionTemplatePtr),
		Body:     CreateV8ConstructorBody(data),
	}
}

func CreateV8ConstructorBody(data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	builder := NewConstructorBuilder()
	constructor := v8FunctionTemplate{g.NewValue("constructor")}

	createWrapperFunction := g.NewValue(fmt.Sprintf("new%s", naming.PrototypeWrapperBaseName()))

	statements := g.StatementList(
		builder.v8Iso.Assign(scriptHost.Field("iso")),
		g.Assign(builder.Wrapper, createWrapperFunction.Call(scriptHost)),
		g.Assign(constructor, builder.NewFunctionTemplateOfWrappedMethod("Constructor")),
		g.Line,
		g.Assign(builder.InstanceTmpl, constructor.GetInstanceTemplate()),
		builder.InstanceTmpl.SetInternalFieldCount(1),
		g.Line,
		// g.Assign(builder.Proto, constructor.GetPrototypeTemplate()),
		builder.Wrapper.Field("installPrototype").Call(constructor.GetPrototypeTemplate()),
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
	naming := V8NamingStrategy{data}
	receiver := WrapperInstance{g.NewValue(naming.Receiver())}
	if data.Constructor == nil {
		return CreateV8IllegalConstructorBody(data)
	}
	var readArgsResult V8ReadArguments
	op := *data.Constructor
	readArgsResult = ReadArguments(data, op)
	statements := g.StatementList(
		AssignArgs(data, op),
		readArgsResult)
	statements.Append(V8RequireContext(receiver))
	baseFunctionName := "CreateInstance"
	var CreateCall = func(functionName string, argnames []g.Generator, op ESOperation) g.Generator {
		return g.StatementList(
			g.Return(
				g.Raw(jen.Id(naming.Receiver()).Dot(functionName).CallFunc(func(grp *jen.Group) {
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
			data,
			op,
			baseFunctionName,
			readArgsResult.Args,
			nil,
			CreateCall,
			false,
		),
	)
	return statements
}

func CreateV8WrapperMethodInstanceInvocations(
	prototype ESConstructorData,
	op ESOperation,
	baseFunctionName string,
	args []V8ReadArg,
	instanceErr g.Generator,
	createCallInstance func(string, []g.Generator, ESOperation) g.Generator,
	extraError bool,
) g.Generator {
	// arguments := op.Arguments
	statements := g.StatementList()
	missingArgsConts := fmt.Sprintf("%s.%s: Missing arguments", prototype.Name(), op.Name)
	for i := len(args); i >= 0; i-- {
		functionName := baseFunctionName
		for j, arg := range args {
			if j < i {
				if arg.Argument.OptionalInGo() {
					functionName += idlNameToGoName(arg.Argument.Name)
				}
			}
		}
		currentArgs := args[0:i]
		ei := i
		if extraError {
			ei++
		}
		errNames := make([]g.Generator, 0, i+1)
		if instanceErr != nil {
			errNames = append(errNames, instanceErr)
		}
		for _, a := range currentArgs {
			errNames = append(errNames, a.ErrName)
		}

		callArgs := make([]g.Generator, i)
		for idx, a := range currentArgs {
			callArgs[idx] = a.ArgName
		}
		callInstance := createCallInstance(functionName, callArgs, op)
		if i > 0 {
			arg := args[i-1].Argument
			statements.Append(g.StatementList(
				g.IfStmt{
					Condition: g.Raw(jen.Id("args").Dot("noOfReadArguments").Op(">=").Lit(i)),
					Block: g.StatementList(
						ReturnOnAnyError(errNames),
						callInstance,
					),
				}))
			if !arg.OptionalInGo() {
				statements.Append(
					g.Return(
						g.Nil,
						g.Raw(jen.Qual("errors", "New").Call(jen.Lit(missingArgsConts))),
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
		wrapper.MustGetContext(info),
		// wrapper.GetScriptHost().Method("mustGetContext").Call(info.GetV8Context()),
	)
}

type V8InstanceInvocation struct {
	Name     string
	Args     []g.Generator
	Op       ESOperation
	Instance *g.Value
	Receiver WrapperInstance
}

type V8InstanceInvocationResult struct {
	Generator      g.Generator
	HasValue       bool
	HasError       bool
	RequireContext bool
}

func (c V8InstanceInvocation) PerformCall() (genRes V8InstanceInvocationResult) {
	args := []g.Generator{}
	genRes.HasError = c.Op.GetHasError()
	genRes.HasValue = c.Op.HasResult() // != "undefined"
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
	list := g.StatementListStmt{}
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

func (c V8InstanceInvocation) GetGenerator() V8InstanceInvocationResult {
	genRes := c.PerformCall()
	list := g.StatementList()
	list.Append(genRes.Generator)
	if !genRes.HasValue {
		if genRes.HasError {
			list.Append(g.Return(g.Nil, g.Id("callErr")))
		} else {
			list.Append(g.Return(g.Nil, g.Nil))
		}
	} else {
		retType := c.Op.RetType
		if retType.IsNode() {
			genRes.RequireContext = true
			valueReturn := (g.Return(g.Raw(jen.Id("ctx").Dot("getInstanceForNode").Call(jen.Id("result")))))
			if genRes.HasError {
				list.Append(g.IfStmt{
					Condition: g.Neq{Lhs: g.Id("callErr"), Rhs: g.Nil},
					Block:     g.Return(g.Nil, g.Id("callErr")),
					Else:      valueReturn,
				})
			} else {
				list.Append(valueReturn)
			}
		} else {
			converter := c.Op.Encoder()
			genRes.RequireContext = true
			valueReturn := g.Return(c.Receiver.Method(converter).Call(g.Id("ctx"), g.Id("result")))
			if genRes.HasError {
				list.Append(g.IfStmt{
					Condition: g.Neq{Lhs: g.Id("callErr"), Rhs: g.Nil},
					Block:     g.Return(g.Nil, g.Id("callErr")),
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
	naming := V8NamingStrategy{data}
	return g.Return(g.Nil, g.NewValuePackage("NewTypeError", v8).
		Call(g.NewValue(naming.Receiver()).Field("scriptHost").Field("iso"),
			g.Lit("Illegal Constructor")))
}

type V8ReadArg struct {
	Argument ESOperationArgument
	ArgName  g.Generator
	ErrName  g.Generator
	Index    int
}

type V8ReadArguments struct {
	Args      []V8ReadArg
	Generator g.Generator
}

func (r V8ReadArguments) Generate() *jen.Statement {
	if r.Generator != nil {
		return r.Generator.Generate()
	} else {
		return g.Noop.Generate()
	}
}

func AssignArgs(data ESConstructorData, op ESOperation) g.Generator {
	if len(op.Arguments) == 0 {
		return g.Noop
	}
	naming := V8NamingStrategy{data}
	return g.Assign(
		g.Id("args"),
		g.NewValue("newArgumentHelper").Call(
			g.NewValue(naming.Receiver()).Field("scriptHost"),
			g.Id("info")),
	)
}

func ReadArguments(data ESConstructorData, op ESOperation) (res V8ReadArguments) {
	naming := V8NamingStrategy{data}
	argCount := len(op.Arguments)
	res.Args = make([]V8ReadArg, 0, argCount)
	statements := g.StatementList()
	for i, arg := range op.Arguments {
		argName := g.Id(sanitizeVarName(arg.Name))
		errName := g.Id(fmt.Sprintf("err%d", i+1))
		if arg.Ignore {
			continue
		}
		res.Args = append(res.Args, V8ReadArg{
			Argument: arg,
			ArgName:  argName,
			ErrName:  errName,
			Index:    i,
		})

		var convertNames []string
		if arg.Type != "" {
			convertNames = []string{fmt.Sprintf("decode%s", idlNameToGoName(arg.Type))}
		} else {
			types := arg.IdlType.IdlType.IType.Types
			convertNames = make([]string, len(types))
			for i, t := range types {
				convertNames[i] = fmt.Sprintf("decode%s", t.IType.TypeName)
			}
		}

		gConverters := []g.Generator{g.Id("args"), g.Lit(i)}
		defaultName, hasDefault := arg.DefaultValueInGo()
		if hasDefault {
			gConverters = append(gConverters, g.NewValue(naming.Receiver()).Field(defaultName))
		}
		for _, n := range convertNames {
			gConverters = append(gConverters, g.NewValue(naming.Receiver()).Field(n))
		}
		if hasDefault {
			statements.Append(g.AssignMany(g.List(argName, errName),
				g.NewValue("tryParseArgWithDefault").Call(gConverters...)))
		} else {
			statements.Append(g.AssignMany(
				g.List(argName, errName),
				g.NewValue("tryParseArg").Call(gConverters...)))
		}
	}
	res.Generator = statements
	return
}

func GetInstanceAndError(id g.Generator, errId g.Generator, data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	return g.AssignMany(
		g.List(id, errId),
		g.NewValue(naming.Receiver()).Field("getInstance").Call(g.Id("info")),
	)
}
