package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type ESElement struct {
	ESWrapper[Element]
}

func (e ESElement) ClassList(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	tokenList := e.host.globals.namedGlobals["DOMTokenList"]
	ctx := e.host.MustGetContext(info.Context())
	instance, err := tokenList.GetInstanceTemplate().NewInstance(ctx.v8ctx)
	if err != nil {
		return nil, err
	}
	element, err := e.GetInstance(info)
	if err != nil {
		return nil, err
	}
	value, err := v8.NewValue(e.host.iso, element.ObjectId())
	if err != nil {
		return nil, err
	}
	instance.SetInternalField(0, value)
	return instance.Value, nil
}

func CreateElement(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := ESElement{NewESWrapper[Element](host)}
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
	prototypeTemplate := helper.proto
	prototypeTemplate.SetAccessorPropertyCallback("classList", wrapper.ClassList, nil, v8.ReadOnly)
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
