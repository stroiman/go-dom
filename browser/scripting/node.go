package scripting

import (
	"github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type NodeV8Wrapper struct {
	NodeV8WrapperBase[dom.Node]
}

func NewNodeV8Wrapper(host *ScriptHost) NodeV8Wrapper {
	return NodeV8Wrapper{NewNodeV8WrapperBase[dom.Node](host)}
}

func (n NodeV8Wrapper) NodeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	return v8.NewValue(n.host.iso, int32(instance.NodeType()))
}
