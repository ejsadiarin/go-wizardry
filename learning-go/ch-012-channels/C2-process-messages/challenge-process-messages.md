# Process Messages

Textio needs to speed up message processing using concurrency.

# Assignment

Implement the `processMessages` function,

- it takes a slice of messages.
- It should process each message concurrently within a [goroutine](https://gobyexample.com/goroutines).
- Use the `process` function to simulate the benefit of using goroutines.
- Use a channel to ensure that all messages are processed and collected correctly
  - then return the slice of _processed_ messages.

Example output:

````go
messages: []string{"Here are some messages", "Here is another", "and another"}
processedMessages := processMessages(messages)
fmt.Printf(processedMessages)
// prints ["Here are some messages-processed", "Here is another-processed", "and another-processed"]
    ```

````
