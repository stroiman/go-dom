// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateHTMLTemplateElementPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESHTMLTemplateElement(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.NewInstance)

	instanceTmpl := constructor.GetInstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()

	prototypeTmpl.SetAccessorProperty("content",
		v8.NewFunctionTemplateWithError(iso, wrapper.Content),
		nil,
		v8.ReadOnly)
	prototypeTmpl.SetAccessorProperty("shadowRootMode",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootMode),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootMode),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootDelegatesFocus",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootDelegatesFocus),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootDelegatesFocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootClonable",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootClonable),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootClonable),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootSerializable",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootSerializable),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootSerializable),
		v8.None)

	return constructor
}

func (e ESHTMLTemplateElement) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.MustGetContext(info.Context())
	return e.CreateInstance(ctx, info.This())
}

func (e ESHTMLTemplateElement) Content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.MustGetContext(info.Context())
	instance, err := e.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return e.ToDocumentFragment(ctx, result)
}

func (e ESHTMLTemplateElement) GetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootMode")
}

func (e ESHTMLTemplateElement) SetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootMode")
}

func (e ESHTMLTemplateElement) GetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootDelegatesFocus")
}

func (e ESHTMLTemplateElement) SetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootDelegatesFocus")
}

func (e ESHTMLTemplateElement) GetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootClonable")
}

func (e ESHTMLTemplateElement) SetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootClonable")
}

func (e ESHTMLTemplateElement) GetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootSerializable")
}

func (e ESHTMLTemplateElement) SetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootSerializable")
}
