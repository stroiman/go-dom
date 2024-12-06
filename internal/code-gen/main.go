package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

func getWriter(output string) io.Writer {
	if output == "stdout" {
		return os.Stdout
	}
	file, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating output file")
		os.Exit(1)
	}
	return file
}

func main() {
	debug := flag.Bool("d", false, "Debug")
	outputFile := flag.String("o", "", "Output file to write")
	generatorType := flag.String("g", "", "Generator type")
	flag.Parse()
	if *outputFile == "" || *generatorType == "" {
		fmt.Println("Internal code generator from IDL definitions")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *debug {
		fmt.Println(strings.Join(os.Args, " "))
		fmt.Println("--------")
	}

	file := newBuilder(getWriter(*outputFile))

	switch *generatorType {
	case "html-elements":
		generateHtmlElements(file)
	case "xhr":
		generateXhr(file)
	default:
		fmt.Println("Unrecognised generator type")
		os.Exit(1)
	}
}

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

func generateXhr(b *builder) {
	WriteHeader(b)
}

func WriteHeader(b *builder) {
	b.Printf("package scripting\n\n")
	b.Printf("// This file is generated. Do not edit.\n\n")
}

type builder struct {
	io.Writer
	indentLvl int
}

func newBuilder(w io.Writer) *builder {
	return &builder{w, 0}
}

func (b builder) Printf(format string, args ...interface{}) {
	for i := 0; i < b.indentLvl; i++ {
		fmt.Fprint(b.Writer, "\t")
	}
	fmt.Fprintf(b.Writer, format, args...)
}

func (b *builder) indent() {
	b.indentLvl++
}

func (b *builder) indentF(format string, args ...interface{}) {
	b.indent()
	b.Printf(format, args...)
}

func (b *builder) unIndent() {
	b.indentLvl--
	if b.indentLvl < 0 {
		panic("More unindentation than indentation")
	}
}

func (b *builder) unIndentF(format string, args ...interface{}) {
	b.unIndent()
	b.Printf(format, args...)
}
