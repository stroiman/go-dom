package scripting

import (
	"errors"
	"fmt"

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
	stringIterator := NewIterator(
		host,
		func(instance string, ctx *ScriptContext) (*v8.Value, error) {
			fmt.Println("Create instance", instance)
			return v8.NewValue(ctx.host.iso, instance)
		},
	)
	entryIterator := NewIterator(
		host,
		func(instance browser.FormDataEntry, ctx *ScriptContext) (*v8.Value, error) {
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
	builder.constructor.GetInstanceTemplate().SetSymbol(
		v8.SymbolIterator(iso),
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			fmt.Println("\n\n*** Iterator")
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
	protoBuilder.CreateFunction(
		"keys",
		func(instance *browser.FormData, args argumentHelper) (result *v8.Value, err error) {
			return stringIterator.NewIteratorInstance(args.ctx, instance.Keys())
		},
	)

	protoBuilder.CreateFunction(
		"entries",
		func(instance *browser.FormData, args argumentHelper) (result *v8.Value, err error) {
			return entryIterator.NewIteratorInstance(args.ctx, instance.Entries)
		},
	)
	return builder.constructor
}
