package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/dom"
)

func (w nodeWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func (w nodeWrapper) textContent(c g.FunctionCall) g.Value {
	i := w.getInstance(c)
	return w.ctx.vm.ToValue(i.TextContent())
}
func (w nodeWrapper) setTextContent(c g.FunctionCall) g.Value {
	i := w.getInstance(c)
	i.SetTextContent(c.Argument(0).String())
	return g.Undefined()
}

func (w nodeWrapper) nodeType(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	return w.toUnsignedShort(int(instance.NodeType()))
}

func (w nodeWrapper) decodeGetRootNodeOptions(v g.Value) (result dom.GetRootNodeOptions) {
	if o, ok := v.(*g.Object); ok {
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
func (l nodeListDynamicArray) Get(idx int) g.Value {
	result := l.Item(idx)
	if result == nil {
		return g.Null()
	}
	return l.ctx.cachedNodes[result.ObjectId()]
}

func (l nodeListDynamicArray) Set(_ int, _ g.Value) bool { return false }
func (l nodeListDynamicArray) SetLen(_ int) bool         { return false }

func (w nodeWrapper) toNodeList(l dom.NodeList) g.Value {
	if result := w.getCachedObject(l); result != nil {
		return result
	}
	result := w.ctx.vm.NewDynamicArray(nodeListDynamicArray{l, w.ctx})
	result.SetPrototype(w.ctx.globals["NodeList"].Prototype)
	w.storeInternal(l, result)
	return result
}
