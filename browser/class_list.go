package browser

import (
	"slices"
	"strings"
)

type DOMTokenList struct {
	element Element
}

func NewClassList(element Element) DOMTokenList {
	return DOMTokenList{element}
}

func (l DOMTokenList) Add(tokens ...string) error {
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

func (l DOMTokenList) Length() int {
	class := l.element.GetAttribute("class")
	if class == "" {
		return 0
	}
	return len(strings.Split(class, " "))
}

func (l DOMTokenList) GetValue() string {
	return l.element.GetAttribute("class")
}

func (l DOMTokenList) SetValue(val string) {
	l.element.SetAttribute("class", val)
}
