package v8host_test

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/http"
	"github.com/gost-dom/browser/internal/test"
	suite "github.com/gost-dom/browser/internal/test/script-test-suite"
	"github.com/gost-dom/browser/logger"
	. "github.com/gost-dom/browser/scripting/v8host"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *V8ScriptHost

var scriptTestSuite *suite.ScriptTestSuite

func OpenTestWindowFromHandler(location string, handler http.Handler) (html.Window, error) {
	win, err := html.OpenWindowFromLocation(location, html.WindowOptions{
		ScriptHost: host,
		HttpClient: NewHttpClientFromHandler(handler),
	})
	DeferCleanup(func() {
		if win != nil {
			win.Close()
		}
	})
	return win, err
}

func init() {
	logger.SetDefault(test.CreateTestLogger(slog.LevelWarn))

	host = New()
	scriptTestSuite = suite.NewScriptTestSuite(host, "v8")
	scriptTestSuite.CreateAllGinkgoTests()

	BeforeSuite(func() {
	})

	AfterSuite(func() {
		host.Close()
	})
}
