package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type ESDOMTokenList struct {
	ESWrapper[dom.Element]
	Iterator Iterator[string]
}

func NewESDOMTokenList(host *ScriptHost) ESDOMTokenList {
	return ESDOMTokenList{
		NewESWrapper[dom.Element](host),
		NewIterator(host, func(item string, ctx *ScriptContext) (*v8.Value, error) {
			return v8.NewValue(host.iso, item)
		}),
	}
}

func (l ESDOMTokenList) GetInstance(
	info *v8.FunctionCallbackInfo,
) (result dom.DOMTokenList, err error) {
	element, err := l.ESWrapper.GetInstance(info)
	if err == nil {
		result = dom.NewClassList(element)
	}
	return
}

func (l ESDOMTokenList) CustomInitialiser(constructor *v8.FunctionTemplate) {
	constructor.InstanceTemplate().SetSymbol(
		v8.SymbolIterator(l.host.iso),
		v8.NewFunctionTemplateWithError(l.host.iso, l.GetIterator),
	)
}

func (l ESDOMTokenList) GetIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := l.host.MustGetContext(info.Context())
	instance, err := l.GetInstance(info)
	if err != nil {
		return nil, err
	}
	return l.Iterator.NewIteratorInstanceOfIterable(ctx, instance)
}

func (l ESDOMTokenList) Toggle(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(l.host, info)
	token, err0 := TryParseArg(args, 0, l.DecodeUSVString)
	force, err1 := TryParseArg(args, 1, l.DecodeBoolean)
	instance, errInstance := l.GetInstance(info)
	if args.noOfReadArguments >= 2 {
		if err := errors.Join(err0, err1, errInstance); err != nil {
			return nil, err
		}
		if force {
			instance.Add(token)
			return v8.NewValue(l.host.iso, true)
		} else {
			instance.Remove(token)
			return v8.NewValue(l.host.iso, false)
		}
	}
	if err := errors.Join(err0, errInstance); err != nil {
		return nil, err
	}
	return v8.NewValue(l.host.iso, instance.Toggle(token))
}

type ESHTMLTemplateElement struct {
	ESWrapper[html.HTMLTemplateElement]
}

func NewESHTMLTemplateElement(host *ScriptHost) ESHTMLTemplateElement {
	return ESHTMLTemplateElement{NewESWrapper[html.HTMLTemplateElement](host)}
}

func (e ESHTMLTemplateElement) CreateInstance(
	ctx *ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}

func (e ESHTMLTemplateElement) ToDocumentFragment(
	ctx *ScriptContext,
	fragment dom.DocumentFragment,
) (*v8.Value, error) {
	return ctx.GetInstanceForNode(fragment)
}
