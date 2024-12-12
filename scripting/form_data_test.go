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
	It("Returns keys", func() {
		c := NewTestContext()
		Expect(c.RunTestScript(`
			data = new FormData();
			data.append("key1", "value");
			data.append("key2", "value");
			Array.from(data.keys()).join(",")
			`)).To(Equal("key1,key2"))
	})

	It("Returns entries", func() {
		c := NewTestContext()
		Expect(c.RunTestScript(`
			data = new FormData();
			data.append("key1", "value1");
			data.append("key2", "value2");
			Array.from(data.entries()).map(x => x.join(";")).join(",")
			`)).To(Equal("key1;value1,key2;value2"))
	})

	It("Should support forEach", func() {
		c := NewTestContext()
		Expect(c.RunTestScript(`
			const result = [];
			data = new FormData();
			data.append("key1", "value1");
			data.append("key2", "value2");
			data.forEach(([k,v]) => { result.push(k + ": " + v) })
			result.join(", ")
			`)).To(Equal("key1: value1, key2: value2"))
	})

	It("Implements iterable", func() {
		c := NewTestContext()
		Expect(c.RunTestScript(`
			data = new FormData();
			typeof data[Symbol.iterator]`)).To(Equal("function"))
	})
	It("Is itself iterable entries", func() {
		c := NewTestContext()
		Expect(c.RunTestScript(`
			data = new FormData();
			data.append("key1", "value1");
			data.append("key2", "value2");
			Array.from(data).map(x => x.join(";")).join(",")
			`)).To(Equal("key1;value1,key2;value2"))
	})
})
