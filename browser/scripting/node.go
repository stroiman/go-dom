package scripting

import (
	"github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type NodeV8Wrapper struct {
	nodeV8WrapperBase[dom.Node]
}

func NewNodeV8Wrapper(host *ScriptHost) NodeV8Wrapper {
	return NodeV8Wrapper{newNodeV8WrapperBase[dom.Node](host)}
}

func (n NodeV8Wrapper) NodeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	return v8.NewValue(n.host.iso, int32(instance.NodeType()))
}

func (n NodeV8Wrapper) decodeGetRootNodeOptions(
	ctx *ScriptContext,
	value *v8.Value,
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n NodeV8Wrapper) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}
