package gojahost

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

func (w nodeWrapper) constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func (w nodeWrapper) NodeType(c FunctionCall) Value {
	instance := w.getInstance(c)
	return w.toUnsignedShort(int(instance.NodeType()))
}

func (w nodeWrapper) decodeGetRootNodeOptions(v Value) (result dom.GetRootNodeOptions) {
	if o, ok := v.(*Object); ok {
		return dom.GetRootNodeOptions(o.Get("composed").ToBoolean())
	} else {
		return false
	}
}

// nodeListDynamicArray implements [g.DynamicArray] on top of a node list
type nodeListDynamicArray struct {
	dom.NodeList
	ctx *GojaContext
}

func (l nodeListDynamicArray) Len() int { return l.Length() }
func (l nodeListDynamicArray) Get(idx int) Value {
	result := l.Item(idx)
	if result == nil {
		return Null()
	}
	return l.ctx.cachedNodes[result.ObjectId()]
}

func (l nodeListDynamicArray) Set(_ int, _ Value) bool { return false }
func (l nodeListDynamicArray) SetLen(_ int) bool       { return false }

func (w nodeWrapper) toNodeList(l dom.NodeList) Value {
	if result := w.getCachedObject(l); result != nil {
		return result
	}
	result := w.instance.vm.NewDynamicArray(nodeListDynamicArray{l, w.instance})
	result.SetPrototype(w.instance.globals["NodeList"].Prototype)
	w.storeInternal(l, result)
	return result
}
