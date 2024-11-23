package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateNamedNodeMap(host *ScriptHost) *v8.FunctionTemplate {
	builder := NewIllegalConstructorBuilder[NamedNodeMap](host)
	return builder.constructor
}
