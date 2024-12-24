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
	classes := l.getTokens()
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
	l.setTokens(classes)
	return nil
}

func (l DOMTokenList) Contains(token string) bool {
	classes := l.getTokens()
	return slices.Contains(classes, token)
}

func (l DOMTokenList) Length() int {
	return len(l.getTokens())
}

func (l DOMTokenList) GetValue() string {
	return l.element.GetAttribute("class")
}

func (l DOMTokenList) SetValue(val string) {
	l.element.SetAttribute("class", val)
}

func (l DOMTokenList) Item(index int) *string {
	classes := l.getTokens()
	if index >= len(classes) {
		return nil
	}
	return &classes[index]
}

func (l DOMTokenList) Remove(token string) {
	tokens := l.getTokens()
	itemIndex := slices.Index(tokens, token)
	if itemIndex >= 0 {
		newList := slices.Delete(tokens, itemIndex, itemIndex+1)
		l.setTokens(newList)
	}
}

func (l DOMTokenList) getTokens() []string {
	class := l.element.GetAttribute("class")
	if class == "" {
		return []string{}
	}
	return strings.Split(class, " ")
}

func (l DOMTokenList) setTokens(tokens []string) {
	l.element.SetAttribute("class", strings.Join(tokens, " "))
}
