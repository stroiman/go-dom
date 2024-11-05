package go_dom_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoDom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoDom Suite")
}
