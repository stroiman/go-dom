package htmlelements

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
)

func writeFile(s FileGeneratorSpec) error {
	jf := jen.NewFilePath(s.Package)
	jf.HeaderComment("This file is generated. Do not edit.")
	jf.Add(s.Generator.Generate())
	outputFileName := fmt.Sprintf("%s_generated.go", s.Name)
	if writer, err := os.Create(outputFileName); err != nil {
		return err
	} else {
		defer writer.Close()
		if err = jf.Render(writer); err != nil {
			return err
		}
	}
	return nil
}

func GenerateHTMLElements() error {
	files, err := CreateHTMLElementGenerators()
	if err != nil {
		return err
	}
	for _, f := range files {
		if err = writeFile(f); err != nil {
			return err
		}
	}
	return nil
}

func GenerateDOMTypes() error {
	files, err := CreateDOMGenerators()
	if err != nil {
		return err
	}
	for _, f := range files {
		if err = writeFile(f); err != nil {
			return err
		}
	}
	return nil
}
