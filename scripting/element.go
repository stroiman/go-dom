package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateElement(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewIllegalConstructorBuilder[Element](host)
	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (Element, error) {
		element, ok := ctx.GetCachedNode(this)
		if e, e_ok := element.(Element); e_ok && ok {
			return e, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of Element")
		}
	}
	helper := builder.NewPrototypeBuilder()
	helper.CreateReadonlyProp("outerHTML", Element.OuterHTML)
	helper.CreateReadonlyProp("tagName", Element.TagName)
	helper.CreateFunctionStringToString("getAttribute", Element.GetAttribute)
	helper.CreateReadonlyProp2(
		"attributes",
		func(element Element, ctx *ScriptContext) (*v8.Value, error) {
			return ctx.GetInstanceForNodeByName("NamedNodeMap", element.GetAttributes())
		},
	)

	helper.CreateFunction(
		"insertAdjacentHTML",
		func(element Element, info argumentHelper) (val *v8.Value, err error) {
			position, e1 := info.GetStringArg(0)
			html, e2 := info.GetStringArg(1)
			err = errors.Join(e1, e2)
			if err == nil {
				element.InsertAdjacentHTML(position, html)
				val, err = v8.NewValue(iso, element.OuterHTML())
			}
			return
		},
	)
	return builder.constructor
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
