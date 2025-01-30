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
	// if req.Host != "" {
	// 	// TODO: Not tested, nowhere near the case where we can test this yet, but
	// 	// the idea is if we are serving content directly from an http.Handler, any
	// 	// external script or CSS reference (when implemented) should be downloaded,
	// 	// so the test still works if you reference content from CDN.
	// 	return http.DefaultTransport.RoundTrip(req)
	// }
	rec := httptest.NewRecorder()
	serverReq := new(http.Request)
	*serverReq = *req
	if serverReq.Body == nil {
		serverReq.Body = nullReader{}
	}
	h.ServeHTTP(rec, serverReq)
	return rec.Result(), nil
}

type nullReader struct{}

func (_ nullReader) Read(b []byte) (int, error) {
	return 0, io.EOF
}
func (_ nullReader) Close() error { return nil }
