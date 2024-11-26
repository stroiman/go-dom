package scripting

import (
	"fmt"

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
	builder.NewInstanceBuilder().proto.SetIndexedHandler(
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			fmt.Println("Indexed property handler!!!")
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
