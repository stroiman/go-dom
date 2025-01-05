package main

import (
	_ "embed"
	"io"
)

//go:embed webref/curated/idlparsed/url.json
var urlData []byte

func generateUrl(writer io.Writer) error {
	wrapper := ESClassWrapper{
		TypeName: "URL",
		Receiver: "u",
	}
	wrapper.MarkMembersAsNotImplemented(
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
	)
	data := createData(urlData, wrapper)

	return writeGenerator(writer, data)
}
