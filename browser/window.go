package browser

import "strings"

type WindowEvent = string

const (
	WindowEventDOMContentLoaded WindowEvent = "DOMContentLoaded"
	WindowEventLoad             WindowEvent = "load"
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
}

func NewWindow() Window {
	return &window{
		eventTarget: newEventTarget(),
		document:    NewDocument(),
	}
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) {
	parseStream(w, strings.NewReader(html))
	w.DispatchEvent(NewCustomEvent(WindowEventDOMContentLoaded))
	// 'load' is emitted when css and images are loaded, not relevant yet, so
	// just emit it right await
	w.DispatchEvent(NewCustomEvent(WindowEventLoad))
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
