# Currying

Function [currying](https://en.wikipedia.org/wiki/Currying) is a concept from functional programming and involves [partial application](https://en.wikipedia.org/wiki/Partial_application) of functions.

- It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.

Let's simulate this behavior. For example:

```go
func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)

  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
```

In the example above, the `selfMath` function takes in a function as its parameter

- and returns a function that itself returns the value of running that input function on its parameter.

# Assignment

The Mailio API needs a very robust error-logging system so we can see when things are going awry in the back-end system.

We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

Complete the `getLogger` function.

- It should take as input a `formatter` function and `return` a new function.
- The new logger function takes as input two strings and passes them to the `formatter,` then prints the result.
- Keep the order of the strings.
