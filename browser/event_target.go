package browser

type Event interface {
	Type() string
	Target() EventTarget
}

type CustomEvent interface{}

type EventHandler interface {
	HandleEvent(event Event)
}

type EventHandlerFunc func(Event)

func (l EventHandlerFunc) HandleEvent(e Event) { l(e) }

type EventTarget interface {
	// ObjectId is used internally for the scripting engine to associate a v8
	// object with the Go object it wraps.
	ObjectId() ObjectId
	AddEventListener(eventType string, listener EventHandler /* TODO: options */)
	RemoveEventListener(eventType string, listener EventHandler)
	DispatchEvent(event Event)
}

type eventTarget struct {
	base
	listeners []EventHandler
}

func newEventTarget() eventTarget {
	return eventTarget{
		base: newBase(),
	}
}

func NewEventTarget() EventTarget {
	res := newEventTarget()
	return &res
}

func (e *eventTarget) AddEventListener(eventType string, listener EventHandler) {
	e.listeners = append(e.listeners, listener)
}

func (e *eventTarget) RemoveEventListener(eventType string, listener EventHandler) {}

func (e *eventTarget) DispatchEvent(event Event) {
	for _, l := range e.listeners {
		l.HandleEvent(event)
	}
}
