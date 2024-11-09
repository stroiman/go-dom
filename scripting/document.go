package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type V8Document struct {
	Value    *v8.Value
	document Document
}

func NewV8Document(val *v8.Value, doc Document) *V8Document {
	return &V8Document{val, doc}
}

func CreateDocumentPrototype(iso *v8.Isolate) *v8.FunctionTemplate {
	res := v8.NewFunctionTemplate(
		iso,
		func(args *v8.FunctionCallbackInfo) *v8.Value {
			return v8.Undefined(iso)
		},
	)
	res.GetInstanceTemplate().SetInternalFieldCount(1)
	return res
}
