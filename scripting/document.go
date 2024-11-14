package scripting

import (
	"runtime"
	"unsafe"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type CachedElement[T Node] struct {
	Value    *v8.Value
	document T
	pinner   *runtime.Pinner
}

func NewCachedValue[T Node](val *v8.Value, doc T) *CachedElement[T] {
	result := &CachedElement[T]{
		Value:    val,
		document: doc,
		pinner:   new(runtime.Pinner),
	}
	result.pinner.Pin(result.pinner)
	result.pinner.Pin(result.document)
	result.pinner.Pin(result.Value)
	return result
}

func CreateDocumentPrototype(iso *v8.Isolate) *v8.FunctionTemplate {
	res := v8.NewFunctionTemplate(
		iso,
		func(args *v8.FunctionCallbackInfo) *v8.Value {
			doc := NewDocument()
			v8Doc := NewCachedValue(args.This().Value, doc)
			args.This().SetInternalField(0, v8.NewExternalValue(iso, unsafe.Pointer(v8Doc)))
			return v8.Undefined(iso)
		},
	)
	instanceTemplate := res.GetInstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	proto := res.PrototypeTemplate()
	proto.Set("createElement", v8.NewFunctionTemplate(iso,
		func(info *v8.FunctionCallbackInfo) *v8.Value {
			return v8.Undefined(iso)
		}))
	proto.SetAccessorProperty("outerHTML", v8.AccessProp{
		Get: func(info *v8.FunctionCallbackInfo) *v8.Value {
			tmp := info.This().GetInternalField(0)
			V8Document := (*CachedElement[Document])(tmp.External())
			v, _ := v8.NewValue(iso, V8Document.document.DocumentElement().OuterHTML())
			return v
		}})
	return res
}
