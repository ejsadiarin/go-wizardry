package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	// "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// need go get github.com/docker/docker/client
func dockerSDK() {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	// List all running containers
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Fatalf("Unable to list containers: %s", err)
	}

	// Print the list of containers
	for _, container := range containers {
		fmt.Printf("Container ID: %s, Image: %s, Names: %v\n", container.ID[:10], container.Image, container.Names)
	}
}

func main() {
	// Docker socket path (usually /var/run/docker.sock on Unix systems)
	socketPath := "/var/run/docker.sock"

	// Create a new HTTP client that communicates with the Docker socket
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}

	// Make a GET request to the Docker API to list containers
	resp, err := client.Get("http://localhost/containers/json") // NOTE: same as "docker ps"
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
