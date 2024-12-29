package dom

import (
	"errors"
	"log/slog"
	"net/http"
	netURL "net/url"

	. "github.com/stroiman/go-dom/browser/internal/http"
)

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. This used to be the case for _some_ identity providers, but I'm
// not sure if that even work anymore because of browser sercurity.
type Browser struct {
	Client              http.Client
	ScriptEngineFactory ScriptEngineFactory
	windows             []*window
}

func (b *Browser) Open(location string) (window Window, err error) {
	slog.Debug("Browser: OpenWindow", "URL", location)
	resp, err := b.Client.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non-ok Response")
	}
	var url *netURL.URL
	url, err = netURL.Parse(location)
	if err == nil {
		window, err = NewWindowReader(resp.Body, b.createOptions(url))
	}
	return
}

// Creates a new window containing an empty document
func (b *Browser) NewWindow(baseUrl string) (window Window, err error) {
	var url *netURL.URL
	url, err = netURL.Parse(baseUrl)
	if err == nil {
		window = NewWindow(b.createOptions(url))
	}
	return
}

func NewBrowserFromHandler(handler http.Handler) *Browser {
	return &Browser{
		Client: NewHttpClientFromHandler(handler),
	}
}

func (b *Browser) createOptions(url *netURL.URL) WindowOptions {
	return WindowOptions{
		ScriptEngineFactory: b.ScriptEngineFactory,
		HttpClient:          b.Client,
		URL:                 url,
	}
}

func (b *Browser) Dispose() {
	slog.Debug("Browser: Dispose")
	for _, win := range b.windows {
		slog.Debug("Browser: Dispose window")
		win.Dispose()
	}
}
