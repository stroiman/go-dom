// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

type hTMLFormElementV8Wrapper struct {
	nodeV8WrapperBase[html.HTMLFormElement]
}

func newHTMLFormElementV8Wrapper(scriptHost *V8ScriptHost) *hTMLFormElementV8Wrapper {
	return &hTMLFormElementV8Wrapper{newNodeV8WrapperBase[html.HTMLFormElement](scriptHost)}
}

func init() {
	registerJSClass("HTMLFormElement", "HTMLElement", createHTMLFormElementPrototype)
}

func createHTMLFormElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLFormElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (e hTMLFormElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := e.scriptHost.iso
	prototypeTmpl.Set("submit", v8.NewFunctionTemplateWithError(iso, e.submit))
	prototypeTmpl.Set("requestSubmit", v8.NewFunctionTemplateWithError(iso, e.requestSubmit))
	prototypeTmpl.Set("reset", v8.NewFunctionTemplateWithError(iso, e.reset))
	prototypeTmpl.Set("checkValidity", v8.NewFunctionTemplateWithError(iso, e.checkValidity))
	prototypeTmpl.Set("reportValidity", v8.NewFunctionTemplateWithError(iso, e.reportValidity))

	prototypeTmpl.SetAccessorProperty("acceptCharset",
		v8.NewFunctionTemplateWithError(iso, e.acceptCharset),
		v8.NewFunctionTemplateWithError(iso, e.setAcceptCharset),
		v8.None)
	prototypeTmpl.SetAccessorProperty("action",
		v8.NewFunctionTemplateWithError(iso, e.action),
		v8.NewFunctionTemplateWithError(iso, e.setAction),
		v8.None)
	prototypeTmpl.SetAccessorProperty("autocomplete",
		v8.NewFunctionTemplateWithError(iso, e.autocomplete),
		v8.NewFunctionTemplateWithError(iso, e.setAutocomplete),
		v8.None)
	prototypeTmpl.SetAccessorProperty("enctype",
		v8.NewFunctionTemplateWithError(iso, e.enctype),
		v8.NewFunctionTemplateWithError(iso, e.setEnctype),
		v8.None)
	prototypeTmpl.SetAccessorProperty("encoding",
		v8.NewFunctionTemplateWithError(iso, e.encoding),
		v8.NewFunctionTemplateWithError(iso, e.setEncoding),
		v8.None)
	prototypeTmpl.SetAccessorProperty("method",
		v8.NewFunctionTemplateWithError(iso, e.method),
		v8.NewFunctionTemplateWithError(iso, e.setMethod),
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, e.target),
		v8.NewFunctionTemplateWithError(iso, e.setTarget),
		v8.None)
	prototypeTmpl.SetAccessorProperty("rel",
		v8.NewFunctionTemplateWithError(iso, e.rel),
		v8.NewFunctionTemplateWithError(iso, e.setRel),
		v8.None)
	prototypeTmpl.SetAccessorProperty("relList",
		v8.NewFunctionTemplateWithError(iso, e.relList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("elements",
		v8.NewFunctionTemplateWithError(iso, e.elements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, e.length),
		nil,
		v8.None)
}

func (e hTMLFormElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e hTMLFormElementV8Wrapper) submit(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.submit")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Submit()
	return nil, callErr
}

func (e hTMLFormElementV8Wrapper) requestSubmit(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.requestSubmit")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	submitter, err1 := tryParseArgWithDefault(args, 0, e.defaultHTMLElement, e.decodeHTMLElement)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.RequestSubmit(submitter)
		return nil, callErr
	}
	return nil, errors.New("HTMLFormElement.requestSubmit: Missing arguments")
}

func (e hTMLFormElementV8Wrapper) reset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.reset")
	return nil, errors.New("HTMLFormElement.reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) checkValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.checkValidity")
	return nil, errors.New("HTMLFormElement.checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) reportValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.reportValidity")
	return nil, errors.New("HTMLFormElement.reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) acceptCharset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.acceptCharset")
	return nil, errors.New("HTMLFormElement.acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) setAcceptCharset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setAcceptCharset")
	return nil, errors.New("HTMLFormElement.setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) action(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLFormElement.action")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Action()
	return e.toUSVString(ctx, result)
}

func (e hTMLFormElementV8Wrapper) setAction(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setAction")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetAction(val)
		return nil, nil
	}
	return nil, errors.New("HTMLFormElement.setAction: Missing arguments")
}

func (e hTMLFormElementV8Wrapper) autocomplete(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.autocomplete")
	return nil, errors.New("HTMLFormElement.autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) setAutocomplete(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setAutocomplete")
	return nil, errors.New("HTMLFormElement.setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) enctype(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.enctype")
	return nil, errors.New("HTMLFormElement.enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) setEnctype(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setEnctype")
	return nil, errors.New("HTMLFormElement.setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) encoding(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.encoding")
	return nil, errors.New("HTMLFormElement.encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) setEncoding(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setEncoding")
	return nil, errors.New("HTMLFormElement.setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) method(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLFormElement.method")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Method()
	return e.toDOMString(ctx, result)
}

func (e hTMLFormElementV8Wrapper) setMethod(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setMethod")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetMethod(val)
		return nil, nil
	}
	return nil, errors.New("HTMLFormElement.setMethod: Missing arguments")
}

func (e hTMLFormElementV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.target")
	return nil, errors.New("HTMLFormElement.target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) setTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setTarget")
	return nil, errors.New("HTMLFormElement.setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) rel(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.rel")
	return nil, errors.New("HTMLFormElement.rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) setRel(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.setRel")
	return nil, errors.New("HTMLFormElement.setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) relList(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.relList")
	return nil, errors.New("HTMLFormElement.relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e hTMLFormElementV8Wrapper) elements(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLFormElement.elements")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Elements()
	return e.toHTMLFormControlsCollection(ctx, result)
}

func (e hTMLFormElementV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLFormElement.length")
	return nil, errors.New("HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
