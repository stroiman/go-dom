package scripting

import (
	"runtime"
	"unsafe"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type V8Window struct {
	host   *ScriptHost
	window Window
}

func NewV8Window(host *ScriptHost, w Window) *V8Window {
	return &V8Window{
		host:   host,
		window: w,
	}
}

func (w *V8Window) V8Document(ctx *v8.Context) *v8.Value {
	domDocument := w.window.Document()
	if cached, ok := w.host.contexts[ctx].v8nodes[domDocument.ObjectId()]; ok {
		return cached
	}
	object, err := w.host.document.GetInstanceTemplate().NewInstance(ctx)
	if err != nil {
		panic(err) // TODO
	}
	value := object.Value
	docNode := w.window.Document()
	id := docNode.ObjectId()
	myContext := w.host.contexts[ctx]
	myContext.v8nodes[id] = value
	myContext.domNodes[id] = docNode
	object.SetInternalField(
		0,
		v8.NewExternalValue(ctx.Isolate(), unsafe.Pointer(id)),
	)
	return value
}

type ScriptHost struct {
	iso            *v8.Isolate
	windowTemplate *v8.ObjectTemplate
	document       *v8.FunctionTemplate
	node           *v8.FunctionTemplate
	eventTarget    *v8.FunctionTemplate
	contexts       map[*v8.Context]*ScriptContext
}

type ScriptContext struct {
	host     *ScriptHost
	v8ctx    *v8.Context
	window   Window
	pinner   runtime.Pinner
	v8nodes  map[ObjectId]*v8.Value
	domNodes map[ObjectId]Node
}

func CreateWindowTemplate(host *ScriptHost) *v8.ObjectTemplate {
	iso := host.iso
	windowTemplate := v8.NewObjectTemplate(iso)
	windowTemplate.SetInternalFieldCount(1)
	windowTemplate.SetAccessorProperty(
		"window",
		v8.AccessProp{
			Get: func(i *v8.FunctionCallbackInfo) *v8.Value {
				return i.This().Value
			},
			Attributes: v8.ReadOnly,
		},
	)
	windowTemplate.SetAccessorProperty(
		"document",
		v8.AccessProp{
			Get: func(info *v8.FunctionCallbackInfo) *v8.Value {
				v8window := (*V8Window)(info.This().GetInternalField(0).External())
				return v8window.V8Document(info.Context())
			},
		})
	windowTemplate.Set("Document", host.document)
	windowTemplate.Set("Node", host.node)
	windowTemplate.Set("EventTarget", host.eventTarget)
	return windowTemplate
}

func NewScriptHost() *ScriptHost {
	host := &ScriptHost{iso: v8.NewIsolate()}
	host.document = CreateDocumentPrototype(host)
	host.node = CreateNode(host.iso)
	host.eventTarget = CreateEventTarget(host)
	host.windowTemplate = CreateWindowTemplate(host)
	host.document.Inherit(host.node)
	host.node.Inherit(host.eventTarget)
	host.contexts = make(map[*v8.Context]*ScriptContext)
	return host
}

func (host *ScriptHost) Dispose() {
	host.iso.Dispose()
}

var global *v8.Object

func (host *ScriptHost) NewContext() *ScriptContext {
	window := NewWindow()
	v8window := NewV8Window(host, window)
	context := &ScriptContext{
		host:     host,
		v8ctx:    v8.NewContext(host.iso, host.windowTemplate),
		window:   window,
		v8nodes:  make(map[ObjectId]*v8.Value),
		domNodes: make(map[ObjectId]Node),
	}
	global = context.v8ctx.Global()
	host.contexts[context.v8ctx] = context
	ptr := unsafe.Pointer(v8window)
	context.pinner.Pin(ptr)
	context.pinner.Pin(v8window.window)
	context.pinner.Pin(v8window.host)
	global.SetInternalField(0, v8.NewExternalValue(host.iso, ptr))

	return context
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (host *ScriptHost) createPrototypeChains() {
	host.document.Inherit(host.node)
}

func (ctx *ScriptContext) Dispose() {
	delete(ctx.host.contexts, ctx.v8ctx)
	ctx.v8ctx.Close()
}

func (ctx *ScriptContext) RunScript(script string) (*v8.Value, error) {
	return ctx.v8ctx.RunScript(script, "")
}

func (ctx *ScriptContext) Run(script string) (interface{}, error) {
	return ctx.v8ctx.RunScript(script, "")
}

func (ctx *ScriptContext) Window() Window {
	return ctx.window
}
