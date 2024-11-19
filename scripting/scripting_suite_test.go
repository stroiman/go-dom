package scripting_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser"
	. "github.com/stroiman/go-dom/scripting"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *ScriptHost
var mainBrowser browser.Browser

func NewTestBrowserFromHandler(handler http.Handler) browser.Browser {
	result := browser.NewBrowserFromHandler(handler)
	result.ScriptEngineFactory = (*Wrapper)(host)
	return result
}

func init() {
	BeforeSuite(func() {
		host = NewScriptHost()
		mainBrowser = browser.NewBrowserFromHandler(nil)
		mainBrowser.ScriptEngineFactory = (*Wrapper)(host)
	})

	AfterSuite(func() {
		host.Dispose()
	})
}
