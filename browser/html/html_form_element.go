package html

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
)

type GetReader interface {
	GetReader() io.Reader
}

type HTMLFormElement interface {
	HTMLElement
	GetMethod() string
	SetMethod(value string)
	Submit() error
}

type htmlFormElement struct {
	*htmlElement
}

func NewHtmlFormElement(ownerDocument HTMLDocument) HTMLFormElement {
	result := &htmlFormElement{
		newHTMLElement("form", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlFormElement) Submit() error {
	inputs, err := e.QuerySelectorAll("input")
	if err != nil {
		return err
	}
	formData := dom.NewFormData()
	for _, input := range inputs.All() {
		if inputElement, ok := input.(HTMLInputElement); ok {
			name := inputElement.GetAttribute("name")
			value := inputElement.GetAttribute("value")
			formData.Append(name, dom.NewFormDataValueString(value))
		}
	}
	window := e.htmlDocument.getWindow()
	getReader := GetReader(formData)
	reader := getReader.GetReader()
	method := e.GetAttribute("method")
	if method == "" {
		method = "GET"
	}
	action := e.GetAttribute("action")
	var target dom.URL = window.Location()
	if action != "" {
		var err error
		if target, err = dom.NewUrlBase(action, window.Location().GetHref()); err != nil {
			return err
		}
	}
	targetURL := target.GetHref()
	if method == "GET" {
		searchParams := formData.QueryString()
		targetURL = replaceSearchParams(target, searchParams)
	}
	req, err := http.NewRequest(method, targetURL, reader)
	if err != nil {
		return err
	}
	return window.fetchRequest(req)
}

func (e *htmlFormElement) GetMethod() string {
	if strings.ToLower(e.GetAttribute("method")) == "post" {
		return "post"
	} else {
		return "get"
	}
}

func (e *htmlFormElement) SetMethod(value string) {
	e.SetAttribute("method", value)
}

func replaceSearchParams(location dom.URL, searchParams string) string {
	if searchParams == "" {
		return fmt.Sprintf("%s%s", location.Origin(), location.GetPathname())
	} else {
		return fmt.Sprintf("%s%s?%s", location.Origin(), location.GetPathname(), searchParams)
	}
}
