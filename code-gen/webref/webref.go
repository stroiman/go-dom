package webref

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed definitions/ed/elements/html.json
var Html_defs []byte

//go:embed definitions/curated/idlparsed/*.json
var WebRef embed.FS

func OpenIdlParsed(name string) (fs.File, error) {
	filename := fmt.Sprintf("definitions/curated/idlparsed/%s.json", name)
	return WebRef.Open(filename)
}
