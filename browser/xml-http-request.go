package browser

import (
	"bytes"
	"net/http"
)

// TODO: Events for async
// abort
// error
// load
// loadend
// loadstart
// progress
// readystatechange
// timeout

// TODO: Type URL (or is it in v8 already?)

type XmlHttpRequest struct {
	client   http.Client
	async    bool
	status   int
	url      string
	response []byte
}

func NewXmlHttpRequest(client http.Client) *XmlHttpRequest {
	return &XmlHttpRequest{
		client: client,
	}
}

type RequestOption = func(req *XmlHttpRequest)

func (req *XmlHttpRequest) Open(
	method string,
	// TODO: Should this be a `string` or a stringer? The JS object should accept
	// stringable objects, e.g., a URL, but should we convert here; or on the JS
	// binding layer? Or different methods?
	url string,
	options ...RequestOption) {
	// TODO: if (req.open) { req.Abort() }
}

func (req *XmlHttpRequest) Send() error {
	res, err := req.client.Get(req.url)
	if err != nil {
		return err
	}
	req.status = res.StatusCode
	b := new(bytes.Buffer) // TODO, branch out depending on content-type
	_, err = b.ReadFrom(res.Body)
	req.response = b.Bytes()
	return err
}

func (req *XmlHttpRequest) Status() int { return req.status }

// StatusText implements the [statusText] property
// [statusText]: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/statusText
// TODO: Should this exist? It's just a wrapper around [http.StatusText], could
// be in JS wrapper layer
func (req *XmlHttpRequest) StatusText() string { return http.StatusText(req.status) }

func (req *XmlHttpRequest) ResponseText() string { return string(req.response) }

/* -------- Options -------- */

func RequestOptionAsync(
	val bool,
) RequestOption {
	return func(req *XmlHttpRequest) { req.async = val }
}

func RequestOptionUserName(_ string) RequestOption {
	return func(req *XmlHttpRequest) {
		// TODO
		panic("Not implemented")
	}
}

func RequestOptionPassword(_ string) RequestOption {
	return func(req *XmlHttpRequest) {
		// TODO
		panic("Not implemented")
	}
}
