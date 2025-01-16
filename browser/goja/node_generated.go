// This file is generated. Do not edit.

package goja

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
	prototype.Set("contains", w.contains)
	prototype.Set("insertBefore", w.insertBefore)
	prototype.Set("appendChild", w.appendChild)
	prototype.Set("removeChild", w.removeChild)
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
