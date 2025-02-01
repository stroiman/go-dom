package http

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// A TestRoundTripper is an implementation of the [http.RoundTripper] interface
// that communicates directly with an [http.Handler] instance.
type TestRoundTripper struct{ http.Handler }

func (h TestRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// You could possibly test on req.Host and apply different behaviour, e.g.
	// forwarding to external site, or have mocked external sites, such as IDPs
	rec := httptest.NewRecorder()
	serverReq, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		return nil, err
	}
	serverReq.Header = req.Header
	serverReq.Trailer = req.Trailer
	if serverReq.Body == nil {
		serverReq.Body = nullReader{}
	}
	h.ServeHTTP(rec, serverReq)
	return rec.Result(), nil
}

// nullReader is just a reader with no content. When _sending_ an HTTP request,
// a _nil_ body is allowed, but when receiving; there _is_ a body. This fixes
// the request so the valid output request body is also a valid incoming request
// body.
type nullReader struct{}

func (_ nullReader) Read(b []byte) (int, error) { return 0, io.EOF }
func (_ nullReader) Close() error               { return nil }
