package main

import "fmt"

type Message struct {
	Content string
}

type Actor struct {
	mailbox chan Message
	state   string
}

func (a *Actor) Start() {
	go func() {
		for msg := range a.mailbox {
			a.handleMessage(msg) // process messages (e.g. some operations, etc.)
			a.ShowState()
		}
	}()
}

func (a *Actor) SendMessage(msg Message) {
	a.mailbox <- msg
}

func (a *Actor) handleMessage(msg Message) {
	fmt.Printf("Received message: %s\n", msg.Content)
	a.state = msg.Content
}

func (a *Actor) ShowState() {
	fmt.Printf("Current state: %s\n", a.state)
}

func simple_actor_main() {
	fmt.Println("Actor model")
	actor := Actor{
		mailbox: make(chan Message, 10),
		state:   "init",
	}

	actor.Start()

	// this can be many actors
	actor.SendMessage(Message{Content: "First message"})
	actor.SendMessage(Message{Content: "Second message updated"})

	actor.ShowState()
	fmt.Scanln() // add delay to process messages
}
