// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateHTMLTemplateElementPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewHTMLTemplateElementV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()

	prototypeTmpl.SetAccessorProperty("content",
		v8.NewFunctionTemplateWithError(iso, wrapper.Content),
		nil,
		v8.None)
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

func (e HTMLTemplateElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.host.iso, "Illegal Constructor")
}

func (e HTMLTemplateElementV8Wrapper) Content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.MustGetContext(info.Context())
	instance, err := e.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return ctx.GetInstanceForNode(result)
}

func (e HTMLTemplateElementV8Wrapper) GetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootMode")
}

func (e HTMLTemplateElementV8Wrapper) SetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootMode")
}

func (e HTMLTemplateElementV8Wrapper) GetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootDelegatesFocus")
}

func (e HTMLTemplateElementV8Wrapper) SetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootDelegatesFocus")
}

func (e HTMLTemplateElementV8Wrapper) GetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootClonable")
}

func (e HTMLTemplateElementV8Wrapper) SetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootClonable")
}

func (e HTMLTemplateElementV8Wrapper) GetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootSerializable")
}

func (e HTMLTemplateElementV8Wrapper) SetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootSerializable")
}
