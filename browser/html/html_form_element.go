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
	GetAction() string
	SetAction(val string)
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
		return err // Shouldn't happen, only an invalid DOM string should generate this error
	}
	formData := NewFormData()
	for _, input := range inputs.All() {
		if inputElement, ok := input.(HTMLInputElement); ok {
			name := inputElement.GetAttribute("name")
			value := inputElement.GetAttribute("value")
			formData.Append(name, NewFormDataValueString(value))
		}
	}

	var req *http.Request
	if e.GetMethod() == "get" {
		searchParams := formData.QueryString()
		targetURL := replaceSearchParams(e.getAction(), searchParams)
		req, err = http.NewRequest("GET", targetURL, nil)
	} else {
		getReader := GetReader(formData)
		req, err = http.NewRequest("POST", e.GetAction(), getReader.GetReader())
	}
	if err != nil {
		return err
	}
	return e.htmlDocument.getWindow().fetchRequest(req)
}

func (e *htmlFormElement) GetMethod() string {
	if strings.ToLower(e.GetAttribute("method")) == "post" {
		return "post"
	} else {
		return "get"
	}
}

func (e *htmlFormElement) SetAction(val string) { e.SetAttribute("action", val) }

func (e *htmlFormElement) getAction() dom.URL {
	window := e.getWindow()
	action := e.GetAttribute("action")
	target := dom.URL(window.Location())
	var err error
	if action != "" {
		if target, err = dom.NewUrlBase(action, window.Location().GetHref()); err != nil {
			// This _shouldn't_ happen. But let's refactor code, so err isn't a
			// possible return value
			panic(err)
		}
	}
	return target
}
func (e *htmlFormElement) GetAction() string {
	return e.getAction().GetHref()
	// window := e.getWindow()
	// action := e.GetAttribute("action")
	// target := dom.URL(window.Location())
	// var err error
	// if action != "" {
	// 	if target, err = dom.NewUrlBase(action, window.Location().GetHref()); err != nil {
	// 		// This _shouldn't_ happen. But let's refactor code, so err isn't a
	// 		// possible return value
	// 		panic(err)
	// 	}
	// }
	// return target.GetHref()
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
