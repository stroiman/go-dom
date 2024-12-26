package scripting

import v8 "github.com/tommie/v8go"

func CreateMembers(iso *v8.Isolate, ft *v8.FunctionTemplate) {
	proto := ft.PrototypeTemplate()
	proto.SetAccessorProperty(
		"content",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				return info.This().Value, nil
			},
		),
		nil,
		v8.ReadOnly,
	)
}
