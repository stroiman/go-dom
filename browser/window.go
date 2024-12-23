package browser

import (
	"errors"
	"io"
	"net/http"
	netURL "net/url"
	"strings"
)

type Window interface {
	EventTarget
	Document() Document
	// TODO: Remove, for testing
	LoadHTML(string) error
	Eval(string) (any, error)
	Run(string) error
	SetScriptRunner(ScriptEngine)
	Location() Location
	NewXmlHttpRequest() XmlHttpRequest
}

type ScriptEngine interface {
	// Run a script, and convert the result to a Go type. This will result in an
	// error if the returned value cannot be represented as a Go type.
	Eval(script string) (any, error)
	// Run a script, ignoring any returned value
	Run(script string) error
	Dispose()
}

type window struct {
	eventTarget
	document     Document
	scriptEngine ScriptEngine
	httpClient   http.Client
	url          *netURL.URL
}

func NewWindow(url *netURL.URL) Window {
	result := &window{
		eventTarget: newEventTarget(),
		url:         url,
	}
	result.document = NewDocument(result)
	return result
}

func newWindow(httpClient http.Client, url *netURL.URL) *window {
	result := &window{
		eventTarget: newEventTarget(),
		httpClient:  httpClient,
		url:         url,
	}
	result.document = NewDocument(result)
	return result
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) error {
	return w.loadReader(strings.NewReader(html))
}

func (w *window) loadReader(r io.Reader) (err error) {
	w.document, err = parseStream(w, r)
	if err == nil {
		w.Document().DispatchEvent(NewCustomEvent(DocumentEventDOMContentLoaded))
		// 'load' is emitted when css and images are loaded, not relevant yet, so
		// just emit it right await
		w.Document().DispatchEvent(NewCustomEvent(DocumentEventLoad))
	}
	return
}

func (w *window) Run(script string) error {
	if w.scriptEngine != nil {
		return w.scriptEngine.Run(script)
	}
	return errors.New("Script engine not initialised")
}

func (w *window) Eval(script string) (any, error) {
	if w.scriptEngine != nil {
		return w.scriptEngine.Eval(script)
	}
	return nil, errors.New("Script engine not initialised")
}

func (w *window) SetScriptRunner(r ScriptEngine) {
	w.scriptEngine = r
}

func (w *window) Location() Location {
	u := w.url
	if u == nil {
		u = new(netURL.URL)
	}
	return NewLocationFromNetURL(u)
}

func (w *window) Dispose() {
	if w.scriptEngine != nil {
		w.scriptEngine.Dispose()
	}
}

func (w *window) NewXmlHttpRequest() XmlHttpRequest {
	return *NewXmlHttpRequest(w.httpClient)
}
