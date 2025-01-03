package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func generateHtmlElements(writer io.Writer) error {
	file := newBuilder(writer)
	output := ElementsJSON{}
	json.Unmarshal(html_defs, &output)
	WriteHeader(file)
	fmt.Fprint(file, "var htmlElements = map[string]string {\n")
	file.indent()
	defer file.unIndentF("}\n")
	for _, element := range output.Elements {
		file.Printf("\"%s\": \"%s\",\n", element.Name, element.Interface)
	}
	return nil
}
