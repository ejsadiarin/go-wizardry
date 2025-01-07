package main

import (
	"fmt"
	"time"
)

// Define a message structure
type Message2 struct {
	Sender   string
	Content  string
	Receiver string
}

// Actor2 structure
type Actor2 struct {
	ID      string
	mailbox chan Message2
	state   string
}

// Function to start the actor's behavior loop
func (a *Actor2) Start(actors map[string]*Actor2) {
	go func() {
		for msg := range a.mailbox {
			a.handleMessage2(msg, actors)
		}
	}()
}

// Function to handle messages
func (a *Actor2) handleMessage2(msg Message2, actors map[string]*Actor2) {
	fmt.Printf("Actor2 %s received message: %s\n", a.ID, msg.Content)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Actor2 %s state [a.state]: %s\n", a.ID, a.state)
	fmt.Printf("Actor2 %s state [actors[\"%s\"].state: %s\n", a.ID, a.ID, actors[a.ID].state)
	fmt.Println("-----------------------------------------")

	switch msg.Content {
	case "Process Order":
		// Update state
		a.state = "In Progress"
		fmt.Printf("[IN handleMessage2] Actor2 %s updated state to: %s\n", a.ID, a.state)

		// Send a message to Actor2 B
		if actorB, ok := actors["B"]; ok {
			actorB.mailbox <- Message2{Sender: a.ID, Content: "Order Received", Receiver: "B"}
		}

		// Dynamically create a new Actor2 C
		actorC := &Actor2{ID: "C", mailbox: make(chan Message2, 10), state: "Idle"}
		actors["C"] = actorC
		actorC.Start(actors)

		// Send a message to Actor2 C
		actorC.mailbox <- Message2{Sender: a.ID, Content: "Process This Order", Receiver: "C"}

		// Complete Actor2 A's behavior
		fmt.Printf("Actor2 %s has completed processing.\n", a.ID)
	}
}

func many_actor_model_main() {
	// Create a map to hold all actors
	actors := make(map[string]*Actor2)

	// Create Actor2 A and Actor2 B
	actorA := &Actor2{ID: "A", mailbox: make(chan Message2, 10), state: "Idle"}
	actorB := &Actor2{ID: "B", mailbox: make(chan Message2, 10), state: "Idle"}
	actors["A"] = actorA
	actors["B"] = actorB

	// Start Actor2 A and Actor2 B
	actorA.Start(actors)
	actorB.Start(actors)

	// Send a message to Actor2 A
	actorA.mailbox <- Message2{Sender: "System", Content: "Process Order", Receiver: "A"}

	// Allow time for processing
	time.Sleep(1 * time.Second)
}
