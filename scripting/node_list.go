package scripting

import (
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

/*
Instance properties

    length

Instance methods

    entries()  // returns an iterator
    forEach()  // calls a callback
    item()
    keys()
    values()


## Foreach parameters

    // foreach Behaviour in FF (by experimenting)
    // for (i = 0; i < length; i++) {
    //   element = item(i)
    //   if (element) { callback(element, i) }
    // }
    // Inserting an element _before_ current element will iterate current
// element twice (it has a new index), but last element isn't iterated.
    // Removing an element, and it doesn't iterate _past_ the end of the element

callback

    A function to execute on each element of someNodeList. It accepts 3 parameters:

    currentValue

        The current element being processed in someNodeList.
    currentIndex Optional

        The index of the currentValue being processed in someNodeList.
    listObj Optional

        The someNodeList that forEach() is being applied to.

thisArg Optional

    Value to use as this when executing callback.
*/

func CreateNodeList(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	iteratorResultTemplate := v8.NewObjectTemplate(iso)
	iteratorTemplate := v8.NewObjectTemplate(iso)
	iteratorTemplate.SetInternalFieldCount(2)
	createDoneIteratorResult := func(ctx *v8.Context) (*v8.Value, error) {
		result, err := iteratorResultTemplate.NewInstance(ctx)
		if err != nil {
			return nil, err
		}
		result.Set("done", true)
		return result.Value, nil
	}
	createNotDoneIteratorResult := func(ctx *v8.Context, value interface{}) (*v8.Value, error) {
		result, err := iteratorResultTemplate.NewInstance(ctx)
		if err != nil {
			return nil, err
		}
		result.Set("done", false)
		result.Set("value", value)
		return result.Value, nil
	}
	builder := NewIllegalConstructorBuilder[browser.NodeList](host)
	builder.SetDefaultInstanceLookup()
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp2(
		"length",
		func(instance browser.NodeList, ctx *ScriptContext) (*v8.Value, error) {
			return v8.NewValue(iso, uint32(instance.Length()))
		},
	)
	proto.CreateFunction(
		"item",
		func(instance browser.NodeList, info argumentHelper) (*v8.Value, error) {
			index, err := info.GetInt32Arg(0)
			if err != nil {
				return nil, v8.NewTypeError(iso, "Index must be an integer")
			}
			result := instance.Item(int(index))
			if result == nil {
				return v8.Null(iso), nil
			}
			return info.ctx.GetInstanceForNodeByName("Element", result)
		},
	)
	instanceTemplate := builder.NewInstanceBuilder().proto
	instanceTemplate.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				instance, err := iteratorTemplate.NewInstance(info.Context())
				instance.SetInternalField(0, info.This().GetInternalField(0))
				instance.Set(
					"next",
					*v8.NewFunctionTemplateWithError(iso,
						func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
							ctx := host.MustGetContext(info.Context())
							nodeList, err := getInstanceFromThis[browser.NodeList](ctx, info.This())
							if err != nil {
								return nil, err
							}
							index := (info.This().GetInternalField(1).Int32())
							item := nodeList.Item(int(index))
							if item == nil {
								return createDoneIteratorResult(ctx.v8ctx)
							} else {
								item := nodeList.Item(int(index))
								item_instance, err := ctx.GetInstanceForNode(item)
								if err != nil {
									return nil, err
								}
								result, err := createNotDoneIteratorResult(ctx.v8ctx, item_instance)
								return result, info.This().SetInternalField(1, index+1)
							}
						},
					).GetFunction(info.Context()),
				)
				return instance.Value, err
			},
		),
	)
	instanceTemplate.SetIndexedHandler(
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			instance, ok := ctx.GetCachedNode(info.This())
			nodemap, ok_2 := instance.(browser.NodeList)
			if ok && ok_2 {
				index := int(info.Index())
				item := nodemap.Item(index)
				if item == nil {
					return v8.Undefined(iso), nil
				}
				return ctx.GetInstanceForNode(item)
			}
			return nil, v8.NewTypeError(iso, "dunno")
		},
	)

	return builder.constructor
}
