package dom_test

import (
	"errors"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
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

	Describe("Event propagation", func() {
		var (
			window html.Window
			target Element
		)

		BeforeEach(func() {
			var err error
			window, err = html.NewWindowReader(
				strings.NewReader(`<body><div id="target"></div></body>`),
			)
			Expect(err).ToNot(HaveOccurred())
			target = window.Document().GetElementById("target")
		})

		It("Should not propagate the event to the parent by default", func() {
			called := false
			var l EventHandler = NewEventHandlerFunc(func(e Event) error {
				called = true
				return nil
			})
			window.Document().Body().AddEventListener("custom", l)
			target.DispatchEvent(NewCustomEvent("custom"))
			Expect(called).To(BeFalse())
		})

		It("Should propagate the event to the window", func() {
			called := false

			// window.Document()
			var l EventHandler = NewEventHandlerFunc(func(e Event) error {
				called = true
				return nil
			})
			window.AddEventListener("custom", l)
			target.DispatchEvent(NewCustomEvent("custom", EventBubbles(true)))
			Expect(called).To(BeTrue())
		})

		It("Should propagate the event to the window if 'bubbles' is set", func() {
			called := false

			// window.Document()
			var l EventHandler = NewEventHandlerFunc(func(e Event) error {
				called = true
				return nil
			})
			window.AddEventListener("custom", l)
			target.DispatchEvent(NewCustomEvent("custom", EventBubbles(true)))
			Expect(called).To(BeTrue())
		})

		It("Should not propagate if the handler calls `StopPropagation()`", func() {
			calledA := false
			calledB := false
			window.Document().
				Body().
				AddEventListener("custom", NewEventHandlerFunc(func(e Event) error {
					calledA = true
					e.StopPropagation()
					return nil
				}))
			window.AddEventListener("custom", NewEventHandlerFunc(func(e Event) error {
				calledB = true
				return nil
			}))
			target.DispatchEvent(NewCustomEvent("custom", EventBubbles(true)))
			Expect(calledA).To(BeTrue(), "Event dispatched on body")
			Expect(calledB).To(BeFalse(), "Event dispatched on window")
		})

		Describe("DispatchEvent return value", func() {
			It("Should return true", func() {
				Expect(target.DispatchEvent(NewCustomEvent("custom"))).To(BeTrue())
			})

			It(
				"Should return false when the handler calls PreventDefault() on a cancelable event",
				func() {
					target.AddEventListener("custom", NewEventHandlerFunc(func(e Event) error {
						e.PreventDefault()
						return nil
					}))
					Expect(
						target.DispatchEvent(NewCustomEvent("custom", EventCancelable(true))),
					).To(BeFalse())
				},
			)

			It(
				"Should return true when the handler calls PreventDefault() on a non-cancelable event",
				func() {
					target.AddEventListener("custom", NewEventHandlerFunc(func(e Event) error {
						e.PreventDefault()
						return nil
					}))
					Expect(target.DispatchEvent(NewCustomEvent("custom"))).To(BeTrue())
				},
			)
		})

		Describe("The event handler generates an error", func() {
			BeforeEach(func() {
				target.AddEventListener("custom", NewEventHandlerFunc(func(e Event) error {
					return errors.New("dummy")
				}))
			})

			It("Should be reported on Window", func() {
				var errorOnWindow bool
				window.AddEventListener("error", NewEventHandlerFunc(func(e Event) error {
					errorOnWindow = true
					return nil
				}))
				target.DispatchEvent(NewCustomEvent("custom"))
				Expect(errorOnWindow).To(BeTrue())
			})

			It("Should not be reported on target", func() {
				var errorOnTarget bool
				target.AddEventListener("error", NewEventHandlerFunc(func(e Event) error {
					errorOnTarget = true
					return nil
				}))
				target.DispatchEvent(NewCustomEvent("custom"))
				Expect(errorOnTarget).To(BeFalse())
			})
		})
	})
})
