// This file is generated. Do not edit.

package scripting

import (
	"errors"
	html "github.com/stroiman/go-dom/browser/html"
	v8 "github.com/tommie/v8go"
)

type windowV8Wrapper struct {
	nodeV8WrapperBase[html.Window]
}

func newWindowV8Wrapper(host *ScriptHost) *windowV8Wrapper {
	return &windowV8Wrapper{newNodeV8WrapperBase[html.Window](host)}
}

func init() {
	RegisterJSClass("Window", "EventTarget", createWindowPrototype)
}

func createWindowPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newWindowV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("close", v8.NewFunctionTemplateWithError(iso, wrapper.Close))
	prototypeTmpl.Set("stop", v8.NewFunctionTemplateWithError(iso, wrapper.Stop))
	prototypeTmpl.Set("focus", v8.NewFunctionTemplateWithError(iso, wrapper.Focus))
	prototypeTmpl.Set("blur", v8.NewFunctionTemplateWithError(iso, wrapper.Blur))
	prototypeTmpl.Set("open", v8.NewFunctionTemplateWithError(iso, wrapper.Open))
	prototypeTmpl.Set("alert", v8.NewFunctionTemplateWithError(iso, wrapper.Alert))
	prototypeTmpl.Set("confirm", v8.NewFunctionTemplateWithError(iso, wrapper.Confirm))
	prototypeTmpl.Set("prompt", v8.NewFunctionTemplateWithError(iso, wrapper.Prompt))
	prototypeTmpl.Set("print", v8.NewFunctionTemplateWithError(iso, wrapper.Print))
	prototypeTmpl.Set("postMessage", v8.NewFunctionTemplateWithError(iso, wrapper.PostMessage))

	prototypeTmpl.SetAccessorProperty("window",
		v8.NewFunctionTemplateWithError(iso, wrapper.Window),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("self",
		v8.NewFunctionTemplateWithError(iso, wrapper.Self),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("document",
		v8.NewFunctionTemplateWithError(iso, wrapper.Document),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("name",
		v8.NewFunctionTemplateWithError(iso, wrapper.Name),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("history",
		v8.NewFunctionTemplateWithError(iso, wrapper.History),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigation",
		v8.NewFunctionTemplateWithError(iso, wrapper.Navigation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("customElements",
		v8.NewFunctionTemplateWithError(iso, wrapper.CustomElements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("locationbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.Locationbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("menubar",
		v8.NewFunctionTemplateWithError(iso, wrapper.Menubar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("personalbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.Personalbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("scrollbars",
		v8.NewFunctionTemplateWithError(iso, wrapper.Scrollbars),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.Statusbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("toolbar",
		v8.NewFunctionTemplateWithError(iso, wrapper.Toolbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		v8.NewFunctionTemplateWithError(iso, wrapper.Status),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetStatus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("closed",
		v8.NewFunctionTemplateWithError(iso, wrapper.Closed),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("frames",
		v8.NewFunctionTemplateWithError(iso, wrapper.Frames),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.Length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("top",
		v8.NewFunctionTemplateWithError(iso, wrapper.Top),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("opener",
		v8.NewFunctionTemplateWithError(iso, wrapper.Opener),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetOpener),
		v8.None)
	prototypeTmpl.SetAccessorProperty("frameElement",
		v8.NewFunctionTemplateWithError(iso, wrapper.FrameElement),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigator",
		v8.NewFunctionTemplateWithError(iso, wrapper.Navigator),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientInformation",
		v8.NewFunctionTemplateWithError(iso, wrapper.ClientInformation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("originAgentCluster",
		v8.NewFunctionTemplateWithError(iso, wrapper.OriginAgentCluster),
		nil,
		v8.None)

	return constructor
}

func (w windowV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.host.iso, "Illegal Constructor")
}

func (w windowV8Wrapper) Close(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.close")
}

func (w windowV8Wrapper) Stop(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.stop")
}

func (w windowV8Wrapper) Focus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.focus")
}

func (w windowV8Wrapper) Blur(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.blur")
}

func (w windowV8Wrapper) Open(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.open")
}

func (w windowV8Wrapper) Alert(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.alert")
}

func (w windowV8Wrapper) Confirm(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.confirm")
}

func (w windowV8Wrapper) Prompt(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.prompt")
}

func (w windowV8Wrapper) Print(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.print")
}

func (w windowV8Wrapper) PostMessage(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.postMessage")
}

func (w windowV8Wrapper) Self(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Self")
}

func (w windowV8Wrapper) Document(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.host.MustGetContext(info.Context())
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return ctx.getInstanceForNode(result)
}

func (w windowV8Wrapper) Name(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Name")
}

func (w windowV8Wrapper) SetName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.SetName")
}

func (w windowV8Wrapper) History(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.History")
}

func (w windowV8Wrapper) Navigation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Navigation")
}

func (w windowV8Wrapper) CustomElements(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.CustomElements")
}

func (w windowV8Wrapper) Locationbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Locationbar")
}

func (w windowV8Wrapper) Menubar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Menubar")
}

func (w windowV8Wrapper) Personalbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Personalbar")
}

func (w windowV8Wrapper) Scrollbars(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Scrollbars")
}

func (w windowV8Wrapper) Statusbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Statusbar")
}

func (w windowV8Wrapper) Toolbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Toolbar")
}

func (w windowV8Wrapper) Status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Status")
}

func (w windowV8Wrapper) SetStatus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.SetStatus")
}

func (w windowV8Wrapper) Closed(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Closed")
}

func (w windowV8Wrapper) Frames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Frames")
}

func (w windowV8Wrapper) Length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Length")
}

func (w windowV8Wrapper) Top(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Top")
}

func (w windowV8Wrapper) Opener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Opener")
}

func (w windowV8Wrapper) SetOpener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.SetOpener")
}

func (w windowV8Wrapper) FrameElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.FrameElement")
}

func (w windowV8Wrapper) Navigator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.Navigator")
}

func (w windowV8Wrapper) ClientInformation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.ClientInformation")
}

func (w windowV8Wrapper) OriginAgentCluster(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: Window.OriginAgentCluster")
}
