package scripting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/scripting"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *ScriptHost

func init() {
	BeforeSuite(func() {
		host = NewScriptHost()
	})

	AfterSuite(func() {
		host.Dispose()
	})
}
