package scripting_test

import (
	"log/slog"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/dom"
	domHTTP "github.com/stroiman/go-dom/browser/internal/http"
	"github.com/stroiman/go-dom/browser/internal/test"
	"github.com/stroiman/go-dom/browser/scripting"
	. "github.com/stroiman/go-dom/browser/scripting"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *ScriptHost

func OpenTestWindowFromHandler(location string, handler http.Handler) (dom.Window, error) {
	win, err := dom.OpenWindowFromLocation(location, dom.WindowOptions{
		ScriptEngineFactory: (*scripting.Wrapper)(host),
		HttpClient:          domHTTP.NewHttpClientFromHandler(handler),
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

	BeforeSuite(func() {
		host = NewScriptHost()
	})

	AfterSuite(func() {
		host.Dispose()
	})
}
