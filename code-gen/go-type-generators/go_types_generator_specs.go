package Generators

import (
	"fmt"
	"io"
	"io/fs"

	"github.com/dave/jennifer/jen"
)

type GoTypeGenerators struct {
	IdlSources fs.FS
	Specs      BrowserModuleSpecs
}

type BrowserModuleSpecs map[string](*BrowserModuleSpec)

type BrowserModuleSpec struct {
	Name  string
	Types map[string]BrowserTypeSpec
}

type BrowserTypeSpec struct {
	Name string
}

func (gen GoTypeGenerators) GenerateType(module string, typeName string, writer io.Writer) error {
	moduleName := fmt.Sprintf("github.com/stroiman/go-dom/browser/%s", module)
	file := jen.NewFilePath(moduleName)
	return file.Render(writer)
}

func NewGoTypeGenerators(idlSources fs.FS) GoTypeGenerators {
	return GoTypeGenerators{
		IdlSources: idlSources,
	}
}
