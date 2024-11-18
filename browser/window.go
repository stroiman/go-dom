package browser

import (
	"io"
	"net/http"
	"strings"
)

type Window interface {
	EventTarget
	Document() Document
	// TODO: Remove, for testing
	LoadHTML(string)
	Eval(string) (any, error)
	SetScriptRunner(ScriptEngine)
}

type ScriptEngine interface {
	Run(script string) (any, error)
}

type window struct {
	eventTarget
	document     Document
	scriptEngine ScriptEngine
	httpClient   http.Client
}

func NewWindow() Window {
	return &window{
		eventTarget: newEventTarget(),
		document:    NewDocument(),
	}
}

func newWindow(httpClient http.Client) *window {
	return &window{
		eventTarget: newEventTarget(),
		document:    NewDocument(),
		httpClient:  httpClient,
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
	return nil, nil // or ErrNo
}

func (w *window) SetScriptRunner(r ScriptEngine) {
	w.scriptEngine = r
}
