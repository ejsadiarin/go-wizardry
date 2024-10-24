package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n < max; n++ {
		if n == 2 {
			fmt.Println(n)
		} else if n%2 == 0 {
			continue
		}
		for i := 3; i < int(math.Sqrt(float64(n))); i++ {
			if i*i <= i {
				break
			}
		}
		fmt.Println(n)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
