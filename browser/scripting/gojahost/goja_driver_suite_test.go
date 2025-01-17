package gojahost_test

import (
	"log/slog"
	"testing"

	"github.com/dop251/goja"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	. "github.com/stroiman/go-dom/browser/scripting/gojahost"
	"github.com/stroiman/go-dom/browser/html"
	"github.com/stroiman/go-dom/browser/internal/test"
	. "github.com/stroiman/go-dom/browser/internal/test/script-test-suite"
)

var testSuite = NewScriptTestSuite(New(), "goja", SkipDOM)

func newCtx() html.ScriptContext { return testSuite.NewContext() }

func FormatException(value any) (result string, ok bool) {
	var exception *goja.Exception
	if exception, ok = value.(*goja.Exception); ok {
		result = exception.String()
	}
	return
}

func init() {
	testSuite.CreateAllGinkgoTests()
	var logLevel = test.InstallDefaultTextLogger()
	logLevel.Set(slog.LevelInfo)
	// logLevel.Set(slog.LevelDebug)

	format.RegisterCustomFormatter(FormatException)
}

func TestGojaDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GojaDriver Suite")
}
