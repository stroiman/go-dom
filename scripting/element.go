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

type argumentHelper struct {
	*v8.FunctionCallbackInfo
	ctx *ScriptContext
}

func newArgumentHelper(host *ScriptHost, info *v8.FunctionCallbackInfo) argumentHelper {
	ctx := host.MustGetContext(info.Context())
	return argumentHelper{info, ctx}
}

func (h argumentHelper) GetFunctionArg(index int) (*v8.Function, error) {
	args := h.Args()
	if index >= len(args) {
		return nil, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsFunction() {
		return arg.AsFunction()
	}
	return nil, ErrIncompatibleType
}

func (h argumentHelper) GetInt32Arg(index int) (int32, error) {
	args := h.Args()
	if index >= len(args) {
		return 0, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsNumber() {
		return arg.Int32(), nil
	}
	return 0, ErrIncompatibleType
}

func (h argumentHelper) GetStringArg(index int) (string, error) {
	args := h.Args()
	if index >= len(args) {
		return "", ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsString() {
		return arg.String(), nil
	}
	return "", ErrIncompatibleType
}

func (h argumentHelper) GetNodeArg(index int) (Node, error) {
	args := h.Args()
	if index >= len(args) {
		return nil, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsObject() {
		o := arg.Object()
		cached, ok_1 := h.ctx.GetCachedNode(o)
		if node, ok_2 := cached.(Node); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(h.ctx.host.iso, "Must be a node")
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
