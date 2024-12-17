package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"unicode"

	g "github.com/stroiman/go-dom/internal/code-gen/generators"

	"github.com/dave/jennifer/jen"
)

var (
	v8FunctionTemplatePtr = g.NewTypePackage("FunctionTemplate", v8).Pointer()
	v8Value               = g.NewTypePackage("Value", v8).Pointer()
	scriptHostPtr         = g.NewType("ScriptHost").Pointer()
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
	ops := []*tmp{}
	for _, member := range idlName.Members {
		if -1 != slices.IndexFunc(
			ops,
			func(op *tmp) bool { return op.Op.Name == member.Name },
		) {
			slog.Warn("Function overloads", "Name", member.Name)
			continue
		}
		returnType, nullable := FindIdlType(member.IdlType)
		operation := &tmp{ESOperation{
			Name:       member.Name,
			ReturnType: returnType,
			Nullable:   nullable,
			Arguments:  []ESOperationArgument{},
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
			ops = append(ops, operation)
		}
		if member.Type == "constructor" {
			constructor = &operation.Op
		}
	}
	// fmt.Println(operations)
	// os.Exit(1)

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
		Constructor:      constructor,
		CreatesInnerType: true,
		IdlName:          idlName,
	}, nil
}

type ValueType struct {
	Value ValueTypes `json:"value"`
}

type ValueTypes struct {
	Values    []ValueType
	Value     *ValueType
	ValueName string
}

type ExtAttr struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Rhs  struct {
		Type  string     `json:"type"`
		Value ValueTypes `json:"value"`
	} `json:"rhs"`
}

type IdlType struct {
	Type     string    `json:"type"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Generic  string    `json:"generic"`
	Nullable bool      `json:"nullable"`
	Union    bool      `json:"union"`
	IType    IdlTypes  `json:"idlType"`
}

// Hmmm, can we find this in the IDL somewhere? It's specified in prose, but I
// can't find it in easy consumable JSON.
var hasNoError = map[string]bool{
	"setRequestHeader":  true,
	"open":              true,
	"getResponseHeader": true,
}

type IdlTypes struct {
	Types    []IdlType
	IdlType  *IdlType
	TypeName string
}

func (i IdlTypes) String() string {
	if len(i.Types) > 0 {
		return fmt.Sprintf("%v", i.Types)
	}
	if i.IdlType != nil {
		return fmt.Sprintf("%v", *i.IdlType)
	}
	return i.TypeName
}

func (t *ValueTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Values)
	if err != nil {
		val := new(ValueType)
		err = json.Unmarshal(bytes, val)
		if err == nil {
			t.Value = val
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.ValueName)
	}
	return err
}

func (t *IdlTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Types)
	if err != nil {
		typ := new(IdlType)
		err = json.Unmarshal(bytes, &typ)
		if err == nil {
			t.IdlType = typ
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.TypeName)
	}
	return err

}

type Stuff struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	IdlType  IdlTypes  `json:"idlType"`
}

type ArgumentType struct {
	Stuff
	Default  any  `json:"default"`
	Optional bool `json:"optional"`
	Variadic bool `json:"variadic"`
}

type IdlNameMember struct {
	Stuff
	Arguments []ArgumentType `json:"arguments"`
	Special   string         `json:"special"`
	Readonly  bool           `json:"readOnly"`
	Href      string         `json:"href"`
}

type IdlName struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Members []IdlNameMember `json:"members"`
	Partial bool            `json:"partial"`
	Href    string          `json:"href"`
}

type IdlParsed struct {
	IdlNames map[string]IdlName
}

type ParsedIdlFile struct {
	IdlParsed `json:"idlParsed"`
}

func FindIdlType(idl IdlTypes) (string, bool) {
	types := idl.Types
	if len(types) == 0 && idl.IdlType != nil {
		types = []IdlType{*idl.IdlType}
	}
	for _, t := range types {
		if t.Type == "return-type" {
			return t.IType.TypeName, t.Nullable
		}
	}
	return "", false

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
	Name       string
	ReturnType string
	Nullable   bool
	Arguments  []ESOperationArgument
}

type ESConstructorData struct {
	CreatesInnerType bool
	InnerTypeName    string
	WrapperTypeName  string
	Receiver         string
	Operations       []ESOperation
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

func (c JSConstructor) JSConstructorImpl(grp *jen.Group, data ESConstructorData) {
	grp.Add(c.varScriptContext).
		Op(":=").
		Add(c.argHost).
		Dot("MustGetContext").
		Call(c.argInfo.Clone().Dot("Context").Call())
	buildInstance := jen.Id("wrapper").Dot("CreateInstance").Call(c.varScriptContext)
	grp.Add(c.varInstance).Op(":=").Add(buildInstance)
	grp.List(jen.Id("_"), jen.Id("err")).
		Op(":=").
		Add(c.varScriptContext.Clone().Dot("CacheNode").Call(
			c.getThis,
			c.varInstance,
		))
	grp.Return(jen.Nil(), jen.Id("err"))
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
	errorT := jen.Id("error")
	gr := Helper{&jen.Group{}}
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
			Stmt{
				jen.Qual(v8, "NewFunctionTemplateWithError").
					Call(jen.Id("iso"), jen.Func().
						Params(c.argInfo.Clone().Add(gr.v8FunctionCallbackInfoPtr())).
						Params(v8Value.Generate(), errorT).
						BlockFunc(func(grp *jen.Group) { c.JSConstructorImpl(grp, data) })),
			},
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
	f := jen.Id("wrapper").Dot(camelCase(op.Name))
	ft := NewFunctionTemplate{Stmt{c.varIso}, Stmt{f}}
	return Stmt{(jen.Id("prototype").Dot("Set").Call(jen.Lit(op.Name), ft.Generate()))}
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
			parserName := fmt.Sprintf("Get%sFrom%s", camelCase(s.Arg.Name), t.IType.TypeName)
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

func (s *StatementListStmt) Append(stmt JenGenerator) {
	s.Statements = append(s.Statements, stmt)
}
func (s *StatementListStmt) AppendJen(stmt *jen.Statement) {
	s.Statements = append(s.Statements, Stmt{stmt})
}

func (s StatementListStmt) Generate() *jen.Statement {
	result := []jen.Code{}
	g := jen.Group{}
	for i, s := range s.Statements {
		if i > 0 {
			result = append(result, jen.Line())
		}
		g.Add(s.Generate())
		jenStatement := s.Generate()
		if jenStatement != nil {
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
	genRes.HasError = !hasNoError[c.Op.Name]
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
		evaluation = jen.Id(camelCase(c.Name)).Call(args...)
	} else {
		evaluation = jen.Id(instanceName).Dot(camelCase(c.Name)).Call(args...)
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

func (c CallInstance) GetGenerator(instanceName string) GetGeneratorResult {
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
		converter += c.Op.ReturnType
		genRes.RequireContext = true
		valueReturn := Stmt{jen.Return(jen.Id(converter).Call(jen.Id("ctx"), jen.Id("result")))}
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
	opName = opName + camelCase(arg.Name)
	statements.Append(ifArgs)
	processOptionalArgs(data, args, opName, from+1, innerStatements, argNames, op, requireContext)
	genResult := CallInstance{
		Name: opName,
		Args: argNames,
		Op:   op,
	}.GetGenerator("instance")
	innerStatements.Append(genResult.Generator)
}

func (c JSConstructor) FunctionTemplateCallbackBody(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	return Stmt{jen.BlockFunc(func(grp *jen.Group) {
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
				Getter:   "GetArg" + arg.Type,
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
		}.GetGenerator("instance")
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
		grp.Add(statements.Generate())
	})}
}

func CreateWrapperMethods(c JSConstructor, data ESConstructorData) JenGenerator {
	generators := make([]JenGenerator, len(data.Operations))
	for i, op := range data.Operations {
		generators[i] = c.CreateWrapperMethod(data, op)
	}
	return StatementList(generators...)
}

func (c JSConstructor) CreateWrapperMethod(
	data ESConstructorData,
	op ESOperation,
) JenGenerator {
	v8Value := jen.Op("*").Qual(v8, "Value")
	errorT := jen.Id("error")
	v8FunctionCallbackInfoPtr := jen.Op("*").Qual(v8, "FunctionCallbackInfo")
	f := c.FunctionTemplateCallbackBody(data, op).Generate()
	return StatementList(
		NewLine(),
		Stmt{
			jen.Func().
				Params(jen.Id(data.Receiver).Id(data.WrapperTypeName)).
				Id(camelCase(op.Name)).
				Params(c.argInfo.Clone().Add(v8FunctionCallbackInfoPtr)).
				Params(v8Value, errorT).
				BlockFunc(func(grp *jen.Group) {
					grp.Add(f)
				}),
		})
}

type NewFunctionTemplate struct {
	iso JenGenerator
	f   JenGenerator
}

func (t NewFunctionTemplate) Generate() *jen.Statement {
	return jen.Qual(v8, "NewFunctionTemplateWithError").Call(t.iso.Generate(), t.f.Generate())
}

func camelCase(s string) string {
	buffer := make([]rune, 0, len(s))
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
		Block:     g.Raw(jen.Return(jen.Nil(), jErr)),
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
