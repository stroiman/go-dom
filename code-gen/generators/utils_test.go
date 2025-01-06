package generators_test

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/code-gen/generators"
)

func TestGenerateReceiverName(t *testing.T) {
	expect := gomega.NewWithT(t).Expect
	expect(DefaultReceiverName("HTMLFormElement")).To(Equal("e"))
}
