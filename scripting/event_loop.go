package scripting

import (
	v8 "github.com/tommie/v8go"
)

type workItem struct {
	fn *v8.Function
}

type EventLoop struct {
	workItems    chan workItem
	globalObject *v8.Object
	errorCb      func(error)
}

func newWorkItem(fn *v8.Function) workItem {
	return workItem{fn}
}

// dispatch places an item on the event loop to be executed immediately
func (l *EventLoop) dispatch(w workItem) {
	go func() {
		l.workItems <- w
	}()
}

func NewEventLoop(global *v8.Object, cb func(error)) *EventLoop {
	return &EventLoop{make(chan workItem), global, cb}
}

func (l *EventLoop) Start() {
	go func() {
		for i := range l.workItems {
			_, err := i.fn.Call(l.globalObject)
			if err != nil {
				l.errorCb(err)
			}
		}
	}()
}

func (l *EventLoop) Stop() {
	close(l.workItems)
}

func installEventLoopGlobals(host *ScriptHost, globalObjectTemplate *v8.ObjectTemplate) {
	iso := host.iso

	globalObjectTemplate.Set(
		"setTimeout",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				f, err1 := helper.GetFunctionArg(0)
				// delay, err2 := helper.GetInt32Arg(1)
				if err1 == nil {
					ctx.eventLoop.dispatch(newWorkItem(f))
				}
				// TODO: Return a cancel token
				return v8.Undefined(iso), err1
			},
		),
	)
}
