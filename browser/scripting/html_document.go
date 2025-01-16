package scripting

import (
	"github.com/stroiman/go-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type ESHTMLDocumentWrapper struct {
	documentV8Wrapper
}

func NewHTMLDocumentWrapper(host *V8ScriptHost) ESHTMLDocumentWrapper {
	return ESHTMLDocumentWrapper{newDocumentV8Wrapper(host)}
}

func createHTMLDocumentPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newDocumentV8Wrapper(host)
	builder := NewIllegalConstructorBuilder[html.HTMLDocument](host)
	constructor := builder.constructor
	instanceTemplate := constructor.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	wrapper.BuildInstanceTemplate(constructor)
	return constructor
}
