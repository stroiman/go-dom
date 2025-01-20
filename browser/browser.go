package browser

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/internal/http"
	"github.com/stroiman/go-dom/browser/scripting/v8host"
)

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. This used to be the case for _some_ identity providers, but I'm
// not sure if that even work anymore because of browser security.
type Browser struct {
	Client     http.Client
	ScriptHost ScriptHost
	windows    []Window
}

// Open will open a new [http.Window], loading the specified location. If the
// server does not respons with a 200 status code, an error is returned.
//
// See [html.NewWindowReader] about the return value, and when the window
// returns.
func (b *Browser) Open(location string) (window Window, err error) {
	// slog.Debug("Browser: OpenWindow", "URL", location)
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

// NewFromHandler initialises a new [Browser] with the default script engine and
// sets up the internal [http.Client] used with an [http.Roundtripper] that
// bypasses the TCP stack and calls directly into the
//
// Note: There is a current limitation that NO requests from the browser will be
// sent when using this. So sites will not work if they
//   - Depend on content from CDN
//   - Depend on an external service, e.g., an identity provider.
//
// That is a limitation that was the result of prioritising more important, and
// higher risk features.
func NewFromHandler(handler http.Handler) *Browser {
	return &Browser{
		ScriptHost: v8host.New(),
		Client:     NewHttpClientFromHandler(handler),
	}
}

// New initialises a new [Browser] with the default script engine.
func New() *Browser {
	return &Browser{
		ScriptHost: v8host.New(),
		Client:     NewHttpClient(),
	}
}

// NewBrowser should not be called. Call New instead.
//
// This method will selfdestruct in 10 commits
func NewBrowser() *Browser {
	return New()
}

// NewBrowserFromHandler should not be called, call, NewFromHandler instead.
//
// This method will selfdestruct in 10 commits
func NewBrowserFromHandler(handler http.Handler) *Browser {
	return NewFromHandler(handler)
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
