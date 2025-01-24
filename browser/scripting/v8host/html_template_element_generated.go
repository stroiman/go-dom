// This file is generated. Do not edit.

package v8host

import (
	"errors"
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

	prototypeTmpl := constructor.PrototypeTemplate()

	prototypeTmpl.SetAccessorProperty("content",
		v8.NewFunctionTemplateWithError(iso, wrapper.content),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootMode",
		v8.NewFunctionTemplateWithError(iso, wrapper.shadowRootMode),
		v8.NewFunctionTemplateWithError(iso, wrapper.setShadowRootMode),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootDelegatesFocus",
		v8.NewFunctionTemplateWithError(iso, wrapper.shadowRootDelegatesFocus),
		v8.NewFunctionTemplateWithError(iso, wrapper.setShadowRootDelegatesFocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootClonable",
		v8.NewFunctionTemplateWithError(iso, wrapper.shadowRootClonable),
		v8.NewFunctionTemplateWithError(iso, wrapper.setShadowRootClonable),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootSerializable",
		v8.NewFunctionTemplateWithError(iso, wrapper.shadowRootSerializable),
		v8.NewFunctionTemplateWithError(iso, wrapper.setShadowRootSerializable),
		v8.None)

	return constructor
}

func (e htmlTemplateElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e htmlTemplateElementV8Wrapper) content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return ctx.getInstanceForNode(result)
}

func (e htmlTemplateElementV8Wrapper) shadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) setShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) shadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) setShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) shadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) setShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) shadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e htmlTemplateElementV8Wrapper) setShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}
