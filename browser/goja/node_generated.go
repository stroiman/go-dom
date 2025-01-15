// This file is generated. Do not edit.

package goja

import dom "github.com/stroiman/go-dom/browser/dom"

type nodeWrapper struct {
	baseInstanceWrapper[dom.Node]
}

func newNodeWrapper(instance *GojaContext) wrapper {
	return nodeWrapper{newBaseInstanceWrapper[dom.Node](instance)}
}
