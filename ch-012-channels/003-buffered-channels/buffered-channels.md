# Buffered Channels

Channels can optionally be buffered.

## Creating a channel with a buffer

You can provide a buffer length as the second argument to `make()` to create a buffered channel:

```go
ch := make(chan int, 100)
```

A buffer allows the channel to hold a fixed number of values before sending blocks.

This means sending on a buffered channel only blocks when the buffer is full,

- and receiving blocks only when the buffer is empty.

# Assignment

We want to be able to send emails in _batches_.

A _writing goroutine_ will write an entire batch of email messages to a buffered channel,

- and later, once the channel is full, a _reading goroutine_ will read all of the messages from the channel
- and send them out to our clients.

Complete the `addEmailsToQueue` function.

- It should create a buffered channel with a buffer large enough to store all of the `emails` it's given.
- It should then write the emails to the channel in order, and finally return the channel.
