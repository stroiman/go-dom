package scripting

import (
	v8 "github.com/tommie/v8go"
)

func installGlobals(
	windowFnTemplate *v8.FunctionTemplate,
	host *ScriptHost,
	globalInstalls []globalInstall,
) {
	windowTemplate := windowFnTemplate.InstanceTemplate()
	for _, globalInstall := range globalInstalls {
		windowTemplate.Set(globalInstall.name, globalInstall.constructor)
	}
	location := host.globals.namedGlobals["Location"]
	windowTemplate.Set("location", location.InstanceTemplate())
}

func (w *WindowV8Wrapper) Window(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return info.This().Value, nil
}
