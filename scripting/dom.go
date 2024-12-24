package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type ESDOMTokenList struct {
	ESWrapper[browser.Element]
}

func NewESDOMTokenList(host *ScriptHost) ESDOMTokenList {
	return ESDOMTokenList{NewESWrapper[browser.Element](host)}
}

func (l ESDOMTokenList) GetInstance(
	info *v8.FunctionCallbackInfo,
) (result browser.DOMTokenList, err error) {
	element, err := l.ESWrapper.GetInstance(info)
	if err == nil {
		result = browser.NewClassList(element)
	}
	return
}

func (l ESDOMTokenList) Toggle(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(l.host, info)
	token, err0 := TryParseArg(args, 0, l.DecodeUSVString)
	force, err1 := TryParseArg(args, 1, l.DecodeBoolean)
	instance, errInstance := l.GetInstance(info)
	if args.noOfReadArguments >= 2 {
		if err := errors.Join(err0, err1, errInstance); err != nil {
			return nil, err
		}
		if force {
			instance.Add(token)
			return v8.NewValue(l.host.iso, true)
		} else {
			instance.Remove(token)
			return v8.NewValue(l.host.iso, false)
		}
	}
	if err := errors.Join(err0, errInstance); err != nil {
		return nil, err
	}
	return v8.NewValue(l.host.iso, instance.Toggle(token))
}
