package browser

import (
	"log/slog"
)

type EventTarget interface {
	// ObjectId is used internally for the scripting engine to associate a v8
	// object with the Go object it wraps.
	ObjectId() ObjectId
	AddEventListener(eventType string, listener EventHandler /* TODO: options */)
	RemoveEventListener(eventType string, listener EventHandler)
	DispatchEvent(event Event) bool
	SetCatchAllHandler(listener EventHandler)
	// Unexported
	dispatchError(err ErrorEvent)
}

type eventTarget struct {
	base
	parentTarget    EventTarget
	lmap            map[string][]EventHandler
	catchAllHandler EventHandler
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
	slog.Debug("AddEventListener", "EventType", eventType)
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

func (e *eventTarget) SetCatchAllHandler(handler EventHandler) {
	e.catchAllHandler = handler
}

func (e *eventTarget) DispatchEvent(event Event) bool {
	event.reset()
	if e.catchAllHandler != nil {
		if err := e.catchAllHandler.HandleEvent(event); err != nil {
			slog.Debug("Error occurred", "error", err.Error())
			e.dispatchError(NewErrorEvent(err))
		}
	}
	slog.Debug("Dispatch event", "EventType", event.Type())
	listeners := e.lmap[event.Type()]

	for _, l := range listeners {
		if err := l.HandleEvent(event); err != nil {
			slog.Debug("Error occurred", "error", err.Error())
			e.dispatchError(NewErrorEvent(err))
		}
	}

	if e.parentTarget != nil && event.shouldPropagate() {
		e.parentTarget.DispatchEvent(event)
	}
	return !event.isCancelled()
}

func (e *eventTarget) dispatchError(event ErrorEvent) {
	if e.parentTarget == nil {
		e.DispatchEvent(event)
	} else {
		e.parentTarget.dispatchError(event)
	}
}

/* -------- Event & CustomEvent -------- */

type Event interface {
	Entity
	PreventDefault()
	Type() string
	StopPropagation()
	// Unexported
	reset()
	isCancelled() bool
	shouldPropagate() bool
}

type ErrorEvent interface {
	Event
	Err() error
	Error() string
}

type CustomEvent interface {
	Event
}

type event struct {
	base
	cancelable bool
	cancelled  bool
	eventType  string
	bubbles    bool
	propagate  bool
}

type errorEvent struct {
	event
	err error
}

type CustomEventOption interface {
	updateEvent(*event)
}

type eventOptionFunc func(*event)

func (f eventOptionFunc) updateEvent(e *event) { f(e) }

func EventBubbles(bubbles bool) CustomEventOption {
	return eventOptionFunc(func(e *event) { e.bubbles = bubbles })
}
func EventCancelable(cancelable bool) CustomEventOption {
	return eventOptionFunc(func(e *event) { e.cancelable = cancelable })
}

func newEvent(eventType string) event {
	return event{
		base:      newBase(),
		eventType: eventType,
		bubbles:   false,
		propagate: false,
	}
}

type customEvent struct {
	event
}

func NewCustomEvent(eventType string, options ...CustomEventOption) CustomEvent {
	e := &customEvent{newEvent(eventType)}
	for _, o := range options {
		o.updateEvent(&e.event)
	}
	return e
}

func (e *event) Type() string          { return e.eventType }
func (e *event) StopPropagation()      { e.propagate = false }
func (e *event) PreventDefault()       { e.cancelled = true }
func (e *event) shouldPropagate() bool { return e.propagate }
func (e *event) isCancelled() bool     { return e.cancelable && e.cancelled }

func (e *event) reset() {
	e.propagate = e.bubbles
	e.cancelled = false
}

func NewErrorEvent(err error) ErrorEvent {
	return &errorEvent{newEvent("error"), err}
}

func (e *errorEvent) Err() error    { return e.err }
func (e *errorEvent) Error() string { return e.err.Error() }

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
