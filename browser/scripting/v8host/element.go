package v8host

import (
	"errors"

	"github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

type esElement struct {
	esElementContainerWrapper[Element]
}

type elementV8Wrapper struct {
	esElementContainerWrapper[dom.Element]
}

func newElementV8Wrapper(host *V8ScriptHost) *elementV8Wrapper {
	return &elementV8Wrapper{newESContainerWrapper[dom.Element](host)}
}

func (e *elementV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	iso := e.host.iso
	e.Install(constructor)
	prototype := constructor.PrototypeTemplate()
	prototype.Set(
		"insertAdjacentHTML",
		v8.NewFunctionTemplateWithError(iso, e.insertAdjacentHTML),
	)
	prototype.SetAccessorProperty(
		"outerHTML",
		v8.NewFunctionTemplateWithError(iso, e.outerHTML),
		nil,
		v8.None,
	)
	prototype.SetAccessorProperty(
		"textContent",
		nil,
		v8.NewFunctionTemplateWithError(iso, e.setTextContent),
		v8.None,
	)
}

func (e *elementV8Wrapper) insertAdjacentHTML(
	info *v8.FunctionCallbackInfo,
) (val *v8.Value, err error) {
	iso := e.host.iso
	arg := newArgumentHelper(e.host, info)
	element, e0 := e.getInstance(info)
	position, e1 := arg.getStringArg(0)
	html, e2 := arg.getStringArg(1)
	err = errors.Join(e0, e1, e2)
	if err == nil {
		element.InsertAdjacentHTML(position, html)
		val, err = v8.NewValue(iso, element.OuterHTML())
	}
	return
}

func (e *elementV8Wrapper) outerHTML(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	if i, err := e.getInstance(info); err == nil {
		return v8.NewValue(e.host.iso, i.OuterHTML())
	} else {
		return nil, err
	}
}

func (w *elementV8Wrapper) setTextContent(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	e, err := w.getInstance(info)
	if err == nil {
		e.SetTextContent(info.Args()[0].String())
	}
	return nil, err
}

func (e elementV8Wrapper) ClassList(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	tokenList := e.host.globals.namedGlobals["DOMTokenList"]
	ctx := e.host.mustGetContext(info.Context())
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

func (e *elementV8Wrapper) toNamedNodeMap(
	ctx *V8ScriptContext,
	n dom.NamedNodeMap,
) (*v8.Value, error) {
	return ctx.getInstanceForNodeByName("NamedNodeMap", n)
}

func (w elementV8Wrapper) getAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	helper := newArgumentHelper(w.host, info)
	element, e0 := w.getInstance(info)
	name, e1 := helper.getStringArg(0)
	err := errors.Join(e0, e1)
	if err != nil {
		return nil, err
	}
	if r, ok := element.GetAttribute(name); ok {
		return v8.NewValue(w.host.iso, r)
	} else {
		return v8.Null(w.host.iso), nil
	}
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
