package main

import (
	_ "embed"
	"io"

	"github.com/dave/jennifer/jen"
)

//go:embed webref/curated/idlparsed/dom.json
var domData []byte

func generateDOMTypes(writer io.Writer) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")

	domTokenList := ESClassWrapper{
		TypeName:      "DOMTokenList",
		Receiver:      "u",
		RunCustomCode: true,
	}
	domTokenList.Method("item").SetNoError()
	domTokenList.Method("contains").SetNoError()
	domTokenList.Method("remove").SetNoError()
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("replace").SetNoError()
	domTokenList.Method("supports").SetNotImplemented()

	domTokenListData := createData(domData, domTokenList)

	WriteGenerator(file, domTokenListData)
	return file.Render(writer)
}
