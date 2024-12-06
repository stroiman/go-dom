package browser

import (
	"slices"
)

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

func (d *FormData) Append(name string, value string) {
	d.Entries = append(d.Entries, FormDataEntry{name, value})
}

type Predicate[T any] func(T) bool

func elementByName(name string) Predicate[FormDataEntry] {
	return func(e FormDataEntry) bool { return e.Name == name }
}

func (d *FormData) Set(name string, value string) {
	predicate := elementByName(name)
	i := slices.IndexFunc(d.Entries, predicate)
	if i == -1 {
		d.Append(name, value)
		return
	} else {
		d.Delete(name)
		d.Entries = slices.Insert(d.Entries, i, FormDataEntry{
			Name:  name,
			Value: value,
		})
	}
}

func (d *FormData) Keys() []string {
	result := make([]string, len(d.Entries))
	for i, e := range d.Entries {
		result[i] = e.Name
	}
	return result
}

func (d *FormData) Delete(name string) {
	d.Entries = slices.DeleteFunc(
		d.Entries,
		elementByName(name),
	)
}
