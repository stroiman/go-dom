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
	LoadHTML(string)
	Eval(string) (any, error)
	SetScriptRunner(ScriptEngine)
	Location() Location
}

type ScriptEngine interface {
	Run(script string) (any, error)
}

type window struct {
	eventTarget
	document     Document
	scriptEngine ScriptEngine
	httpClient   http.Client
	url          *url.URL
}

func NewWindow(url *url.URL) Window {
	return &window{
		eventTarget: newEventTarget(),
		document:    NewDocument(),
		url:         url,
	}
}

func newWindow(httpClient http.Client, url *url.URL) *window {
	return &window{
		eventTarget: newEventTarget(),
		document:    NewDocument(),
		httpClient:  httpClient,
		url:         url,
	}
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) {
	w.loadReader(strings.NewReader(html))
}

func (w *window) loadReader(r io.Reader) error {
	parseStream(w, r)
	w.Document().DispatchEvent(NewCustomEvent(DocumentEventDOMContentLoaded))
	// 'load' is emitted when css and images are loaded, not relevant yet, so
	// just emit it right await
	w.Document().DispatchEvent(NewCustomEvent(DocumentEventLoad))
	return nil
}

func (w *window) Eval(script string) (any, error) {
	if w.scriptEngine != nil {
		return w.scriptEngine.Run(script)
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
