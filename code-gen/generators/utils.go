package generators

import (
	"regexp"
	"strings"
)

var findLastWordRegExp *regexp.Regexp = regexp.MustCompile("([A-Z])[a-z]*$")

func DefaultReceiverName(typeName string) string {
	matches := findLastWordRegExp.FindStringSubmatch(typeName)
	if matches == nil || len(matches) < 2 {
		return ""
	}
	return strings.ToLower(matches[1])
}
