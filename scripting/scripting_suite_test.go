package scripting_test

import (
	"log/slog"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser"
	"github.com/stroiman/go-dom/internal/test"
	. "github.com/stroiman/go-dom/scripting"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *ScriptHost

func NewTestBrowserFromHandler(handler http.Handler) *browser.Browser {
	result := browser.NewBrowserFromHandler(handler)
	DeferCleanup(result.Dispose)
	result.ScriptEngineFactory = (*Wrapper)(host)
	return result
}

func init() {
	var logLevel = test.InstallDefaultTextLogger()
	logLevel.Set(slog.LevelInfo)
	// logLevel.Set(slog.LevelDebug)

	BeforeSuite(func() {
		host = NewScriptHost()
	})

	AfterSuite(func() {
		host.Dispose()
	})
}
