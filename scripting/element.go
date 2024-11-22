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
	helper.CreateFunction("getAttribute", Element.GetAttribute)
	helper.proto.Set(
		"insertAdjacentHTML",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				this, ok := ctx.GetCachedNode(info.This())
				if e, e_ok := this.(Element); e_ok && ok {
					args := info.Args()
					if len(args) < 2 {
						return nil, v8.NewTypeError(iso, "Not enough argument")
					}
					position := args[0].String()
					html := args[1].String()
					e.InsertAdjacentHTML(position, html)
					return v8.NewValue(iso, e.OuterHTML())
				} else {
					return nil, v8.NewTypeError(iso, "Not an instance of Element")
				}
			},
		),
	)
	return builder.constructor
}
