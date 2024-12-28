package scripting

import (
	"runtime/cgo"

	"github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

type ESURL struct {
	HandleReffedObject[dom.URL]
	Converters
}

func NewESURL(host *ScriptHost) ESURL {
	return ESURL{HandleReffedObject[dom.URL]{host}, Converters{}}
}

type HandleDisposable cgo.Handle

func (h HandleDisposable) Dispose() { cgo.Handle(h).Delete() }

func (u ESURL) CreateInstance(ctx *ScriptContext, this *v8.Object, url string) (*v8.Value, error) {
	value, err := dom.NewUrl(url)
	if err != nil {
		return nil, err
	}
	u.Store(value, ctx, this)
	return nil, nil
}

func (u ESURL) CreateInstanceBase(
	ctx *ScriptContext,
	this *v8.Object,
	url string,
	base string,
) (*v8.Value, error) {
	value, err := dom.NewUrlBase(url, base)
	if err != nil {
		return nil, err
	}
	u.Store(value, ctx, this)
	return nil, nil
}
