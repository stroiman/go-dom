// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	v8 "github.com/tommie/v8go"
)

type hTMLAnchorElementV8Wrapper struct {
	nodeV8WrapperBase[html.HTMLAnchorElement]
}

func newHTMLAnchorElementV8Wrapper(scriptHost *V8ScriptHost) *hTMLAnchorElementV8Wrapper {
	return &hTMLAnchorElementV8Wrapper{newNodeV8WrapperBase[html.HTMLAnchorElement](scriptHost)}
}

func init() {
	registerJSClass("HTMLAnchorElement", "HTMLElement", createHTMLAnchorElementPrototype)
}

func createHTMLAnchorElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLAnchorElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()

	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, wrapper.target),
		v8.NewFunctionTemplateWithError(iso, wrapper.setTarget),
		v8.None)

	return constructor
}

func (e hTMLAnchorElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e hTMLAnchorElementV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return e.toDOMString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetTarget(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setTarget: Missing arguments")
}
