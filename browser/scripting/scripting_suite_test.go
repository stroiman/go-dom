package scripting_test

import (
	"log/slog"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/internal/http"
	"github.com/stroiman/go-dom/browser/internal/test"
	suite "github.com/stroiman/go-dom/browser/internal/test/script-test-suite"
	"github.com/stroiman/go-dom/browser/scripting"
	. "github.com/stroiman/go-dom/browser/scripting"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *ScriptHost

var scriptTestSuite *suite.ScriptTestSuite

func OpenTestWindowFromHandler(location string, handler http.Handler) (html.Window, error) {
	win, err := html.OpenWindowFromLocation(location, html.WindowOptions{
		ScriptEngineFactory: (*scripting.Wrapper)(host),
		HttpClient:          NewHttpClientFromHandler(handler),
	})
	DeferCleanup(func() {
		if win != nil {
			win.Dispose()
		}
	})
	return win, err
}

func init() {
	var logLevel = test.InstallDefaultTextLogger()
	logLevel.Set(slog.LevelInfo)
	// logLevel.Set(slog.LevelDebug)

	host = NewScriptHost()
	scriptTestSuite = suite.NewScriptTestSuite((*scripting.Wrapper)(host), "v8")
	scriptTestSuite.CreateAllGinkgoTests()

	BeforeSuite(func() {
	})

	AfterSuite(func() {
		host.Dispose()
	})
}
