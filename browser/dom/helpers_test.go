package dom_test

import (
	"strings"

	"github.com/gost-dom/browser/browser/dom"
	"github.com/gost-dom/browser/browser/html"
)

func ParseHtmlString(s string) dom.Document {
	win, err := html.NewWindowReader(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	return win.Document()
}

func CreateHTMLDocument() dom.Document {
	return html.NewHTMLDocument(nil)
}
