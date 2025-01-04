package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Pick choice 1 - simple actor example, 2 - many actors example")

	r := bufio.NewReader(os.Stdin)

	choice, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading string input from stdin: %v", err)
		os.Exit(1)
	}
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		simple_actor_main()
	case "2":
		many_actor_model_main()
	default:
		fmt.Println("Unsupported choice.")
	}
}
