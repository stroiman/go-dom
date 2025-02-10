// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
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

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (e hTMLInputElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := e.scriptHost.iso
	prototypeTmpl.Set("checkValidity", v8.NewFunctionTemplateWithError(iso, e.checkValidity))

	prototypeTmpl.SetAccessorProperty("type",
		v8.NewFunctionTemplateWithError(iso, e.type_),
		v8.NewFunctionTemplateWithError(iso, e.setType),
		v8.None)
}

func (e hTMLInputElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e hTMLInputElementV8Wrapper) checkValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLInputElement.checkValidity")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return e.toBoolean(ctx, result)
}

func (e hTMLInputElementV8Wrapper) type_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLInputElement.type")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return e.toDOMString(ctx, result)
}

func (e hTMLInputElementV8Wrapper) setType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLInputElement.setType")
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
