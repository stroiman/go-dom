package html

import (
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
	action := e.GetAttribute("action")
	req, err := http.NewRequest(method, action, reader)
	if err != nil {
		return err
	}
	return window.fetchRequest(req)
}
