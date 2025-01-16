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

func (w nodeWrapper) Contains(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	other := w.decodeNode(c.Arguments[0])
	result := instance.Contains(other)
	return w.toBoolean(result)
}

func (w nodeWrapper) InsertBefore(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	child := w.decodeNode(c.Arguments[1])
	result, err := instance.InsertBefore(node, child)
	if err != nil {
		panic(err)
	}
	return w.toNode(result)
}

func (w nodeWrapper) AppendChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	result, err := instance.AppendChild(node)
	if err != nil {
		panic(err)
	}
	return w.toNode(result)
}

func (w nodeWrapper) RemoveChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	child := w.decodeNode(c.Arguments[0])
	result, err := instance.RemoveChild(child)
	if err != nil {
		panic(err)
	}
	return w.toNode(result)
}
