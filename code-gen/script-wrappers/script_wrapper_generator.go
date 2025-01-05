package wrappers

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/dave/jennifer/jen"
	g "github.com/stroiman/go-dom/code-gen/generators"
	"github.com/stroiman/go-dom/code-gen/idl"
)

// WrapperGeneratorsSpec is a list of specifications for generating ES wrapper
// code. Each key in the map correspond to a specific IDL file
type WrapperGeneratorsSpec map[string](*WrapperGeneratorFileSpec)

func NewWrapperGeneratorsSpec() WrapperGeneratorsSpec {
	return make(WrapperGeneratorsSpec)
}

type WrapperTypeSpec = *ESClassWrapper

type WrapperGeneratorFileSpec struct {
	Name  string
	Types map[string]WrapperTypeSpec
}

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

func writeGenerator(writer io.Writer, generator g.Generator) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	file.Add(generator.Generate())
	return file.Render(writer)
}

type ScriptWrapperModulesGenerator struct {
	IdlSources fs.FS
	Specs      WrapperGeneratorsSpec
}

func (gen ScriptWrapperModulesGenerator) writeModule(
	writer io.Writer,
	spec *WrapperGeneratorFileSpec,
) error {
	filename := fmt.Sprintf("webref/curated/idlparsed/%s.json", spec.Name)
	file, err := gen.IdlSources.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := idl.ParseIdlJsonReader(file)
	if err != nil {
		return err
	}
	generators := StatementList()
	for _, specType := range spec.Types {
		generators.Append(createData(data, *specType))
		generators.Append(g.Line)
	}
	return writeGenerator(writer, generators)
}

func (gen ScriptWrapperModulesGenerator) writeModules(specs WrapperGeneratorsSpec) error {
	errs := make([]error, 0, len(specs))
	for name, spec := range specs {
		outputFileName := fmt.Sprintf("%s_generated.go", name)
		err := func() error {
			if writer, err := os.Create(outputFileName); err != nil {
				return err
			} else {
				defer writer.Close()
				return gen.writeModule(writer, spec)
			}
		}()
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func (s *WrapperGeneratorFileSpec) Type(typeName string) WrapperTypeSpec {
	if result, ok := s.Types[typeName]; ok {
		return result
	}
	result := new(ESClassWrapper)
	result.ensureMap()
	result.TypeName = typeName
	s.Types[typeName] = result
	return result
}

func NewScriptWrapperModulesGenerator(idlSources fs.FS) ScriptWrapperModulesGenerator {
	specs := NewWrapperGeneratorsSpec()
	xhrModule := specs.Module("xhr")
	xhr := xhrModule.Type("XMLHttpRequest")

	xhr.InnerTypeName = "XmlHttpRequest"
	xhr.WrapperTypeName = "ESXmlHttpRequest"
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

	urlSpecs := specs.Module("url")
	url := urlSpecs.Type("URL")
	url.Receiver = "u"
	url.MarkMembersAsNotImplemented(
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
	)

	domSpecs := specs.Module("dom")
	domTokenList := domSpecs.Type("DOMTokenList")
	domTokenList.Receiver = "u"
	domTokenList.RunCustomCode = true
	domTokenList.Method("item").SetNoError()
	domTokenList.Method("contains").SetNoError()
	domTokenList.Method("remove").SetNoError()
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("replace").SetNoError()
	domTokenList.Method("supports").SetNotImplemented()

	htmlSpecs := specs.Module("html")
	htmlTemplateElement := htmlSpecs.Type("HTMLTemplateElement")
	htmlTemplateElement.Receiver = "e"
	htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()

	return ScriptWrapperModulesGenerator{
		IdlSources: idlSources,
		Specs:      specs,
	}
}

func (gen ScriptWrapperModulesGenerator) GenerateScriptWrappers() error {
	return gen.writeModules(gen.Specs)
}
