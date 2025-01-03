package html_test

import (
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
)

func ParseHtmlString(s string) (res dom.Document) {
	html.NewDOMParser().ParseReader(nil, &res, strings.NewReader(s))
	return
}
