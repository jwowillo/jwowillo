# trimming

API Documentation and Examples: https://godoc.org/github.com/jwowillo/trim/trimming

## Description

Package `trimming` consists of implementations of `Trimming`s. A `Trimming` is a
way to add extra functionality to any `handler.Handler` without having to change
or manipulate the actual `handler.Handler` logic. Imagine a `Trimming` working
like the diagram:

The `handler.Handler` is the base functionality provided in the inner black box.
It already has an input and output. The `Trimming` sits in front of the black
box's input and output and imitates its interface. The input and output are
manipulated by the `Trimming`.

Since `Trimming`s pretend to be `handler.Handler`s, the extend the
`handler.Handler` interface and provided their own interface method,
`Apply(handler.Handler)`. This method substitutes the `trimming.Trimming`s black
box for a new one.

Any type which has `Trimming`s is called `Trimmed`. This interface has a single
method with signature `Trimmings() []Trimming` which returns all the `Trimming`s
on a `Trimmed` type.

The package also consists of several `Trimming` implementations documented at
the linked "API Documentation and Examples".
