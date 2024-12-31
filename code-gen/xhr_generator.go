package main

import (
	_ "embed"

	"github.com/dave/jennifer/jen"
)

//go:embed webref/curated/idlparsed/xhr.json
var xhrData []byte

func generateXhr(b *builder) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	classWrapper := ESClassWrapper{
		TypeName:        "XMLHttpRequest",
		InnerTypeName:   "XmlHttpRequest",
		WrapperTypeName: "ESXmlHttpRequest",
		Receiver:        "xhr",
	}
	classWrapper.MarkMembersAsNotImplemented(
		"readyState",
		"responseType",
		"responseXML",
	)

	classWrapper.Method("open").SetCustomImplementation()
	classWrapper.Method("upload").SetCustomImplementation()
	classWrapper.Method("getResponseHeader").HasNoError = true
	classWrapper.Method("setRequestHeader").HasNoError = true

	data := createData(xhrData, classWrapper)

	WriteGenerator(file, data)
	return file.Render(b)
}
