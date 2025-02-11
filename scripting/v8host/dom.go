package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type domTokenListV8Wrapper struct {
	nodeV8WrapperBase[dom.Element]
	Iterator iterator[string]
}

func newDOMTokenListV8Wrapper(host *V8ScriptHost) domTokenListV8Wrapper {
	return domTokenListV8Wrapper{
		newNodeV8WrapperBase[dom.Element](host),
		newIterator(host, func(item string, ctx *V8ScriptContext) (*v8.Value, error) {
			return v8.NewValue(host.iso, item)
		}),
	}
}

func (l domTokenListV8Wrapper) getInstance(
	info *v8.FunctionCallbackInfo,
) (result dom.DOMTokenList, err error) {
	element, err := l.nodeV8WrapperBase.getInstance(info)
	if err == nil {
		result = dom.NewClassList(element)
	}
	return
}

func (l domTokenListV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	constructor.InstanceTemplate().SetSymbol(
		v8.SymbolIterator(l.scriptHost.iso),
		v8.NewFunctionTemplateWithError(l.scriptHost.iso, l.GetIterator),
	)
}

func (l domTokenListV8Wrapper) GetIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := l.scriptHost.mustGetContext(info.Context())
	instance, err := l.getInstance(info)
	if err != nil {
		return nil, err
	}
	return l.Iterator.newIteratorInstanceOfIterable(ctx, instance)
}

func (l domTokenListV8Wrapper) toggle(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(l.scriptHost, info)
	token, err0 := tryParseArg(args, 0, l.decodeUSVString)
	force, err1 := tryParseArg(args, 1, l.decodeBoolean)
	instance, errInstance := l.getInstance(info)
	if args.noOfReadArguments >= 2 {
		if err := errors.Join(err0, err1, errInstance); err != nil {
			return nil, err
		}
		if force {
			instance.Add(token)
			return v8.NewValue(l.scriptHost.iso, true)
		} else {
			instance.Remove(token)
			return v8.NewValue(l.scriptHost.iso, false)
		}
	}
	if err := errors.Join(err0, errInstance); err != nil {
		return nil, err
	}
	return v8.NewValue(l.scriptHost.iso, instance.Toggle(token))
}

type htmlTemplateElementV8Wrapper struct {
	nodeV8WrapperBase[html.HTMLTemplateElement]
}

func newHTMLTemplateElementV8Wrapper(host *V8ScriptHost) htmlTemplateElementV8Wrapper {
	return htmlTemplateElementV8Wrapper{newNodeV8WrapperBase[html.HTMLTemplateElement](host)}
}

func (e htmlTemplateElementV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}

func (e htmlTemplateElementV8Wrapper) ToDocumentFragment(
	ctx *V8ScriptContext,
	fragment dom.DocumentFragment,
) (*v8.Value, error) {
	return ctx.getInstanceForNode(fragment)
}
