package channels

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, pongs, numPings)
}

// don't touch below this line

func pinger(pings, pongs chan struct{}, numPings int) {
	go func() {
		sleepTime := 50 * time.Millisecond
		for i := 0; i < numPings; i++ {
			fmt.Printf("sending ping %v\n", i)
			pings <- struct{}{}
			time.Sleep(sleepTime)
			sleepTime *= 2
		}
		close(pings)
	}()
	i := 0
	for range pongs {
		fmt.Println("got pong", i)
		i++
	}
	fmt.Println("pongs done")
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}
