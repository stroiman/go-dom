package browser

type EventTarget interface {
	// ObjectId is used internally for the scripting engine to associate a v8
	// object with the Go object it wraps.
	ObjectId() ObjectId
	AddEventListener(eventType string, listener EventHandler /* TODO: options */)
	RemoveEventListener(eventType string, listener EventHandler)
	DispatchEvent(event Event) error
}

type eventTarget struct {
	base
	lmap map[string][]EventHandler
}

func newEventTarget() eventTarget {
	return eventTarget{
		base: newBase(),
		lmap: make(map[string][]EventHandler),
	}
}

func NewEventTarget() EventTarget {
	res := newEventTarget()
	return &res
}

func (e *eventTarget) AddEventListener(eventType string, listener EventHandler) {
	// TODO: Handle options
	// - capture
	// - once
	// - passive. Defaults to false,
	// - signal - TODO: Implement AbortSignal
	// Browser specific
	// - Safari
	//   - passive defaults to true for `wheel`, `mousewheel` `touchstart`, `tourchmove` events
	// - Firefox (Gecko), receives an extra boolean argument, `wantsUntrusted`
	//   - https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener#wantsuntrusted
	listeners := e.lmap[eventType]
	for _, l := range listeners {
		if l.Equals(listener) {
			return
		}
	}
	e.lmap[eventType] = append(listeners, listener)
}

func (e *eventTarget) RemoveEventListener(eventType string, listener EventHandler) {
	listeners := e.lmap[eventType]
	for i, l := range listeners {
		if l.Equals(listener) {
			e.lmap[eventType] = append(listeners[:i], listeners[i+1:]...)
			return
		}
	}
}

func (e *eventTarget) DispatchEvent(event Event) error {
	listeners := e.lmap[event.Type()]
	for _, l := range listeners {
		err := l.HandleEvent(event)
		// TODO: Aggregate errors
		if err != nil {
			return err
		}
	}
	return nil
}

/* -------- Event & CustomEvent -------- */

type Event interface {
	Type() string
}

type CustomEvent interface {
	Event
}

type event struct {
	eventType string
}

func newEvent(eventType string) event {
	return event{eventType}
}

func (e *event) Type() string { return e.eventType }

type customEvent struct {
	event
}

func NewCustomEvent(eventType string) CustomEvent {
	e := &customEvent{newEvent(eventType)}
	return e
}

/* -------- EventHandler -------- */

// EventHandler is the interface for an event handler. In JavaScript; an event
// handler can be a function; or an object with a `handleEvent` function.
//
// Duplicate detection during _add_, or removal is based on equality. JavaScript
// equality does not translate natively to Go, so a handler must be able to
// detect equality by itself
type EventHandler interface {
	HandleEvent(event Event) error
	// The interface for removing event handlers requires the caller to pass in
	// the same handler to `RemoveEventListener`. In Go; functions cannot be
	// compared for equality; so we need to have some kind of mechanism to
	// identify if two handlers are identical.
	Equals(other EventHandler) bool
}

type HandlerFuncWithoutError = func(Event)
type HandlerFunc = func(Event) error

type eventHandlerFunc struct {
	handlerFunc func(Event) error
	id          ObjectId
}

// NewEventHandlerFunc creates an EventHandler wrapping a function with the
// right signature.
// Calling this twice for the same Go-function will be treated as different
// event handlers; as Go functions do not support equality.
func NewEventHandlerFunc(handler HandlerFunc) EventHandler {
	return eventHandlerFunc{handler, NewObjectId()}
}

func NewEventHandlerFuncWithoutError(handler HandlerFuncWithoutError) EventHandler {
	return eventHandlerFunc{func(event Event) error {
		handler(event)
		return nil
	}, NewObjectId()}
}

func (e eventHandlerFunc) HandleEvent(event Event) error {
	return e.handlerFunc(event)
}

func (e eventHandlerFunc) Equals(handler EventHandler) bool {
	x, ok := handler.(Entity)
	return ok && x.ObjectId() == e.id
}

// ObjectId makes the eventHandlerFunc type implement the Entity interface.
// While the code will still compile without this function; equality check will
// fail. This will be caught by tests verifying EventTarget behaviour.
func (e eventHandlerFunc) ObjectId() ObjectId { return e.id }
