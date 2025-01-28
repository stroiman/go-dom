package generators_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	. "github.com/onsi/gomega"
	. "github.com/gost-dom/browser/code-gen/generators"
)

func generatePackage(packageName string, generator Generator) (string, error) {
	file := jen.NewFilePath(packageName)
	file.Add(generator.Generate())
	builder := new(strings.Builder)
	err := file.Render(builder)
	return builder.String(), err
}

func render(generator Generator) (string, error) {
	builder := new(strings.Builder)
	err := generator.Generate().Render(builder)
	return builder.String(), err
}

func TestStructGenerator(t *testing.T) {
	expect := NewWithT(t).Expect
	s := NewStruct(Id("FooBar"))
	expect(render(s)).To(Equal("type FooBar struct{}"))

	s2 := NewStruct(Id("StructWithFields"))
	s2.Field(Id("Name"), Id("string"))
	s2.Field(Id("Age"), Id("int"))
	expect(render(s2)).To(Equal(
		`type StructWithFields struct {
	Name string
	Age  int
}`))
	s3 := NewStruct(Id("StructWithEmbeds"))
	s3.Embed(Id("EmbeddedType1"))
	s3.Embed(Id("EmbeddedType2"))
	s3.Field(Id("StringValue"), Id("string"))
	s3.Field(Id("IntValue"), Id("int"))
	expect(render(s3)).To(Equal(
		`type StructWithEmbeds struct {
	EmbeddedType1
	EmbeddedType2
	StringValue string
	IntValue    int
}`))
}

func TestStructMethodGenerators(t *testing.T) {
	expect := NewWithT(t).Expect
	s := NewStruct(Id("StructWithMembers"))
	s.SetDefaultReceiver("rec")
	foo := s.PointerMethodName("Foo").AddArgument(FunctionArgument{
		Name: Id("str"),
		Type: Id("string"),
	}).WithReturnValue(Id("string")).WithBody(Return(Lit("Foo")))
	expect(render(foo)).To(Equal(`func (rec *StructWithMembers) Foo(str string) string {
	return "Foo"
}`))
}
