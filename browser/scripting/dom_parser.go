package scripting

import (
	"errors"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/html"
	v8 "github.com/tommie/v8go"
)

func CreateDOMParserPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	constructor := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, nil
		},
	)
	prototype := constructor.PrototypeTemplate()
	prototype.Set(
		"parseFromString",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				window := ctx.Window()
				args := newArgumentHelper(host, info)
				html, err0 := args.GetStringArg(0)
				contentType, err1 := args.GetStringArg(1)
				if err := errors.Join(err0, err1); err != nil {
					return nil, err
				}
				if contentType != "text/html" {
					return nil, v8.NewTypeError(host.iso,
						"DOMParser.parseFromString only supports text/html yet",
					)
				}
				domParser := NewDOMParser()
				var doc dom.Document
				if err := domParser.ParseReader(window, &doc, strings.NewReader(html)); err == nil {
					return ctx.GetInstanceForNode(doc)
				} else {
					return nil, err
				}
			},
		),
	)
	return constructor
}
