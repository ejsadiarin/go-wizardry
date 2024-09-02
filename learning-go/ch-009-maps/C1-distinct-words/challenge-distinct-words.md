# Distinct Words

Complete the `countDistinctWords` function using a `map`.

- It should take a slice of strings and return the total count of distinct words across all the strings.
- Assume words are separated by spaces.
- Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

```go
messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)
```

`count` should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.

# Hint

Go's [strings package](https://pkg.go.dev/strings) can be very helpful here.

Specifically the [Fields](https://pkg.go.dev/strings#Fields) method and the [ToLower](https://pkg.go.dev/strings#ToLower) method.
