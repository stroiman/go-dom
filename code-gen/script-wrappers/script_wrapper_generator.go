package wrappers

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stroiman/go-dom/code-gen/generators"
	g "github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/webref/idl"
)

// WrapperGeneratorsSpec is a list of specifications for generating ES wrapper
// code. Each key in the map correspond to a specific IDL file
type WrapperGeneratorsSpec map[string](*WrapperGeneratorFileSpec)

func NewWrapperGeneratorsSpec() WrapperGeneratorsSpec {
	return make(WrapperGeneratorsSpec)
}

type WrapperTypeSpec = *ESClassWrapper

func (s *ESClassWrapper) CreateWrapper() {
	s.WrapperStruct = true
}

type WrapperGeneratorFileSpec struct {
	Name          string
	MultipleFiles bool
	Types         map[string]WrapperTypeSpec
}

func (spec WrapperGeneratorFileSpec) GetTypesSorted() []WrapperTypeSpec {
	types := make([]WrapperTypeSpec, len(spec.Types))
	idx := 0
	for _, t := range spec.Types {
		types[idx] = t
		idx++
	}
	slices.SortFunc(types, func(x, y WrapperTypeSpec) int {
		return cmp.Compare(x.TypeName, y.TypeName)
	})
	return types
}

func (spec WrapperGeneratorFileSpec) UseMultipleFiles() bool {
	return spec.MultipleFiles == true
}

func (spec *WrapperGeneratorFileSpec) SetMultipleFiles(value bool) { spec.MultipleFiles = value }

func (g WrapperGeneratorsSpec) Module(spec string) *WrapperGeneratorFileSpec {
	if mod, ok := g[spec]; ok {
		return mod
	}
	mod := &WrapperGeneratorFileSpec{
		Name:  spec,
		Types: make(map[string]WrapperTypeSpec),
	}
	g[spec] = mod
	return mod
}

func writeGenerator(writer io.Writer, packagePath string, generator g.Generator) error {
	file := jen.NewFilePath(packagePath)
	file.HeaderComment("This file is generated. Do not edit.")
	// file.ImportName(dom, "browser")
	file.ImportAlias(v8, "v8")
	file.ImportAlias(gojaSrc, "g")
	file.Add(generator.Generate())
	return file.Render(writer)
}

type TargetGenerators interface {
	CreateJSConstructorGenerator(data ESConstructorData) g.Generator
}

type ScriptWrapperModulesGenerator struct {
	Specs            WrapperGeneratorsSpec
	PackagePath      string
	TargetGenerators TargetGenerators
}

func (gen ScriptWrapperModulesGenerator) writeModule(
	writer io.Writer,
	spec *WrapperGeneratorFileSpec,
) error {
	data, err := idl.LoadIdlParsed(spec.Name)
	if err != nil {
		return err
	}
	generators := g.StatementList()
	for _, specType := range spec.GetTypesSorted() {
		typeGenerationInformation := createData(data, specType)
		generators.Append(
			gen.TargetGenerators.CreateJSConstructorGenerator(typeGenerationInformation),
		)
		generators.Append(g.Line)
	}
	return writeGenerator(writer, gen.PackagePath, generators)
}

func (gen ScriptWrapperModulesGenerator) writeModuleTypes(
	spec *WrapperGeneratorFileSpec,
) error {
	data, err := idl.LoadIdlParsed(spec.Name)
	if err != nil {
		return err
	}
	// generators := g.StatementList()
	types := spec.GetTypesSorted()
	errs := make([]error, len(types))
	for i, specType := range types {
		outputFileName := fmt.Sprintf("%s_generated.go", typeNameToFileName(specType.TypeName))
		if writer, err := os.Create(outputFileName); err != nil {
			errs[i] = err
		} else {
			typeGenerationInformation := createData(data, specType)
			errs[i] = writeGenerator(writer, gen.PackagePath, gen.TargetGenerators.CreateJSConstructorGenerator(typeGenerationInformation))
		}
	}
	return errors.Join(errs...)
}

var matchKnownWord = regexp.MustCompile("(HTML|URL|DOM)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")

func typeNameToFileName(name string) string {
	snake := matchKnownWord.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (gen ScriptWrapperModulesGenerator) writeModules(specs WrapperGeneratorsSpec) error {
	errs := make([]error, len(specs))
	i := 0
	for _, spec := range specs {
		if spec.UseMultipleFiles() {
			errs[i] = gen.writeModuleTypes(spec)
		} else {
			errs[i] = gen.writeModuleSingleFile(spec)
		}
		i++
	}
	return errors.Join(errs...)
}

func (gen ScriptWrapperModulesGenerator) writeModuleSingleFile(
	spec *WrapperGeneratorFileSpec,
) error {
	outputFileName := fmt.Sprintf("%s_generated.go", spec.Name)
	if writer, err := os.Create(outputFileName); err != nil {
		return err
	} else {
		defer writer.Close()
		return gen.writeModule(writer, spec)
	}
}

func (s *WrapperGeneratorFileSpec) Type(typeName string) WrapperTypeSpec {
	if result, ok := s.Types[typeName]; ok {
		return result
	}
	result := &ESClassWrapper{
		DomSpec:  s,
		TypeName: typeName,
		Receiver: generators.DefaultReceiverName(typeName),
	}
	result.ensureMap()
	s.Types[typeName] = result
	return result
}

func CreateSpecs() WrapperGeneratorsSpec {
	specs := NewWrapperGeneratorsSpec()
	domSpecs := specs.Module("dom")
	domSpecs.SetMultipleFiles(true)
	domNode := domSpecs.Type("Node")
	domNode.Method("nodeType").SetCustomImplementation()
	domNode.Method("contains").SetNoError()
	domNode.Method("getRootNode").SetNoError().Argument("options").HasDefault()
	domNode.Method("previousSibling").SetNoError()
	domNode.Method("nextSibling").SetNoError()

	domNode.Method("hasChildNodes").Ignore()
	domNode.Method("normalize").Ignore()
	domNode.Method("cloneNode").Ignore()
	domNode.Method("isEqualNode").Ignore()
	domNode.Method("isSameNode").Ignore()
	domNode.Method("compareDocumentPosition").Ignore()
	domNode.Method("lookupPrefix").Ignore()
	domNode.Method("lookupNamespaceURI").Ignore()
	domNode.Method("isDefaultNamespace").Ignore()
	domNode.Method("replaceChild").Ignore()
	domNode.Method("baseURI").Ignore()
	domNode.Method("parentNode").Ignore()
	domNode.Method("lastChild").Ignore()
	domNode.Method("nodeValue").Ignore()
	domNode.Method("textContent").Ignore()

	return specs
}

func NewScriptWrapperModulesGenerator() ScriptWrapperModulesGenerator {
	specs := CreateSpecs()
	xhrModule := specs.Module("xhr")
	xhr := xhrModule.Type("XMLHttpRequest")
	xhr.SkipPrototypeRegistration = true
	xhr.InnerTypeName = "XmlHttpRequest"
	xhr.Receiver = "xhr"

	xhr.MarkMembersAsNotImplemented(
		"readyState",
		"responseType",
		"responseXML",
	)
	xhr.Method("open").SetCustomImplementation()
	xhr.Method("upload").SetCustomImplementation()
	xhr.Method("getResponseHeader").HasNoError = true
	xhr.Method("setRequestHeader").HasNoError = true
	xhr.Method("onreadystatechange").Ignore()

	urlSpecs := specs.Module("url")
	url := urlSpecs.Type("URL")
	url.InnerTypeName = "Url"
	url.Receiver = "u"
	url.MarkMembersAsNotImplemented(
		"setHref",
		"setProtocol",
		"username",
		"password",
		"setHost",
		"setPort",
		"setHostname",
		"setPathname",
		"searchParams",
		"setHash",
		"setSearch",
	)

	domSpecs := specs.Module("dom")

	event := domSpecs.Type("Event")
	event.CreateWrapper()
	event.Method("constructor").Argument("eventInitDict").HasDefault()
	event.Method("initEvent").Ignore()
	event.Method("composed").Ignore()
	event.Method("composedPath").Ignore()
	event.Method("stopPropagation").SetNoError()
	event.Method("stopImmediatePropagation").Ignore()
	event.Method("preventDefault").SetNoError()
	event.Method("isTrusted").Ignore()
	event.Method("CancelBubble").Ignore()
	event.Method("cancelBubble").Ignore()
	event.Method("EventPhase").Ignore()
	event.Method("eventPhase").Ignore()
	event.Method("TimeStamp").Ignore()
	event.Method("timeStamp").Ignore()
	event.Method("ReturnValue").Ignore()
	event.Method("returnValue").Ignore()
	event.Method("srcElement").Ignore()
	event.Method("defaultPrevented").Ignore()

	domElement := domSpecs.Type("Element")
	domElement.RunCustomCode = true
	domElement.Method("getAttribute").SetCustomImplementation()
	domElement.Method("setAttribute").SetNoError()
	domElement.Method("hasAttribute").SetNoError()
	domElement.Method("classList").SetCustomImplementation()
	domElement.Method("matches")

	domElement.MarkMembersAsNotImplemented(
		"hasAttributes",
		"hasAttributeNS",
		"getAttributeNames",
		"getAttributeNS",
		"setAttributeNS",
		"removeAttributeNode",
		"removeAttribute",
		"removeAttributeNS",
		"toggleAttribute",
		"toggleAttributeForce",
		"setAttributeNode",
		"setAttributeNodeNS",
		"getAttributeNode",
		"getAttributeNodeNS",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getElementsByClassName",
		"insertAdjacentElement",
		"insertAdjacentText",
		"namespaceURI",
		"prefix",
		"localName",
		"id",
		"shadowRoot",
		"slot",
		"className",
		"decodeShadowRootInit",
		"attachShadow",
	)

	domElement.MarkMembersAsIgnored(
		// HTMX fails if these exist but throw
		"webkitMatchesSelector",
		"closest",
	)

	domTokenList := domSpecs.Type("DOMTokenList")
	domTokenList.InnerTypeName = "DomTokenList"
	domTokenList.Receiver = "u"
	domTokenList.RunCustomCode = true
	domTokenList.Method("item").SetNoError()
	domTokenList.Method("contains").SetNoError()
	domTokenList.Method("remove").SetNoError()
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("replace").SetNoError()
	domTokenList.Method("supports").SetNotImplemented()

	htmlSpecs := specs.Module("html")
	htmlSpecs.SetMultipleFiles(true)

	htmlTemplateElement := htmlSpecs.Type("HTMLTemplateElement")
	htmlTemplateElement.InnerTypeName = "HtmlTemplateElement"
	htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()

	window := htmlSpecs.Type("Window")
	window.InnerTypeName = "Window"
	window.CreateWrapper()

	window.Method("window").SetCustomImplementation()
	window.Method("location").Ignore()
	window.Method("parent").Ignore() // On `Node`

	window.Method("prompt").SetNotImplemented()
	window.Method("close").SetNotImplemented()
	window.Method("stop").SetNotImplemented()
	window.Method("focus").SetNotImplemented()
	window.Method("blur").SetNotImplemented()
	window.Method("open").SetNotImplemented()
	window.Method("alert").SetNotImplemented()
	window.Method("confirm").SetNotImplemented()
	window.Method("postMessage").SetNotImplemented()
	window.Method("print").SetNotImplemented()
	window.Method("self").SetNotImplemented()
	window.Method("name").SetNotImplemented()
	window.Method("personalbar").SetNotImplemented()
	window.Method("locationbar").SetNotImplemented()
	window.Method("menubar").SetNotImplemented()
	window.Method("scrollbars").SetNotImplemented()
	window.Method("statusbar").SetNotImplemented()
	window.Method("status").SetNotImplemented()
	window.Method("toolbar").SetNotImplemented()
	window.Method("history").SetNotImplemented()
	window.Method("navigation").SetNotImplemented()
	window.Method("customElements").SetNotImplemented()
	window.Method("closed").SetNotImplemented()
	window.Method("frames").SetNotImplemented()
	window.Method("navigator").SetNotImplemented()
	window.Method("frames").SetNotImplemented()
	window.Method("top").SetNotImplemented()
	window.Method("opener").SetNotImplemented()
	window.Method("frameElement").SetNotImplemented()
	window.Method("clientInformation").SetNotImplemented()
	window.Method("originAgentCluster").SetNotImplemented()
	window.Method("length").SetNotImplemented()

	history := htmlSpecs.Type("History")
	history.Method("go").Argument("delta").HasDefaultValue("defaultDelta")
	history.Method("pushState").Argument("url").HasDefaultValue("defaultUrl")
	history.Method("pushState").Argument("unused").Ignore()
	history.Method("replaceState").Argument("url").HasDefaultValue("defaultUrl")
	history.Method("replaceState").Argument("unused").Ignore()
	history.Method("scrollRestoration").Ignore()
	history.Method("state").SetEncoder("toJSON")

	return ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      v8host,
		TargetGenerators: V8TargetGenerators{},
	}
}

func (gen ScriptWrapperModulesGenerator) GenerateScriptWrappers() error {
	return gen.writeModules(gen.Specs)
}
