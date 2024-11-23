package scripting

import (
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

	helper.CreateFunction(
		"insertAdjacentHTML",
		func(element Element, info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			args := info.Args()
			if len(args) < 2 {
				return nil, v8.NewTypeError(iso, "Not enough argument")
			}
			position := args[0].String()
			html := args[1].String()
			element.InsertAdjacentHTML(position, html)
			return v8.NewValue(iso, element.OuterHTML())
		},
	)
	return builder.constructor
}
