package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	goTypes "github.com/stroiman/go-dom/code-gen/go-type-generators"
	wrappers "github.com/stroiman/go-dom/code-gen/script-wrappers"
)

//go:embed webref/ed/elements/html.json
var html_defs []byte

//go:embed webref/*/idlparsed/*.json
var idlParsedFS embed.FS

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

var generators = map[string]func(io.Writer) error{
	"html-elements": generateHtmlElements,
}

func main() {
	debug := flag.Bool("d", false, "Debug")
	outputFile := flag.String("o", "", "Output file to write")
	generatorType := flag.String("g", "", "Generator type")
	flag.Parse()
	fmt.Println("Generateor?", *generatorType, *debug)
	switch *generatorType {
	case "scripting":
		gen := wrappers.NewScriptWrapperModulesGenerator(idlParsedFS)
		err := gen.GenerateScriptWrappers()
		exit(err)
	case "browser-types":
		gen := goTypes.NewGoTypeGenerators(idlParsedFS)
		err := gen.GenerateType("html", "HTMLInputElement", os.Stdout)
		exit(err)
		return
	}
	if *outputFile == "" || *generatorType == "" {
		fmt.Println("Internal code generator from IDL definitions")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *debug {
		fmt.Println(strings.Join(os.Args, " "))
		fmt.Println("--------")
	}

	file := getWriter(*outputFile)

	generator, ok := generators[*generatorType]
	if !ok {
		os.Exit(1)
	}
	err := generator(file)
	exit(err)
}

func exit(err error) {
	if err != nil {
		fmt.Println("Error running generator")
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
