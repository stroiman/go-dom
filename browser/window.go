package browser

import (
	"errors"
	"io"
	"net/http"
	"net/url"
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
}

type ScriptEngine interface {
	// Run a script, and convert the result to a Go type. This will result in an
	// error if the returned value cannot be represented as a Go type.
	Eval(script string) (any, error)
	// Run a script, ignoring any returned value
	Run(script string) error
}

type window struct {
	eventTarget
	document     Document
	scriptEngine ScriptEngine
	httpClient   http.Client
	url          *url.URL
}

func NewWindow(url *url.URL) Window {
	result := &window{
		eventTarget: newEventTarget(),
		url:         url,
	}
	result.document = NewDocument(result)
	return result
}

func newWindow(httpClient http.Client, url *url.URL) *window {
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

func (w *window) loadReader(r io.Reader) error {
	w.document = parseStream(w, r)
	err := w.Document().DispatchEvent(NewCustomEvent(DocumentEventDOMContentLoaded))
	// 'load' is emitted when css and images are loaded, not relevant yet, so
	// just emit it right await
	if err == nil {
		w.Document().DispatchEvent(NewCustomEvent(DocumentEventLoad))
	}
	return err
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
		u = new(url.URL)
	}
	return NewLocation(u)
}
