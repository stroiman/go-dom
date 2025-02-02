// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	v8 "github.com/tommie/v8go"
)

type hTMLInputElementV8Wrapper struct {
	nodeV8WrapperBase[html.HTMLInputElement]
}

func newHTMLInputElementV8Wrapper(scriptHost *V8ScriptHost) *hTMLInputElementV8Wrapper {
	return &hTMLInputElementV8Wrapper{newNodeV8WrapperBase[html.HTMLInputElement](scriptHost)}
}

func init() {
	registerJSClass("HTMLInputElement", "HTMLElement", createHTMLInputElementPrototype)
}

func createHTMLInputElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLInputElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("checkValidity", v8.NewFunctionTemplateWithError(iso, wrapper.checkValidity))

	prototypeTmpl.SetAccessorProperty("type",
		v8.NewFunctionTemplateWithError(iso, wrapper.type_),
		v8.NewFunctionTemplateWithError(iso, wrapper.setType),
		v8.None)

	return constructor
}

func (e hTMLInputElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e hTMLInputElementV8Wrapper) checkValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return e.toBoolean(ctx, result)
}

func (e hTMLInputElementV8Wrapper) type_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return e.toDOMString(ctx, result)
}

func (e hTMLInputElementV8Wrapper) setType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetType(val)
		return nil, nil
	}
	return nil, errors.New("HTMLInputElement.setType: Missing arguments")
}
