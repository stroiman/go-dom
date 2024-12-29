package browser

import (
	"errors"
	"log/slog"
	"net/http"

	. "github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/internal/http"
	"github.com/stroiman/go-dom/browser/scripting"
)

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. This used to be the case for _some_ identity providers, but I'm
// not sure if that even work anymore because of browser security.
type Browser struct {
	Client              http.Client
	ScriptEngineFactory ScriptEngineFactory
	windows             []Window
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
	window, err = NewWindowReader(resp.Body, b.createOptions(location))
	b.windows = append(b.windows, window)
	return
}

func NewBrowserFromHandler(handler http.Handler) *Browser {
	host := scripting.NewScriptHost()
	return &Browser{
		ScriptEngineFactory: (*scripting.Wrapper)(host),
		Client:              NewHttpClientFromHandler(handler),
	}
}

func (b *Browser) createOptions(location string) WindowOptions {
	return WindowOptions{
		ScriptEngineFactory: b.ScriptEngineFactory,
		HttpClient:          b.Client,
		BaseLocation:        location,
	}
}

func (b *Browser) Dispose() {
	slog.Debug("Browser: Dispose")
	for _, win := range b.windows {
		win.Dispose()
	}
	if b.ScriptEngineFactory != nil {
		b.ScriptEngineFactory.Dispose()
	}
}
