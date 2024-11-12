package main

import (
	"fmt"
	"log"
	"os"
)

func ls() {
	// get current path
	cwd, err := os.Getwd()
	fmt.Println("Path:", cwd)
	if err != nil {
		log.Fatal(err)
	}

	// get current path
	dirpath, err := os.ReadDir(cwd)
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
