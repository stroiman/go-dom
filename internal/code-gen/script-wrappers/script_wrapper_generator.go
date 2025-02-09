package wrappers

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

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
	Specs            configuration.WrapperGeneratorsSpec
	PackagePath      string
	TargetGenerators TargetGenerators
}

func (gen ScriptWrapperModulesGenerator) writeModule(
	writer io.Writer,
	spec *configuration.WrapperGeneratorFileSpec,
) error {
	data, err := idl.Load(spec.Name)
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
	spec *configuration.WrapperGeneratorFileSpec,
) error {
	data, err := idl.Load(spec.Name)
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

func (gen ScriptWrapperModulesGenerator) writeModules(
	specs configuration.WrapperGeneratorsSpec,
) error {
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
	spec *configuration.WrapperGeneratorFileSpec,
) error {
	outputFileName := fmt.Sprintf("%s_generated.go", spec.Name)
	if writer, err := os.Create(outputFileName); err != nil {
		return err
	} else {
		defer writer.Close()
		return gen.writeModule(writer, spec)
	}
}

func (gen ScriptWrapperModulesGenerator) GenerateScriptWrappers() error {
	return gen.writeModules(gen.Specs)
}

func NewScriptWrapperModulesGenerator() ScriptWrapperModulesGenerator {
	specs := configuration.CreateV8Specs()

	return ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      v8host,
		TargetGenerators: V8TargetGenerators{},
	}
}
