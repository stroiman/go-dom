package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

var _ = Describe("FormData", func() {
	It("Should be empty when new", func() {
		Expect(NewFormData()).To(BeEmptyFormData())
	})

	Describe("Append", func() {
		It("Should allow multiple values with the same key", func() {
			formData := NewFormData()
			formData.Append("Key1", "Value1")
			formData.Append("Key2", "Value2")
			formData.Append("Key1", "Value3")
			Expect(formData).To(HaveEntries("Key1", "Value1", "Key2", "Value2", "Key1", "Value3"))
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
			Value: entries[j+1],
		}
	}
	return WithTransform(func(data *FormData) []FormDataEntry { return data.Entries },
		ContainElements(expected),
	)
}
