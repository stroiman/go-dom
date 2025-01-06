package wrappers

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"unicode"

	g "github.com/stroiman/go-dom/code-gen/generators"
	. "github.com/stroiman/go-dom/code-gen/idl"

	"github.com/dave/jennifer/jen"
)

var (
	v8FunctionTemplatePtr     = g.NewTypePackage("FunctionTemplate", v8).Pointer()
	v8FunctionCallbackInfoPtr = g.NewTypePackage("FunctionCallbackInfo", v8).Pointer()
	v8Value                   = g.NewTypePackage("Value", v8).Pointer()
	v8ReadOnly                = g.Raw(jen.Qual(v8, "ReadOnly"))
	v8None                    = g.Raw(jen.Qual(v8, "None"))
	scriptHostPtr             = g.NewType("ScriptHost").Pointer()
)

func createData(spec ParsedIdlFile, dataData ESClassWrapper) ESConstructorData {
	var constructor *ESOperation
	idlName := spec.IdlNames[dataData.TypeName]
	type tmp struct {
		Op ESOperation
		Ok bool
	}
	ops := []*tmp{}
	attributes := []ESAttribute{}
	for _, member := range idlName.Members {
		if member.Special == "static" {
			continue
		}
		if -1 != slices.IndexFunc(
			ops,
			func(op *tmp) bool { return op.Op.Name == member.Name },
		) {
			slog.Warn("Function overloads", "Name", member.Name)
			continue
		}
		returnType, nullable := FindMemberReturnType(member)
		methodCustomization := dataData.GetMethodCustomization(member.Name)
		operation := &tmp{ESOperation{
			Name:                 member.Name,
			NotImplemented:       methodCustomization.NotImplemented,
			CustomImplementation: methodCustomization.CustomImplementation,
			ReturnType:           returnType,
			Nullable:             nullable,
			MethodCustomization:  methodCustomization,
			Arguments:            []ESOperationArgument{},
		}, true}
		if member.Type == "operation" && member.Name != "" {
			// Empty name seems to indicate a named property getter. Not sure yet.
			operation.Op.HasError = !operation.Op.MethodCustomization.HasNoError
			ops = append(ops, operation)
		}
		for _, arg := range member.Arguments {
			esArg := ESOperationArgument{
				Name:     arg.Name,
				Optional: arg.Optional,
				IdlType:  arg.IdlType,
			}
			if len(arg.IdlType.Types) > 0 {
				slog.Warn(
					"Multiple argument types",
					"Operation",
					member.Name,
					"Argument",
					arg.Name,
				)
				operation.Ok = false
				break
			}
			if arg.IdlType.IdlType != nil {
				esArg.Type = arg.IdlType.IdlType.IType.TypeName
			}
			operation.Op.Arguments = append(operation.Op.Arguments, esArg)
		}
		if member.Type == "constructor" {
			constructor = &operation.Op
		}
		if IsAttribute(member) {
			op := operation.Op
			var (
				getter *ESOperation
				setter *ESOperation
			)
			rtnType, nullable := FindMemberAttributeType(member)
			getter = new(ESOperation)
			*getter = op
			getterName := idlNameToGoName(op.Name)
			if !member.Readonly {
				getterName = fmt.Sprintf("Get%s", getterName)
			}
			getter.Name = getterName
			getter.ReturnType = rtnType
			getter.Nullable = nullable
			getterCustomization := dataData.GetMethodCustomization(getter.Name)
			getter.NotImplemented = getterCustomization.NotImplemented || op.NotImplemented
			getter.CustomImplementation = getterCustomization.CustomImplementation ||
				op.CustomImplementation
			if !member.Readonly {
				setter = new(ESOperation)
				*setter = op
				setter.Name = fmt.Sprintf("Set%s", idlNameToGoName(op.Name))
				methodCustomization := dataData.GetMethodCustomization(setter.Name)
				setter.NotImplemented = methodCustomization.NotImplemented ||
					op.NotImplemented
				setter.CustomImplementation = methodCustomization.CustomImplementation ||
					op.CustomImplementation
				setter.ReturnType = "undefined"
				setter.Arguments = []ESOperationArgument{{
					Name:     "val",
					Type:     idlNameToGoName(rtnType),
					Optional: false,
					Variadic: false,
					// IdlType  IdlTypes
				}}
			}
			attributes = append(attributes, ESAttribute{op.Name, getter, setter})
		}
	}

	operations := make([]ESOperation, 0, len(ops))
	for _, op := range ops {
		if op.Ok {
			operations = append(operations, op.Op)
		}
	}
	wrappedTypeName := dataData.InnerTypeName
	if wrappedTypeName == "" {
		wrappedTypeName = idlName.Name
	}
	wrapperTypeName := dataData.WrapperTypeName
	if wrapperTypeName == "" {
		wrapperTypeName = "ES" + wrappedTypeName
	}
	return ESConstructorData{
		InnerTypeName:    wrappedTypeName,
		WrapperTypeName:  wrapperTypeName,
		Receiver:         dataData.Receiver,
		Operations:       operations,
		Attributes:       attributes,
		Constructor:      constructor,
		CreatesInnerType: true,
		IdlName:          idlName,
		RunCustomCode:    dataData.RunCustomCode,
	}
}

const br = "github.com/stroiman/go-dom/browser/dom"
const sc = "github.com/stroiman/go-dom/browser/scripting"
const v8 = "github.com/tommie/v8go"

type ESOperationArgument struct {
	Name     string
	Type     string
	Optional bool
	Variadic bool
	IdlType  IdlTypes
}

type ESOperation struct {
	Name                 string
	NotImplemented       bool
	ReturnType           string
	Nullable             bool
	HasError             bool
	CustomImplementation bool
	MethodCustomization  ESMethodWrapper
	Arguments            []ESOperationArgument
}

func (op ESOperation) GetHasError() bool {
	return op.HasError
}

type ESAttribute struct {
	Name   string
	Getter *ESOperation
	Setter *ESOperation
}

type ESConstructorData struct {
	CreatesInnerType bool
	InnerTypeName    string
	WrapperTypeName  string
	Receiver         string
	Operations       []ESOperation
	Attributes       []ESAttribute
	Constructor      *ESOperation
	RunCustomCode    bool
	IdlName
}

type Imports = [][][2]string

func IllegalConstructor(data ESConstructorData) g.Generator {
	return g.Return(g.Nil,
		g.Raw(jen.Qual(v8, "NewTypeError").Call(
			jen.Id(data.Receiver).Dot("host").Dot("iso"), jen.Lit("Illegal Constructor"),
		)),
	)
}

func ReturnOnAnyError(errNames []g.Generator) g.Generator {
	if len(errNames) == 0 {
		return g.Noop
	}
	if len(errNames) == 1 {
		return ReturnOnError{err: errNames[0]}
	}
	return Statements(
		g.Assign(g.Id("err"),
			g.Raw(jen.Qual("errors", "Join").CallFunc(func(g *jen.Group) {
				for _, e := range errNames {
					g.Add(e.Generate())
				}
			})),
		),
		ReturnOnError{},
	)
}

func WrapperCalls(
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

func RequireContext(wrapper WrapperInstance) g.Generator {
	info := v8ArgInfo(g.NewValue("info"))
	return g.Assign(
		g.Id("ctx"),
		wrapper.GetScriptHost().Method("MustGetContext").Call(info.GetV8Context()),
	)
}

func JSConstructorImpl(data ESConstructorData) g.Generator {
	receiver := WrapperInstance{g.NewValue(data.Receiver)}
	if data.Constructor == nil {
		return IllegalConstructor(data)
	}
	var readArgsResult ReadArgumentsResult
	op := *data.Constructor
	readArgsResult = ReadArguments(data, op)
	statements := StatementList(
		AssignArgs(data, op),
		readArgsResult)
	statements.Append(RequireContext(receiver))
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
		WrapperCalls(
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

func (data ESConstructorData) Generate() *jen.Statement {
	return StatementList(
		CreateConstructor(data),
		CreateConstructorWrapper(data),
		CreateWrapperMethods(data),
	).Generate()
}

func CreateConstructor(data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name:     fmt.Sprintf("Create%sPrototype", data.InnerTypeName),
		Args:     g.Arg(g.Id("host"), scriptHostPtr),
		RtnTypes: g.List(v8FunctionTemplatePtr),
		Body:     CreateConstructorBody(data),
	}
}

func CreateConstructorBody(data ESConstructorData) g.Generator {
	builder := NewConstructorBuilder()
	scriptHost := g.NewValue("host")
	constructor := v8FunctionTemplate{g.NewValue("constructor")}

	createWrapperFunction := g.NewValue(fmt.Sprintf("New%s", data.WrapperTypeName))

	statements := StatementList(
		builder.v8Iso.Assign(scriptHost.Field("iso")),
		g.Assign(builder.Wrapper, createWrapperFunction.Call(scriptHost)),
		g.Assign(constructor, builder.NewFunctionTemplateOfWrappedMethod("NewInstance")),
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

type JenGenerator = g.Generator

type IfStmt struct {
	Condition JenGenerator
	Block     JenGenerator
	Else      JenGenerator
}

type StatementListStmt struct {
	Statements []JenGenerator
}

func StatementList(statements ...JenGenerator) StatementListStmt {
	return StatementListStmt{statements}
}

func (s *StatementListStmt) Prepend(stmt JenGenerator) {
	s.Statements = slices.Insert(s.Statements, 0, stmt)
}

func (s IfStmt) Generate() *jen.Statement {
	result := jen.If(s.Condition.Generate()).Block(s.Block.Generate())
	if s.Else != nil {
		result.Else().Block(s.Else.Generate())
	}
	return result
}

func GetSliceLength(gen JenGenerator) JenGenerator {
	return Stmt{jen.Len(gen.Generate())}
}

func Statements(stmts ...JenGenerator) JenGenerator {
	return StatementListStmt{stmts}
}

func (s *StatementListStmt) Append(stmt ...JenGenerator) {
	s.Statements = append(s.Statements, stmt...)
}
func (s *StatementListStmt) AppendJen(stmt *jen.Statement) {
	s.Statements = append(s.Statements, Stmt{stmt})
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

type CallInstance struct {
	Name     string
	Args     []g.Generator
	Op       ESOperation
	Instance *g.Value
	Receiver WrapperInstance
}

type GetGeneratorResult struct {
	Generator      JenGenerator
	HasValue       bool
	HasError       bool
	RequireContext bool
}

func (c CallInstance) PerformCall() (genRes GetGeneratorResult) {
	args := []g.Generator{}
	genRes.HasError = c.Op.GetHasError()
	genRes.HasValue = c.Op.ReturnType != "undefined"
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
		converter := "To"
		if c.Op.Nullable {
			converter += "Nullable"
		}
		converter += idlNameToGoName(c.Op.ReturnType)
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
	genRes.Generator = list
	return genRes
}

type ReadArgumentsResult struct {
	ArgNames  []g.Generator
	ErrNames  []g.Generator
	Generator g.Generator
}

func (r ReadArgumentsResult) Generate() *jen.Statement {
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
	return g.Assign(
		g.Id("args"),
		g.Raw(
			jen.Id("newArgumentHelper").
				Call(jen.Id(data.Receiver).Dot("host"), jen.Id("info")),
		),
	)
}

func ReadArguments(data ESConstructorData, op ESOperation) (res ReadArgumentsResult) {
	argCount := len(op.Arguments)
	res.ArgNames = make([]g.Generator, argCount)
	res.ErrNames = make([]g.Generator, argCount)
	statements := &StatementListStmt{}
	for i, arg := range op.Arguments {
		argName := g.Id(arg.Name)
		errName := g.Id(fmt.Sprintf("err%d", i+1))
		res.ArgNames[i] = argName
		res.ErrNames[i] = errName

		var convertNames []string
		if arg.Type != "" {
			convertNames = []string{fmt.Sprintf("Decode%s", idlNameToGoName(arg.Type))}
		} else {
			types := arg.IdlType.IdlType.IType.Types
			convertNames = make([]string, len(types))
			for i, t := range types {
				convertNames[i] = fmt.Sprintf("Decode%s", t.IType.TypeName)
			}
		}

		converters := make([]jen.Code, 0)
		converters = append(converters, jen.Id("args"))
		converters = append(converters, jen.Lit(i))
		for _, n := range convertNames {
			converters = append(converters, g.Raw(jen.Id(data.Receiver).Dot(n)).Generate())
		}
		statements.Append(g.Assign(
			g.Raw(jen.List(argName.Generate(), errName.Generate())),
			Stmt{jen.Id("TryParseArg").Call(converters...)}))
	}
	res.Generator = statements
	return
}

func GetInstanceAndError(id g.Generator, errId g.Generator, data ESConstructorData) g.Generator {
	return g.AssignMany(
		g.List(id, errId),
		g.Raw(jen.Id(data.Receiver).Dot("GetInstance").Call(jen.Id("info"))),
	)
}

func FunctionTemplateCallbackBody(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	if op.NotImplemented {
		errMsg := fmt.Sprintf("Not implemented: %s.%s", data.Name, op.Name)
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
		WrapperCalls(
			op,
			idlNameToGoName(op.Name),
			readArgsResult.ArgNames,
			errNames,
			CreateCall,
			true,
		),
	)
	if requireContext {
		statements.Prepend(RequireContext(receiver))
	}
	return statements
}

func CreateConstructorWrapper(data ESConstructorData) JenGenerator {
	return StatementList(
		g.Line,
		g.FunctionDefinition{
			Name: "NewInstance",
			Receiver: g.FunctionArgument{
				Name: g.Id(data.Receiver),
				Type: g.Id(data.WrapperTypeName),
			},
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     JSConstructorImpl(data),
		},
	)
}

func CreateWrapperMethods(data ESConstructorData) JenGenerator {
	list := StatementList()
	for _, op := range data.Operations {
		list.Append(CreateWrapperMethod(data, op))
	}
	for _, attr := range data.Attributes {
		if attr.Getter != nil {
			list.Append(CreateWrapperMethod(data, *attr.Getter))
		}
		if attr.Setter != nil {
			list.Append(CreateWrapperMethod(data, *attr.Setter))
		}
	}
	return list
}

func CreateWrapperMethod(
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
			Body:     FunctionTemplateCallbackBody(data, op),
		})
}

type NewFunctionTemplate struct {
	iso JenGenerator
	f   JenGenerator
}

func (t NewFunctionTemplate) Generate() *jen.Statement {
	return jen.Qual(v8, "NewFunctionTemplateWithError").Call(t.iso.Generate(), t.f.Generate())
}

func idlNameToGoName(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = upperCaseFirstLetter(word)
	}
	return strings.Join(words, "")
}

func upperCaseFirstLetter(s string) string {
	strLen := len(s)
	if strLen == 0 {
		slog.Warn("Passing empty string to upperCaseFirstLetter")
		return ""
	}
	buffer := make([]rune, 0, strLen)
	buffer = append(buffer, unicode.ToUpper([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}

type Stmt struct{ *jen.Statement }

func (s Stmt) Generate() *jen.Statement { return s.Statement }

type ReturnOnError struct {
	err g.Generator
}

func (ret ReturnOnError) Generate() *jen.Statement {
	err := ret.err
	if err == nil {
		err = g.Id("err")
	}
	return IfStmt{
		Condition: g.Neq{Lhs: err, Rhs: g.Nil}, //g.Raw(err.Generate().Op("!=").Nil()),
		Block:     g.Return(g.Raw(jen.Nil()), err),
	}.Generate()
}
