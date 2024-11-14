package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateDocumentPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplate(
		iso,
		func(args *v8.FunctionCallbackInfo) *v8.Value {
			v8Context := args.Context()
			scriptContext := host.contexts[v8Context]
			doc := NewDocument()
			id := doc.ObjectId()
			scriptContext.v8nodes[id] = args.This().Value
			scriptContext.domNodes[id] = doc
			internal, err := v8.NewValue(iso, id)
			if err != nil {
				// TODO
				panic(err)
			}
			args.This().SetInternalField(0, internal)
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
	proto.SetAccessorPropertyWithError("outerHTML", v8.AccessPropWithError{
		Get: func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			v8Context := info.Context()
			context := host.contexts[v8Context]
			id := info.This().GetInternalField(0).Int32()
			val := context.domNodes[id]
			if doc, ok := val.(Document); ok {
				v, _ := v8.NewValue(iso, doc.DocumentElement().OuterHTML())
				return v, nil
			} else {
				return nil, errors.New("Not a document")
			}
		}})
	return res
}
