package html

import (
	"bytes"
	"io"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/log"
)

type htmlScriptElement struct{ *htmlElement }

func NewHTMLScriptElement(ownerDocument HTMLDocument) HTMLElement {
	result := &htmlScriptElement{newHTMLElement("script", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlScriptElement) Connected() {
	var script = ""
	if src, hasSrc := e.GetAttribute("src"); !hasSrc {
		script = e.TextContent()
	} else {
		window, _ := e.htmlDocument.getWindow().(*window)
		resp, err := window.httpClient.Get(src)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			log.Error("Error from server", "body", string(body), "src", src)
			panic("Bad response")
		}

		buf := bytes.NewBuffer([]byte{})
		buf.ReadFrom(resp.Body)
		script = string(buf.Bytes())

	}
	e.window().Run(script)
}

func (e *htmlScriptElement) AppendChild(n dom.Node) (dom.Node, error) {
	return e.htmlElement.AppendChild(n)
}
