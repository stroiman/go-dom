package v8host

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	. "github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting"

	"github.com/tommie/v8go"
	v8 "github.com/tommie/v8go"
)

type globalInstall struct {
	name        string
	constructor *v8.FunctionTemplate
}

type globals struct {
	namedGlobals map[string]*v8.FunctionTemplate
}

type V8ScriptHost struct {
	mu              *sync.Mutex
	iso             *v8.Isolate
	inspector       *v8.Inspector
	inspectorClient *v8.InspectorClient
	windowTemplate  *v8.ObjectTemplate
	globals         globals
	contexts        map[*v8.Context]*V8ScriptContext
}

func (h *V8ScriptHost) netContext(v8ctx *v8.Context) (*V8ScriptContext, bool) {
	ctx, ok := h.contexts[v8ctx]
	return ctx, ok
}

func (h *V8ScriptHost) mustGetContext(v8ctx *v8.Context) *V8ScriptContext {
	if ctx, ok := h.netContext(v8ctx); ok {
		return ctx
	}
	panic("Unknown v8 context")
}

type V8ScriptContext struct {
	htmxLoaded    bool
	mu            *sync.RWMutex
	workItemCount *atomic.Int64
	host          *V8ScriptHost
	v8ctx         *v8.Context
	window        html.Window
	pinner        runtime.Pinner
	v8nodes       map[entity.ObjectId]*v8.Value
	domNodes      map[entity.ObjectId]entity.Entity
	eventLoop     *eventLoop
	disposers     []disposable
	disposing     bool
	disposed      bool
	closer        chan bool
}

func (c *V8ScriptContext) inc() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.workItemCount.Add(1) > 0
}

func (c *V8ScriptContext) dec() { c.decN(1) }

func (c *V8ScriptContext) decN(n int64) {
	if c.workItemCount.Add(-n) < 0 {
		go func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			c.closer <- true
		}()
	}
}

// begin/end close are called when we want to close the context
func (c *V8ScriptContext) beginClose() {
	c.closer = make(chan bool)
	c.dec()
	<-c.closer
}

func (c *V8ScriptContext) endClose() {}

// begin/end callback are called when an event handler calls into script code
func (c *V8ScriptContext) beginCallback() bool { return c.inc() }
func (c *V8ScriptContext) endCallback()        { c.dec() }

// begin/end script are called when Go code directly calls eval or run
func (c *V8ScriptContext) beginScript() bool { return c.inc() }
func (c *V8ScriptContext) endScript()        { c.dec() }

// begin/end dispatch are called when something is about to be dispatched to the
// event loop. Returns true if we are allowed to execute.
func (c *V8ScriptContext) beginDispatch() bool { return c.inc() }
func (c *V8ScriptContext) endDispatch()        {}

// begin/end process are called when picking something from the event loop.
// Returns true if we are allowed to execute.
func (c *V8ScriptContext) beginProcess() bool { return c.inc() }

func (c *V8ScriptContext) endProcess() {
	// Event loop items count as two, one for when they were added, and one for
	// whey were started. Theoretically, they probably don't need to mark
	// themselves as started, and probably don't need to return a bool.
	c.decN(2)
}

func (c *V8ScriptContext) cacheNode(obj *v8.Object, node entity.Entity) (*v8.Value, error) {
	val := obj.Value
	objectId := node.ObjectId()
	c.v8nodes[objectId] = val
	c.domNodes[objectId] = node
	internal, err := v8.NewValue(c.host.iso, objectId)
	if err != nil {
		return nil, err
	}
	obj.SetInternalField(0, internal)
	return val, nil
}

func (c *V8ScriptContext) getInstanceForNode(
	node entity.Entity,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	switch n := node.(type) {
	case CustomEvent:
		return c.getInstanceForNodeByName("CustomEvent", n)
	case Event:
		return c.getInstanceForNodeByName("Event", n)
	case Element:
		if constructor, ok := scripting.HtmlElements[strings.ToLower(n.TagName())]; ok {
			return c.getInstanceForNodeByName(constructor, n)
		}
		return c.getInstanceForNodeByName("Element", n)
	case html.HTMLDocument:
		return c.getInstanceForNodeByName("HTMLDocument", n)
	case Document:
		return c.getInstanceForNodeByName("Document", n)
	case DocumentFragment:
		return c.getInstanceForNodeByName("DocumentFragment", n)
	case Node:
		return c.getInstanceForNodeByName("Node", n)
	case Attr:
		return c.getInstanceForNodeByName("Attr", n)
	default:
		panic("Cannot lookup node")
	}
}

func (c *V8ScriptContext) getInstanceForNodeByName(
	constructor string,
	node entity.Entity,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	prototype, ok := c.host.globals.namedGlobals[constructor]
	if !ok {
		panic("Bad constructor name")
	}
	value, err := prototype.InstanceTemplate().NewInstance(c.v8ctx)
	if err == nil {
		objectId := node.ObjectId()
		if cached, ok := c.v8nodes[objectId]; ok {
			return cached, nil
		}
		return c.cacheNode(value, node)
	}
	return nil, err
}

func (c *V8ScriptContext) getCachedNode(this *v8.Object) (entity.Entity, bool) {
	result, ok := c.domNodes[this.GetInternalField(0).Int32()]
	return result, ok
}

type jsConstructorFactory = func(*V8ScriptHost) *v8.FunctionTemplate

type class struct {
	globalIdentifier string
	constructor      jsConstructorFactory
	subClasses       []class
}

// createGlobals returns an ordered list of constructors to be created in global
// scope. They must be installed in "order", as base classes must be installed
// before subclasses
func createGlobals(host *V8ScriptHost) []globalInstall {
	result := make([]globalInstall, 0)
	var iter func(class classSpec) *v8.FunctionTemplate
	uniqueNames := make(map[string]*v8.FunctionTemplate)
	iter = func(class classSpec) *v8.FunctionTemplate {
		if constructor, found := uniqueNames[class.name]; found {
			return constructor
		}
		var superClassConstructor *v8.FunctionTemplate
		if class.superClassName != "" {
			superClassSpec, found := classes[class.superClassName]
			if !found {
				panic(
					"Missing super class spec. Class: " + class.name + ". Super: " + class.superClassName,
				)
			}
			superClassConstructor = iter(superClassSpec)
		}
		constructor := class.factory(host)
		if superClassConstructor != nil {
			constructor.Inherit(superClassConstructor)
		}
		uniqueNames[class.name] = constructor
		result = append(result, globalInstall{class.name, constructor})
		return constructor
	}
	for _, class := range classes {
		iter(class)
	}

	// if htmlElement, ok := uniqueNames["HTMLElement"]; ok {
	// 	for _, cls := range htmlElements {
	// 		if _, ok := uniqueNames[cls]; !ok {
	// 			fn := NewIllegalConstructorBuilder[Element](host).constructor
	// 			fn.Inherit(htmlElement)
	// 			uniqueNames[cls] = fn
	// 			result = append(result, globalInstall{cls, fn})
	// 		}
	// 	}
	// }
	return result
}

// consoleAPIMessageFunc represents a function that can receive javascript
// console messages and implements the [v8.consoleAPIMessageFunc] interface.
//
// This type is a simple solution to avoid exporting the consoleAPIMessage
// function.
type consoleAPIMessageFunc func(message v8.ConsoleAPIMessage)

func (f consoleAPIMessageFunc) ConsoleAPIMessage(message v8.ConsoleAPIMessage) {
	f(message)
}

func (host *V8ScriptHost) consoleAPIMessage(message v8.ConsoleAPIMessage) {
	switch message.ErrorLevel {
	case v8.ErrorLevelDebug:
		log.Debug(message.Message)
	case v8.ErrorLevelInfo:
	case v8.ErrorLevelLog:
		log.Info(message.Message)
	case v8.ErrorLevelWarning:
		log.Warn(message.Message)
	case v8.ErrorLevelError:
		log.Error(message.Message)
	}
}

type classSpec struct {
	name           string
	superClassName string
	factory        jsConstructorFactory
}

var classes map[string]classSpec = make(map[string]classSpec)

func registerJSClass(
	className string,
	superClassName string,
	constructorFactory jsConstructorFactory,
) {
	spec := classSpec{
		className, superClassName, constructorFactory,
	}
	if _, ok := classes[className]; ok {
		panic("Same class added twice: " + className)
	}
	if superClassName == "" {
		classes[className] = spec
		return
	}
	parent, parentFound := classes[superClassName]
	for parentFound {
		if parent.superClassName == className {
			panic("Recursive class parents" + className)
		}
		parent, parentFound = classes[parent.superClassName]
	}
	classes[className] = spec
}
func createFile(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	return v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(iso, "Illegal constructor")
		},
	)
}

func init() {
	registerJSClass("File", "", createCustomEvent)
	registerJSClass("CustomEvent", "Event", createCustomEvent)
	registerJSClass("NamedNodeMap", "", createNamedNodeMap)
	registerJSClass("Location", "", createLocationPrototype)
	registerJSClass("NodeList", "", createNodeList)
	registerJSClass("EventTarget", "", createEventTarget)

	registerJSClass("XMLHttpRequest", "EventTarget", createXmlHttpRequestPrototype)

	registerJSClass("Document", "Node", createDocumentPrototype)
	registerJSClass("HTMLDocument", "Document", createHTMLDocumentPrototype)
	registerJSClass("DocumentFragment", "Node", createDocumentFragmentPrototype)
	registerJSClass("ShadowRoot", "DocumentFragment", createShadowRootPrototype)
	registerJSClass("HTMLElement", "Element", createHtmlElement)
	registerJSClass("Attr", "Node", createAttr)

	registerJSClass("FormData", "", createFormData)
	registerJSClass("DOMParser", "", createDOMParserPrototype)

	for _, cls := range scripting.HtmlElements {
		if _, found := classes[cls]; !found {
			registerJSClass(cls, "HTMLElement", createIllegalConstructor)
		}
	}
}

func New() *V8ScriptHost {
	host := &V8ScriptHost{
		mu:  new(sync.Mutex),
		iso: v8.NewIsolate(),
	}
	host.inspectorClient = v8.NewInspectorClient(consoleAPIMessageFunc(host.consoleAPIMessage))
	host.inspector = v8.NewInspector(host.iso, host.inspectorClient)

	globalInstalls := createGlobals(host)
	host.globals = globals{make(map[string]*v8.FunctionTemplate)}
	for _, globalInstall := range globalInstalls {
		host.globals.namedGlobals[globalInstall.name] = globalInstall.constructor
	}
	constructors := host.globals.namedGlobals
	window := constructors["Window"]
	host.windowTemplate = window.InstanceTemplate()
	host.contexts = make(map[*v8.Context]*V8ScriptContext)
	installGlobals(window, host, globalInstalls)
	installEventLoopGlobals(host, host.windowTemplate)
	return host
}

func (host *V8ScriptHost) Close() {
	// host.mu.Lock()
	// defer host.mu.Unlock()
	var undiposedContexts []*V8ScriptContext
	for _, ctx := range host.contexts {
		undiposedContexts = append(undiposedContexts, ctx)
	}
	undisposedCount := len(undiposedContexts)

	if undisposedCount > 0 {
		log.Warn("Script host shutdown: Not all contexts disposed", "count", len(host.contexts))
		for _, ctx := range undiposedContexts {
			ctx.Close()
		}
	}
	host.inspectorClient.Dispose()
	host.inspector.Dispose()
	host.iso.Dispose()
}

var global *v8.Object

func (host *V8ScriptHost) NewContext(w html.Window) html.ScriptContext {
	context := &V8ScriptContext{
		mu:            new(sync.RWMutex),
		workItemCount: new(atomic.Int64),
		host:          host,
		v8ctx:         v8.NewContext(host.iso, host.windowTemplate),
		window:        w,
		v8nodes:       make(map[entity.ObjectId]*v8.Value),
		domNodes:      make(map[entity.ObjectId]entity.Entity),
	}
	host.inspector.ContextCreated(context.v8ctx)
	err := installPolyfills(context)
	if err != nil {
		// TODO: Handle
		panic(err)
	}
	global = context.v8ctx.Global()
	errorCallback := func(err error) {
		w.DispatchEvent(NewErrorEvent(err))
	}
	context.eventLoop = newEventLoop(context, global, errorCallback)
	host.contexts[context.v8ctx] = context
	context.cacheNode(global, w)
	context.addDisposer(context.eventLoop.Start())

	return context
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (ctx *V8ScriptContext) Close() {
	ctx.beginClose()
	defer ctx.endClose()
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.disposed {
		panic("Context already disposed")
	}
	ctx.disposed = true
	ctx.host.inspector.ContextDestroyed(ctx.v8ctx)
	log.Debug("ScriptContext: Dispose")
	for _, dispose := range ctx.disposers {
		dispose.dispose()
	}
	ctx.pinner.Unpin()
	// TODO: Synchronize
	delete(ctx.host.contexts, ctx.v8ctx)
	ctx.v8ctx.Close()
}

func (ctx *V8ScriptContext) addDisposer(disposer disposable) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.disposers = append(ctx.disposers, disposer)
}

func (ctx *V8ScriptContext) runScript(script string) (res *v8.Value, err error) {
	defer ctx.endScript()
	if ctx.beginScript() {
		res, err = ctx.v8ctx.RunScript(script, "")
	}
	return
}

func (ctx *V8ScriptContext) Run(script string) error {
	_, err := ctx.runScript(script)
	return err
}

func (ctx *V8ScriptContext) Eval(script string) (interface{}, error) {
	result, err := ctx.runScript(script)
	if err == nil {
		return v8ValueToGoValue(result)
	}
	return nil, err
}

func (ctx *V8ScriptContext) EvalCore(script string) (any, error) {
	return ctx.runScript(script)
}

func (ctx *V8ScriptContext) RunFunction(script string, arguments ...any) (res any, err error) {
	var (
		v  *v8.Value
		f  *v8.Function
		ok bool
	)
	if v, err = ctx.runScript(script); err == nil {
		f, err = v.AsFunction()
	}
	if err == nil {
		args := make([]v8.Valuer, len(arguments))
		for i, a := range arguments {
			if args[i], ok = a.(v8.Valuer); !ok {
				err = fmt.Errorf("V8ScriptContext.RunFunction: Arguments is not a V8 value: %d", i)
			}
		}
		return f.Call(ctx.v8ctx.Global(), args...)
	}
	return
}

func (ctx *V8ScriptContext) Export(val any) (any, error) {
	if res, ok := val.(*v8.Value); ok {
		return v8ValueToGoValue(res)
	} else {
		return nil, errors.New("V8ScriptContext.Export: value not a V8 value")
	}
}

func v8ValueToGoValue(result *v8go.Value) (interface{}, error) {
	if result == nil {
		return nil, nil
	}
	if result.IsBoolean() {
		return result.Boolean(), nil
	}
	if result.IsInt32() {
		return result.Int32(), nil
	}
	if result.IsString() {
		return result.String(), nil
	}
	if result.IsNull() {
		return nil, nil
	}
	if result.IsUndefined() {
		return nil, nil
	}
	if result.IsArray() {
		obj, _ := result.AsObject()
		length, err := obj.Get("length")
		l := length.Uint32()
		errs := make([]error, l+1)
		errs[0] = err
		result := make([]any, l)
		for i := uint32(0); i < l; i++ {
			val, err := obj.GetIdx(i)
			if err == nil {
				result[i], err = v8ValueToGoValue(val)
			}
			errs[i+1] = err
		}
		return result, errors.Join(errs...)
	}
	return nil, fmt.Errorf("Value not yet supported: %v", *result)
}
