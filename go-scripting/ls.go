package main

import (
	"fmt"
	"log"
	"os"
)

func ls() {
	// get current path
	dir := os.Args[1]
	if len(os.Args) > 2 && dir != "" {
		log.Fatal("Cannot have more than 1 arguments.")
	}
	if dir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dir = cwd
	}
	// cwd, err := os.Getwd()
	fmt.Println("Path:", dir)

	// get current path
	dirpath, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range dirpath {
		fmt.Println(v)
	}
}

func main() {
	ls()
}
