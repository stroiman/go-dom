package gojahost_test

import (
	"log/slog"
	"testing"

	"github.com/dop251/goja"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/gost-dom/browser/browser/html"
	"github.com/gost-dom/browser/browser/internal/test"
	. "github.com/gost-dom/browser/browser/internal/test/script-test-suite"
	"github.com/gost-dom/browser/browser/logger"
	. "github.com/gost-dom/browser/browser/scripting/gojahost"
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
	logger.SetDefault(test.CreateTestLogger(slog.LevelInfo))

	format.RegisterCustomFormatter(FormatException)
}

func TestGojaDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GojaDriver Suite")
}
