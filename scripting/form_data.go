package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateFormData(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewConstructorBuilder[*browser.FormData](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			var e browser.Entity = browser.NewFormData()
			return ctx.CacheNode(info.This(), e)
		},
	)
	builder.SetDefaultInstanceLookup()
	protoBuilder := builder.NewPrototypeBuilder()
	protoBuilder.CreateFunction(
		"append",
		func(instance *browser.FormData, args argumentHelper) (res *v8.Value, err error) {
			key, err1 := args.GetStringArg(0)
			value, err2 := args.GetStringArg(1)
			err = errors.Join(err1, err2)
			if err != nil {
				return
			}
			instance.Append(key, browser.FormDataValue(value))
			return
		},
	)
	protoBuilder.CreateFunction(
		"get",
		func(instance *browser.FormData, args argumentHelper) (result *v8.Value, err error) {
			var key string
			key, err = args.GetStringArg(0)
			if err != nil {
				return
			}
			val := string(instance.Get(key))
			return v8.NewValue(iso, val)
		},
	)
	return builder.constructor
}
