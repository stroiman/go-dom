package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type domTokenListV8Wrapper struct {
	nodeV8WrapperBase[dom.Element]
	Iterator Iterator[string]
}

func newDomTokenListV8Wrapper(host *V8ScriptHost) domTokenListV8Wrapper {
	return domTokenListV8Wrapper{
		newNodeV8WrapperBase[dom.Element](host),
		NewIterator(host, func(item string, ctx *V8ScriptContext) (*v8.Value, error) {
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
		v8.SymbolIterator(l.host.iso),
		v8.NewFunctionTemplateWithError(l.host.iso, l.GetIterator),
	)
}

func (l domTokenListV8Wrapper) GetIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := l.host.MustGetContext(info.Context())
	instance, err := l.getInstance(info)
	if err != nil {
		return nil, err
	}
	return l.Iterator.NewIteratorInstanceOfIterable(ctx, instance)
}

func (l domTokenListV8Wrapper) Toggle(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(l.host, info)
	token, err0 := TryParseArg(args, 0, l.decodeUSVString)
	force, err1 := TryParseArg(args, 1, l.decodeBoolean)
	instance, errInstance := l.getInstance(info)
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

type htmlTemplateElementV8Wrapper struct {
	nodeV8WrapperBase[html.HTMLTemplateElement]
}

func newHtmlTemplateElementV8Wrapper(host *V8ScriptHost) htmlTemplateElementV8Wrapper {
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
