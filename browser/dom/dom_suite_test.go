package dom_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/gost-dom/browser/browser/internal/test"
	"github.com/gost-dom/browser/browser/logger"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func init() {
	fmt.Println("Set debug level")
	logger.SetDefault(test.CreateTestLogger(slog.LevelInfo))
}

func TestDomTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Browser Suite")
}
