package scripting

import (
	v8 "github.com/tommie/v8go"
)

func CreateEventTarget(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplate(
		iso,
		func(args *v8.FunctionCallbackInfo) *v8.Value {
			return v8.Undefined(iso)
		},
	)
	instanceTemplate := res.GetInstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
