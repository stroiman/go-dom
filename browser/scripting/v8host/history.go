package v8host

import (
	"github.com/gost-dom/browser/browser/html"
	v8 "github.com/tommie/v8go"
)

type historyV8Wrapper struct {
	handleReffedObject[html.History]
}

func newHistoryV8Wrapper(host *V8ScriptHost) *historyV8Wrapper {
	return &historyV8Wrapper{newHandleReffedObject[html.History](host)}
}

func (w historyV8Wrapper) defaultDelta() int {
	return 0
}

func (w historyV8Wrapper) defaultUrl() string {
	return ""
}

func (w historyV8Wrapper) decodeAny(
	ctx *V8ScriptContext,
	val *v8.Value,
) (html.HistoryState, error) {
	r, err := v8.JSONStringify(ctx.v8ctx, val)
	return html.HistoryState(r), err
}

func (w historyV8Wrapper) toJSON(ctx *V8ScriptContext, val html.HistoryState) (*v8.Value, error) {
	return v8.JSONParse(ctx.v8ctx, string(val))
}
