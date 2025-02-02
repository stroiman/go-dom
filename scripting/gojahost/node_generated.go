// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	dom "github.com/gost-dom/browser/dom"
)

func init() {
	installClass("Node", "EventTarget", newNodeWrapper)
}

type nodeWrapper struct {
	baseInstanceWrapper[dom.Node]
}

func newNodeWrapper(instance *GojaContext) wrapper {
	return nodeWrapper{newBaseInstanceWrapper[dom.Node](instance)}
}
func (w nodeWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("getRootNode", w.getRootNode)
	prototype.Set("cloneNode", w.cloneNode)
	prototype.Set("isSameNode", w.isSameNode)
	prototype.Set("contains", w.contains)
	prototype.Set("insertBefore", w.insertBefore)
	prototype.Set("appendChild", w.appendChild)
	prototype.Set("removeChild", w.removeChild)
	prototype.DefineAccessorProperty("nodeType", w.ctx.vm.ToValue(w.nodeType), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("nodeName", w.ctx.vm.ToValue(w.nodeName), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("isConnected", w.ctx.vm.ToValue(w.isConnected), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("ownerDocument", w.ctx.vm.ToValue(w.ownerDocument), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("parentElement", w.ctx.vm.ToValue(w.parentElement), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("childNodes", w.ctx.vm.ToValue(w.childNodes), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("firstChild", w.ctx.vm.ToValue(w.firstChild), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("previousSibling", w.ctx.vm.ToValue(w.previousSibling), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("nextSibling", w.ctx.vm.ToValue(w.nextSibling), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w nodeWrapper) getRootNode(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	options := w.decodeGetRootNodeOptions(c.Arguments[0])
	result := instance.GetRootNode(options)
	return w.toNode(result)
}

func (w nodeWrapper) cloneNode(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	subtree := w.decodeboolean(c.Arguments[0])
	result := instance.CloneNode(subtree)
	return w.toNode(result)
}

func (w nodeWrapper) isSameNode(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	otherNode := w.decodeNode(c.Arguments[0])
	result := instance.IsSameNode(otherNode)
	return w.toBoolean(result)
}

func (w nodeWrapper) contains(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	other := w.decodeNode(c.Arguments[0])
	result := instance.Contains(other)
	return w.toBoolean(result)
}

func (w nodeWrapper) insertBefore(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	child := w.decodeNode(c.Arguments[1])
	result, err := instance.InsertBefore(node, child)
	if err != nil {
		panic(err)
	}
	return w.toNode(result)
}

func (w nodeWrapper) appendChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	result, err := instance.AppendChild(node)
	if err != nil {
		panic(err)
	}
	return w.toNode(result)
}

func (w nodeWrapper) removeChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	child := w.decodeNode(c.Arguments[0])
	result, err := instance.RemoveChild(child)
	if err != nil {
		panic(err)
	}
	return w.toNode(result)
}

func (w nodeWrapper) nodeName(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.NodeName()
	return w.toDOMString(result)
}

func (w nodeWrapper) isConnected(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.IsConnected()
	return w.toBoolean(result)
}

func (w nodeWrapper) ownerDocument(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.OwnerDocument()
	return w.toDocument(result)
}

func (w nodeWrapper) parentElement(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.ParentElement()
	return w.toElement(result)
}

func (w nodeWrapper) childNodes(c g.FunctionCall) g.Value {
	panic("Node.childNodes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w nodeWrapper) firstChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.FirstChild()
	return w.toNode(result)
}

func (w nodeWrapper) previousSibling(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.PreviousSibling()
	return w.toNode(result)
}

func (w nodeWrapper) nextSibling(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.NextSibling()
	return w.toNode(result)
}
