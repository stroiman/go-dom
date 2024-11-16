package browser

}

}

type EventTarget interface {
	// ObjectId is used internally for the scripting engine to associate a v8
	// object with the Go object it wraps.
	ObjectId() ObjectId
}

type eventTarget struct {
	base
}

func newEventTarget() eventTarget {
	return eventTarget{newBase()}
}

