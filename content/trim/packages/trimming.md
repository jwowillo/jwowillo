# trimming

API Documentation and Examples: https://godoc.org/github.com/jwowillo/trim/trimming

## Description

Package `trimming` consists of implementations of `trim.Trimming`s. A
`trim.Trimming` is a way to add extra functionality to any `trim.Handler`
without having to change or manipulate the actual `trim.Handler` logic. Imagine
a `trim.Trimming` working like the diagram:

<center>![Diagram]({{ static }}/trim/packages/trimming.png)</center>

The `trim.Handler` is the base functionality provided in the inner black box.
It already has an input and output. The `trim.Trimming` sits in front of the
black box's input and output and imitates its interface. The input and output
are manipulated by the `trim.Trimming`.

Since `trim.Trimming`s pretend to be `trim.Handler`s, they extend the
`trim.Handler` interface and provide their own interface method,
`Apply(trim.Handler)`. This method substitutes the `trim.Trimming`s black
box for a new one.

Any type which has `trim.Trimming`s is called `trim.Trimmed`. This interface has
a single method with signature `Trimmings() []trim.Trimming` which returns all
the `trim.Trimming`s on a `trim.Trimmed` type.

The package contains implementations of several `trim.Trimming`.
