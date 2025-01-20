package html_test

import (
	"net/http"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
	domHttp "github.com/stroiman/go-dom/browser/internal/http"
)

func ParseHtmlString(s string) (res dom.Document) {
	html.NewDOMParser().ParseReader(nil, &res, strings.NewReader(s))
	return
}

func NewWindowFromHandler(handler http.Handler) html.Window {
	return html.NewWindow(html.WindowOptions{HttpClient: domHttp.NewHttpClientFromHandler(handler)})
}
