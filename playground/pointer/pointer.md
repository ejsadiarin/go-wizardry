# Pointers
- gives write access
## Pointer Dereferencing
- originally, a pointer "points" to the memory address of a variable
  - when it is "dereferenced", it points to the `data` inside the memory address
- in other words: when `pointer` is dereferenced, the `value` is accessed
```go
func main() {
	number := 7
	ptr := &number

  fmt.Printf("Original value of number: %d\n", number)

	*ptr = 10 // dereferencing ptr and assign new value or data to which it points to (number)

  fmt.Printf("New value of number: %d\n", number)
}
```
## Pointer Type Inference
- happens when assigning an address to a variable (the variable becomes a pointer)
```go
func main() {
	number := 7
  // pointer type inference
	ptr := &number

	// the same:
	fmt.Printf("Value of number: %d", *ptr)
	fmt.Printf("Value of number: %d", ptr)
}
```
## Pointer Panics
<!-- TODO: add link to article -->
- see panics in `panic/panic.md`

- "pointer panic" happens when a pointer is dereferenced with a `nil` value like so:
```go
number := 7
ptr := &number
*ptr := nil // this is a runtime panic!
```

## Pointer Receiver VS Value Receiver
- ref: 
