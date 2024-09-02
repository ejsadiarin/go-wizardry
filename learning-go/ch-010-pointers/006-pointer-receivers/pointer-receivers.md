# Pointer Receivers

A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points.

- Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

However, methods with pointer receivers don't require that a pointer is used to call the method.

- The pointer will automatically be derived from the value.

## Pointer receiver

```go
type car struct {
	color string
}
```

```go
func (c *car) setColor(color string) {
	c.color = color
}
```

```go
func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
```

## Non-pointer receiver

```go
type car struct {
	color string
}
```

```go
func (c car) setColor(color string) {
	c.color = color
}
```

```go
func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "white"
}
```
