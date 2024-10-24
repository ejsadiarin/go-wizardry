package main

import "fmt"

func fizzbuzz() {
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
		i++
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
