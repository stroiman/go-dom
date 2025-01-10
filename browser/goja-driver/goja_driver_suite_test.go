package goja_driver_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser/goja-driver"
	. "github.com/stroiman/go-dom/browser/internal/test/script-test-suite"
)

var testSuite = NewScriptTestSuite(NewGojaScriptEngine(), "goja")

func init() {
	testSuite.CreateAllGinkgoTests()
}

func TestGojaDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GojaDriver Suite")
}
