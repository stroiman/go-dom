package html

import (
	"io"
	"slices"
)

type FormDataValue string // TODO Blob/file

func NewFormDataValueString(value string) FormDataValue { return FormDataValue(value) }

type FormDataEntry struct {
	Name  string
	Value FormDataValue
}

type FormData struct {
	Entries []FormDataEntry
}

func NewFormData() *FormData {
	return &FormData{nil}
}

func (d *FormData) Append(name string, value FormDataValue) {
	d.Entries = append(d.Entries, FormDataEntry{name, value})
}

type Predicate[T any] func(T) bool

func elementByName(name string) Predicate[FormDataEntry] {
	return func(e FormDataEntry) bool { return e.Name == name }
}

func (d *FormData) Set(name string, value FormDataValue) {
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

func (d *FormData) Values() []FormDataValue {
	result := make([]FormDataValue, len(d.Entries))
	for i, e := range d.Entries {
		result[i] = e.Value
	}
	return result
}

func (d *FormData) Delete(name string) {
	d.Entries = slices.DeleteFunc(
		d.Entries,
		elementByName(name),
	)
}

func (d *FormData) Get(name string) FormDataValue {
	for _, e := range d.Entries {
		if e.Name == name {
			return e.Value
		}
	}
	return ""
}

func (d *FormData) GetAll(name string) []FormDataValue {
	var result []FormDataValue
	for _, e := range d.Entries {
		if e.Name == name {
			result = append(result, e.Value)
		}
	}
	return result
}

func (d *FormData) Has(name string) bool {
	return slices.IndexFunc(d.Entries, elementByName(name)) != -1
}

func (d *FormData) GetReader() io.Reader {
	data := NewXHRRequestBodyOfFormData(d)
	return data.getReader()
}

// QueryString returns the formdata as a &-separated URL encoded key-value pair.
func (d *FormData) QueryString() string {
	data := NewXHRRequestBodyOfFormData(d)
	return data.string()
}
