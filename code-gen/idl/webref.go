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

func OpenIdlParsed(name string) (fs.File, error) {
	filename := fmt.Sprintf("webref/curated/idlparsed/%s.json", name)
	return WebRef.Open(filename)
}
