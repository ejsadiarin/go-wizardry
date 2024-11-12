package main

import (
	"fmt"
	"os"
)

func scripting() {
	fmt.Println("first args:", os.Args[1])
}

func main() {
	scripting()
}
