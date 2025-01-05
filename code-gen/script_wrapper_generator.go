package main

import (
	"io"

	"github.com/dave/jennifer/jen"
	g "github.com/stroiman/go-dom/code-gen/generators"
)

func writeGenerator(writer io.Writer, generator g.Generator) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	file.Add(generator.Generate())
	return file.Render(writer)
}
