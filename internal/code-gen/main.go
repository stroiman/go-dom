package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	generatorType := flag.String("g", "", "Generator type")
	flag.Parse()
	if *outputFile == "" || *generatorType == "" {
		fmt.Println("Internal code generator from IDL definitions")
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error creating output file")
		os.Exit(1)
	}

	switch *generatorType {
	case "html-elements":
		generateHtmlElements(file)
	default:
		fmt.Println("Unrecognised generator type")
		os.Exit(1)
	}
}

func generateHtmlElements(file io.Writer) {
	output := ElementsJSON{}
	json.Unmarshal(html_defs, &output)
	fmt.Fprint(file, "package scripting\n\n")
	fmt.Fprint(file, "// This file is generated. Do not edit.\n\n")
	fmt.Fprint(file, "var htmlElements = map[string]string {\n")
	for _, element := range output.Elements {
		fmt.Fprintf(file, "\t\"%s\": \"%s\",\n", element.Name, element.Interface)
	}
	fmt.Fprint(file, "}\n")
}
