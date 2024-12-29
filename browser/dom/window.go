package dom

import (
	"errors"
	"io"
	"net/http"
	netURL "net/url"
	"strings"
)

type ScriptEngineFactory interface {
	NewScriptEngine(window Window) ScriptEngine
}

type ScriptEngine interface {
	// Run a script, and convert the result to a Go type. This will result in an
	// error if the returned value cannot be represented as a Go type.
	Eval(script string) (any, error)
	// Run a script, ignoring any returned value
	Run(script string) error
	Dispose()
}

type Window interface {
	EventTarget
	Document() Document
	Dispose()
	// TODO: Remove, for testing
	LoadHTML(string) error
	Eval(string) (any, error)
	Run(string) error
	SetScriptRunner(ScriptEngine)
	Location() Location
	NewXmlHttpRequest() XmlHttpRequest
}

type window struct {
	eventTarget
	document     Document
	scriptEngine ScriptEngine
	httpClient   http.Client
	baseLocation string
}

func NewWindow(windowOptions ...WindowOption) Window {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	result := &window{
		eventTarget:  newEventTarget(),
		httpClient:   options.HttpClient,
		baseLocation: options.BaseLocation,
	}
	if options.ScriptEngineFactory != nil {
		result.scriptEngine = options.ScriptEngineFactory.NewScriptEngine(result)
	}
	result.document = NewDocument(result)
	return result
}

func OpenWindowFromLocation(location string, windowOptions ...WindowOption) (Window, error) {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	if options.BaseLocation != "" {
		u, err := netURL.Parse(options.BaseLocation)
		if err == nil {
			location = u.JoinPath(location).String()
		}
	} else {
		options.BaseLocation = location
	}
	result := &window{
		eventTarget:  newEventTarget(),
		httpClient:   options.HttpClient,
		baseLocation: options.BaseLocation,
	}
	if options.ScriptEngineFactory != nil {
		result.scriptEngine = options.ScriptEngineFactory.NewScriptEngine(result)
	}
	result.document = NewDocument(result)
	resp, err := result.httpClient.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non-ok Response")
	}
	err = result.loadReader(resp.Body)
	return result, err
}

func NewWindowReader(reader io.Reader, windowOptions ...WindowOption) (window Window, err error) {
	window = NewWindow(windowOptions...)
	document := window.Document()
	err = parseIntoDocument(window, document, reader)
	if err == nil {
		document.DispatchEvent(NewCustomEvent(DocumentEventDOMContentLoaded))
		// 'load' is emitted when css and images are loaded, not relevant yet, so
		// just emit it right await
		document.DispatchEvent(NewCustomEvent(DocumentEventLoad))
	}
	return
}

type WindowOptions struct {
	ScriptEngineFactory
	HttpClient   http.Client
	BaseLocation string
}

type WindowOption interface {
	Apply(options *WindowOptions)
}

type WindowOptionFunc func(*WindowOptions)

func (f WindowOptionFunc) Apply(options *WindowOptions) { f(options) }

func WindowOptionLocation(location string) WindowOptionFunc {
	return func(options *WindowOptions) {
		options.BaseLocation = location
	}
}

func (o WindowOptions) Apply(options *WindowOptions) {
	*options = o
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) error {
	return w.loadReader(strings.NewReader(html))
}

func (w *window) loadReader(r io.Reader) (err error) {
	w.document = NewDocument(w)
	err = parseIntoDocument(w, w.document, r)
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
	var u *netURL.URL
	if w.baseLocation != "" {
		u, _ = netURL.Parse(w.baseLocation)
	} else {
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
	return NewXmlHttpRequest(w.httpClient)
}
