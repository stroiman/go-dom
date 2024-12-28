package dom_test

import (
	"log/slog"
	"testing"

	"github.com/stroiman/go-dom/internal/test"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func init() {
	var logLevel = test.InstallDefaultTextLogger()
	logLevel.Set(slog.LevelInfo)
	// logLevel.Set(slog.LevelDebug)
}

func TestDomTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Browser Suite")
}
