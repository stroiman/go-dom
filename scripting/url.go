package scripting

import (
	"errors"
	"runtime/cgo"

	"github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type ESURL struct {
	Converters
	host *ScriptHost
}

func NewESURL(host *ScriptHost) ESURL { return ESURL{Converters{}, host} }

type HandleDisposable cgo.Handle

func (h HandleDisposable) Dispose() { cgo.Handle(h).Delete() }

func (u ESURL) CreateInstance(ctx *ScriptContext, this *v8.Object, url string) (*v8.Value, error) {
	value, err := browser.NewUrl(url)
	if err != nil {
		return nil, err
	}
	handle := cgo.NewHandle(value)
	internalField, err := v8.NewValue(u.host.iso, (uintptr)(handle))
	ctx.AddDisposer(HandleDisposable(handle))

	if err != nil {
		return nil, err
	}
	this.SetInternalField(0, internalField)
	return nil, nil
}

func (u ESURL) CreateInstanceBase(
	ctx *ScriptContext,
	this *v8.Object,
	url string,
	base string,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}

func (u ESURL) GetInstance(info *v8.FunctionCallbackInfo) (browser.URL, error) {
	h := info.This().GetInternalField(0)
	handle := cgo.Handle(h.Uint32())
	result := handle.Value().(browser.URL)
	return result, nil
}
