package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed webref/ed/elements/html.json
var html_defs []byte

type ElementJSON struct {
	Name      string `json:"name"`
	Interface string `json:"interface"`
}

type ElementsJSON struct {
	Elements []ElementJSON `json:"elements"`
}

func main() {
	output := ElementsJSON{}
	json.Unmarshal(html_defs, &output)
	fmt.Println("DEFS", string(html_defs))
	fmt.Print("package scripting\n\n")
	fmt.Print("var htmlElements = map[string]string {\n")
	for _, element := range output.Elements {
		fmt.Printf("\t\"%s\": \"%s\",\n", element.Name, element.Interface)
	}
	fmt.Println("}")
}
