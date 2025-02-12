// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("Window", "EventTarget", createWindowPrototype)
}

type windowV8Wrapper struct {
	nodeV8WrapperBase[html.Window]
}

func newWindowV8Wrapper(scriptHost *V8ScriptHost) *windowV8Wrapper {
	return &windowV8Wrapper{newNodeV8WrapperBase[html.Window](scriptHost)}
}

func createWindowPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newWindowV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w windowV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("close", v8.NewFunctionTemplateWithError(iso, w.close))
	prototypeTmpl.Set("stop", v8.NewFunctionTemplateWithError(iso, w.stop))
	prototypeTmpl.Set("focus", v8.NewFunctionTemplateWithError(iso, w.focus))
	prototypeTmpl.Set("blur", v8.NewFunctionTemplateWithError(iso, w.blur))
	prototypeTmpl.Set("open", v8.NewFunctionTemplateWithError(iso, w.open))
	prototypeTmpl.Set("alert", v8.NewFunctionTemplateWithError(iso, w.alert))
	prototypeTmpl.Set("confirm", v8.NewFunctionTemplateWithError(iso, w.confirm))
	prototypeTmpl.Set("prompt", v8.NewFunctionTemplateWithError(iso, w.prompt))
	prototypeTmpl.Set("print", v8.NewFunctionTemplateWithError(iso, w.print))
	prototypeTmpl.Set("postMessage", v8.NewFunctionTemplateWithError(iso, w.postMessage))

	prototypeTmpl.SetAccessorProperty("window",
		v8.NewFunctionTemplateWithError(iso, w.window),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("self",
		v8.NewFunctionTemplateWithError(iso, w.self),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("document",
		v8.NewFunctionTemplateWithError(iso, w.document),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("name",
		v8.NewFunctionTemplateWithError(iso, w.name),
		v8.NewFunctionTemplateWithError(iso, w.setName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("history",
		v8.NewFunctionTemplateWithError(iso, w.history),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigation",
		v8.NewFunctionTemplateWithError(iso, w.navigation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("customElements",
		v8.NewFunctionTemplateWithError(iso, w.customElements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("locationbar",
		v8.NewFunctionTemplateWithError(iso, w.locationbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("menubar",
		v8.NewFunctionTemplateWithError(iso, w.menubar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("personalbar",
		v8.NewFunctionTemplateWithError(iso, w.personalbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("scrollbars",
		v8.NewFunctionTemplateWithError(iso, w.scrollbars),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusbar",
		v8.NewFunctionTemplateWithError(iso, w.statusbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("toolbar",
		v8.NewFunctionTemplateWithError(iso, w.toolbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		v8.NewFunctionTemplateWithError(iso, w.status),
		v8.NewFunctionTemplateWithError(iso, w.setStatus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("closed",
		v8.NewFunctionTemplateWithError(iso, w.closed),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("frames",
		v8.NewFunctionTemplateWithError(iso, w.frames),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("top",
		v8.NewFunctionTemplateWithError(iso, w.top),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("opener",
		v8.NewFunctionTemplateWithError(iso, w.opener),
		v8.NewFunctionTemplateWithError(iso, w.setOpener),
		v8.None)
	prototypeTmpl.SetAccessorProperty("frameElement",
		v8.NewFunctionTemplateWithError(iso, w.frameElement),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigator",
		v8.NewFunctionTemplateWithError(iso, w.navigator),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientInformation",
		v8.NewFunctionTemplateWithError(iso, w.clientInformation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("originAgentCluster",
		v8.NewFunctionTemplateWithError(iso, w.originAgentCluster),
		nil,
		v8.None)
}

func (w windowV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w windowV8Wrapper) close(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.close")
	return nil, errors.New("Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) stop(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.stop")
	return nil, errors.New("Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) focus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.focus")
	return nil, errors.New("Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) blur(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.blur")
	return nil, errors.New("Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) open(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.open")
	return nil, errors.New("Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) alert(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.alert")
	return nil, errors.New("Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) confirm(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.confirm")
	return nil, errors.New("Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) prompt(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.prompt")
	return nil, errors.New("Window.prompt: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) print(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.print")
	return nil, errors.New("Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) postMessage(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.postMessage")
	return nil, errors.New("Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) self(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.self")
	return nil, errors.New("Window.self: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) document(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: Window.document")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return ctx.getInstanceForNode(result)
}

func (w windowV8Wrapper) name(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.name")
	return nil, errors.New("Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) setName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.setName")
	return nil, errors.New("Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) navigation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.navigation")
	return nil, errors.New("Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) customElements(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.customElements")
	return nil, errors.New("Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) locationbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.locationbar")
	return nil, errors.New("Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) menubar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.menubar")
	return nil, errors.New("Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) personalbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.personalbar")
	return nil, errors.New("Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) scrollbars(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.scrollbars")
	return nil, errors.New("Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) statusbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.statusbar")
	return nil, errors.New("Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) toolbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.toolbar")
	return nil, errors.New("Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.status")
	return nil, errors.New("Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) setStatus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.setStatus")
	return nil, errors.New("Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) closed(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.closed")
	return nil, errors.New("Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) frames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.frames")
	return nil, errors.New("Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.length")
	return nil, errors.New("Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) top(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.top")
	return nil, errors.New("Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) opener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.opener")
	return nil, errors.New("Window.opener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) setOpener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.setOpener")
	return nil, errors.New("Window.setOpener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) frameElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.frameElement")
	return nil, errors.New("Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) navigator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.navigator")
	return nil, errors.New("Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) clientInformation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.clientInformation")
	return nil, errors.New("Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) originAgentCluster(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: Window.originAgentCluster")
	return nil, errors.New("Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
