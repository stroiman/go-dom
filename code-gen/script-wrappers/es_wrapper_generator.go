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

const (
	dom  = "github.com/stroiman/go-dom/browser/dom"
	html = "github.com/stroiman/go-dom/browser/html"
	sc   = "github.com/stroiman/go-dom/browser/scripting"
	v8   = "github.com/tommie/v8go"
)

func createData(spec IdlSpec, dataData WrapperTypeSpec) ESConstructorData {
	idlName, ok := spec.GetType(dataData.TypeName)
	if !ok {
		panic("Missing type")
	}
	wrappedTypeName := dataData.InnerTypeName
	if wrappedTypeName == "" {
		wrappedTypeName = idlName.Type.Name
	}
	wrapperTypeName := dataData.WrapperTypeName
	if wrapperTypeName == "" {
		wrapperTypeName = fmt.Sprintf("%sV8Wrapper", wrappedTypeName)
	}
	return ESConstructorData{
		Spec:            dataData,
		InnerTypeName:   wrappedTypeName,
		WrapperTypeName: wrapperTypeName,
		Receiver:        dataData.Receiver,
		RunCustomCode:   dataData.RunCustomCode,
		Constructor:     CreateConstructor(dataData, idlName),
		Operations:      CreateInstanceMethods(dataData, idlName),
		Attributes:      CreateAttributes(dataData, idlName),
	}
}

func CreateConstructor(
	dataData WrapperTypeSpec,
	idlName IdlTypeSpec) *ESOperation {
	if c, ok := idlName.Constructor(); ok {
		result := createOperation(dataData, c)
		return &result
	} else {
		return nil
	}
}

func CreateInstanceMethods(
	dataData WrapperTypeSpec,
	idlName IdlTypeSpec) (result []ESOperation) {
	for instanceMethod := range idlName.InstanceMethods() {
		op := createOperation(dataData, instanceMethod)
		result = append(result, op)
	}
	return
}

func CreateAttributes(
	dataData WrapperTypeSpec,
	idlName IdlTypeSpec,
) (res []ESAttribute) {
	for attribute := range idlName.Attributes() {
		methodCustomization := dataData.GetMethodCustomization(attribute.Name)
		if methodCustomization.Ignored {
			continue
		}

		var (
			getter *ESOperation
			setter *ESOperation
		)
		r := attribute.AttributeType()
		rtnType := r.TypeName
		getter = &ESOperation{
			Name:                 attribute.Name,
			NotImplemented:       methodCustomization.NotImplemented,
			CustomImplementation: methodCustomization.CustomImplementation,
			RetType:              attribute.AttributeType(),
			MethodCustomization:  methodCustomization,
		}
		name := idlNameToGoName(getter.Name)
		if attribute.Readonly {
			getter.Name = name
		} else {
			getter.Name = fmt.Sprintf("Get%s", name)

			setter = new(ESOperation)
			*setter = *getter
			setter.Name = fmt.Sprintf("Set%s", name)
			methodCustomization := dataData.GetMethodCustomization(setter.Name)
			setter.NotImplemented = setter.NotImplemented || methodCustomization.NotImplemented
			setter.CustomImplementation = setter.CustomImplementation || methodCustomization.CustomImplementation
			setter.RetType = NewRetTypeUndefined()
			setter.Arguments = []ESOperationArgument{{
				Name:     "val",
				Type:     idlNameToGoName(rtnType),
				Optional: false,
				Variadic: false,
			}}
		}
		getterCustomization := dataData.GetMethodCustomization(getter.Name)
		getter.NotImplemented = getterCustomization.NotImplemented || getter.NotImplemented
		getter.CustomImplementation = getterCustomization.CustomImplementation ||
			getter.CustomImplementation
		res = append(res, ESAttribute{attribute.Name, getter, setter})
	}
	return
}

func createOperation(typeSpec WrapperTypeSpec, member MemberSpec) ESOperation {
	methodCustomization := typeSpec.GetMethodCustomization(member.Name)
	op := ESOperation{
		Name:                 member.Name,
		NotImplemented:       methodCustomization.NotImplemented,
		CustomImplementation: methodCustomization.CustomImplementation,
		RetType:              member.ReturnType(),
		MethodCustomization:  methodCustomization,
		HasError:             !methodCustomization.HasNoError,
		Arguments:            []ESOperationArgument{},
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
		}
		if arg.IdlType.IdlType != nil {
			esArg.Type = arg.IdlType.IdlType.IType.TypeName
		}
		op.Arguments = append(op.Arguments, esArg)
	}
	return op
}

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
	RetType              RetType
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
	Spec            *ESClassWrapper
	InnerTypeName   string
	WrapperTypeName string
	Receiver        string
	Operations      []ESOperation
	Attributes      []ESAttribute
	Constructor     *ESOperation
	RunCustomCode   bool
}

func (d ESConstructorData) Name() string { return d.Spec.TypeName }

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

type JenGenerator = g.Generator

type IfStmt struct {
	Condition g.Generator
	Block     g.Generator
	Else      g.Generator
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

func IsNodeType(typeName string) bool {
	loweredName := strings.ToLower(typeName)
	switch loweredName {
	case "node":
		return true
	case "document":
		return true
	case "documentfragment":
		return true
	}
	if strings.HasSuffix(loweredName, "element") {
		return true
	}
	return false
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

func ReadArguments(data ESConstructorData, op ESOperation) (res V8ReadArguments) {
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
		Block:     g.Return(g.Nil, err),
	}.Generate()
}
