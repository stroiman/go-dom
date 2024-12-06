package browser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

type XHREvent = string

const (
	XHREventLoad      XHREvent = "load"
	XHREventLoadstart XHREvent = "loadstart"
	XHREventLoadend   XHREvent = "loadend"
)

// TODO: Type URL (or is it in v8 already?)

type XmlHttpRequest struct {
	eventTarget
	client   http.Client
	async    bool
	status   int
	method   string
	url      string
	response []byte
	headers  map[string][]string
}

func NewXmlHttpRequest(client http.Client) *XmlHttpRequest {
	return &XmlHttpRequest{
		eventTarget: newEventTarget(),
		client:      client,
		headers:     make(map[string][]string),
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

	req.method = method
	req.url = url
	for _, o := range options {
		o(req)
	}
	// TODO: if (req.open) { req.Abort() }
}

func (req *XmlHttpRequest) send(body io.Reader) error {
	fmt.Println("Create request", req.method, req.url)
	httpRequest, err := http.NewRequest(req.method, req.url, body)
	if err != nil {
		return err
	}
	httpRequest.Header = req.headers
	res, err := req.client.Do(httpRequest)
	if err != nil {
		return err
	}
	req.status = res.StatusCode
	b := new(bytes.Buffer) // TODO, branch out depending on content-type
	_, err = b.ReadFrom(res.Body)
	req.response = b.Bytes()
	req.DispatchEvent(NewCustomEvent(XHREventLoad))
	return err
}

func (req *XmlHttpRequest) Send() error {
	return req.SendBody(nil)
}

func (req *XmlHttpRequest) SendBody(body *XHRRequestBody) error {
	var reader io.Reader
	if body != nil {
		// TODO: Set content type or not?
		req.headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		reader = body.getReader()
	}
	if req.async {
		req.DispatchEvent(NewCustomEvent((XHREventLoadstart)))
		go req.send(reader)
		return nil
	}
	return req.send(reader)
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

/* -------- XHRRequestBody -------- */

type XHRRequestBody struct {
	data []byte // Temporary solution, should probably be an io.Reader
}

func NewXHRRequestBodyOfFormData(data *FormData) *XHRRequestBody {
	sb := strings.Builder{}
	for i, e := range data.Entries {
		if i != 0 {
			sb.WriteString("&")
		}

		sb.WriteString(url.QueryEscape(e.Name))
		sb.WriteString("=")
		sb.WriteString(url.QueryEscape(string(e.Value)))
	}
	sb.WriteString("foo")

	return &XHRRequestBody{
		data: []byte(sb.String()),
	}
}

// Bytes retrieves the RAW bytes.
//
// Deprecated: This is added for testing purposes only, will probably be
// removed. File an issue if you believe there's a valid case for this.
func (b XHRRequestBody) Bytes() []byte {
	return b.data
}

func (b XHRRequestBody) getReader() io.Reader {
	return bytes.NewReader(b.data)
}
