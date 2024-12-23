package scripting

import (
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
