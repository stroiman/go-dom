package idl

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed webref/ed/elements/html.json
var Html_defs []byte

//go:embed webref/*/idlparsed/*.json
var WebRef embed.FS

func openIdlParsed(name string) (fs.File, error) {
	filename := fmt.Sprintf("webref/curated/idlparsed/%s.json", name)
	return WebRef.Open(filename)
}

// LoadIdlParsed loads a files from the /curated/idlpased directory containing
// specifications of the interfaces.
func LoadIdlParsed(name string) (IdlSpec, error) {
	file, err := openIdlParsed(name)
	if err != nil {
		return IdlSpec{}, err
	}
	defer file.Close()
	return ParseIdlJsonReader(file)
}

// Sanitize converts property names into valid Go identifiers, e.g., adding an
// underscore to reserved words like go and type.
func SanitizeName(name string) string {
	switch name {
	case "go":
		return "go_"
	case "type":
		return "type_"
	default:
		return name
	}
}
