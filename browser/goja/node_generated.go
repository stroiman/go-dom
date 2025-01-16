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
	instance.Contains(other)
	return nil
}

func (w nodeWrapper) InsertBefore(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	child := w.decodeNode(c.Arguments[1])
	instance.InsertBefore(node, child)
	return nil
}

func (w nodeWrapper) AppendChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	instance.AppendChild(node)
	return nil
}

func (w nodeWrapper) RemoveChild(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	child := w.decodeNode(c.Arguments[0])
	instance.RemoveChild(child)
	return nil
}
