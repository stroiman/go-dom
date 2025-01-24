# Generators

This package is part of the [Go-DOM](https://github.com/stroiman/go-dom) to
support code generation from web specification IDL files.

However, the code is of general use, and exposed publicly.

> [!WARNING]
>
> This package is part of the code generator itself, which is of no interest
> outside Go-DOM, and makes no guarantees about backwards compatibility.
>
> If you find it usefule, I suggest copy/paste the code to your own project.

## Build on top of Jennifer

This library is a wrapper on top of [Jennifer](https://github.com/dave/jennifer)

Jennifer is a very complete code generator, but I did not agree with the design
choices in the interface.

- Jennifer has a more imperative API
- IMHO, the API Jennifer exposes is too low-level

### Declarative vs. Imperative

This type of problem lends itself very well to a declarative composition.
Jennifer however, is designed around mutation.

This library provides a more composition oriented API.

### Level of abstraction

Jennifer has a model centered around the output tokens generated in code, e.g.
there is a `Func` representing the `func` keyword. But the keword has different
uses.

- Declaring a function type
- Declaring a function literal

Likewise, `*` is created using `Op("*")`, but this has mutliple uses as well:

- Declaring a pointer type
- Dereferencing a pointer variable

Index represents `[]`, which can be used to index a slice or map, as well as
providing a type parameter to a generic type of function.

This library uses an abstraction level of "Variable assignment", "Pointer",
"Reference", "Equals", rather than `Op(":=")`, `Op("*")`, `Op("&")`, `Op("==")`.

## Examples

This is not a full documentation, just to get you started.

Everything in this library is a `Generator`, and interface representing the
`Generate` function that can return a `*jen.Statement`.

```go
type Generator interface {
	Generate() *jen.Statement
}
```

### Types and Values

Two types, `Type` and `Value` are simple wrappers on top of `Generator` to
provide easy access to constructs 

E.g., on a `Value`, you can access fields, or call them. On a 

```Go
v := g.NewValue("MyStruct")
return StatementList(
    g.Assign(v, g.NewValue("NewMyStruct").Call()),
    g.Field("Initialize").Call(),
    g.Return(v),
    )
```

### Conditions

Simple `if`/`else` is implemented by `IfStmt`. `Eq` and `Neq` provides `==` and
`!=` support. `Gte`, greater-than or equal, is `>=`

```Go
IfStmt{
    Condition: Eq{ Lhs: value1, Rhs: Value2 },
    Block: someFunctionValue.Call(),
    // Else is optional
    Else: someOtherFunctionValue.Call(),
}
```

### Escape Hatch

Not everything is supported. To deal with that, the `Raw` generator can wrap a
native Jennifer statement. E.g., at the moment, `append` is not supported, so
here `Raw` is used:

```Go
items := NewValue("items")
item := NewValue("item")
functionBody := StatementList(
    Assign(item, NewValue("NewItem").Call()),
    Reassign(items,
        Raw(jen.Append(items.Generate(), item.Generate())),
    ),
)
// Generates
// item := NewItem()
// items = append(items, item)
```

> [!NOTE]
>
> If you create custom support for new constructs, please add them

### Creating a file

This package does not handle the overall file creation, package specification,
and import aliases. So you need to use Jen directly here

```go
func WriteGenerator(g Generator, w io.Writer) (error) {
    // Fully qualified package path
    file := jen.NewFilePath("example.com/my/package")
    file.HeaderComment("This file is generated. Do not edit.")
    // Potentially, create aliases for imports
    file.ImportAlias("github.com/tommie/v8go", "v8")
    file.Add(generator.Generate())
    return file.Render(w)
}
```

## Philosophy: Embrace composition

The philosophy is to be able to compose larger structures out of individual
parts. Each level of composition adds a higher level of abstraction.

As an example

- Compose high level file generators from application specific generators
- Compose application specific generators from high level general purpose generators
- Compose high level general purpose generators from low-level general purpose
generators.

### High-level file generators

At the highest level, compose that parts that need to be in the file:

```Go
func FileContents() g.Generator {
    return StatementList(
        TypeGenerator(),
        ConstructorGenerator(),
        MethodsGenerator(),
        )
}
```

### High-level application specific generators

This would be generators created specifically for the types that exist in your
application:

```go
type MyTypeInstance struct {
    g.Generator
}

func (i MyTypeInstance) CallMethod1(arg1 Generator, arg2 Generator) Generator {
    v := g.Value{i.Generator}
    return v.Field("Method1").Call(arg1, arg2)
}

func Body() g.Generator {
    v := MyTypeInstance{g.NewValue("t")}
    return StatementList {
        Assign(v, g.Value("NewMyType").Call()),
        v.CallMethod1(g.Lit("foo"), g.Lit("bar")),
        g.Return(v),
    }
}
```

### High-level general purpose generators

A high-level general purpose generator could be a "Getter", retrieving a private
read-only field:

```Go
type Getter struct {
    FieldName    string
    FieldType    Generator
    ReceiverName Generator
    ReceiverType Generator
}

func (gg Getter) Generate() *jen.Statement {
    return Function{
        Name: g.Id(fmt.Sprintf("Get%s", gg.FieldName),
        Receiver: g.Arg(gg.ReceiverName, gg.ReceiverType),
        RetType: gg.FieldType,
        Body: g.Return(gg.ReceiverName.Field(gg.FieldName)),
    }
}
```

