package html

import (
	"errors"
	"io"
	"net/http"
	netURL "net/url"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/dom"
)

type ScriptEngineFactory interface {
	NewScriptEngine(window Window) ScriptEngine
	Dispose()
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
	Entity
	Document() Document
	Dispose()
	Navigate(string) error // TODO: Remove, perhaps? for testing
	LoadHTML(string) error // TODO: Remove, for testing
	Eval(string) (any, error)
	Run(string) error
	SetScriptRunner(ScriptEngine)
	GetScriptEngine() ScriptEngine
	Location() Location
	HTTPClient() http.Client
	ParseFragment(ownerDocument Document, reader io.Reader) (dom.DocumentFragment, error)
	// unexported
	fetchRequest(req *http.Request) error
}

type window struct {
	EventTarget
	document            Document
	scriptEngineFactory ScriptEngineFactory
	scriptEngine        ScriptEngine
	httpClient          http.Client
	baseLocation        string
	domParser           domParser
}

func newWindow(windowOptions ...WindowOption) *window {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	result := &window{
		EventTarget:         NewEventTarget(),
		httpClient:          options.HttpClient,
		baseLocation:        options.BaseLocation,
		scriptEngineFactory: options.ScriptEngineFactory,
	}
	result.domParser = domParser{}
	result.initScriptEngine()
	result.document = NewHTMLDocument(result)
	return result
}

func NewWindow(windowOptions ...WindowOption) Window {
	return newWindow(windowOptions...)
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
	result := newWindow(options)
	resp, err := result.httpClient.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non-ok Response")
	}
	err = result.parseReader(resp.Body)
	return result, err
}

func (w *window) initScriptEngine() {
	factory := w.scriptEngineFactory
	engine := w.scriptEngine
	if engine != nil {
		engine.Dispose()
	}
	if factory != nil {
		w.scriptEngine = factory.NewScriptEngine(w)
	}
}

func (w *window) ParseFragment(
	ownerDocument Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	return w.domParser.ParseFragment(ownerDocument, reader)
}
func NewWindowReader(reader io.Reader, windowOptions ...WindowOption) (Window, error) {
	window := newWindow(windowOptions...)
	err := window.parseReader(reader)
	return window, err
}

func (w *window) parseReader(reader io.Reader) error {
	err := w.domParser.ParseReader(w, &w.document, reader)
	if err == nil {
		w.document.DispatchEvent(NewCustomEvent(DocumentEventDOMContentLoaded))
		// 'load' is emitted when css and images are loaded, not relevant yet, so
		// just emit it right await
		w.document.DispatchEvent(NewCustomEvent(DocumentEventLoad))
	}
	return err
}

func (w *window) HTTPClient() http.Client { return w.httpClient }

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

func (w *window) handleResponse(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return errors.New("Non-ok Response")
	}
	return w.parseReader(resp.Body)

}

func (w *window) Navigate(href string) error {
	w.initScriptEngine()
	resp, err := w.httpClient.Get(href)
	if err != nil {
		return err
	}
	return w.handleResponse(resp)
}

func (w *window) fetchRequest(req *http.Request) error {
	resp, err := w.httpClient.Do(req)
	if err == nil {
		err = w.handleResponse(resp)
	}
	return err
}

func (w *window) LoadHTML(html string) error {
	return w.parseReader(strings.NewReader(html))
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

func (w *window) GetScriptEngine() ScriptEngine { return w.scriptEngine }

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

func (w *window) ObjectId() int32 { return -1 }
