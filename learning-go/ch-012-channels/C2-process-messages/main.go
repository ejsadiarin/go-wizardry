package channels

import "time"

func processMessages(messages []string) []string {
	// ?
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
