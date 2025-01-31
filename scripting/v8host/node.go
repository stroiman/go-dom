package v8host

import (
	"github.com/gost-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type nodeV8Wrapper struct {
	nodeV8WrapperBase[dom.Node]
}

func newNodeV8Wrapper(host *V8ScriptHost) nodeV8Wrapper {
	return nodeV8Wrapper{newNodeV8WrapperBase[dom.Node](host)}
}

func (n nodeV8Wrapper) nodeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	return v8.NewValue(n.scriptHost.iso, int32(instance.NodeType()))
}

func (n nodeV8Wrapper) decodeGetRootNodeOptions(
	ctx *V8ScriptContext,
	value *v8.Value,
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n nodeV8Wrapper) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}

func (w nodeV8Wrapper) defaultboolean() bool {
	return false
}
