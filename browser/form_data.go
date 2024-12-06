package browser

type FormDataEntry struct {
	Name  string
	Value string // TODO Blob/file
	// TODO FileName string
}

type FormData struct {
	Entries []FormDataEntry
}

func NewFormData() *FormData {
	return &FormData{}
}

func (d *FormData) Append(key string, value string) {
	d.Entries = append(d.Entries, FormDataEntry{key, value})
}
