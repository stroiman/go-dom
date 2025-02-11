// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("HTMLTemplateElement", "HTMLElement", createHtmlTemplateElementPrototype)
}

func createHtmlTemplateElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHtmlTemplateElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlTemplateElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("content",
		v8.NewFunctionTemplateWithError(iso, w.content),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootMode",
		v8.NewFunctionTemplateWithError(iso, w.shadowRootMode),
		v8.NewFunctionTemplateWithError(iso, w.setShadowRootMode),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootDelegatesFocus",
		v8.NewFunctionTemplateWithError(iso, w.shadowRootDelegatesFocus),
		v8.NewFunctionTemplateWithError(iso, w.setShadowRootDelegatesFocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootClonable",
		v8.NewFunctionTemplateWithError(iso, w.shadowRootClonable),
		v8.NewFunctionTemplateWithError(iso, w.setShadowRootClonable),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootSerializable",
		v8.NewFunctionTemplateWithError(iso, w.shadowRootSerializable),
		v8.NewFunctionTemplateWithError(iso, w.setShadowRootSerializable),
		v8.None)
}

func (w htmlTemplateElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w htmlTemplateElementV8Wrapper) content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLTemplateElement.content")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return ctx.getInstanceForNode(result)
}

func (w htmlTemplateElementV8Wrapper) shadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.shadowRootMode")
	return nil, errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.setShadowRootMode")
	return nil, errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) shadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.shadowRootDelegatesFocus")
	return nil, errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.setShadowRootDelegatesFocus")
	return nil, errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) shadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.shadowRootClonable")
	return nil, errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.setShadowRootClonable")
	return nil, errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) shadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.shadowRootSerializable")
	return nil, errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLTemplateElement.setShadowRootSerializable")
	return nil, errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
