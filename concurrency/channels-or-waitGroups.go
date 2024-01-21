package main

import (
	"fmt"
	"time"
)

func doworkThis(d time.Duration) {
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println("work is done!")
}

func start() {
	start := time.Now()
	// dowork(time.Second * 2)
	// dowork(time.Second * 4)
	fmt.Printf("work took %v seconds\n", time.Since(start))
}
