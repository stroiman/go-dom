package v8host

import (
	v8 "github.com/tommie/v8go"
)

type workItem struct {
	fn *v8.Function
}

type eventLoop struct {
	ctx          *V8ScriptContext
	workItems    chan workItem
	globalObject *v8.Object
	errorCb      func(error)
}

func newWorkItem(fn *v8.Function) workItem {
	return workItem{fn}
}

// dispatch places an item on the event loop to be executed immediately
func (l *eventLoop) dispatch(w workItem) {
	go func() {
		// Have seen issues with scripts being executed while trying to shut down
		// v8. I want a better, channel based, message passing, solution. But if
		// this works, it will for for a v 0.1 release
		l.ctx.host.mu.Lock()
		defer l.ctx.host.mu.Unlock()
		l.workItems <- w
	}()
}

func newEventLoop(context *V8ScriptContext, global *v8.Object, cb func(error)) *eventLoop {
	return &eventLoop{context, make(chan workItem), global, cb}
}

type disposable interface {
	dispose()
}

type disposeFunc func()

func (fn disposeFunc) dispose() { fn() }

func (l *eventLoop) Start() disposable {
	closer := make(chan bool)
	go func() {
		for i := range l.workItems {
			if l.ctx.disposed {
				return
			}
			_, err := i.fn.Call(l.globalObject)
			if err != nil {
				l.errorCb(err)
			}
		}
		closer <- true
	}()
	// There is some logic here that isn't tested specifically (but HTMX test
	// fails if not implemented properly).
	// When we shut down, we must be sure that we don't have any running scripts
	// when disposing the v8 Isolate, otherwise that will cause a panic.
	// That is why the close function waits for a channel event before proceeding
	return disposeFunc(func() {
		close(l.workItems)
		<-closer
	})
}

func installEventLoopGlobals(host *V8ScriptHost, globalObjectTemplate *v8.ObjectTemplate) {
	iso := host.iso

	globalObjectTemplate.Set(
		"setTimeout",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				f, err1 := helper.getFunctionArg(0)
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
