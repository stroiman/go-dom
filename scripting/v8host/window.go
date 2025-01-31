package v8host

import (
	"runtime/cgo"

	v8 "github.com/tommie/v8go"
)

func installGlobals(
	windowFnTemplate *v8.FunctionTemplate,
	host *V8ScriptHost,
	globalInstalls []globalInstall,
) {
	windowTemplate := windowFnTemplate.InstanceTemplate()
	for _, globalInstall := range globalInstalls {
		windowTemplate.Set(globalInstall.name, globalInstall.constructor)
	}
	location := host.globals.namedGlobals["Location"]
	windowTemplate.Set("location", location.InstanceTemplate())
}

func (w *windowV8Wrapper) window(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return info.This().Value, nil
}

func (w *windowV8Wrapper) history(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	win, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	ctx := w.mustGetContext(info)
	history, err := w.scriptHost.globals.namedGlobals["History"].InstanceTemplate().
		NewInstance(ctx.v8ctx)
	if err != nil {
		return nil, err
	}
	handle := cgo.NewHandle(win.History())
	ctx.addDisposer(handleDisposable(handle))
	internal := v8.NewValueExternalHandle(w.iso(), handle)
	history.SetInternalField(0, internal)
	return history.Value, nil
}
