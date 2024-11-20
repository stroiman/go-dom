package browser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/browser"
)

var _ = Describe("EventTarget", func() {
	It("Should call a handler of the right type", func() {
		callCount := 0
		handler := NewEventHandlerFuncWithoutError(func(e Event) { callCount++ })
		target := NewEventTarget()
		target.AddEventListener("click", handler)
		target.DispatchEvent(NewCustomEvent("click"))
		Expect(callCount).To(Equal(1))
	})

	It("Should not call a handler of a different type", func() {
		callCount := 0
		handler := NewEventHandlerFuncWithoutError(func(e Event) { callCount++ })
		target := NewEventTarget()
		target.AddEventListener("click", handler)
		target.DispatchEvent(NewCustomEvent("keyDown"))
		Expect(callCount).To(Equal(0))
	})

	It("Should not call a handler that was removed type", func() {
		callCount := 0
		handler := NewEventHandlerFuncWithoutError(func(e Event) { callCount++ })
		target := NewEventTarget()
		target.AddEventListener("click", handler)
		target.RemoveEventListener("click", handler)
		target.DispatchEvent(NewCustomEvent("click"))
		Expect(callCount).To(Equal(0))
	})

	It("Should only call a handler once, even if added twice", func() {
		callCount := 0
		handler := NewEventHandlerFuncWithoutError(func(e Event) { callCount++ })
		target := NewEventTarget()
		target.AddEventListener("click", handler)
		target.AddEventListener("click", handler)
		target.DispatchEvent(NewCustomEvent("click"))
		Expect(callCount).To(Equal(1))
	})

	It("Should call the event listeners in the order they are added", func() {
		// This is a very silly test, and there are probably many _wrong_
		// implementations that would by accident make this test pass. The test
		// mostly exists here to document required behaviour.
		callCount := 0
		eventHandler := func(e Event) { callCount++ }
		target := NewEventTarget()
		NewEventHandlerFuncWithoutError(eventHandler)
		target.AddEventListener("click", NewEventHandlerFuncWithoutError(func(e Event) {
			Expect(callCount).To(Equal(0), "First handler")
			callCount++
		}))
		target.AddEventListener("click", NewEventHandlerFuncWithoutError(func(e Event) {
			Expect(callCount).To(Equal(1), "Second handler")
			callCount++
		}))
		target.DispatchEvent(NewCustomEvent("click"))
		Expect(callCount).To(Equal(2), "Final state")
	})
})
