package v8host

import (
	v8 "github.com/tommie/v8go"
)

func createHtmlElement(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplateWithError(
		iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(iso, "Illegal Constructor")
		},
	)
	instanceTemplate := res.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
