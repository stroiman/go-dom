package scripting

import (
	"github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type nodeV8Wrapper struct {
	nodeV8WrapperBase[dom.Node]
}

func newNodeV8Wrapper(host *ScriptHost) nodeV8Wrapper {
	return nodeV8Wrapper{newNodeV8WrapperBase[dom.Node](host)}
}

func (n nodeV8Wrapper) NodeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	return v8.NewValue(n.host.iso, int32(instance.NodeType()))
}

func (n nodeV8Wrapper) decodeGetRootNodeOptions(
	ctx *ScriptContext,
	value *v8.Value,
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n nodeV8Wrapper) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}
