package idl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIdl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Idl Suite")
}
