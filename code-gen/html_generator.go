package main

import (
	_ "embed"
	"io"

	"github.com/dave/jennifer/jen"
)

//go:embed webref/curated/idlparsed/html.json
var htmlData []byte

func generateHTMLTypes(writer io.Writer) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")

	htmlTemplateElement := ESClassWrapper{
		TypeName: "HTMLTemplateElement",
		Receiver: "e",
	}
	htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()

	htmlTemplateData := createData(htmlData, htmlTemplateElement)

	WriteGenerator(file, htmlTemplateData)
	return file.Render(writer)
}
