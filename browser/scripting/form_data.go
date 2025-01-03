package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

func CreateFormData(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewConstructorBuilder[*dom.FormData](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			var e dom.Entity = dom.NewFormData()
			return ctx.CacheNode(info.This(), e)
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
		func(instance dom.FormDataEntry, ctx *ScriptContext) (*v8.Value, error) {
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
	builder.constructor.InstanceTemplate().SetSymbol(
		v8.SymbolIterator(iso),
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			data, err := builder.instanceLookup(ctx, info.This())
			if err != nil {
				return nil, err
			}
			return stringIterator.NewIteratorInstance(ctx, data.Keys())
		},
	)
	protoBuilder.CreateFunction(
		"append",
		func(instance *dom.FormData, args argumentHelper) (res *v8.Value, err error) {
			key, err1 := args.GetStringArg(0)
			value, err2 := args.GetStringArg(1)
			err = errors.Join(err1, err2)
			if err != nil {
				return
			}
			instance.Append(key, dom.FormDataValue(value))
			return
		},
	)
	protoBuilder.CreateFunction(
		"get",
		func(instance *dom.FormData, args argumentHelper) (result *v8.Value, err error) {
			var key string
			key, err = args.GetStringArg(0)
			if err != nil {
				return
			}
			val := string(instance.Get(key))
			return v8.NewValue(iso, val)
		},
	)
	protoBuilder.CreateFunction(
		"keys",
		func(instance *dom.FormData, args argumentHelper) (result *v8.Value, err error) {
			return stringIterator.NewIteratorInstance(args.ctx, instance.Keys())
		},
	)

	getEntries := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			instance, err := builder.instanceLookup(ctx, info.This())
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
