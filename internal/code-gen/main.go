package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
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
	args := os.Args
	if len(args) != 3 || args[1] != "-o" {
		fmt.Println("Usage: code-gen -o <output_file>")
		fmt.Println(args[0])
		fmt.Println(args[1])
		fmt.Println(args[2])
		os.Exit(1 + len(args))
	}
	output := ElementsJSON{}
	json.Unmarshal(html_defs, &output)

	file, err := os.Create(args[2])
	if err != nil {
		fmt.Println("Error creating output file")
		os.Exit(1)
	}
	fmt.Fprint(file, "package scripting\n\n")
	fmt.Fprint(file, "var htmlElements = map[string]string {\n")
	for _, element := range output.Elements {
		fmt.Fprintf(file, "\t\"%s\": \"%s\",\n", element.Name, element.Interface)
	}
	fmt.Fprint(file, "}\n")
}
