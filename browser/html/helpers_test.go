package html_test

import (
	"net/http"
	"strings"

	"github.com/gost-dom/browser/browser/dom"
	"github.com/gost-dom/browser/browser/html"
	. "github.com/gost-dom/browser/browser/internal/http"
	domHttp "github.com/gost-dom/browser/browser/internal/http"
)

func ParseHtmlString(s string) (res dom.Document) {
	html.NewDOMParser().ParseReader(nil, &res, strings.NewReader(s))
	return
}

func NewWindowFromHandler(handler http.Handler) html.Window {
	return html.NewWindow(html.WindowOptions{HttpClient: domHttp.NewHttpClientFromHandler(handler)})
}
func windowOptionHandler(h http.Handler) html.WindowOption {
	return html.WindowOptionFunc(func(o *html.WindowOptions) {
		o.HttpClient = NewHttpClientFromHandler(h)
	})
}
