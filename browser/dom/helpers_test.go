package dom_test

import (
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
)

func ParseHtmlString(s string) dom.Document {
	win, err := html.NewWindowReader(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	return win.Document()
}

func CreateDocument() dom.Document {
	return dom.NewDocument(nil)
}
