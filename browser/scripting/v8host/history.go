package v8host

import "github.com/stroiman/go-dom/browser/html"

type historyV8Wrapper struct {
	handleReffedObject[html.History]
}

func newHistoryV8Wrapper(host *V8ScriptHost) *historyV8Wrapper {
	return &historyV8Wrapper{newHandleReffedObject[html.History](host)}
}
