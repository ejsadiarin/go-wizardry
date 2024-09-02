# References

It's possible to define an empty pointer.

For example, an empty pointer to an integer:

```go
var p *int
```

Its zero value is `nil`, which means it doesn't point to any memory address.

It's common to immediately create a new pointer by instead using the `&` operator to get a pointer to its operand:

```go
myString := "hello" // myString is just a string
myStringPtr := &myString // myStringPtr is a pointer to myString's address
```

## Dereference

The `*` operator dereferences a pointer to get the original value.

```go
fmt.Println(*myStringPtr) // read myString through the pointer
    *myStringPtr = "world"    // set myString through the pointer
```

Unlike C, Go has no [pointer arithmetic](https://www.tutorialspoint.com/cprogramming/c_pointer_arithmetic.htm).

# Assignment

Complete the `removeProfanity` function.

It should use the [strings.ReplaceAll](https://pkg.go.dev/strings#ReplaceAll) function to replace all instances of the following words in the input `message` with asterisks.

| Word  | Replacement |
| ----- | ----------- |
| fubb  | `****`      |
| shiz  | `****`      |
| witch | `*****`     |

It should _update_ the value in the pointer and return nothing.

Do _not_ alter the function signature.
