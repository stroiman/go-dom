package main

import (
	_ "embed"

	"github.com/dave/jennifer/jen"
)

//go:embed webref/curated/idlparsed/url.json
var urlData []byte

func generateUrl(b *builder) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	data, err := createData(urlData, "URL", CreateDataData{
		// InnerTypeName:   "XmlHttpRequest",
		//WrapperTypeName: "ESXmlHttpRequest",
		Receiver: "u",
	})
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}
