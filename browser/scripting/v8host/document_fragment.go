package v8host

import (
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

// createDocumentFragmentPrototype currently only exists to allow code to check
// for inheritence, i.e., `node instanceof DocumentFragment`
func createDocumentFragmentPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	builder := newIllegalConstructorBuilder[Location](host)
	wrapper := newESContainerWrapper[DocumentFragment](host)
	wrapper.Install(builder.constructor)
	return builder.constructor
}
