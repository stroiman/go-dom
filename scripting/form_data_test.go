package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 FormData", func() {
	It("Should inherit from Object", func() {
		c := NewTestContext()
		Expect(
			c.RunTestScript("Object.getPrototypeOf(FormData.prototype) === Object.prototype"),
		).To(BeTrue())
	})

	It("Allows adding/getting", func() {
		c := NewTestContext()
		Expect(c.RunTestScript(`
			data = new FormData();
			data.append("key", "value");
			data.get("key");
			`)).To(Equal("value"))
	})
})
