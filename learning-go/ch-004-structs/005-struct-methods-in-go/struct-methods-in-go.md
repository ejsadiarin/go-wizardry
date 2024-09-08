# Struct methods in Go

While Go is not object-oriented, it does support methods that can be defined on structs.

Methods are just functions that have a receiver.

A receiver is a special parameter that syntactically goes before the name of the function.

```go
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50
```

A receiver is just a special kind of function parameter.

In the example above, the `r` in `(r rect)` could just as easily have been rec or even `x,` `y` or `z.`

- By convention, Go code will often use the first letter of the struct's name.

Receivers are important because they will, as you'll learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.

# Assignment

Let's clean up Textio's authentication logic.

- We store our user's authentication data inside an `authenticationInfo` struct.
- We need a method that can take that data and return a basic authorization string.

The format of the string should be:

```txt
Authorization: Basic USERNAME:PASSWORD
```

Create a method on the `authenticationInfo` struct called `getBasicAuth` that returns the formatted string.
