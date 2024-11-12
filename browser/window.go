package browser

import "strings"

type Window interface {
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
	document     Document
	scriptEngine ScriptEngine
}

func NewWindow() Window {
	return &window{
		NewDocument(),
		nil,
	}
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) {
	parseStream(w, nil, strings.NewReader(html))
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
