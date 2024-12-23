package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"unicode"

	g "github.com/stroiman/go-dom/internal/code-gen/generators"

	"github.com/dave/jennifer/jen"
)

var (
	v8FunctionTemplatePtr     = g.NewTypePackage("FunctionTemplate", v8).Pointer()
	v8FunctionCallbackInfoPtr = g.NewTypePackage("FunctionCallbackInfo", v8).Pointer()
	v8Value                   = g.NewTypePackage("Value", v8).Pointer()
	scriptHostPtr             = g.NewType("ScriptHost").Pointer()
)

type CreateDataData struct {
	InnerTypeName   string
	WrapperTypeName string
	Receiver        string
}

func createData(data []byte, iName string, dataData CreateDataData) (ESConstructorData, error) {
	spec := ParsedIdlFile{}
	var constructor *ESOperation
	err := json.Unmarshal(data, &spec)
	if err != nil {
		panic(err)
	}
	idlName := spec.IdlNames[iName]
	type tmp struct {
		Op ESOperation
		Ok bool
	}
	missingOps := notImplementedFunctions[iName]
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
		notImplemented := slices.Index(missingOps, member.Name) != -1
		operation := &tmp{ESOperation{
			Name:           member.Name,
			NotImplemented: notImplemented,
			ReturnType:     returnType,
			Nullable:       nullable,
			Arguments:      []ESOperationArgument{},
		}, true}
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
		if member.Type == "operation" {
			operation.Op.HasError = !hasNoError[member.Name]
			ops = append(ops, operation)
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
			getter.NotImplemented = slices.Index(missingOps, getter.Name) != -1 || op.NotImplemented
			if !member.Readonly {
				setter = new(ESOperation)
				*setter = op
				setter.Name = fmt.Sprintf("Set%s", idlNameToGoName(op.Name))
				setter.NotImplemented = slices.Index(missingOps, setter.Name) != -1 ||
					op.NotImplemented
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
	}, nil
}

// Hmmm, can we find this in the IDL somewhere? It's specified in prose, but I
// can't find it in easy consumable JSON.
var hasNoError = map[string]bool{
	"setRequestHeader":  true,
	"open":              true,
	"getResponseHeader": true,
}

var notImplementedFunctions = map[string][]string{
	"XMLHttpRequest": {
		"readyState",
		"timeout",
		"withCredentials",
		"upload",
		"responseURL",
		"response", // TODO, just because of the return value
		"responseType",
		"responseXML",
	},
	"URL": {
		"SetHref",
		"SetProtocol",
		"username",
		"password",
		"SetHost",
		"SetPort",
		"SetHostname",
		"SetPathname",
		"searchParams",
		"SetHash",
		"SetSearch",
	},
	"DOMTokenList": {
		"item",
		"contains",
		"remove",
		"toggle",
		"replace",
		"supports",
		"value",
		"length",
	},
}

const br = "github.com/stroiman/go-dom/browser"
const sc = "github.com/stroiman/go-dom/scripting"
const v8 = "github.com/tommie/v8go"

type ESOperationArgument struct {
	Name     string
	Type     string
	Optional bool
	Variadic bool
	IdlType  IdlTypes
}

type ESOperation struct {
	Name           string
	NotImplemented bool
	ReturnType     string
	Nullable       bool
	HasError       bool
	Arguments      []ESOperationArgument
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
	IdlName
}

type Imports = [][][2]string

func WriteImports(b *builder, imports Imports) {
	b.Printf("import (\n")
	b.indent()
	defer b.unIndentF(")\n\n")
	for _, grp := range imports {
		for _, imp := range grp {
			alias := imp[0]
			pkg := imp[1]
			if alias == "" {
				b.Printf("\"%s\"\n", pkg)
			} else {
				b.Printf("%s \"%s\"\n", imp[0], imp[1])
			}
		}
	}
}

type st = *jen.Statement

type JSConstructor struct {
	argHost          st
	argInfo          st
	getThis          st
	varInstance      st
	varIso           st
	varScriptContext st
}

func CreateJSConstructor() JSConstructor {
	argInfo := jen.Id("info")
	argHost := jen.Id("host")
	getThis := argInfo.Clone().Dot("This").Call()
	varInstance := jen.Id("instance")
	varIso := jen.Id("iso")
	varScriptContext := jen.Id("ctx")
	return JSConstructor{
		argHost,
		argInfo,
		getThis,
		varInstance,
		varIso,
		varScriptContext,
	}
}

func IllegalConstructor(data ESConstructorData) g.Generator {
	return g.Return(g.Nil,
		g.Raw(jen.Qual(v8, "NewTypeError").Call(
			jen.Id(data.Receiver).Dot("host").Dot("iso"), jen.Lit("Illegal Constructor"),
		)),
	)
}

func (c JSConstructor) JSConstructorImpl(data ESConstructorData) g.Generator {
	if data.Constructor == nil {
		return IllegalConstructor(data)
	}
	var readArgsResult ReadArgumentsResult
	var constructorArguments []ESOperationArgument
	readArgsResult = ReadArguments(data, *data.Constructor)
	constructorArguments = data.Constructor.Arguments
	statements := StatementList(readArgsResult)
	statements.Append(
		g.Assign(g.Id("ctx"), Stmt{jen.Id(data.Receiver).Dot("host").Dot("MustGetContext").Call(
			jen.Id("info").Dot("Context").Call(),
		)}),
	)
	for i := len(constructorArguments); i >= 0; i-- {
		functionName := "CreateInstance"
		for j, arg := range constructorArguments {
			if j < i {
				if arg.Optional {
					functionName += idlNameToGoName(arg.Name)
				}
			}
		}
		argnames := readArgsResult.ArgNames[0:i]
		errNames := readArgsResult.ErrNames[0:i]
		construction := StatementList(
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
		if i > 0 {
			arg := constructorArguments[i-1]
			argErrorCheck := StatementList(
				g.Assign(g.Id("err"),
					g.Raw(jen.Qual("errors", "Join").CallFunc(func(g *jen.Group) {
						for _, e := range errNames {
							g.Add(e.Generate())
						}
					})),
				),
				GenReturnOnError(),
			)
			statements.Append(StatementList(
				IfStmt{
					Condition: g.Raw(jen.Id("args").Dot("noOfReadArguments").Op(">=").Lit(i)),
					Block: StatementList(
						argErrorCheck,
						construction,
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
			statements.Append(construction)
		}
	}
	return statements
}

func CreateInstance(typeName string, params ...jen.Code) JenGenerator {
	constructorName := fmt.Sprintf("New%s", typeName)
	return Stmt{
		jen.Id(constructorName).Call(params...),
	}
}

func Id(id string) JenGenerator { return Stmt{jen.Id(id).Clone()} }

func (c JSConstructor) Run(f *jen.File, data ESConstructorData) {
	gen := StatementList(
		CreateConstructor(c, data),
		CreateConstructorWrapper(c, data),
		CreateWrapperMethods(c, data),
	)
	f.Add(gen.Generate())
}

func CreateConstructor(c JSConstructor, data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name:     fmt.Sprintf("Create%sPrototype", data.InnerTypeName),
		Args:     g.Arg(g.Id("host"), scriptHostPtr),
		RtnTypes: g.List(v8FunctionTemplatePtr),
		Body:     CreateConstructorBody(c, data),
	}
}

func CreateConstructorBody(c JSConstructor, data ESConstructorData) g.Generator {
	constructor := g.Id("constructor")
	return StatementList(
		g.Assign(Id("iso"),
			Stmt{jen.Id("host").Dot("iso")},
		),
		g.Assign(
			Id("wrapper"),
			CreateInstance(data.WrapperTypeName, jen.Id("host")),
		),
		g.Assign(constructor,
			NewFunctionTemplate{g.Id("iso"), Stmt{jen.Id("wrapper").Dot("NewInstance")}},
		),
		g.Raw(jen.Id("constructor").
			Dot("GetInstanceTemplate").
			Call().
			Dot("SetInternalFieldCount").
			Call(jen.Lit(1)),
		),
		g.Assign(
			Id("prototype"),
			Stmt{jen.Id("constructor").Dot("PrototypeTemplate").Call()},
		),
		NewLine(),
		InstallFunctionHandlers(c, data),
		InstallAttributeHandlers(c, data),
		g.Return(constructor),
	)
}

func InstallFunctionHandlers(c JSConstructor, data ESConstructorData) JenGenerator {
	generators := make([]JenGenerator, len(data.Operations))
	for i, op := range data.Operations {
		generators[i] = InstallFunctionHandler(c, op)
	}
	return StatementList(generators...)
}

func InstallFunctionHandler(c JSConstructor, op ESOperation) JenGenerator {
	f := jen.Id("wrapper").Dot(idlNameToGoName(op.Name))
	ft := NewFunctionTemplate{Stmt{c.varIso}, Stmt{f}}
	return Stmt{(jen.Id("prototype").Dot("Set").Call(jen.Lit(op.Name), ft.Generate()))}
}

func InstallAttributeHandlers(c JSConstructor, data ESConstructorData) g.Generator {
	length := len(data.Attributes)
	if length == 0 {
		return g.Noop
	}
	generators := make([]JenGenerator, length+1)
	generators[0] = g.Line
	for i, op := range data.Attributes {
		generators[i+1] = InstallAttributeHandler(c, op)
	}
	return StatementList(generators...)
}

func InstallAttributeHandler(c JSConstructor, op ESAttribute) g.Generator {
	getter := op.Getter
	setter := op.Setter
	list := StatementList()
	if getter != nil {
		f := jen.Id("wrapper").Dot(idlNameToGoName(getter.Name))
		ft := NewFunctionTemplate{Stmt{c.varIso}, Stmt{f}}
		var setterFt g.Generator
		var Attributes = "ReadOnly"
		if setter != nil {
			f := Stmt{jen.Id("wrapper").Dot(idlNameToGoName(setter.Name))}
			setterFt = NewFunctionTemplate{Stmt{c.varIso}, f}
			Attributes = "None"
		} else {
			setterFt = g.Nil
		}

		list.Append(Stmt{
			(jen.Id("prototype").Dot("SetAccessorProperty").Call(jen.Lit(op.Name), jen.Line().Add(ft.Generate()), jen.Line().Add(setterFt.Generate()), jen.Line().Add(jen.Qual(v8, Attributes)))),
		})
	}
	// if setter != nil {
	// 	f := jen.Id("wrapper").Dot(idlNameToGoName(setter.Name))
	// 	ft := NewFunctionTemplate{Stmt{c.varIso}, Stmt{f}}
	// 	list.Append(Stmt{
	// 		(jen.Id("prototype").Dot("SetAccessorProperty").Call(jen.Lit(op.Name), ft.Generate())),
	// 	})
	// }
	return list
}

type JenGenerator = g.Generator

type GetArgStmt struct {
	Name     string
	Receiver string
	ErrName  string
	Getter   string
	Index    int
	Arg      ESOperationArgument
}

type IfStmt struct {
	Condition JenGenerator
	Block     JenGenerator
	Else      JenGenerator
}

func (s GetArgStmt) Generate() *jen.Statement {
	if s.Arg.Type != "" {
		return AssignmentStmt{
			VarNames: []string{s.Name, s.ErrName},
			Expression: Stmt{
				jen.Id(s.Receiver).Dot(s.Getter).Call(jen.Id("args"), jen.Lit(s.Index)),
			},
		}.Generate()
	} else {
		statements := []jen.Code{jen.Id("ctx"), jen.Id("args"), jen.Lit(s.Index)}
		for _, t := range s.Arg.IdlType.IdlType.IType.Types {
			parserName := fmt.Sprintf("Get%sFrom%s", idlNameToGoName(s.Arg.Name), t.IType.TypeName)
			statements = append(statements, jen.Id(parserName))
		}
		return AssignmentStmt{
			VarNames:   []string{s.Name, s.ErrName},
			Expression: Stmt{jen.Id("TryParseArgs").Call(statements...)},
		}.Generate()
	}
}

type AssignmentStmt struct {
	VarNames   []string
	Expression JenGenerator
	NoNewVars  bool
}

type StatementListStmt struct {
	Statements []JenGenerator
}

func StatementList(statements ...JenGenerator) StatementListStmt {
	return StatementListStmt{statements}
}

func NewLine() JenGenerator { return Stmt{jen.Line()} }

func (s AssignmentStmt) Generate() *jen.Statement {
	list := make([]jen.Code, 0, len(s.VarNames))
	for _, n := range s.VarNames {
		list = append(list, jen.Id(n))
	}
	operator := ":="
	if s.NoNewVars {
		operator = "="
	}
	return jen.List(list...).Op(operator).Add(s.Expression.Generate())
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

type GeneratorFunc func() *jen.Statement

func (g GeneratorFunc) Generate() *jen.Statement {
	return g.Generate()
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
	Name string
	Args []string
	Op   ESOperation
}

type GetGeneratorResult struct {
	Generator      JenGenerator
	HasValue       bool
	HasError       bool
	RequireContext bool
}

func (c CallInstance) PerformCall(instanceName string) (genRes GetGeneratorResult) {
	args := []jen.Code{}
	genRes.HasError = c.Op.HasError
	genRes.HasValue = c.Op.ReturnType != "undefined"
	var stmt *jen.Statement
	if genRes.HasValue {
		stmt = jen.Id("result")
	}
	if genRes.HasError {
		if stmt != nil {
			stmt = stmt.Op(",").Id("err")
		} else {
			stmt = jen.Id("err")
		}
	}
	if stmt != nil {
		if genRes.HasValue {
			stmt = stmt.Op(":=")
		} else {
			stmt = stmt.Op("=")
		}
	}

	for _, a := range c.Args {
		args = append(args, jen.Id(a))
	}
	list := StatementListStmt{}
	var evaluation *jen.Statement
	if instanceName == "" {
		evaluation = jen.Id(idlNameToGoName(c.Name)).Call(args...)
	} else {
		evaluation = jen.Id(instanceName).Dot(idlNameToGoName(c.Name)).Call(args...)
	}
	if stmt == nil {
		list.Append(Stmt{evaluation})
	} else {
		list.Append(Stmt{
			stmt.Add(evaluation),
		})
	}
	genRes.Generator = list
	return
}

func (c CallInstance) GetGenerator(receiver string, instanceName string) GetGeneratorResult {
	genRes := c.PerformCall(instanceName)
	list := StatementListStmt{}
	list.Append(genRes.Generator)
	if !genRes.HasValue {
		if genRes.HasError {
			list.Append(Stmt{jen.Return(jen.Nil(), jen.Id("err"))})
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
		valueReturn := Stmt{jen.Return(jen.Id(receiver).Dot(converter).Call(jen.Id("ctx"), jen.Id("result")))}
		if genRes.HasError {
			list.Append(IfStmt{
				Condition: Stmt{jen.Id("err").Op("!=").Nil()},
				Block:     Stmt{jen.Return(jen.Nil(), jen.Id("err"))},
				Else:      valueReturn,
			})
		} else {
			list.Append(valueReturn)
		}
	}
	genRes.Generator = list
	return genRes
}

func getInstance(receiver string) JenGenerator {
	return AssignmentStmt{
		VarNames:   []string{"instance", "err"},
		Expression: Stmt{jen.Id(receiver).Dot("GetInstance").Call(jen.Id("info"))},
	}
}

func processOptionalArgs(
	data ESConstructorData,
	args []ESOperationArgument,
	opName string,
	from int,
	statements *StatementListStmt,
	argNames []string,
	op ESOperation,
	requireContext *bool,
) {
	if from >= len(args) {
		return
	}
	arg := args[from]
	if len(arg.IdlType.IdlType.IType.Types) > 0 {
		*requireContext = true
	}
	innerStatements := &StatementListStmt{}
	ifArgs := IfStmt{
		Condition: Stmt{jen.Id("argsLen").Op(">=").Lit(from + 1)},
		Block:     innerStatements,
	}
	innerStatements.Append(GetArgStmt{
		Name:     arg.Name,
		Receiver: data.Receiver,
		ErrName:  "err",
		Getter:   "GetArg" + arg.Type,
		Index:    from,
		Arg:      arg,
	})
	innerStatements.Append(GenReturnOnError())

	argNames = append(argNames, arg.Name)
	opName = opName + idlNameToGoName(arg.Name)
	statements.Append(ifArgs)
	processOptionalArgs(
		data,
		args,
		opName,
		from+1,
		innerStatements,
		argNames,
		op,
		requireContext,
	)
	genResult := CallInstance{
		Name: opName,
		Args: argNames,
		Op:   op,
	}.GetGenerator(data.Receiver, "instance")
	innerStatements.Append(genResult.Generator)
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

func ReadArguments(data ESConstructorData, op ESOperation) (res ReadArgumentsResult) {
	argCount := len(op.Arguments)
	res.ArgNames = make([]g.Generator, argCount)
	res.ErrNames = make([]g.Generator, argCount)
	statements := &StatementListStmt{}
	if argCount > 0 {
		statements.Append(
			g.Assign(
				g.Id("args"),
				g.Raw(
					jen.Id("newArgumentHelper").
						Call(jen.Id(data.Receiver).Dot("host"), jen.Id("info")),
				),
			),
		)
	}
	for i, arg := range op.Arguments {
		argName := g.Id(arg.Name)
		errName := g.Id(fmt.Sprintf("err%d", i))
		res.ArgNames[i] = argName
		res.ErrNames[i] = errName
		converterName := fmt.Sprintf("Decode%s", arg.Type)
		converter := g.Raw(jen.Id(data.Receiver).Dot(converterName))
		statements.Append(g.Assign(
			g.Raw(jen.List(argName.Generate(), errName.Generate())),
			Stmt{jen.Id("TryParseArg").Call(jen.Id("args"), jen.Lit(i), converter)}))
		// Stmt{jen.Id("args").Dot("GetArg").Call(jen.Lit(i), converter)}))
	}
	res.Generator = statements
	return
}

func (c JSConstructor) FunctionTemplateCallbackBody(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	if op.NotImplemented {
		return g.Return(g.Nil, g.Raw(jen.Qual("errors", "New").Call(jen.Lit("Not implemented"))))
	}
	requireContext := new(bool)
	statements := &StatementListStmt{}
	statements.Append(getInstance(data.Receiver))
	statements.Append(GenReturnOnError())

	firstOptionalArg := slices.IndexFunc(
		op.Arguments,
		func(arg ESOperationArgument) bool {
			return arg.Optional
		},
	)
	argCount := len(op.Arguments)
	if firstOptionalArg == -1 {
		firstOptionalArg = argCount
	}
	requiredArgs := op.Arguments[0:firstOptionalArg]
	if argCount > 0 {
		statements.AppendJen(jen.Id("args").Op(":=").Id("info").Dot("Args").Call())
		statements.Append(AssignmentStmt{
			VarNames:   []string{"argsLen"},
			Expression: GetSliceLength(Stmt{jen.Id("args")}),
		})
	}
	if len(requiredArgs) > 0 {
		statements.Append(IfStmt{
			Condition: Stmt{jen.Id("argsLen").Op("<").Lit(len(requiredArgs))},
			Block: Stmt{jen.Return(
				jen.Nil(),
				jen.Qual("errors", "New").Call(jen.Lit("Too few arguments")),
			)},
		})
	}
	argNames := make([]string, 0, len(op.Arguments))
	for i, arg := range requiredArgs {
		var errName string
		if len(requiredArgs) > 1 {
			errName = fmt.Sprintf("err%d", i)
		} else {
			errName = fmt.Sprintf("err")
		}
		stmt := GetArgStmt{
			Name:     arg.Name,
			ErrName:  errName,
			Receiver: data.Receiver,
			Getter:   "GetArg" + idlNameToGoName(arg.Type),
			Index:    i,
			Arg:      arg,
		}
		statements.Append(stmt)
		argNames = append(argNames, arg.Name)
	}
	statements.Append(genErrorHandler(len(requiredArgs)))

	processOptionalArgs(
		data,
		op.Arguments,
		op.Name,
		firstOptionalArg,
		statements,
		argNames,
		op,
		requireContext,
	)

	genResult := CallInstance{
		Name: op.Name,
		Args: argNames,
		Op:   op,
	}.GetGenerator(data.Receiver, "instance")
	if *requireContext || genResult.RequireContext {
		statements.Prepend(Stmt{
			jen.Id("ctx").
				Op(":=").
				Id(data.Receiver).Dot("host").
				Dot("MustGetContext").
				Call(jen.Id("info").Dot("Context").Call()),
		})
	}
	statements.Append(genResult.Generator)
	return statements
}

func CreateConstructorWrapper(c JSConstructor, data ESConstructorData) JenGenerator {
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
			Body:     c.JSConstructorImpl(data),
		},
	)
}

func CreateWrapperMethods(c JSConstructor, data ESConstructorData) JenGenerator {
	generators := make([]JenGenerator, len(data.Operations))
	for i, op := range data.Operations {
		generators[i] = c.CreateWrapperMethod(data, op)
	}
	list := StatementList(generators...)
	for _, attr := range data.Attributes {
		if attr.Getter != nil {
			list.Append(c.CreateWrapperMethod(data, *attr.Getter))
		}
		if attr.Setter != nil {
			list.Append(c.CreateWrapperMethod(data, *attr.Setter))
		}
	}
	return list
}

func (c JSConstructor) CreateWrapperMethod(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	f := c.FunctionTemplateCallbackBody(data, op)
	return StatementList(
		NewLine(),
		g.FunctionDefinition{
			Receiver: g.FunctionArgument{
				Name: g.Id(data.Receiver),
				Type: g.Id(data.WrapperTypeName),
			},
			Name:     idlNameToGoName(op.Name),
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     f,
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

func GenReturnOnError() JenGenerator {
	jErr := jen.Id("err")
	stmt := IfStmt{
		Condition: g.Raw(jErr.Clone().Op("!=").Nil()),
		Block:     g.Return(g.Raw(jen.Nil()), g.Raw(jErr)),
	}
	return stmt
}

func WriteReturnOnError(grp *jen.Group) {
	grp.Add(GenReturnOnError().Generate())
}

func genErrorHandler(count int) JenGenerator {
	if count == 0 {
		return StatementListStmt{}
	}
	jErr := jen.Id("err")
	result := StatementListStmt{}
	if count > 1 {
		var args []jen.Code
		for i := 0; i < count; i++ {
			args = append(args, jen.Id(fmt.Sprintf("err%d", i)))
		}
		s := Stmt{jErr.Clone().Op("=").Qual("errors", "Join").Call(args...)}
		result.Append(s)
	}
	result.Append(GenReturnOnError())
	return result

}

func WriteErrorHandler(grp *jen.Group, count int) {
	grp.Add(genErrorHandler(count).Generate())
}

func writeFactory(f *jen.File, data ESConstructorData) {
	CreateJSConstructor().Run(f, data)
}

type Helper struct{ *jen.Group }

func (h Helper) BuildInstance() *jen.Statement {
	return h.scriptContext().Dot("Window").Call().Dot("NewXmlHttpRequest").Call()
}

func (h Helper) v8FunctionCallbackInfoPtr() *jen.Statement {
	return h.Op("*").Qual(v8, "FunctionCallbackInfo")
}
func (h Helper) hostArg() *jen.Statement {
	return h.Id("info")
}
func (h Helper) infoArg() *jen.Statement {
	return h.Id("info")
}
func (h Helper) scriptContext() *jen.Statement {
	return h.Id("scriptContext")
}
