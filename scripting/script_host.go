package scripting

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/stroiman/go-dom/browser"
	. "github.com/stroiman/go-dom/browser"

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

type ScriptHost struct {
	iso            *v8.Isolate
	windowTemplate *v8.ObjectTemplate
	globals        globals
	contexts       map[*v8.Context]*ScriptContext
}

func (h *ScriptHost) GetContext(v8ctx *v8.Context) (*ScriptContext, bool) {
	ctx, ok := h.contexts[v8ctx]
	return ctx, ok
}

func (h *ScriptHost) MustGetContext(v8ctx *v8.Context) *ScriptContext {
	if ctx, ok := h.GetContext(v8ctx); ok {
		return ctx
	}
	panic("Unknown v8 context")
}

type ScriptContext struct {
	host     *ScriptHost
	v8ctx    *v8.Context
	window   Window
	pinner   runtime.Pinner
	v8nodes  map[ObjectId]*v8.Value
	domNodes map[ObjectId]Entity
}

func (c *ScriptContext) CacheNode(obj *v8.Object, node Entity) (*v8.Value, error) {
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

func (c *ScriptContext) GetInstanceForNode(
	prototype *v8.FunctionTemplate,
	node Node,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	value, err := prototype.GetInstanceTemplate().NewInstance(c.v8ctx)
	if err == nil {
		objectId := node.ObjectId()
		if cached, ok := c.v8nodes[objectId]; ok {
			return cached, nil
		}
		return c.CacheNode(value, node)
	}
	return nil, err
}

func (c *ScriptContext) GetInstanceForNodeByName(
	constructor string,
	node Entity,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	prototype, ok := c.host.globals.namedGlobals[constructor]
	if !ok {
		panic("Bad constructor name")
	}
	value, err := prototype.GetInstanceTemplate().NewInstance(c.v8ctx)
	if err == nil {
		objectId := node.ObjectId()
		if cached, ok := c.v8nodes[objectId]; ok {
			return cached, nil
		}
		return c.CacheNode(value, node)
	}
	return nil, err
}

func (c *ScriptContext) GetCachedNode(this *v8.Object) (Entity, bool) {
	result, ok := c.domNodes[this.GetInternalField(0).Int32()]
	return result, ok
}

type class struct {
	globalIdentifier string
	constructor      func(*ScriptHost) *v8.FunctionTemplate
	subClasses       []class
}

// createGlobals returns an ordered list of constructors to be created in global
// scope. They must be installed in "order", as base classes must be installed
// before subclasses
func createGlobals(host *ScriptHost, classes []class) []globalInstall {
	result := make([]globalInstall, 0)
	var iter func(*v8.FunctionTemplate, []class)
	iter = func(superClass *v8.FunctionTemplate, classes []class) {
		for _, class := range classes {
			constructor := class.constructor(host)
			result = append(result, globalInstall{class.globalIdentifier, constructor})
			if superClass != nil {
				constructor.Inherit(superClass)
			}
			iter(constructor, class.subClasses)
		}
	}
	iter(nil, classes)
	return result
}

func NewScriptHost() *ScriptHost {
	host := &ScriptHost{iso: v8.NewIsolate()}
	classes := []class{
		{"CustomEvent", CreateCustomEvent, nil},
		{"NamedNodeMap", CreateNamedNodeMap, nil},
		{"Location", CreateLocationPrototype, nil},
		{"EventTarget", CreateEventTarget, []class{
			{"Window", CreateWindowTemplate, nil},
			{"Node", CreateNode, []class{
				{"Document", CreateDocumentPrototype, nil},
				{"DocumentFragment", CreateDocumentFragmentPrototype, []class{
					{"ShadowRoot", CreateShadowRootPrototype, nil},
				}},
				{"Element", CreateElement, []class{
					{"HTMLElement", CreateHtmlElement, nil},
				}},
				{"Attr", CreateAttr, nil},
			}},
		}},
	}

	globalInstalls := createGlobals(host, classes)
	host.globals = globals{make(map[string]*v8.FunctionTemplate)}
	for _, globalInstall := range globalInstalls {
		host.globals.namedGlobals[globalInstall.name] = globalInstall.constructor
	}
	constructors := host.globals.namedGlobals
	window := constructors["Window"]
	host.windowTemplate = window.GetInstanceTemplate()
	host.contexts = make(map[*v8.Context]*ScriptContext)
	installGlobals(window, host, globalInstalls)
	return host
}

func (host *ScriptHost) Dispose() {
	host.iso.Dispose()
}

var global *v8.Object

func (host *ScriptHost) NewContext(window Window) *ScriptContext {
	context := &ScriptContext{
		host:     host,
		v8ctx:    v8.NewContext(host.iso, host.windowTemplate),
		window:   window,
		v8nodes:  make(map[ObjectId]*v8.Value),
		domNodes: make(map[ObjectId]Entity),
	}
	global = context.v8ctx.Global()
	host.contexts[context.v8ctx] = context
	context.CacheNode(global, window)

	return context
}

type Wrapper ScriptHost

func (w *Wrapper) NewScriptEngine(window Window) browser.ScriptEngine {
	host := (*ScriptHost)(w)
	return host.NewContext(window)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (ctx *ScriptContext) Dispose() {
	ctx.pinner.Unpin()
	delete(ctx.host.contexts, ctx.v8ctx)
	ctx.v8ctx.Close()
}

func (ctx *ScriptContext) RunScript(script string) (*v8.Value, error) {
	return ctx.v8ctx.RunScript(script, "")
}

func (ctx *ScriptContext) Run(script string) error {
	_, err := ctx.RunScript(script)
	return err
}

func (ctx *ScriptContext) Eval(script string) (interface{}, error) {
	result, err := ctx.RunScript(script)
	if err == nil {
		return v8ValueToGoValue(result)
	}
	return nil, err
}

func (ctx *ScriptContext) Window() Window {
	return ctx.window
}

// TODO: Refactor, deps are totally the wrong way around
func (ctx *ScriptContext) NewBrowserFromHandler(handler http.Handler) Browser {
	browser := NewBrowserFromHandler(handler)
	browser.ScriptEngine = ctx
	return browser
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
	return nil, fmt.Errorf("Value not yet supported: %v", *result)
}
