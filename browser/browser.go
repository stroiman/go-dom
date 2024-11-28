package browser

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
)

type ScriptEngineFactory interface {
	NewScriptEngine(window Window) ScriptEngine
}

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. Particularly if testing login flow; where the login
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
	u, err := url.Parse(location)
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
		panic("err")
	}
	return ParseHtmlStream(resp.Body)
}

type HandlerRoundTripper struct{ http.Handler }

func (h HandlerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// if req.Host != "" {
	// 	// TODO: Not tested, nowhere near the case where we can test this yet, but
	// 	// the idea is if we are serving content directly from an http.Handler, any
	// 	// external script or CSS reference (when implemented) should be downloaded,
	// 	// so the test still works if you reference content from CDN.
	// 	return http.DefaultTransport.RoundTrip(req)
	// }
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Result(), nil
}

func NewBrowserFromHandler(handler http.Handler) *Browser {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Transport: HandlerRoundTripper{handler},
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
