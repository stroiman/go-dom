package html_test

import (
	. "github.com/stroiman/go-dom/browser/html"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

var _ = Describe("FormDataForm", func() {
	It("Should handle empty name/value of input fields", func() {
		Skip("Not yet handled properly")
	})
})

var _ = Describe("FormData", func() {
	It("Should be empty when new", func() {
		Expect(NewFormData()).To(BeEmptyFormData())
	})

	Describe("Multiple values have been appended with the same key", func() {
		var formData *FormData

		BeforeEach(func() {
			formData = NewFormData()
			formData.Append("Key1", "Value1")
			formData.Append("Key2", "Value2")
			formData.Append("Key1", "Value3")
			formData.Append("Key3", "Value4")
		})

		It("Should contain all values", func() {
			Expect(
				formData,
			).To(HaveEntries(
				"Key1", "Value1",
				"Key2", "Value2",
				"Key1", "Value3",
				"Key3", "Value4"))
		})

		It("Delete() should remove all values with the name", func() {
			formData.Delete("Key1")
			Expect(formData).To(HaveEntries(
				"Key2", "Value2",
				"Key3", "Value4"))
		})

		It("Set() should replace all values with the name", func() {
			formData.Set("Key1", "Value5")
			Expect(formData).To(HaveEntries(
				"Key1", "Value5",
				"Key2", "Value2",
				"Key3", "Value4"))

		})

		It("Set() should add a new value when given a new name", func() {
			formData.Set("Key4", "Value5")
			Expect(
				formData,
			).To(HaveEntries(
				"Key1", "Value1",
				"Key2", "Value2",
				"Key1", "Value3",
				"Key3", "Value4",
				"Key4", "Value5"))
		})

		It("Keys() Should return all keys, including duplicates", func() {
			Expect(formData.Keys()).To(HaveExactElements(
				"Key1",
				"Key2",
				"Key1",
				"Key3"))
		})

		It("Values() Should return all values, including duplicates", func() {
			Expect(formData.Values()).To(HaveExactElements(
				BeFormDataStringValue("Value1"),
				BeFormDataStringValue("Value2"),
				BeFormDataStringValue("Value3"),
				BeFormDataStringValue("Value4")))
		})

		It("Get() Should return the first value with the name", func() {
			Expect(formData.Get("Key1")).To(BeEquivalentTo("Value1"))
		})

		It("GetAll() Should return all values with the name", func() {
			Expect(formData.GetAll("Key1")).To(HaveExactElements(
				BeFormDataStringValue("Value1"),
				BeFormDataStringValue("Value3"),
			))
		})

		Describe("Has()", func() {
			It("Should return true for a key that exists", func() {
				Expect(formData.Has("Key1")).To(BeTrue())
			})
			It("Should return false for a key that exists", func() {
				Expect(formData.Has("BadKey")).To(BeFalse())
			})
		})
	})
})

func BeEmptyFormData() types.GomegaMatcher {
	return gcustom.MakeMatcher(func(data *FormData) (bool, error) {
		return len(data.Entries) == 0, nil
	})
}

func HaveEntries(entries ...string) types.GomegaMatcher {
	if len(entries)%2 != 0 {
		panic("Wrong number of entries, must be even")
	}
	noOfEntries := len(entries) / 2
	expected := make([]FormDataEntry, noOfEntries)
	for i := 0; i < noOfEntries; i++ {
		j := i * 2
		expected[i] = FormDataEntry{
			Name:  entries[j],
			Value: FormDataValue(entries[j+1]),
		}
	}
	return WithTransform(
		func(data *FormData) []FormDataEntry { return data.Entries },
		HaveExactElements(expected),
	)
}

func BeFormDataStringValue(value string) types.GomegaMatcher {
	return BeEquivalentTo(value)
}
