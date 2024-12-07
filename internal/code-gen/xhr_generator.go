package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"
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

func generateXhr(b *builder) {
	spec := ParsedIdlFile{}
	err := json.Unmarshal(xhrData, &spec)
	if err != nil {
		panic(err)
	}
	WriteHeader(b)
	WriteImports(b, [][][2]string{
		{{"", "github.com/stroiman/go-dom/browser"}},
		{{"v8", "github.com/tommie/v8go"}},
	})

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
					esArg.Name = arg.IdlType.IdlType.IType.TypeName
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
	fmt.Println(operations)
	os.Exit(1)

	data := ESConstructorData{
		InnerTypeName:      "XmlHttpRequest",
		CreatesInnerType:   true,
		CustomConstruction: "scriptContext.Window().NewXmlHttpRequest()",
		IdlName:            idlName,
	}

	WriteFactoryFunction(b, data)
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

func WriteFactoryFunction(b *builder, data ESConstructorData) {
	funcName := fmt.Sprintf("Create%sPrototype", data.InnerTypeName)
	qualifiedType := fmt.Sprintf("%s.%s", "browser", data.InnerTypeName)
	newInnerFuncName := data.CustomConstruction
	b.Printf("func %s(host *ScriptHost) *v8.FunctionTemplate {\n", funcName)
	b.indent()
	defer b.unIndentF("}")
	b.Printf("// iso := host.iso\n")
	b.Printf("builder := NewConstructorBuilder[%s](\n", qualifiedType)
	b.indent()
	b.Printf("host,\n")
	b.Printf("func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {\n")
	b.indent()
	b.Printf("scriptContext := host.MustGetContext(info.Context())\n")
	b.Printf("instance := %s\n", newInnerFuncName)
	b.Printf("return scriptContext.CacheNode(info.This(), instance)\n")
	b.unIndentF("},\n")
	b.unIndentF(")\n")
	b.Printf("builder.SetDefaultInstanceLookup()\n")
	b.Printf("protoBuilder := builder.NewPrototypeBuilder()\n")
	WriteOperations(b, data)
	b.Printf("return builder.constructor\n")

}

func WriteOperations(b *builder, data ESConstructorData) {
	for _, member := range data.IdlName.Members {
		if member.Type == "operation" {
			b.Printf("%s: %s\n", member.Name, member.Type)

			b.indent()
			for _, a := range member.Arguments {
				b.Printf("%s: %s - %v\n", a.Name, a.Type, a.IdlType)
			}
			b.unIndent()
		}
	}
}
