package main

import (
	_ "embed"
	"io"
)

//go:embed webref/curated/idlparsed/dom.json
var domData []byte

func generateDOMTypes(writer io.Writer) error {
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

	return writeGenerator(writer, domTokenListData)
}
