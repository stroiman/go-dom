package scripting

import (
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

// CreateDocumentFragmentPrototype currently only exists to allow code to check
// for inheritence, i.e., `node instanceof DocumentFragment`
func CreateDocumentFragmentPrototype(host *ScriptHost) *v8.FunctionTemplate {
	builder := NewIllegalConstructorBuilder[Location](host)
	wrapper := NewESContainerWrapper[DocumentFragment](host)
	wrapper.Install(builder.constructor)
	return builder.constructor
}
