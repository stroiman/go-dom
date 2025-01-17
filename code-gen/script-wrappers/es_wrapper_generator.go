package wrappers

import (
	"fmt"
	"iter"
	"log/slog"
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
	scriptHostPtr             = g.NewType("V8ScriptHost").Pointer()
)

const (
	dom      = BASE_PKG + "/browser/dom"
	html     = BASE_PKG + "/browser/html"
	v8host   = BASE_PKG + "/browser/v8host"
	gojahost = BASE_PKG + "/browser/gojahost"
	v8       = "github.com/tommie/v8go"
	gojaSrc  = "github.com/dop251/goja"
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
	wrapperTypeBaseName := dataData.WrapperTypeName
	if wrapperTypeBaseName == "" {
		wrapperTypeBaseName = fmt.Sprintf("%sV8Wrapper", wrappedTypeName)
	}
	return ESConstructorData{
		Spec:                dataData,
		InnerTypeName:       wrappedTypeName,
		WrapperTypeName:     lowerCaseFirstLetter(wrapperTypeBaseName),
		WrapperTypeBaseName: wrapperTypeBaseName,
		Receiver:            dataData.Receiver,
		RunCustomCode:       dataData.RunCustomCode,
		Inheritance:         idlName.Inheritance(),
		Constructor:         CreateConstructor(dataData, idlName),
		Operations:          CreateInstanceMethods(dataData, idlName),
		Attributes:          CreateAttributes(dataData, idlName),
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
		getter.Name = idlNameToGoName(getter.Name)
		if attribute.Readonly {
		} else {
			setter = new(ESOperation)
			*setter = *getter
			setter.Name = fmt.Sprintf("Set%s", getter.Name)
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
		esArgumentSpec := methodCustomization.Argument(arg.Name)
		esArg := ESOperationArgument{
			Name:         arg.Name,
			Optional:     arg.Optional && !esArgumentSpec.required,
			IdlType:      arg.IdlType,
			ArgumentSpec: esArgumentSpec,
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
	Name         string
	Type         string
	Optional     bool
	Variadic     bool
	IdlType      IdlTypes
	ArgumentSpec *ESMethodArgument
}

func (a ESOperationArgument) OptionalInGo() bool {
	hasDefault := a.ArgumentSpec != nil && a.ArgumentSpec.hasDefault
	return a.Optional && !hasDefault
}

func (a ESOperationArgument) DefaultValueInGo() (string, bool) {
	hasDefaultInGo := a.Optional && a.ArgumentSpec != nil && a.ArgumentSpec.hasDefault
	return fmt.Sprintf("default%s", a.Type), hasDefaultInGo
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

func (op ESOperation) HasResult() bool {
	return op.RetType.IsDefined()
}

type ESAttribute struct {
	Name   string
	Getter *ESOperation
	Setter *ESOperation
}

type ESConstructorData struct {
	Spec                *ESClassWrapper
	InnerTypeName       string
	WrapperTypeName     string
	WrapperTypeBaseName string
	Receiver            string
	Inheritance         string
	Operations          []ESOperation
	Attributes          []ESAttribute
	Constructor         *ESOperation
	RunCustomCode       bool
}

func (d ESConstructorData) WrapperFunctionsToInstall() iter.Seq[ESOperation] {
	return func(yield func(ESOperation) bool) {
		for _, op := range d.Operations {
			if !op.MethodCustomization.Ignored && !yield(op) {
				return
			}
		}
	}
}

func (d ESConstructorData) AttributesToInstall() iter.Seq[ESAttribute] {
	return func(yield func(ESAttribute) bool) {
		for _, a := range d.Attributes {
			if !yield(a) {
				return
			}
		}
	}
}

func (d ESConstructorData) WrapperFunctionsToGenerate() iter.Seq[ESOperation] {
	return func(yield func(ESOperation) bool) {
		for op := range d.WrapperFunctionsToInstall() {
			if !op.MethodCustomization.CustomImplementation && !yield(op) {
				return
			}
		}
		for _, a := range d.Attributes {
			if a.Getter != nil && !a.Getter.CustomImplementation {
				yield(*a.Getter)
			}
			if a.Setter != nil && !a.Setter.CustomImplementation {
				yield(*a.Setter)
			}
		}
	}
}

func (d ESConstructorData) Name() string { return d.Spec.TypeName }

func ReturnOnAnyError(errNames []g.Generator) g.Generator {
	if len(errNames) == 0 {
		return g.Noop
	}
	if len(errNames) == 1 {
		return ReturnOnError{err: errNames[0]}
	}
	return g.StatementList(
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

// sanitizeVarName create a valid go variable name from a variable to avoid
// invalid generated code due to
//
//   - The name is a reserved word, e.g. `type`.
//   - The name already an identifiers in scope (not yet implemented)
func sanitizeVarName(name string) string {
	switch name {
	case "type":
		return "type_"
	}
	return name
}

func idlNameToGoName(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = upperCaseFirstLetter(word)
	}
	return strings.Join(words, "")
}

func idlNameToUnexportedGoName(s string) string {
	return lowerCaseFirstLetter(idlNameToGoName(s))
}

func lowerCaseFirstLetter(s string) string {
	strLen := len(s)
	if strLen == 0 {
		slog.Warn("Passing empty string to upperCaseFirstLetter")
		return ""
	}
	buffer := make([]rune, 0, strLen)
	buffer = append(buffer, unicode.ToLower([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
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

type ReturnOnError struct {
	err g.Generator
}

func (ret ReturnOnError) Generate() *jen.Statement {
	err := ret.err
	if err == nil {
		err = g.Id("err")
	}
	return g.IfStmt{
		Condition: g.Neq{Lhs: err, Rhs: g.Nil}, //g.Raw(err.Generate().Op("!=").Nil()),
		Block:     g.Return(g.Nil, err),
	}.Generate()
}
