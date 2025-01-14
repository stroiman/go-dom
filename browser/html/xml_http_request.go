package html

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	netURL "net/url"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
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

type XmlHttpRequest interface {
	dom.EventTarget
	Abort() error
	Open(string, string, ...RequestOption)
	Send() error
	SendBody(body *XHRRequestBody) error
	Status() int
	StatusText() string
	ResponseText() string
	SetRequestHeader(name string, value string)
	GetAllResponseHeaders() (res string, err error)
	OverrideMimeType(mimeType string) error
	GetResponseHeader(headerName string) *string
	SetWithCredentials(val bool) error
	WithCredentials() bool
	ResponseURL() string
	Response() string
	SetTimeout(int) error
	Timeout() int
}

type xmlHttpRequest struct {
	dom.EventTarget
	client   http.Client
	async    bool
	status   int
	method   string
	url      string
	response []byte
	res      *http.Response
	headers  http.Header
}

func NewXmlHttpRequest(client http.Client) XmlHttpRequest {
	return &xmlHttpRequest{
		EventTarget: dom.NewEventTarget(),
		client:      client,
		headers:     make(map[string][]string),
	}
}

type RequestOption = func(req *xmlHttpRequest)

func (req *xmlHttpRequest) Open(
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

func (req *xmlHttpRequest) send(body io.Reader) error {
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
	req.res = res
	b := new(bytes.Buffer) // TODO, branch out depending on content-type
	_, err = b.ReadFrom(res.Body)
	req.response = b.Bytes()
	req.DispatchEvent(dom.NewCustomEvent(XHREventLoad))
	return err
}

func (req *xmlHttpRequest) Send() error {
	return req.SendBody(nil)
}

func (req *xmlHttpRequest) SendBody(body *XHRRequestBody) error {
	var reader io.Reader
	if body != nil {
		// TODO: Set content type or not?
		req.headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		reader = body.getReader()
	}
	if req.async {
		req.DispatchEvent(dom.NewCustomEvent((XHREventLoadstart)))
		go req.send(reader)
		return nil
	}
	return req.send(reader)
}

func (req *xmlHttpRequest) Status() int { return req.status }

// GetStatusText implements the [statusText] property
// [statusText]: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/statusText
// TODO: Should this exist? It's just a wrapper around [http.GetStatusText], could
// be in JS wrapper layer
func (req *xmlHttpRequest) StatusText() string { return http.StatusText(req.status) }

func (req *xmlHttpRequest) ResponseURL() string { return req.url }

func (req *xmlHttpRequest) Response() string { return req.ResponseText() }

func (req *xmlHttpRequest) ResponseText() string { return string(req.response) }

func (req *xmlHttpRequest) SetRequestHeader(name string, value string) {
	req.headers.Add(name, value)
}

func (req *xmlHttpRequest) Abort() error {
	return errors.New("XmlHttpRequest.Abort called - not implemented - ignoring call")
}

func (req *xmlHttpRequest) GetAllResponseHeaders() (res string, err error) {
	if req.res == nil {
		return
	}
	builder := strings.Builder{}
	for k, vs := range req.res.Header {
		key := strings.ToLower(k)
		if key != "set-cookie" {
			for _, v := range vs {
				_, err = builder.WriteString(fmt.Sprintf("%s: %s\r\n", key, v))
				if err != nil {
					return
				}
			}
		}
	}
	return builder.String(), nil
}

func (req *xmlHttpRequest) OverrideMimeType(mimeType string) error {
	// This has no effect at the moment, but has an empty implementation to be
	// compatible with HTMX.
	return nil
}

func (req *xmlHttpRequest) GetResponseHeader(headerName string) *string {
	if req.res == nil {
		return nil
	}
	key := http.CanonicalHeaderKey(headerName)
	if val, ok := req.res.Header[key]; ok && len(val) > 0 {
		res := new(string)
		*res = strings.Join(val, ", ")
		return res

	}
	return nil
}

func (req *xmlHttpRequest) SetWithCredentials(val bool) error {
	return nil
}

func (req *xmlHttpRequest) WithCredentials() bool {
	return false
}

func (req *xmlHttpRequest) SetTimeout(val int) error {
	return nil
}

func (req *xmlHttpRequest) Timeout() int {
	return 0
}

/* -------- Options -------- */

func RequestOptionAsync(
	val bool,
) RequestOption {
	return func(req *xmlHttpRequest) { req.async = val }
}

func RequestOptionUserName(_ string) RequestOption {
	return func(req *xmlHttpRequest) {
		// TODO
		panic("Not implemented")
	}
}

func RequestOptionPassword(_ string) RequestOption {
	return func(req *xmlHttpRequest) {
		// TODO
		panic("Not implemented")
	}
}

/* -------- XHRRequestBody -------- */

type XHRRequestBody struct {
	data []byte // Temporary solution, should probably be an io.Reader
}

func NewXHRRequestBodyOfString(data string) *XHRRequestBody {
	return &XHRRequestBody{[]byte(data)}
}

func NewXHRRequestBodyOfFormData(data *FormData) *XHRRequestBody {
	sb := strings.Builder{}
	for i, e := range data.Entries {
		if i != 0 {
			sb.WriteString("&")
		}

		sb.WriteString(netURL.QueryEscape(e.Name))
		sb.WriteString("=")
		sb.WriteString(netURL.QueryEscape(string(e.Value)))
	}

	return &XHRRequestBody{
		data: []byte(sb.String()),
	}
}

func (b XHRRequestBody) getReader() io.Reader {
	return bytes.NewReader(b.data)
}

func (b XHRRequestBody) string() string {
	return string(b.data)
}
