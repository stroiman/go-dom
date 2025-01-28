package v8host

import (
	. "github.com/gost-dom/browser/browser/dom"

	v8 "github.com/tommie/v8go"
)

// createShadowRootPrototype currently only exists to allow code to check
// for inheritence, i.e., `node instanceof DocumentFragment`.
// This is performed by HTMX; but it doesn't itself _create_ a shadow root.
func createShadowRootPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	builder := newIllegalConstructorBuilder[ShadowRoot](host)
	return builder.constructor
}
