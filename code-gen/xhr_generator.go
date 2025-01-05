package main

import (
	_ "embed"
	"io"
)

//go:embed webref/curated/idlparsed/xhr.json
var xhrData []byte

func generateXhr(writer io.Writer) error {
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

	return writeGenerator(writer, data)
}
