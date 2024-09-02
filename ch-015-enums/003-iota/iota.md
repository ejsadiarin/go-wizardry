# Iota

Go has a language feature, that when used with a type alias (and if you squint really hard), kinda looks like an enum (but it's not).

It's called `iota`.

```go
type sendingChannel int

const (
    Email sendingChannel = iota
    SMS
    Phone
)
```

The `iota` keyword is a special keyword in Go that creates a sequence of numbers.

- It starts at 0 and increments by 1 for each constant in the `const` block.

So in the example above, `Email` is 0, `SMS` is 1, and `Phone` is 2.

Go developers sometimes use `iota` to create a sequence of constants to represent a set of related values,

- much like you would with an enum in other languages.

But remember, it's not an enum. It's just a sequence of numbers.

# Assignment

Define an `emailStatus` type that uses `iota` syntax to represent the following states:

- `emailBounced:` 0
- `emailInvalid:` 1
- `emailDelivered:` 2
- `emailOpened:` 3
