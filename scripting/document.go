package scripting

import (
	"fmt"
	"runtime"
	"unsafe"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type V8Document struct {
	Value    *v8.Value
	document Document
	pinner   runtime.Pinner
}

func NewV8Document(val *v8.Value, doc Document) *V8Document {
	result := &V8Document{
		Value:    val,
		document: doc,
	}
	result.pinner.Pin(result)
	result.pinner.Pin(doc)
	return result
}

func CreateDocumentPrototype(iso *v8.Isolate) *v8.FunctionTemplate {
	res := v8.NewFunctionTemplate(
		iso,
		func(args *v8.FunctionCallbackInfo) *v8.Value {
			doc := NewDocument()
			v8Doc := NewV8Document(args.This().Value, doc)
			args.This().SetInternalField(0, unsafe.Pointer(v8Doc))
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
			fmt.Println("Getting tmp", tmp.IsExternal())
			V8Document := (*V8Document)(tmp.External())
			fmt.Println("Document", V8Document.document)
			// return v8.Undefined(iso)
			v, _ := v8.NewValue(iso, V8Document.document.DocumentElement().OuterHTML())
			return v
		}})
	return res
}
