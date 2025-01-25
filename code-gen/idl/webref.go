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
