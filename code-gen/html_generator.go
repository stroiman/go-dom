package main

import (
	_ "embed"
	"io"
)

//go:embed webref/curated/idlparsed/html.json
var htmlData []byte

func generateHTMLTypes(writer io.Writer) error {
	htmlTemplateElement := ESClassWrapper{
		TypeName: "HTMLTemplateElement",
		Receiver: "e",
	}
	htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()

	htmlTemplateData := createData(htmlData, htmlTemplateElement)

	return writeGenerator(writer, htmlTemplateData)
}
