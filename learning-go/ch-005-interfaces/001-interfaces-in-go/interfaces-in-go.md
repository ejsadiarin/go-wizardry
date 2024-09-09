# Interfaces in Go

[Interfaces](https://go.dev/tour/methods/9) are just collections of method signatures.

- A type "implements" an interface if it has methods that match the interface's method signatures.

In the following example, a "shape" must be able to return its area and perimeter.

- Both `rect` and `circle` fulfill the interface.

```go
type shape interface {
    area() float64
    perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

When a type implements an interface, it can then be used as that interface type.

```go
func printShapeData(shapes []shape) {
	fmt.Println("The areas and perimeters of the shapes are:")
	for i, shape := range shapes {
		fmt.Printf("Index: %v - Area: %v - Perimeter: %v\n", i, shape.area(), shape.perimeter())
	}
}
```

Because we say the input is of type `shape`, we know that each element of the slice must implement the `.area()` and `.perimeter()` methods.

As an example, because [the empty interface](https://go.dev/tour/methods/14) doesn't require a type to have any methods at all, every type automatically implements the empty interface, written as:

```go
interface{}
```

# Assignment

- The `birthdayMessage` and `sendingReport` structs already have implementations of the getMessage method.
- The `getMessage` method returns a string,
- and any type that implements the method can be considered a `message` (meaning it implements the `message` interface).

1. Add the `getMessage()` method signature as a requirement on the message interface.
2. Complete the `sendMessage` function. It should return:
   1. The content of the message.
   2. The cost of the message, which is the length of the message multiplied by 3.

Notice that your code doesn't care at all about whether a specific message is a `birthdayMessage` or a `sendingReport!`

# Tip

The length of a string can be obtained using the len function, which returns the number of bytes.

```go
s := "Hello, World!"
fmt.Println(len(s))
// 13
```
