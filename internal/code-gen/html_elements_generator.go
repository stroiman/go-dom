package main

import (
	"encoding/json"
	"fmt"
)

func generateHtmlElements(file *builder) {
	output := ElementsJSON{}
	json.Unmarshal(html_defs, &output)
	WriteHeader(file)
	fmt.Fprint(file, "package scripting\n\n")
	fmt.Fprint(file, "// This file is generated. Do not edit.\n\n")
	fmt.Fprint(file, "var htmlElements = map[string]string {\n")
	file.indent()
	defer file.unIndentF("}\n")
	for _, element := range output.Elements {
		file.Printf("\"%s\": \"%s\",\n", element.Name, element.Interface)
	}
}
