package scripting

import (
	"github.com/stroiman/go-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type ESHTMLDocumentWrapper struct {
	DocumentV8Wrapper
}

func NewHTMLDocumentWrapper(host *ScriptHost) ESHTMLDocumentWrapper {
	return ESHTMLDocumentWrapper{NewDocumentV8Wrapper(host)}
}

func CreateHTMLDocumentPrototype(host *ScriptHost) *v8.FunctionTemplate {
	wrapper := NewDocumentV8Wrapper(host)
	builder := NewIllegalConstructorBuilder[html.HTMLDocument](host)
	constructor := builder.constructor
	instanceTemplate := constructor.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	wrapper.BuildInstanceTemplate(constructor)
	return constructor
}
