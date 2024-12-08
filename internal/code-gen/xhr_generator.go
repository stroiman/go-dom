package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"unicode"

	"github.com/dave/jennifer/jen"
)

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

//go:embed webref/curated/idlparsed/xhr.json
var xhrData []byte

func createData() (ESConstructorData, error) {
	spec := ParsedIdlFile{}
	err := json.Unmarshal(xhrData, &spec)
	if err != nil {
		panic(err)
	}
	idlName := spec.IdlNames["XMLHttpRequest"]
	operations := []ESOperation{}
	for _, member := range idlName.Members {
		if member.Type == "operation" {
			if -1 != slices.IndexFunc(
				operations,
				func(op ESOperation) bool { return op.Name == member.Name },
			) {
				slog.Warn("Function overloads", "Name", member.Name)
				continue
			}
			arguments := []ESOperationArgument{}
			argumentsOk := true
			for _, arg := range member.Arguments {
				esArg := ESOperationArgument{
					Name: arg.Name,
				}
				if len(arg.IdlType.Types) > 0 {
					slog.Warn(
						"Multiple argument types",
						"Operation",
						member.Name,
						"Argument",
						arg.Name,
					)
					argumentsOk = false
					break
				}
				if arg.IdlType.IdlType != nil {
					esArg.Type = arg.IdlType.IdlType.IType.TypeName
				}
				arguments = append(arguments, esArg)
			}
			if argumentsOk {
				operation := ESOperation{
					Name:      member.Name,
					Arguments: arguments,
				}
				operations = append(operations, operation)
			}
		}
	}
	// fmt.Println(operations)
	// os.Exit(1)

	data := ESConstructorData{
		InnerTypeName:      "XmlHttpRequest",
		Operations:         operations,
		CreatesInnerType:   true,
		CustomConstruction: "scriptContext.Window().NewXmlHttpRequest()",
		IdlName:            idlName,
	}
	return data, nil
}

const br = "github.com/stroiman/go-dom/browser"
const sc = "github.com/stroiman/go-dom/scripting"
const v8 = "github.com/tommie/v8go"

func generateXhr(b *builder) error {

	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	data, err := createData()
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}

type ESOperationArgument struct {
	Name     string
	Type     string
	Optional bool
	Variadic bool
}

type ESOperation struct {
	Name      string
	Arguments []ESOperationArgument
}

type ESConstructorData struct {
	CreatesInnerType   bool
	InnerTypeName      string
	CustomConstruction string
	Operations         []ESOperation
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
	varScriptContext := jen.Id("scriptContext")
	return JSConstructor{
		argHost,
		argInfo,
		getThis,
		varInstance,
		varIso,
		varScriptContext,
	}
}

func (c JSConstructor) JSConstructorImpl(grp *jen.Group) {
	buildInstance := c.varScriptContext.Clone().
		Dot("Window").
		Call().
		Dot("NewXmlHttpRequest").
		Call()
	grp.Add(c.varScriptContext).
		Op(":=").
		Add(c.argHost).
		Dot("MustGetContext").
		Call(c.argInfo.Clone().Dot("Context").Call())
	grp.Add(c.varInstance).Op(":=").Add(buildInstance)
	grp.Return(c.varScriptContext.Clone().Dot("CacheNode").Call(
		c.getThis,
		c.varInstance,
	))
}

func (c JSConstructor) Run(f *jen.File, data ESConstructorData) {
	hostType := jen.Id("ScriptHost")
	hostPtr := jen.Add(jen.Op("*"), hostType)
	builder := jen.Id("builder")
	v8FunctionTemplatePtr := jen.Op("*").Qual(v8, "FunctionTemplate")
	v8Value := jen.Op("*").Qual(v8, "Value")
	errorT := jen.Id("error")
	g := Helper{&jen.Group{}}
	f.Func().
		Id(fmt.Sprintf("Create%sPrototype", data.InnerTypeName)).
		Params(c.argHost.Clone().Add(hostPtr)).Add(v8FunctionTemplatePtr).
		BlockFunc(func(grp *jen.Group) {
			grp.Add(c.varIso).Op(":=").Add(jen.Id("host")).Dot("iso").GoString()
			grp.Add(builder).
				Op(":=").
				Id("NewConstructorBuilder").Index(jen.Qual(br, data.InnerTypeName)).
				Call(c.argHost.Clone(), jen.Func().
					Params(c.argInfo.Clone().Add(g.v8FunctionCallbackInfoPtr())).
					Params(v8Value, errorT).
					BlockFunc(c.JSConstructorImpl))
			grp.Id("protoBuilder").Op(":=").Add(builder).Dot("NewPrototypeBuilder").Call()
			grp.Id("prototype").Op(":=").Id("protoBuilder").Dot("proto")
			grp.Line()
			for _, op := range data.Operations {
				c.CreateOperation(grp, op)
			}

			grp.Add(builder).Dot("SetDefaultInstanceLookup").Call()
			grp.Return(builder.Clone().Dot("constructor"))
		})
}

func (c JSConstructor) CreateOperation(grp *jen.Group, op ESOperation) {
	v8Value := jen.Op("*").Qual(v8, "Value")
	errorT := jen.Id("error")
	v8FunctionCallbackInfoPtr := jen.Op("*").Qual(v8, "FunctionCallbackInfo")
	f := jen.Func().
		Params(c.argInfo.Clone().Add(v8FunctionCallbackInfoPtr)).
		Params(v8Value, errorT).
		BlockFunc(func(grp *jen.Group) {
			grp.Id("args").Op(":=").Id("info").Dot("Args").Call()
			argCount := len(op.Arguments)
			var args []jen.Code
			for i, arg := range op.Arguments {
				v := jen.Id(arg.Name) //fmt.Sprintf("arg%d", i))
				args = append(args, v)
				var e *jen.Statement
				if argCount > 1 {
					e = jen.Id(fmt.Sprintf("err%d", i))
				} else {
					e = jen.Id(fmt.Sprintf("err"))
				}
				grp.Add(v).
					Op(",").
					Add(e).
					Op(":=").
					Id("GetArg"+arg.Type).
					Call(jen.Id("args"), jen.Lit(i))
			}
			WriteErrorHandler(grp, len(op.Arguments))
			grp.Add(jen.Id("instance")).Op(",").Id("err").
				Op(":=").
				Id("builder").
				Dot("GetInstance").
				Call(jen.Id("info"))
			WriteReturnOnError(grp)
			grp.Id("instance").Dot(camelCase(op.Name)).Call(args...)
			grp.Return(jen.Nil(), jen.Nil())
		})
	ft := jen.Qual(v8, "NewFunctionTemplateWithError").Call(c.varIso, f)
	grp.Id("prototype").Dot("Set").Call(jen.Lit(op.Name), ft)
	grp.Line()
}

func camelCase(s string) string {
	buffer := make([]rune, 0, len(s))
	buffer = append(buffer, unicode.ToUpper([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}

func WriteReturnOnError(grp *jen.Group) {
	jErr := jen.Id("err")
	grp.If(jErr.Clone().Op("!=").Nil()).Block(
		jen.Return(jen.Nil(), jErr),
	)
}

func WriteErrorHandler(grp *jen.Group, count int) {
	if count == 0 {
		return
	}
	jErr := jen.Id("err")
	if count > 1 {
		var args []jen.Code
		for i := 0; i < count; i++ {
			args = append(args, jen.Id(fmt.Sprintf("err%d", i)))
		}
		grp.Add(jErr).Op(":=").Qual("errors", "Join").Call(args...)
	}
	WriteReturnOnError(grp)
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

// func WriteOperations(b *builder, data ESConstructorData) {
// 	for _, member := range data.IdlName.Members {
// 		if member.Type == "operation" {
// 			b.Printf("%s: %s\n", member.Name, member.Type)
//
// 			b.indent()
// 			for _, a := range member.Arguments {
// 				b.Printf("%s: %s - %v\n", a.Name, a.Type, a.IdlType)
// 			}
// 			b.unIndent()
// 		}
// 	}
// }
