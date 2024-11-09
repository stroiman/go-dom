package browser

type Window interface {
	Document() Document
	// TODO: Remove, for testing
	LoadHTML(string)
}

type window struct {
	document Document
}

func NewWindow() Window {
	return &window{
		NewDocument(),
	}
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) {
	w.document = ParseHtmlString(html)
}
