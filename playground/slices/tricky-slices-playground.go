package slices

import (
	"fmt"
)

func main() {
	a := make([]int, 4)

	a[0] = 1
	a[1] = 1
	a[2] = 1
	a[3] = 1

	b := a[0:1]

	// Makes sense, b is a reference
	// a: [2 1 1 1]
	// b: [2]
	b[0] = 2

	fmt.Println("Before append to b:")
	fmt.Printf("a: %v\n b: %v\n", a, b)

	b = append(b, 3)

	b[0] = 3

	// Um... why did append change a?
	// a: [3 3 1 1]
	// b: [3 3]
	fmt.Println("After append to b:")
	fmt.Printf("a: %v\n b: %v\n", a, b)
}

/*
EXPLANATION
- slices have an underlying array (from go source code):
type slice struct {
  array unsafe.Pointer
  len int
  cap int
}
- a slice created from an existing slice share an underlying array, unless the capacity is not enough
- if capacity is not enough, it will be re-assigned another new array

*/
