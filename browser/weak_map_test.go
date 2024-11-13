package browser_test

import (
	"fmt"
	"runtime"
	"time"

	. "github.com/stroiman/go-dom/browser"
	"golang.org/x/net/html"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// type Wrapped struct {
// 	dummy int // Apparently two pointers compare equal if the struct has no data
// }

type Wrapper struct {
	Inner Node
}

func createAndWait(c chan bool) {
	v := new(Dummy)
	runtime.SetFinalizer(v, func(z *Dummy) {
		fmt.Println("Finalizer run")
		c <- true
	})
}

func GCAndWaitForFinalizer() {
	fin := make(chan bool)
	createAndWait(fin)
	runtime.GC()
	runtime.GC()
	select {
	case <-fin:
	case <-time.After(4 * time.Second):
		panic("finalizer of next string in memory didn't run")
	}
}

func NewWrapped() Node                 { return NewElement("div", &html.Node{}) }
func NewWrapper(wrapped Node) *Wrapper { return &Wrapper{wrapped} }

var _ = Describe("WeakMap", Focus, func() {
	var (
		weakMap *NodeMap[*Wrapper]
	)

	AddWrapped := func() Node {
		wrapped := NewWrapped()
		wrapper := NewWrapper(wrapped)
		// fmt.Printf("Inserting: %v\n", wrapped)
		// fmt.Printf("No of entries before insert: %d\n", weakMap.Length())
		weakMap.Set(wrapped, wrapper)
		// fmt.Printf("No of entries after insert: %d\n", weakMap.Length())
		return wrapped
	}

	BeforeEach(func() {
		weakMap = NewNodeMap[*Wrapper]()
	})

	It("Should be empty initially", func() {
		Expect(weakMap.Length()).To(Equal(0))
	})

	It("Should have two entries when they are kept alive", func() {
		a := AddWrapped()
		Expect(weakMap.Length()).To(Equal(1), "After first add")
		b := AddWrapped()
		count := weakMap.Length()
		runtime.KeepAlive(a)
		runtime.KeepAlive(b)
		Expect(count).To(Equal(2))
	})

	It("Should have two entries when they are kept alive", Focus, func() {
		a := AddWrapped()
		Expect(weakMap.Length()).To(Equal(1), "After first add")
		AddWrapped()
		GCAndWaitForFinalizer()
		GCAndWaitForFinalizer()
		runtime.GC()
		runtime.KeepAlive(a)
		count := weakMap.Length()
		Expect(count).To(Equal(1))
	})
})
