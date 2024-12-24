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
	data, err := createData(urlData, ESClassWrapper{
		TypeName: "URL",
		// InnerTypeName:   "XmlHttpRequest",
		//WrapperTypeName: "ESXmlHttpRequest",
		Receiver: "u",
		Customization: []string{
			"SetHref",
			"SetProtocol",
			"username",
			"password",
			"SetHost",
			"SetPort",
			"SetHostname",
			"SetPathname",
			"searchParams",
			"SetHash",
			"SetSearch",
		},
	})
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}
