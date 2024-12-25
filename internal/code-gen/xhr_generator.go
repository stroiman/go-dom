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
		"timeout",
		"withCredentials",
		"upload",
		"responseURL",
		"response", // TODO, just because of the return value
		"responseType",
		"responseXML",
	)

	classWrapper.Method("open").SetCustomImplementation()
	classWrapper.Method("getResponseHeader").HasNoError = true
	classWrapper.Method("setRequestHeader").HasNoError = true

	data, err := createData(xhrData, classWrapper)
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}
