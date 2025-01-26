package elements

import (
	"encoding/json"

	"github.com/stroiman/go-dom/code-gen/idl"
)

type ElementJSON struct {
	Name      string `json:"name"`
	Interface string `json:"interface"`
}

type ElementsJSON struct {
	Elements []ElementJSON `json:"elements"`
}

type Elements ElementsJSON

func Load() (Elements, error) {
	output := ElementsJSON{}
	err := json.Unmarshal(idl.Html_defs, &output)
	return Elements(output), err
}
