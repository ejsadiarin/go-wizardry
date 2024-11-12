package main

import (
	"fmt"
	"os"
)

func testers() {
	fmt.Println("first args:", os.Args[1])
}

func main() {
	testers()
}
