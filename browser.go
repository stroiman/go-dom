package go_dom

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"

	dom "github.com/stroiman/go-dom/dom-types"
)

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. Particularly if testing login flow; where the login
type Browser struct {
	Client http.Client
}

func (b Browser) Open(url string) dom.Document {
	resp, err := b.Client.Get(url)
	if err != nil {
		panic("err")
	}
	return Parse(resp.Body)
}

type HandlerRoundTripper struct{ http.Handler }

func (h HandlerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Host != "" {
		// TODO: Not tested, nowhere near the case where we can test this yet, but
		// the idea is if we are serving content directly from an http.Handler, any
		// external script or CSS reference (when implemented) should be downloaded,
		// so the test still works if you reference content from CDN.
		return http.DefaultTransport.RoundTrip(req)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Result(), nil
}

func NewBrowserFromHandler(handler http.Handler) Browser {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Transport: HandlerRoundTripper{handler},
		Jar:       cookiejar,
	}
	return Browser{client}
}
