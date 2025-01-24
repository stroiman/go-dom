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

## Example

To come
