package main

import (
	_ "embed"

	"github.com/dave/jennifer/jen"
)

//go:embed webref/curated/idlparsed/dom.json
var domData []byte

func generateDOMTypes(b *builder) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")

	wrapper := ESClassWrapper{
		TypeName: "DOMTokenList",
		Receiver: "u",
	}

	wrapper.Method("item").SetNoError()
	wrapper.Method("contains").SetNoError()
	wrapper.Method("remove").SetNotImplemented()
	wrapper.Method("toggle").SetNotImplemented()
	wrapper.Method("replace").SetNotImplemented()
	wrapper.Method("supports").SetNotImplemented()

	data, err := createData(domData, wrapper)
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}
