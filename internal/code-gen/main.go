package main

import (
	_ "embed"
	"encoding/json"
	"flag"
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
	outputFile := flag.String("o", "", "Output file to write")
	flag.Parse()
	if *outputFile == "" {
		fmt.Println("Internal code generator from IDL definitions")
		flag.PrintDefaults()
		os.Exit(1)
	}
	output := ElementsJSON{}
	json.Unmarshal(html_defs, &output)

	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error creating output file")
		os.Exit(1)
	}
	fmt.Fprint(file, "package scripting\n\n")
	fmt.Fprint(file, "// This file is generated. Do not edit.\n\n")
	fmt.Fprint(file, "var htmlElements = map[string]string {\n")
	for _, element := range output.Elements {
		fmt.Fprintf(file, "\t\"%s\": \"%s\",\n", element.Name, element.Interface)
	}
	fmt.Fprint(file, "}\n")
}
