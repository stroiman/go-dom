package scripting

import (
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

func CreateNodeList(host *ScriptHost) *v8.FunctionTemplate {
	builder := NewIllegalConstructorBuilder[browser.NodeList](host)
	return builder.constructor
}
