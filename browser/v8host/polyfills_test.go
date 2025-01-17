package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Polyfills", func() {
	It("Should have a URL", func() {
		c := NewTestContext()
		Expect(c.RunTestScript("typeof window.URL")).To(Equal("function"))
	})
})
