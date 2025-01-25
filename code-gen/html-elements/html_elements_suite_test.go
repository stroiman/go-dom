package htmlelements_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHtmlElements(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HtmlElements Suite")
}
