package browser

type ObjectId = int32

var idSeq <-chan ObjectId

func init() {
	c := make(chan ObjectId)
	idSeq = c
	go func() {
		var val ObjectId = 1
		for {
			c <- val
			val = val + 1
		}
	}()
}

func NewObjectId() ObjectId {
	return <-idSeq
}

type EventTarget interface {
	// ObjectId is used internally for the scripting engine to associate a v8
	// object with the Go object it wraps.
	ObjectId() ObjectId
}

type eventTarget struct {
	objectId ObjectId
}

func newEventTarget() eventTarget {
	id := NewObjectId()
	return eventTarget{id}
}

func (e *eventTarget) ObjectId() ObjectId { return e.objectId }
