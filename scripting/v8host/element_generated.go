// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("Element", "Node", createElementPrototype)
}

func createElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (e elementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := e.scriptHost.iso
	prototypeTmpl.Set("hasAttributes", v8.NewFunctionTemplateWithError(iso, e.hasAttributes))
	prototypeTmpl.Set("getAttributeNames", v8.NewFunctionTemplateWithError(iso, e.getAttributeNames))
	prototypeTmpl.Set("getAttribute", v8.NewFunctionTemplateWithError(iso, e.getAttribute))
	prototypeTmpl.Set("getAttributeNS", v8.NewFunctionTemplateWithError(iso, e.getAttributeNS))
	prototypeTmpl.Set("setAttribute", v8.NewFunctionTemplateWithError(iso, e.setAttribute))
	prototypeTmpl.Set("setAttributeNS", v8.NewFunctionTemplateWithError(iso, e.setAttributeNS))
	prototypeTmpl.Set("removeAttribute", v8.NewFunctionTemplateWithError(iso, e.removeAttribute))
	prototypeTmpl.Set("removeAttributeNS", v8.NewFunctionTemplateWithError(iso, e.removeAttributeNS))
	prototypeTmpl.Set("toggleAttribute", v8.NewFunctionTemplateWithError(iso, e.toggleAttribute))
	prototypeTmpl.Set("hasAttribute", v8.NewFunctionTemplateWithError(iso, e.hasAttribute))
	prototypeTmpl.Set("hasAttributeNS", v8.NewFunctionTemplateWithError(iso, e.hasAttributeNS))
	prototypeTmpl.Set("getAttributeNode", v8.NewFunctionTemplateWithError(iso, e.getAttributeNode))
	prototypeTmpl.Set("getAttributeNodeNS", v8.NewFunctionTemplateWithError(iso, e.getAttributeNodeNS))
	prototypeTmpl.Set("setAttributeNode", v8.NewFunctionTemplateWithError(iso, e.setAttributeNode))
	prototypeTmpl.Set("setAttributeNodeNS", v8.NewFunctionTemplateWithError(iso, e.setAttributeNodeNS))
	prototypeTmpl.Set("removeAttributeNode", v8.NewFunctionTemplateWithError(iso, e.removeAttributeNode))
	prototypeTmpl.Set("attachShadow", v8.NewFunctionTemplateWithError(iso, e.attachShadow))
	prototypeTmpl.Set("matches", v8.NewFunctionTemplateWithError(iso, e.matches))
	prototypeTmpl.Set("getElementsByTagName", v8.NewFunctionTemplateWithError(iso, e.getElementsByTagName))
	prototypeTmpl.Set("getElementsByTagNameNS", v8.NewFunctionTemplateWithError(iso, e.getElementsByTagNameNS))
	prototypeTmpl.Set("getElementsByClassName", v8.NewFunctionTemplateWithError(iso, e.getElementsByClassName))
	prototypeTmpl.Set("insertAdjacentElement", v8.NewFunctionTemplateWithError(iso, e.insertAdjacentElement))
	prototypeTmpl.Set("insertAdjacentText", v8.NewFunctionTemplateWithError(iso, e.insertAdjacentText))

	prototypeTmpl.SetAccessorProperty("namespaceURI",
		v8.NewFunctionTemplateWithError(iso, e.namespaceURI),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("prefix",
		v8.NewFunctionTemplateWithError(iso, e.prefix),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("localName",
		v8.NewFunctionTemplateWithError(iso, e.localName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tagName",
		v8.NewFunctionTemplateWithError(iso, e.tagName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("id",
		v8.NewFunctionTemplateWithError(iso, e.id),
		v8.NewFunctionTemplateWithError(iso, e.setId),
		v8.None)
	prototypeTmpl.SetAccessorProperty("className",
		v8.NewFunctionTemplateWithError(iso, e.className),
		v8.NewFunctionTemplateWithError(iso, e.setClassName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("classList",
		v8.NewFunctionTemplateWithError(iso, e.classList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("slot",
		v8.NewFunctionTemplateWithError(iso, e.slot),
		v8.NewFunctionTemplateWithError(iso, e.setSlot),
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributes",
		v8.NewFunctionTemplateWithError(iso, e.attributes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRoot",
		v8.NewFunctionTemplateWithError(iso, e.shadowRoot),
		nil,
		v8.None)
}

func (e elementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e elementV8Wrapper) hasAttributes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.hasAttributes")
	return nil, errors.New("Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) getAttributeNames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getAttributeNames")
	return nil, errors.New("Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) getAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getAttributeNS")
	return nil, errors.New("Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) setAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setAttribute")
	args := newArgumentHelper(e.scriptHost, info)
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
	return nil, errors.New("Element.setAttribute: Missing arguments")
}

func (e elementV8Wrapper) setAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setAttributeNS")
	return nil, errors.New("Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) removeAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.removeAttribute")
	return nil, errors.New("Element.removeAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) removeAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.removeAttributeNS")
	return nil, errors.New("Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) toggleAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.toggleAttribute")
	return nil, errors.New("Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) hasAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: Element.hasAttribute")
	args := newArgumentHelper(e.scriptHost, info)
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
	return nil, errors.New("Element.hasAttribute: Missing arguments")
}

func (e elementV8Wrapper) hasAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.hasAttributeNS")
	return nil, errors.New("Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) getAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getAttributeNode")
	return nil, errors.New("Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) getAttributeNodeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getAttributeNodeNS")
	return nil, errors.New("Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) setAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setAttributeNode")
	return nil, errors.New("Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) setAttributeNodeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setAttributeNodeNS")
	return nil, errors.New("Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) removeAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.removeAttributeNode")
	return nil, errors.New("Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) attachShadow(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.attachShadow")
	return nil, errors.New("Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) matches(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: Element.matches")
	args := newArgumentHelper(e.scriptHost, info)
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
	return nil, errors.New("Element.matches: Missing arguments")
}

func (e elementV8Wrapper) getElementsByTagName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getElementsByTagName")
	return nil, errors.New("Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) getElementsByTagNameNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getElementsByTagNameNS")
	return nil, errors.New("Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) getElementsByClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.getElementsByClassName")
	return nil, errors.New("Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) insertAdjacentElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.insertAdjacentElement")
	return nil, errors.New("Element.insertAdjacentElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) insertAdjacentText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.insertAdjacentText")
	return nil, errors.New("Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) namespaceURI(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.namespaceURI")
	return nil, errors.New("Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) prefix(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.prefix")
	return nil, errors.New("Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) localName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.localName")
	return nil, errors.New("Element.localName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) tagName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: Element.tagName")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return e.toDOMString(ctx, result)
}

func (e elementV8Wrapper) id(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.id")
	return nil, errors.New("Element.id: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) setId(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setId")
	return nil, errors.New("Element.setId: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) className(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.className")
	return nil, errors.New("Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) setClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setClassName")
	return nil, errors.New("Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) slot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.slot")
	return nil, errors.New("Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) setSlot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.setSlot")
	return nil, errors.New("Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (e elementV8Wrapper) attributes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: Element.attributes")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return e.toNamedNodeMap(ctx, result)
}

func (e elementV8Wrapper) shadowRoot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Element.shadowRoot")
	return nil, errors.New("Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
