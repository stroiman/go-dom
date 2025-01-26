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
	receiver := g.ValueOf(a.Receiver.Name)
	result := g.Id("result")
	getter := g.FunctionDefinition{
		Receiver: g.FunctionArgument(a.Receiver),
		Name:     upperCaseFirstLetter(a.AttributeName),
		RtnTypes: g.List(attrType),
		Body: g.StatementList(
			g.AssignMany(
				g.List(result, g.Id("_")),
				receiver.Field("GetAttribute").Call(g.Lit(a.AttributeName)),
			),
			g.Return(result),
		),
	}
	l := g.StatementList(
		getter,
	)
	if !a.ReadOnly {
		argument := g.NewValue("val")
		l.Append(
			g.Line,
			g.FunctionDefinition{
				Receiver: getter.Receiver,
				Name:     fmt.Sprintf("Set%s", getter.Name),
				Args:     g.Arg(argument, attrType),
				Body:     receiver.Field("SetAttribute").Call(g.Lit(a.AttributeName), argument),
			})
	}
	return l.Generate()
}
