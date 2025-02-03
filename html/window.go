package html

import (
	"errors"
	"io"
	"net/http"
	netURL "net/url"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
)

type ScriptHost interface {
	NewContext(window Window) ScriptContext
	Close()
}

type ScriptContext interface {
	// Run a script, and convert the result to a Go type. Only use this if you
	// need the return value, otherwise call Run.
	//
	// If the evaluated JS value cannot be converted to a Go value, an error is
	// returned.
	Eval(script string) (any, error)
	// Run a script. This is should be used instead of eval when the return value
	// is not needed, as eval could generate an error if the value cannot be
	// converted to a go-type.
	Run(script string) error
	Close()
}

type Window interface {
	EventTarget
	entity.Entity
	Document() Document
	Close()
	Navigate(string) error // TODO: Remove, perhaps? for testing
	LoadHTML(string) error // TODO: Remove, for testing
	Eval(string) (any, error)
	Run(string) error
	ScriptContext() ScriptContext
	Location() Location
	History() *History
	HTTPClient() http.Client
	ParseFragment(ownerDocument Document, reader io.Reader) (dom.DocumentFragment, error)
	// unexported

	fetchRequest(req *http.Request) error
	resolveHref(string) dom.URL
}

type window struct {
	EventTarget
	document            Document
	history             *History
	scriptEngineFactory ScriptHost
	scriptContext       ScriptContext
	httpClient          http.Client
	baseLocation        string
	domParser           domParser
}

func newWindow(windowOptions ...WindowOption) *window {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	win := &window{
		EventTarget:         NewEventTarget(),
		httpClient:          options.HttpClient,
		baseLocation:        options.BaseLocation,
		scriptEngineFactory: options.ScriptHost,
		history:             new(History),
	}
	if win.baseLocation == "" {
		win.baseLocation = "about:blank"
	}
	win.history.window = win
	win.history.pushLoad(win.baseLocation)
	win.domParser = domParser{}
	win.initScriptEngine()
	win.document = NewHTMLDocument(win)
	dom.SetEventTargetSelf(win)
	return win
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
	w.EventTarget.RemoveAll()
	factory := w.scriptEngineFactory
	engine := w.scriptContext
	if engine != nil {
		engine.Close()
	}
	if factory != nil {
		w.scriptContext = factory.NewContext(w)
	}
}

func (w *window) setBaseLocation(href string) string {
	if href == "" {
		return w.baseLocation
	}
	w.baseLocation = w.resolveHref(href).Href()
	return w.baseLocation
}

func (w *window) History() *History {
	return w.history
}

func (w *window) ParseFragment(
	ownerDocument Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	return w.domParser.ParseFragment(ownerDocument, reader)
}

// NewWindowReader will create a new window and load parse the HTML from the
// reader. If there is an error reading from the stream, or parsing the DOM, an
// error is returned.
//
// If this function returns without an error, the DOM will have been parsed and
// the DOMContentLoaded event has been dispached on the [dom.Document]
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
	ScriptHost
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
	log.Info("Window.navigate:", "href", href)
	w.History().pushLoad(href)
	w.initScriptEngine()
	w.baseLocation = href
	if href == "about:blank" {
		w.document = NewHTMLDocument(w)
		return nil
	} else {
		resp, err := w.httpClient.Get(href)
		if err != nil {
			return err
		}
		return w.handleResponse(resp)
	}
}

// reload is used internally to load a page into the browser, but without
// affecting the history
func (w *window) reload(href string) error {
	log.Debug("Window.reload:", "href", href)
	w.initScriptEngine()
	w.baseLocation = href
	if href == "about:blank" {
		w.document = NewHTMLDocument(w)
		return nil
	} else {
		resp, err := w.httpClient.Get(href)
		if err != nil {
			return err
		}
		return w.handleResponse(resp)
	}
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
	if w.scriptContext != nil {
		return w.scriptContext.Run(script)
	}
	return errors.New("Script engine not initialised")
}

func (w *window) Eval(script string) (any, error) {
	if w.scriptContext != nil {
		return w.scriptContext.Eval(script)
	}
	return nil, errors.New("Script engine not initialised")
}

func (w *window) ScriptContext() ScriptContext { return w.scriptContext }

func (w *window) Location() Location {
	var u *netURL.URL
	if w.baseLocation != "" {
		u, _ = netURL.Parse(w.baseLocation)
	} else {
		u = new(netURL.URL)
	}
	return NewURLFromNetURL(u)
}

func (w *window) Close() {
	if w.scriptContext != nil {
		w.scriptContext.Close()
	}
}

func (w *window) ObjectId() int32 { return -1 }

// resolveHref takes an href from a <a> tag, or action from a <form> tag and
// resolves an absolute URL that must be requested.
func (w *window) resolveHref(href string) dom.URL {
	r, err := dom.NewUrlBase(href, w.Location().Href())
	if err != nil {
		panic(err)
	}
	return r
}
