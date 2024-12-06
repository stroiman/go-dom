package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type ExtAttr struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Rhs  struct {
		Type  string `json:"type"`
		Value []struct {
			Value string `json:"value"`
		} `json:"value"`
	} `json:"rhs"`
}

type IdlType struct {
	Type     string    `json:"type"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Generic  string    `json:"generic"`
	Nullable bool      `json:"nullable"`
	Union    bool      `json:"union"`
	IdlType  string    `json:"idlType"`
}

type Stuff struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	IdlType  `          json:"idlType"`
}

type IdlNameMember struct {
	Stuff
	Arguments []struct {
		Stuff
		Default  any  `json:"default"`
		Optional bool `json:"optional"`
		Variadic bool `json:"variadic"`
	} `json:"arguments"`
	Special  string `json:"special"`
	Readonly bool   `json:"readOnly"`
	Href     string `json:"href"`
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
	json.Unmarshal(xhrData, &spec)
	WriteHeader(b)
	WriteImports(b, [][][2]string{
		{{"", "github.com/stroiman/go-dom/browser"}},
		{{"v8", "github.com/tommie/v8go"}},
	})

	WriteFactoryFunction(b, ESConstructorData{
		InnerTypeName:      "XmlHttpRequest",
		CreatesInnerType:   true,
		CustomConstruction: "scriptContext.Window().NewXmlHttpRequest()",
	})
}

type ESConstructorData struct {
	CreatesInnerType   bool
	InnerTypeName      string
	CustomConstruction string
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
	b.Printf("return builder.constructor\n")

}
