package dom

import (
	"github.com/gost-dom/browser/browser/internal/entity"
	"github.com/gost-dom/browser/browser/internal/log"
)

type EventTarget interface {
	// ObjectId is used internally for the scripting engine to associate a v8
	// object with the Go object it wraps.
	AddEventListener(eventType string, listener EventHandler /* TODO: options */)
	RemoveEventListener(eventType string, listener EventHandler)
	DispatchEvent(event Event) bool
	SetCatchAllHandler(listener EventHandler)
	// Unexported
	dispatchError(err ErrorEvent)
	dispatchEvent(event Event) bool
	setSelf(e EventTarget)
}

type eventTarget struct {
	parentTarget    EventTarget
	lmap            map[string][]EventHandler
	catchAllHandler EventHandler
	self            EventTarget
}

func NewEventTarget() EventTarget {
	res := newEventTarget()
	return &res
}

func SetEventTargetSelf(t EventTarget) {
	t.setSelf(t)
}

func newEventTarget() eventTarget {
	return eventTarget{
		lmap: make(map[string][]EventHandler),
	}
}

func (e *eventTarget) setSelf(self EventTarget) {
	e.self = self
}

func (e *eventTarget) AddEventListener(eventType string, listener EventHandler) {
	log.Debug("AddEventListener", "EventType", eventType)
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
	event.reset(e.self)
	log.Debug("Dispatch event", "EventType", event.Type())
	return e.dispatchEvent(event)
}

func (e *eventTarget) dispatchEvent(event Event) bool {
	event.setCurrentTarget(e.self)
	defer func() { event.setCurrentTarget(nil) }()
	if e.catchAllHandler != nil {
		if err := e.catchAllHandler.HandleEvent(event); err != nil {
			log.Debug("Error occurred", "error", err.Error())
			e.dispatchError(NewErrorEvent(err))
		}
	}
	listeners := e.lmap[event.Type()]

	for _, l := range listeners {
		if err := l.HandleEvent(event); err != nil {
			log.Debug("Error occurred", "error", err.Error())
			e.dispatchError(NewErrorEvent(err))
		}
	}

	if e.parentTarget != nil && event.shouldPropagate() {
		e.parentTarget.dispatchEvent(event)
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
	entity.Entity
	Cancelable() bool
	Bubbles() bool
	PreventDefault()
	Type() string
	StopPropagation()
	Target() EventTarget
	CurrentTarget() EventTarget
	// Unexported
	reset(t EventTarget)
	isCancelled() bool
	shouldPropagate() bool
	setCurrentTarget(t EventTarget)
}

type ErrorEvent interface {
	Event
	Err() error
	Error() string
}

type CustomEvent interface {
	Event
}

type EventOption interface {
	updateEvent(*event)
}

type EventOptions []EventOption

func (o EventOptions) updateEvent(e *event) {
	for _, option := range o {
		option.updateEvent(e)
	}
}

type eventOptionFunc func(*event)

func (f eventOptionFunc) updateEvent(e *event) { f(e) }

func EventBubbles(bubbles bool) EventOption {
	return eventOptionFunc(func(e *event) { e.bubbles = bubbles })
}
func EventCancelable(cancelable bool) EventOption {
	return eventOptionFunc(func(e *event) { e.cancelable = cancelable })
}

/* -------- event -------- */

type event struct {
	entity.Entity
	cancelable    bool
	cancelled     bool
	eventType     string
	bubbles       bool
	propagate     bool
	target        EventTarget
	currentTarget EventTarget
}

func newEvent(eventType string) event {
	return event{
		Entity:    entity.New(),
		eventType: eventType,
		bubbles:   false,
		propagate: false,
	}
}

func NewEvent(eventType string, options ...EventOption) Event {
	e := newEvent(eventType)
	for _, o := range options {
		o.updateEvent(&e)
	}
	return &e
}

func (e *event) Type() string                   { return e.eventType }
func (e *event) StopPropagation()               { e.propagate = false }
func (e *event) PreventDefault()                { e.cancelled = true }
func (e *event) Cancelable() bool               { return e.cancelable }
func (e *event) Bubbles() bool                  { return e.bubbles }
func (e *event) shouldPropagate() bool          { return e.propagate }
func (e *event) isCancelled() bool              { return e.cancelable && e.cancelled }
func (e *event) Target() EventTarget            { return e.target }
func (e *event) CurrentTarget() EventTarget     { return e.currentTarget }
func (e *event) setCurrentTarget(t EventTarget) { e.currentTarget = t }

func (e *event) reset(t EventTarget) {
	e.target = t
	e.propagate = e.bubbles
	e.cancelled = false
}

/* -------- customEvent -------- */

type customEvent struct {
	event
}

func NewCustomEvent(eventType string, options ...EventOption) CustomEvent {
	e := &customEvent{newEvent(eventType)}
	for _, o := range options {
		o.updateEvent(&e.event)
	}
	return e
}

/* -------- errorEvent -------- */

type errorEvent struct {
	event
	err error
}

func NewErrorEvent(err error) ErrorEvent {
	return &errorEvent{newEvent("error"), err}
}

func (e *errorEvent) Err() error    { return e.err }
func (e *errorEvent) Error() string { return e.err.Error() }

/* -------- EventHandler -------- */

// EventHandler is the interface for an event handler. In JavaScript; an event
// handler can be a function; or an object with a `handleEvent` function. In Go
// code, you can provide your own implementation, or use [NewEventHandlerFunc]
// to create a valid handler from a function.
//
// Multiple EventHandler instances can represent the same underlying event
// handler. E.g., when JavaScript code calls RemoveEventListener, a new Go
// struct is created wrapping the same underlying handler.
//
// The Equals function must return true when the other event handler is the same
// as the current value, so event handlers can properly be removed, and avoiding
// duplicates are added by AddEventListener.
type EventHandler interface {
	// HandleEvent is called when the the event occurrs.
	//
	// An non-nil error return value will dispatch an error event on the global
	// object in a normally configured environment.
	HandleEvent(event Event) error
	// Equals must return true, if they underlying event handler of the other
	// handler is the same as this handler.
	Equals(other EventHandler) bool
}

type HandlerFuncWithoutError = func(Event)
type HandlerFunc = func(Event) error

type eventHandlerFunc struct {
	handlerFunc func(Event) error
	id          entity.ObjectId
}

// NewEventHandlerFunc creates an [EventHandler] implementation from a compatible
// function.
//
// Note: Calling this twice for the same Go-function will be treated as
// different event handlers. Be sure to use the same instance returned from this
// function when removing.
func NewEventHandlerFunc(handler HandlerFunc) EventHandler {
	return eventHandlerFunc{handler, entity.NewObjectId()}
}

// NoError takes a function accepting a single argument and has no return value,
// and transforms it into a function that can be used where an error return
// value is expected.
func NoError[T func(U), U any](f T) func(U) error {
	return func(u U) error {
		f(u)
		return nil
	}
}

func NewEventHandlerFuncWithoutError(handler HandlerFuncWithoutError) EventHandler {
	return eventHandlerFunc{func(event Event) error {
		handler(event)
		return nil
	}, entity.NewObjectId()}
}

func (e eventHandlerFunc) HandleEvent(event Event) error {
	return e.handlerFunc(event)
}

func (e eventHandlerFunc) Equals(handler EventHandler) bool {
	x, ok := handler.(entity.Entity)
	return ok && x.ObjectId() == e.id
}

// ObjectId makes the eventHandlerFunc type implement the Entity interface.
// While the code will still compile without this function; equality check will
// fail. This will be caught by tests verifying EventTarget behaviour.
func (e eventHandlerFunc) ObjectId() entity.ObjectId { return e.id }
