// socket-client project main.go

package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	// establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	///send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}

/*
* UPDATE [December 25, 2024 - 7:53 PM]:
* - created "go-tcp" and "go-udp" repos
* - critique on this simple demo:
*   - this is a "client" for a TCP server
*   - note that the buffer is there only for reading the response of the server
*   - why defer on end (move it to line 22)
* */

// ref: https://www.developer.com/languages/intro-socket-programming-go/
