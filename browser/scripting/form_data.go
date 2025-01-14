package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type FormDataV8Wrapper struct {
	handleReffedObject[*html.FormData]
}

func NewFormDataV8Wrapper(host *ScriptHost) FormDataV8Wrapper {
	return FormDataV8Wrapper{newHandleReffedObject[*html.FormData](host)}
}

func (w FormDataV8Wrapper) CreateInstance(
	ctx *ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	var value = html.NewFormData()
	w.store(value, ctx, this)
	return nil, nil
}

func createFormData(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewFormDataV8Wrapper(host)
	builder := NewConstructorBuilder[*html.FormData](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			return wrapper.CreateInstance(ctx, info.This())
		},
	)
	stringIterator := NewIterator(
		host,
		func(instance string, ctx *ScriptContext) (*v8.Value, error) {
			return v8.NewValue(ctx.host.iso, instance)
		},
	)

	entryIterator := NewIterator(
		host,
		func(instance html.FormDataEntry, ctx *ScriptContext) (*v8.Value, error) {
			// TODO, no option to create an array, totally a hack!
			arr, e1 := ctx.RunScript("(k,v) => [k,v]")
			if e1 != nil {
				return nil, e1
			}
			f, e2 := arr.AsFunction()
			k, e3 := v8.NewValue(iso, instance.Name)
			v, e4 := v8.NewValue(iso, string(instance.Value))
			err := errors.Join(e2, e3, e4)
			if err != nil {
				return nil, err
			}
			res, err := f.Call(v8.Null(iso), k, v)
			return res, err
		})
	builder.SetDefaultInstanceLookup()
	protoBuilder := builder.NewPrototypeBuilder()
	prototype := protoBuilder.proto
	builder.constructor.InstanceTemplate().SetSymbol(
		v8.SymbolIterator(iso),
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			data, err := wrapper.getInstance(info)
			if err != nil {
				return nil, err
			}
			return stringIterator.NewIteratorInstance(ctx, data.Keys())
		},
	)
	// protoBuilder.CreateFunction(
	// 	"append",
	// 	func(instance *dom.FormData, args argumentHelper) (res *v8.Value, err error) {
	// 		key, err1 := args.GetStringArg(0)
	// 		value, err2 := args.GetStringArg(1)
	// 		err = errors.Join(err1, err2)
	// 		if err != nil {
	// 			return
	// 		}
	// 		instance.Append(key, dom.FormDataValue(value))
	// 		return
	// 	},
	// )
	prototype.Set(
		"append",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				key, err1 := args.getStringArg(0)
				value, err2 := args.getStringArg(1)
				err := errors.Join(err0, err1, err2)
				if err != nil {
					return nil, err
				}
				instance.Append(key, html.FormDataValue(value))
				return nil, nil
			},
		),
	)

	prototype.Set(
		"get",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				if err0 != nil {
					return nil, err0
				}
				key, err := args.getStringArg(0)
				if err != nil {
					return nil, err
				}
				val := string(instance.Get(key))
				return v8.NewValue(iso, val)
			},
		),
	)
	prototype.Set(
		"keys",
		v8.NewFunctionTemplateWithError(host.iso,
			func(info *v8.FunctionCallbackInfo) (result *v8.Value, err error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				if err0 != nil {
					return nil, err0
				}
				return stringIterator.NewIteratorInstance(args.ctx, instance.Keys())
			}),
	)

	getEntries := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			instance, err := wrapper.getInstance(info)
			if err != nil {
				return nil, err
			}
			return entryIterator.NewIteratorInstance(ctx, instance.Entries)
		},
	)
	protoBuilder.proto.Set("entries", getEntries)
	protoBuilder.proto.SetSymbol(v8.SymbolIterator(iso), getEntries)
	return builder.constructor
}
