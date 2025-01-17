package v8host

import (
	"errors"

	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

type esElement struct {
	esElementContainerWrapper[Element]
}

func (e esElement) ClassList(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	tokenList := e.host.globals.namedGlobals["DOMTokenList"]
	ctx := e.host.MustGetContext(info.Context())
	instance, err := tokenList.InstanceTemplate().NewInstance(ctx.v8ctx)
	if err != nil {
		return nil, err
	}
	element, err := e.getInstance(info)
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

func createElement(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := esElement{NewESContainerWrapper[Element](host)}
	builder := NewIllegalConstructorBuilder[Element](host)
	wrapper.Install(builder.constructor)
	builder.SetDefaultInstanceLookup()
	helper := builder.NewPrototypeBuilder()
	prototypeTemplate := helper.proto
	prototypeTemplate.SetAccessorPropertyCallback("classList", wrapper.ClassList, nil, v8.ReadOnly)
	prototypeTemplate.SetAccessorPropertyCallback(
		"textContent",
		nil,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			e, err := wrapper.getInstance(info)
			if err == nil {
				e.SetTextContent(info.Args()[0].String())
			}
			return nil, err
		},
		v8.None,
	)
	helper.CreateReadonlyProp("outerHTML", Element.OuterHTML)
	helper.CreateReadonlyProp("tagName", Element.TagName)
	prototypeTemplate.Set(
		"getAttribute",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				helper := newArgumentHelper(host, info)
				element, e0 := builder.GetInstance(info)
				name, e1 := helper.getStringArg(0)
				err := errors.Join(e0, e1)
				if err != nil {
					return nil, err
				}
				if r, ok := element.GetAttribute(name); ok {
					return v8.NewValue(iso, r)
				} else {
					return v8.Null(iso), nil
				}
			},
		),
	)
	helper.CreateFunction(
		"hasAttribute",
		func(instance Element, info argumentHelper) (*v8.Value, error) {
			name, e1 := info.getStringArg(0)
			_, ok := instance.GetAttribute(name)
			result, e2 := v8.NewValue(iso, ok)
			return result, errors.Join(e1, e2)
		},
	)
	helper.CreateReadonlyProp2(
		"attributes",
		func(element Element, ctx *V8ScriptContext) (*v8.Value, error) {
			return ctx.GetInstanceForNodeByName("NamedNodeMap", element.GetAttributes())
		},
	)
	helper.CreateFunction("setAttribute",
		func(instance Element, info argumentHelper) (result *v8.Value, err error) {
			name, err0 := info.getStringArg(0)
			value, err1 := info.getStringArg(1)
			if err = errors.Join(err0, err1); err == nil {
				instance.SetAttribute(name, value)
			}
			return
		},
	)
	helper.CreateFunction(
		"insertAdjacentHTML",
		func(element Element, info argumentHelper) (val *v8.Value, err error) {
			position, e1 := info.getStringArg(0)
			html, e2 := info.getStringArg(1)
			err = errors.Join(e1, e2)
			if err == nil {
				element.InsertAdjacentHTML(position, html)
				val, err = v8.NewValue(iso, element.OuterHTML())
			}
			return
		},
	)
	// helper.CreateFunction(
	// 	"querySelector",
	// 	func(instance Element, info argumentHelper) (*v8.Value, error) {
	// 		selector, e1 := info.GetStringArg(0)
	// 		node, e2 := instance.QuerySelector(selector)
	// 		err := errors.Join(e1, e2)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if node == nil {
	// 			return v8.Null(iso), nil
	// 		}
	// 		return info.ctx.GetInstanceForNode(node)
	// 	})
	// helper.CreateFunction(
	// 	"querySelectorAll",
	// 	func(instance Element, info argumentHelper) (*v8.Value, error) {
	// 		selector, e1 := info.GetStringArg(0)
	// 		nodeList, e2 := instance.QuerySelectorAll(selector)
	// 		err := errors.Join(e1, e2)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return info.ctx.GetInstanceForNodeByName("NodeList", nodeList)
	// 	},
	// )
	return builder.constructor
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
