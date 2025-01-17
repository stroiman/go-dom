package browser

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/internal/http"
	"github.com/stroiman/go-dom/browser/v8host"
)

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. This used to be the case for _some_ identity providers, but I'm
// not sure if that even work anymore because of browser security.
type Browser struct {
	Client     http.Client
	ScriptHost ScriptHost
	windows    []Window
}

func (b *Browser) Open(location string) (window Window, err error) {
	// slog.Debug("Browser: OpenWindow", "URL", location)
	// return OpenWindowFromLocation(location, b.createOptions(location))
	resp, err := b.Client.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non-ok Response")
	}
	window, err = html.NewWindowReader(resp.Body, b.createOptions(location))
	b.windows = append(b.windows, window)
	return
}

func NewBrowser() *Browser {
	return &Browser{
		ScriptHost: v8host.NewScriptHost(),
		Client:     NewHttpClient(),
	}
}

func NewBrowserFromHandler(handler http.Handler) *Browser {
	return &Browser{
		ScriptHost: v8host.NewScriptHost(),
		Client:     NewHttpClientFromHandler(handler),
	}
}

func (b *Browser) createOptions(location string) WindowOptions {
	return WindowOptions{
		ScriptHost:   b.ScriptHost,
		HttpClient:   b.Client,
		BaseLocation: location,
	}
}

func (b *Browser) Close() {
	slog.Debug("Browser: Close()")
	for _, win := range b.windows {
		win.Close()
	}
	if b.ScriptHost != nil {
		b.ScriptHost.Close()
	}
}
