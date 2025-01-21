package dom

import (
	"log/slog"
)

func logElement(e Element) slog.Attr {
	if e == nil {
		return slog.String("element", "nil")
	}
	return slog.Group("element", "tagName", e.TagName(), "outerHTML", e.OuterHTML())
}
