package dom

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	netURL "net/url"
)

type ScriptEngineFactory interface {
	NewScriptEngine(window Window) ScriptEngine
}

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. This used to be the case for _some_ identity providers, but I'm
// not sure if that even work anymore because of browser sercurity.
type Browser struct {
	Client              http.Client
	ScriptEngineFactory ScriptEngineFactory
	windows             []*window
}

// TODO: Rename to Open
func (b *Browser) OpenWindow(location string) (Window, error) {
	slog.Debug("Browser: OpenWindow", "URL", location)
	resp, err := b.Client.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non-ok Response")
	}
	u, err := netURL.Parse(location)
	if err != nil {
		return nil, err
	}
	window := newWindow(b.Client, u)
	var scriptEngine ScriptEngine
	b.windows = append(b.windows, window)
	if b.ScriptEngineFactory != nil {
		scriptEngine = b.ScriptEngineFactory.NewScriptEngine(window)
	}
	window.SetScriptRunner(scriptEngine)
	err = window.loadReader(resp.Body)
	return window, err
}

// TODO: Delete
func (b *Browser) Open(url string) Document {
	resp, err := b.Client.Get(url)
	if err != nil {
		panic(err)
	}
	doc, err := ParseHtmlStream(resp.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

// Creates a new window containing an empty document
func (b *Browser) NewWindow(baseUrl string) (Window, error) {
	u, err := netURL.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	window := newWindow(b.Client, u)
	var scriptEngine ScriptEngine
	if b.ScriptEngineFactory != nil {
		scriptEngine = b.ScriptEngineFactory.NewScriptEngine(window)
	}
	window.SetScriptRunner(scriptEngine)
	return window, nil
}

func NewBrowserFromHandler(handler http.Handler) *Browser {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Transport: TestRoundTripper{handler},
		Jar:       cookiejar,
	}
	return &Browser{
		Client: client,
	}
}

func (b *Browser) Dispose() {
	slog.Debug("Browser: Dispose")
	for _, win := range b.windows {
		slog.Debug("Browser: Dispose window")
		win.Dispose()
	}
}
