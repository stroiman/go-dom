package scripting

import (
	"runtime"
	"unsafe"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type V8Window struct {
	host     *ScriptHost
	window   Window
	document *V8Document
}

func NewV8Window(host *ScriptHost, w Window) *V8Window {
	return &V8Window{
		host:   host,
		window: w,
	}
}

func (w *V8Window) V8Document(ctx *v8.Context) *V8Document {
	if w.document == nil {
		document := &V8Document{}
		// v8.FunctionCallbackInfo
		value, err := w.host.document.GetInstanceTemplate().NewInstance(ctx)
		if err != nil {
			panic(err.Error())
		}
		value.SetInternalField(0, unsafe.Pointer(document))
		w.document = document
		document.Value = value.Value
	}
	return w.document
}

type ScriptHost struct {
	iso            *v8.Isolate
	windowTemplate *v8.ObjectTemplate
	document       *v8.FunctionTemplate
}

type ScriptContext struct {
	ctx    *v8.Context
	window Window
	pinner runtime.Pinner
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
				return v8window.V8Document(info.Context()).Value
			},
		})
	windowTemplate.Set("Document", host.document)
	return windowTemplate
}

func NewScriptHost() *ScriptHost {
	host := &ScriptHost{iso: v8.NewIsolate()}
	host.document = CreateDocumentPrototype(host.iso)
	host.windowTemplate = CreateWindowTemplate(host)
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
		ctx:    v8.NewContext(host.iso, host.windowTemplate),
		window: window,
	}
	global = context.ctx.Global()
	ptr := unsafe.Pointer(v8window)
	context.pinner.Pin(ptr)
	context.pinner.Pin(v8window.window)
	context.pinner.Pin(v8window.host)
	global.SetInternalField(0, v8.NewExternalValue(host.iso, ptr))

	return context
}

func (ctx *ScriptContext) Dispose() {
	ctx.ctx.Close()
}

func (ctx *ScriptContext) RunScript(script string) (*v8.Value, error) {
	return ctx.ctx.RunScript(script, "")
}

func (ctx *ScriptContext) Window() Window {
	return ctx.window
}
