package html

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stroiman/go-dom/browser/dom"
)

type GetReader interface {
	GetReader() io.Reader
}

type HTMLFormElement interface {
	HTMLElement
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
	if action == "" {
		window := e.getWindow()
		searchParams := ""
		if method == "GET" {
			searchParams = formData.QueryString()
		}
		action = replaceSearchParams(window.Location(), searchParams)
	} else {
		if u, err := dom.NewUrlBase(action, window.Location().GetHref()); err != nil {
			return err
		} else {
			action = u.GetHref()
		}
	}
	req, err := http.NewRequest(method, action, reader)
	if err != nil {
		return err
	}
	return window.fetchRequest(req)
}

func replaceSearchParams(location dom.Location, searchParams string) string {
	if searchParams == "" {
		return fmt.Sprintf("%s%s", location.Origin(), location.GetPathname())
	} else {
		return fmt.Sprintf("%s%s?%s", location.Origin(), location.GetPathname(), searchParams)
	}
}
