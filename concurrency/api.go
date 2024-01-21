package main

import (
	"fmt"
	"math/rand"
	"time"
)

// using wait groups
// func dowork(d time.Duration, wg *sync.WaitGroup) {
// 	fmt.Println("doing work...")
// 	time.Sleep(d)
// 	fmt.Println("work is done!")
// 	wg.Done()
// }

// func main() {
// 	start := time.Now()
// 	wg := sync.WaitGroup{}
// 	go dowork(time.Second*2, &wg)
// 	go dowork(time.Second*4, &wg)
// 	wg.Wait()
// 	fmt.Printf("work took %v seconds\n", time.Since(start))
// }

// using channels

func dowork(d time.Duration, resch chan string) {
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println("work is done!")
	resch <- fmt.Sprintf("the result of the work -> %d", rand.Intn(100))
}

func main() {
	start := time.Now()
	resultch := make(chan string)
	go dowork(time.Second*2, resultch)
	go dowork(time.Second*4, resultch)

	res1 := <-resultch
	fmt.Println(res1)
	res2 := <-resultch
	fmt.Println(res2)
	fmt.Printf("work took %v seconds\n", time.Since(start))
}
