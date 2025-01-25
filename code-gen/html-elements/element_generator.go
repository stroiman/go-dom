package htmlelements

import (
	"fmt"
	"log/slog"
	"unicode"

	"github.com/dave/jennifer/jen"
	g "github.com/stroiman/go-dom/code-gen/generators"
)

func upperCaseFirstLetter(s string) string {
	strLen := len(s)
	if strLen == 0 {
		slog.Warn("Passing empty string to upperCaseFirstLetter")
		return ""
	}
	buffer := make([]rune, 0, strLen)
	buffer = append(buffer, unicode.ToUpper([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}

type Receiver struct {
	Name g.Generator
	Type g.Generator
}

type IDLAttribute struct {
	AttributeName string
	Receiver      Receiver
	ReadOnly      bool
}

func (a IDLAttribute) Generate() *jen.Statement {
	attrType := g.NewType("string")
	field := g.ValueOf(a.Receiver.Name).Field(a.AttributeName)
	getter := g.FunctionDefinition{
		Receiver: g.FunctionArgument(a.Receiver),
		Name:     upperCaseFirstLetter(a.AttributeName),
		RtnTypes: g.List(attrType),
		Body:     g.Return(field),
	}
	l := g.StatementList(
		getter,
		g.Line,
	)
	if !a.ReadOnly {
		argument := g.NewValue("val")
		l.Append(g.FunctionDefinition{
			Receiver: getter.Receiver,
			Name:     fmt.Sprintf("Set%s", getter.Name),
			Args:     g.Arg(argument, attrType),
			Body:     g.Reassign(field, argument),
		})
	}
	return l.Generate()
}

func GenerateHtmlAnchor() g.Generator {
	return IDLAttribute{
		AttributeName: "target",
		Receiver: Receiver{
			Name: g.NewValue("e"),
			Type: g.NewType("htmlAnchorElement").Pointer(),
		},
	}
}
