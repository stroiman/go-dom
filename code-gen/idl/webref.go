package idl

import (
	"embed"
)

//go:embed webref/ed/elements/html.json
var Html_defs []byte

//go:embed webref/*/idlparsed/*.json
var WebRef embed.FS
