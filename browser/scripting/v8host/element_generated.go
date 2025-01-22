// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("Element", "Node", createElementPrototype)
}

func createElementPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newElementV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("hasAttributes", v8.NewFunctionTemplateWithError(iso, wrapper.HasAttributes))
	prototypeTmpl.Set("getAttributeNames", v8.NewFunctionTemplateWithError(iso, wrapper.GetAttributeNames))
	prototypeTmpl.Set("getAttribute", v8.NewFunctionTemplateWithError(iso, wrapper.GetAttribute))
	prototypeTmpl.Set("getAttributeNS", v8.NewFunctionTemplateWithError(iso, wrapper.GetAttributeNS))
	prototypeTmpl.Set("setAttribute", v8.NewFunctionTemplateWithError(iso, wrapper.SetAttribute))
	prototypeTmpl.Set("setAttributeNS", v8.NewFunctionTemplateWithError(iso, wrapper.SetAttributeNS))
	prototypeTmpl.Set("removeAttribute", v8.NewFunctionTemplateWithError(iso, wrapper.RemoveAttribute))
	prototypeTmpl.Set("removeAttributeNS", v8.NewFunctionTemplateWithError(iso, wrapper.RemoveAttributeNS))
	prototypeTmpl.Set("toggleAttribute", v8.NewFunctionTemplateWithError(iso, wrapper.ToggleAttribute))
	prototypeTmpl.Set("hasAttribute", v8.NewFunctionTemplateWithError(iso, wrapper.HasAttribute))
	prototypeTmpl.Set("hasAttributeNS", v8.NewFunctionTemplateWithError(iso, wrapper.HasAttributeNS))
	prototypeTmpl.Set("getAttributeNode", v8.NewFunctionTemplateWithError(iso, wrapper.GetAttributeNode))
	prototypeTmpl.Set("getAttributeNodeNS", v8.NewFunctionTemplateWithError(iso, wrapper.GetAttributeNodeNS))
	prototypeTmpl.Set("setAttributeNode", v8.NewFunctionTemplateWithError(iso, wrapper.SetAttributeNode))
	prototypeTmpl.Set("setAttributeNodeNS", v8.NewFunctionTemplateWithError(iso, wrapper.SetAttributeNodeNS))
	prototypeTmpl.Set("removeAttributeNode", v8.NewFunctionTemplateWithError(iso, wrapper.RemoveAttributeNode))
	prototypeTmpl.Set("attachShadow", v8.NewFunctionTemplateWithError(iso, wrapper.AttachShadow))
	prototypeTmpl.Set("matches", v8.NewFunctionTemplateWithError(iso, wrapper.Matches))
	prototypeTmpl.Set("getElementsByTagName", v8.NewFunctionTemplateWithError(iso, wrapper.GetElementsByTagName))
	prototypeTmpl.Set("getElementsByTagNameNS", v8.NewFunctionTemplateWithError(iso, wrapper.GetElementsByTagNameNS))
	prototypeTmpl.Set("getElementsByClassName", v8.NewFunctionTemplateWithError(iso, wrapper.GetElementsByClassName))
	prototypeTmpl.Set("insertAdjacentElement", v8.NewFunctionTemplateWithError(iso, wrapper.InsertAdjacentElement))
	prototypeTmpl.Set("insertAdjacentText", v8.NewFunctionTemplateWithError(iso, wrapper.InsertAdjacentText))

	prototypeTmpl.SetAccessorProperty("namespaceURI",
		v8.NewFunctionTemplateWithError(iso, wrapper.NamespaceURI),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("prefix",
		v8.NewFunctionTemplateWithError(iso, wrapper.Prefix),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("localName",
		v8.NewFunctionTemplateWithError(iso, wrapper.LocalName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tagName",
		v8.NewFunctionTemplateWithError(iso, wrapper.TagName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("id",
		v8.NewFunctionTemplateWithError(iso, wrapper.Id),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetId),
		v8.None)
	prototypeTmpl.SetAccessorProperty("className",
		v8.NewFunctionTemplateWithError(iso, wrapper.ClassName),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetClassName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("classList",
		v8.NewFunctionTemplateWithError(iso, wrapper.ClassList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("slot",
		v8.NewFunctionTemplateWithError(iso, wrapper.Slot),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetSlot),
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributes",
		v8.NewFunctionTemplateWithError(iso, wrapper.Attributes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRoot",
		v8.NewFunctionTemplateWithError(iso, wrapper.ShadowRoot),
		nil,
		v8.None)

	wrapper.CustomInitialiser(constructor)
	return constructor
}

func (e elementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.host.iso, "Illegal Constructor")
}

func (e elementV8Wrapper) HasAttributes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.hasAttributes: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) GetAttributeNames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getAttributeNames: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) GetAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getAttributeNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) SetAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(e.host, info)
	instance, err0 := e.getInstance(info)
	qualifiedName, err1 := tryParseArg(args, 0, e.decodeDOMString)
	value, err2 := tryParseArg(args, 1, e.decodeDOMString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.SetAttribute(qualifiedName, value)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (e elementV8Wrapper) SetAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.setAttributeNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) RemoveAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.removeAttribute: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) RemoveAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) ToggleAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.toggleAttribute: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) HasAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.mustGetContext(info.Context())
	args := newArgumentHelper(e.host, info)
	instance, err0 := e.getInstance(info)
	qualifiedName, err1 := tryParseArg(args, 0, e.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.HasAttribute(qualifiedName)
		return e.toBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (e elementV8Wrapper) HasAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) GetAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getAttributeNode: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) GetAttributeNodeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) SetAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.setAttributeNode: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) SetAttributeNodeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) RemoveAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) AttachShadow(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.attachShadow: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) Matches(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.mustGetContext(info.Context())
	args := newArgumentHelper(e.host, info)
	instance, err0 := e.getInstance(info)
	selectors, err1 := tryParseArg(args, 0, e.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.Matches(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return e.toBoolean(ctx, result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (e elementV8Wrapper) GetElementsByTagName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) GetElementsByTagNameNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) GetElementsByClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) InsertAdjacentElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.insertAdjacentElement: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) InsertAdjacentText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) NamespaceURI(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.NamespaceURI: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) Prefix(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.Prefix: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) LocalName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.LocalName: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) TagName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.mustGetContext(info.Context())
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return e.toDOMString(ctx, result)
}

func (e elementV8Wrapper) Id(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.Id: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) SetId(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.SetId: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) ClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.ClassName: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) SetClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.SetClassName: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) Slot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.Slot: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) SetSlot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.SetSlot: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (e elementV8Wrapper) Attributes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.mustGetContext(info.Context())
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return e.toNamedNodeMap(ctx, result)
}

func (e elementV8Wrapper) ShadowRoot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Element.ShadowRoot: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}
