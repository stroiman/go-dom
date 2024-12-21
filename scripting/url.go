package scripting

import (
	"runtime/cgo"
	"unsafe"

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

func (u ESURL) storeHandle(
	value browser.URL,
	this *v8.Object,
	ctx *ScriptContext,
) (*v8.Value, error) {
	handle := cgo.NewHandle(value)
	ctx.AddDisposer(HandleDisposable(handle))

	internalField := v8.NewExternalValue(u.host.iso, unsafe.Pointer(&handle))
	this.SetInternalField(0, internalField)
	return nil, nil
}

func (u ESURL) CreateInstance(ctx *ScriptContext, this *v8.Object, url string) (*v8.Value, error) {
	value, err := browser.NewUrl(url)
	if err != nil {
		return nil, err
	}
	return u.storeHandle(value, this, ctx)
}

func (u ESURL) CreateInstanceBase(
	ctx *ScriptContext,
	this *v8.Object,
	url string,
	base string,
) (*v8.Value, error) {
	value, err := browser.NewUrlBase(url, base)
	if err != nil {
		return nil, err
	}
	return u.storeHandle(value, this, ctx)
}

func (u ESURL) GetInstance(info *v8.FunctionCallbackInfo) (browser.URL, error) {
	h := info.This().GetInternalField(0)
	handle := *(*cgo.Handle)(h.External())
	result := handle.Value().(browser.URL)
	return result, nil
}
