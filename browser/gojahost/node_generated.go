// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	dom "github.com/stroiman/go-dom/browser/dom"
)

type nodeWrapper struct {
	baseInstanceWrapper[dom.Node]
}

func newNodeWrapper(instance *GojaContext) wrapper {
	return nodeWrapper{newBaseInstanceWrapper[dom.Node](instance)}
}
func (w nodeWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("getRootNode", w.getRootNode)
	prototype.Set("contains", w.contains)
	prototype.Set("insertBefore", w.insertBefore)
	prototype.Set("appendChild", w.appendChild)
	prototype.Set("removeChild", w.removeChild)
	prototype.DefineAccessorProperty("nodeType", w.instance.vm.ToValue(w.NodeType), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("nodeName", w.instance.vm.ToValue(w.NodeName), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("isConnected", w.instance.vm.ToValue(w.IsConnected), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("ownerDocument", w.instance.vm.ToValue(w.OwnerDocument), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("childNodes", w.instance.vm.ToValue(w.ChildNodes), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("firstChild", w.instance.vm.ToValue(w.FirstChild), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("previousSibling", w.instance.vm.ToValue(w.PreviousSibling), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("nextSibling", w.instance.vm.ToValue(w.NextSibling), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w nodeWrapper) getRootNode(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	options := w.decodeGetRootNodeOptions(c.Arguments[0])
	result := instance.GetRootNode(options)
	return w.toNode(result)
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

func (w nodeWrapper) NodeName(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.NodeName()
	return w.toDOMString(result)
}

func (w nodeWrapper) IsConnected(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.IsConnected()
	return w.toBoolean(result)
}

func (w nodeWrapper) OwnerDocument(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.OwnerDocument()
	return w.toDocument(result)
}

func (w nodeWrapper) ChildNodes(c g.FunctionCall) g.Value {
	panic("Node.ChildNodes: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w nodeWrapper) FirstChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.FirstChild()
	return w.toNode(result)
}

func (w nodeWrapper) PreviousSibling(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.PreviousSibling()
	return w.toNode(result)
}

func (w nodeWrapper) NextSibling(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.NextSibling()
	return w.toNode(result)
}
