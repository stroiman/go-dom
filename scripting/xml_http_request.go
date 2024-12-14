package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type JSWrapper[T Entity] struct {
	host *ScriptHost
}

type JSXmlHttpRequest JSWrapper[XmlHttpRequest]

func NewJSXmlHttpRequest(host *ScriptHost) JSXmlHttpRequest {
	return JSXmlHttpRequest{host}
}

func (w JSWrapper[T]) GetInstance(info *v8.FunctionCallbackInfo) (result T, err error) {
	if ctx, ok := w.host.GetContext(info.Context()); ok {
		if instance, ok := ctx.GetCachedNode(info.This()); ok {
			if typedInstance, ok := instance.(T); ok {
				return typedInstance, nil
			}
		}
		err = v8.NewTypeError(ctx.host.iso, "Not an instance of NamedNodeMap")
		return
	}
	err = errors.New("Could not get context")
	return
}

func (w JSXmlHttpRequest) GetInstance(info *v8.FunctionCallbackInfo) (XmlHttpRequest, error) {
	return (JSWrapper[XmlHttpRequest](w)).GetInstance(info)
}
