package browser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/browser"
)

var _ = Describe("EventTarget", func() {
	It("Should have a handler attached", func() {
		callCount := 0
		eventHandler := func(e Event) { callCount++ }
		target := NewEventTarget()
		target.AddEventListener("dummy", EventHandlerFunc(eventHandler))
		target.DispatchEvent(nil)
		Expect(callCount).To(Equal(1))
	})
})
