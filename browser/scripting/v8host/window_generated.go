// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/stroiman/go-dom/browser/html"
	v8 "github.com/tommie/v8go"
)

type windowV8Wrapper struct {
	nodeV8WrapperBase[html.Window]
}

func newWindowV8Wrapper(scriptHost *V8ScriptHost) *windowV8Wrapper {
	return &windowV8Wrapper{newNodeV8WrapperBase[html.Window](scriptHost)}
}

func init() {
	registerJSClass("Window", "EventTarget", createWindowPrototype)
}

func createWindowPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newWindowV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("close", v8.NewFunctionTemplateWithError(iso, wrapper.close))
	prototypeTmpl.Set("stop", v8.NewFunctionTemplateWithError(iso, wrapper.stop))
	prototypeTmpl.Set("focus", v8.NewFunctionTemplateWithError(iso, wrapper.focus))
	prototypeTmpl.Set("blur", v8.NewFunctionTemplateWithError(iso, wrapper.blur))
	prototypeTmpl.Set("open", v8.NewFunctionTemplateWithError(iso, wrapper.open))
	prototypeTmpl.Set("alert", v8.NewFunctionTemplateWithError(iso, wrapper.alert))
	prototypeTmpl.Set("confirm", v8.NewFunctionTemplateWithError(iso, wrapper.confirm))
	prototypeTmpl.Set("prompt", v8.NewFunctionTemplateWithError(iso, wrapper.prompt))
	prototypeTmpl.Set("print", v8.NewFunctionTemplateWithError(iso, wrapper.print))
	prototypeTmpl.Set("postMessage", v8.NewFunctionTemplateWithError(iso, wrapper.postMessage))

	prototypeTmpl.SetAccessorProperty("window",
		v8.NewFunctionTemplateWithError(iso, wrapper.window),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("self",
		v8.NewFunctionTemplateWithError(iso, wrapper.self),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("document",
		v8.NewFunctionTemplateWithError(iso, wrapper.document),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("name",
		v8.NewFunctionTemplateWithError(iso, wrapper.name),
		v8.NewFunctionTemplateWithError(iso, wrapper.setName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("history",
		v8.NewFunctionTemplateWithError(iso, wrapper.history),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigation",
		v8.NewFunctionTemplateWithError(iso, wrapper.navigation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("customElements",
		v8.NewFunctionTemplateWithError(iso, wrapper.customElements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("locationbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.locationbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("menubar",
		v8.NewFunctionTemplateWithError(iso, wrapper.menubar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("personalbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.personalbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("scrollbars",
		v8.NewFunctionTemplateWithError(iso, wrapper.scrollbars),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.statusbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("toolbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.toolbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		v8.NewFunctionTemplateWithError(iso, wrapper.status),
		v8.NewFunctionTemplateWithError(iso, wrapper.setStatus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("closed",
		v8.NewFunctionTemplateWithError(iso, wrapper.closed),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("frames",
		v8.NewFunctionTemplateWithError(iso, wrapper.frames),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("top",
		v8.NewFunctionTemplateWithError(iso, wrapper.top),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("opener",
		v8.NewFunctionTemplateWithError(iso, wrapper.opener),
		v8.NewFunctionTemplateWithError(iso, wrapper.setOpener),
		v8.None)
	prototypeTmpl.SetAccessorProperty("frameElement",
		v8.NewFunctionTemplateWithError(iso, wrapper.frameElement),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigator",
		v8.NewFunctionTemplateWithError(iso, wrapper.navigator),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientInformation",
		v8.NewFunctionTemplateWithError(iso, wrapper.clientInformation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("originAgentCluster",
		v8.NewFunctionTemplateWithError(iso, wrapper.originAgentCluster),
		nil,
		v8.None)

	return constructor
}

func (w windowV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.host.iso, "Illegal Constructor")
}

func (w windowV8Wrapper) close(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.close: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) stop(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.stop: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) focus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.focus: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) blur(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.blur: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) open(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.open: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) alert(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.alert: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) confirm(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.confirm: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) prompt(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.prompt: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) print(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.print: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) postMessage(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.postMessage: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) self(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.self: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) document(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return ctx.getInstanceForNode(result)
}

func (w windowV8Wrapper) name(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.name: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) setName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.setName: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) history(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.history: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) navigation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.navigation: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) customElements(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.customElements: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) locationbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.locationbar: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) menubar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.menubar: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) personalbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.personalbar: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) scrollbars(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.scrollbars: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) statusbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.statusbar: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) toolbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.toolbar: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.status: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) setStatus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.setStatus: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) closed(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.closed: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) frames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.frames: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.length: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) top(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.top: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) opener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.opener: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) setOpener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.setOpener: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) frameElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.frameElement: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) navigator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.navigator: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) clientInformation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.clientInformation: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (w windowV8Wrapper) originAgentCluster(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Window.originAgentCluster: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}
