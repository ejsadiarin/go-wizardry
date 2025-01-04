package main

import (
	"context"
	"fmt"
	"time"
)

/**
* Flow
* 1. The context.WithTimeout internally creates a channel (context.Done).
* 2. When the timeout (2 seconds) is reached, the context signals cancellation by sending to the channel.
* 3. The goroutine listening on <-ctx.Done() is notified and executes.
*
* */
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) // creates a context.Done channel internally

	// the producer/writer of the channel in this case
	defer cancel()

	go func() {
		<-ctx.Done() // listener/consumer (read-only channel/stream) thats waits for a signal
		fmt.Printf("Context done: %v\n", ctx.Err())
	}()

	fmt.Println("Doing some work...")
	time.Sleep(3 * time.Second) // simulate work (3 secs), context signals cancellation since it times out after 2 seconds only
	fmt.Println("Work complete")
}
