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

type base struct {
	objectId ObjectId
}

func newBase() base {
	return base{NewObjectId()}
}

func (b base) ObjectId() ObjectId { return b.objectId }
