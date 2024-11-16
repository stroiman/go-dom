package scripting_test

import (
	"github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventTarget", func() {
	ctx := InitializeContext()

	It("Is an EventTarget", func() {
		Expect(ctx.RunTestScript("(new EventTarget()) instanceof EventTarget")).To(BeTrue())
	})

	It("Can call an added event listener", func() {
		Expect(ctx.RunTestScript(`
var callCount = 0
function listener() { callCount++ };
const target = new EventTarget();
target.addEventListener('custom', listener);
target.dispatchEvent(new CustomEvent('custom'));`)).Error().ToNot(HaveOccurred())
		Expect(ctx.RunTestScript("callCount")).To(BeEquivalentTo(1))
	})

	It("Event from Go code will propagate to JS", func() {
		Expect(ctx.RunTestScript(`
var callCount = 0
function listener() { callCount++ };
const target = window;
target.addEventListener('custom', listener);
`)).Error().ToNot(HaveOccurred())
		ctx.Window().DispatchEvent(browser.NewCustomEvent("custom"))
		Expect(ctx.RunTestScript("callCount")).To(BeEquivalentTo(1))
	})
})
