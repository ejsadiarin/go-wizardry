# Pass by Reference

Functions in Go generally pass variables by value, meaning that functions receive a copy of most non-composite types:

```go
func increment(x int) {
    x++
    fmt.Println(x)
    // 6
}
```

```go
func main() {
    x := 5
    increment(x)
    fmt.Println(x)
    // 5
}
```

The `main` function still prints `5` because the `increment` function received a copy of `x`.

One of the most common use cases for pointers in Go is to pass variables by reference, meaning that the function receives the address of the original variable, not a copy of the value.

This allows the function to **update the original variable's value**.

```go
func increment(x *int) {
    *x++
    fmt.Println(*x)
    // 6
}
```

```go
func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    // 6
}
```

# Assignment

Write a `getMessageText` function.

- It should accept a pointer to an `Analytics` struct and a `Message` struct (not a pointer).
- It should look at the `Success` field of the `Message` struct and,
- based on that, increment the `MessagesTotal`, `MessagesSuccess`, or `MessagesFailed` fields of the `Analytics` struct as appropriate.

# Tip

When your function receives a pointer to a struct, you might try to access a field like this and encounter an error:

```go
msgTotal := *analytics.MessagesTotal
```

Instead, use this simple approach using a [selector expression](https://go.dev/ref/spec#Selectors).

```go
msgTotal := analytics.MessagesTotal
```

This approach is shorthand for `(*analytics).MessagesTotal` and is the recommended, simplest way to access struct fields in Go.
