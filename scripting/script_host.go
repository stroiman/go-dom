package scripting

import (
	v8 "github.com/tommie/v8go"
)

type Window struct {
}

func NewWindow() *Window {
	return &Window{}
}

type ScriptHost struct {
	iso            *v8.Isolate
	windowTemplate *v8.ObjectTemplate
}

type ScriptContext struct {
	ctx    *v8.Context
	window *Window
}

func CreateWindowTemplate(iso *v8.Isolate) *v8.ObjectTemplate {
	windowTemplate := v8.NewObjectTemplate(iso)
	windowTemplate.SetInternalFieldCount(1)
	windowTemplate.SetAccessorProperty(
		"window",
		*v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
			return info.This().Value
		}),
		*v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
			return v8.Undefined(iso)
		}),
	)
	return windowTemplate
}

func NewScriptHost() *ScriptHost {
	iso := v8.NewIsolate()

	return &ScriptHost{iso, CreateWindowTemplate(iso)}
}

func (host *ScriptHost) Dispose() {
	host.iso.Dispose()
}

func (host *ScriptHost) NewContext() *ScriptContext {
	window := NewWindow()
	context := &ScriptContext{
		ctx:    v8.NewContext(host.iso, host.windowTemplate),
		window: window,
	}
	context.ctx.Global().SetInternalField(0, window)

	return context
}

func (ctx *ScriptContext) Dispose() {
	ctx.ctx.Close()
}

func (ctx *ScriptContext) RunScript(script string) (*v8.Value, error) {
	return ctx.ctx.RunScript(script, "")
}
