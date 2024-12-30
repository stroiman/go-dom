package dom_test

import (
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
)

func ParseHtmlString(s string) dom.Document {
	parser := dom.NewDOMParser()
	var doc dom.Document
	err := parser.ParseReader(nil, &doc, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	return doc
}
