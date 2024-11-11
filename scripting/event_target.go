package scripting

import (
	v8 "github.com/tommie/v8go"
)

func CreateEventTarget(iso *v8.Isolate) *v8.FunctionTemplate {
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
