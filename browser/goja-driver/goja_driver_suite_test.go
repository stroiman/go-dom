package goja_driver_test

import (
	"log/slog"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser/goja-driver"
	"github.com/stroiman/go-dom/browser/internal/test"
	. "github.com/stroiman/go-dom/browser/internal/test/script-test-suite"
)

var testSuite = NewScriptTestSuite(NewGojaScriptEngine(), "goja")

func init() {
	testSuite.CreateAllGinkgoTests()
	var logLevel = test.InstallDefaultTextLogger()
	logLevel.Set(slog.LevelInfo)
	// logLevel.Set(slog.LevelDebug)
}

func TestGojaDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GojaDriver Suite")
}
