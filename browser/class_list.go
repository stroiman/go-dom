package browser

import (
	"slices"
	"strings"
)

type DOMTokenList interface {
	Add(...string) error
}

type classList struct {
	element Element
}

func NewClassList(element Element) DOMTokenList {
	return &classList{element}
}

func (l *classList) Add(tokens ...string) error {
	var classes []string
	class := l.element.GetAttribute("class")
	if class != "" {
		classes = strings.Split(class, " ")
	}
	for _, token := range tokens {
		if token == "" {
			return newDomErrorCode("Empty token", domErrorSyntaxError)
		}
		if strings.Contains(token, " ") {
			return newDomErrorCode("Empty token", domErrorInvalidCharacter)
		}
		if !slices.Contains(classes, token) {
			classes = append(classes, token)
		}
	}
	l.element.SetAttribute("class", strings.Join(classes, " "))
	return nil
}
