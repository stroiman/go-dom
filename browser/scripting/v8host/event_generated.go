// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type eventV8Wrapper struct {
	nodeV8WrapperBase[dom.Event]
}

func newEventV8Wrapper(scriptHost *V8ScriptHost) *eventV8Wrapper {
	return &eventV8Wrapper{newNodeV8WrapperBase[dom.Event](scriptHost)}
}

func init() {
	registerJSClass("Event", "", createEventPrototype)
}

func createEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newEventV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("stopPropagation", v8.NewFunctionTemplateWithError(iso, wrapper.stopPropagation))
	prototypeTmpl.Set("preventDefault", v8.NewFunctionTemplateWithError(iso, wrapper.preventDefault))

	prototypeTmpl.SetAccessorProperty("type",
		v8.NewFunctionTemplateWithError(iso, wrapper.type_),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, wrapper.target),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("currentTarget",
		v8.NewFunctionTemplateWithError(iso, wrapper.currentTarget),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("bubbles",
		v8.NewFunctionTemplateWithError(iso, wrapper.bubbles),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("cancelable",
		v8.NewFunctionTemplateWithError(iso, wrapper.cancelable),
		nil,
		v8.None)

	return constructor
}

func (e eventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(e.scriptHost, info)
	type_, err1 := tryParseArg(args, 0, e.decodeDOMString)
	eventInitDict, err2 := tryParseArgWithDefault(args, 1, e.defaultEventInit, e.decodeEventInit)
	ctx := e.mustGetContext(info)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return e.CreateInstance(ctx, info.This(), type_, eventInitDict)
	}
	return nil, errors.New("Missing arguments")
}

func (e eventV8Wrapper) stopPropagation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func (e eventV8Wrapper) preventDefault(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func (e eventV8Wrapper) type_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return e.toDOMString(ctx, result)
}

func (e eventV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return e.toNullableEventTarget(ctx, result)
}

func (e eventV8Wrapper) currentTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget()
	return e.toNullableEventTarget(ctx, result)
}

func (e eventV8Wrapper) bubbles(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles()
	return e.toBoolean(ctx, result)
}

func (e eventV8Wrapper) cancelable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable()
	return e.toBoolean(ctx, result)
}
