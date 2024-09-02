# Channels Continuation

[Empty structs](https://dave.cheney.net/2014/03/25/the-empty-struct) are often used as a [unary](https://en.wikipedia.org/wiki/Unary_operation) value.

- Sometimes, we don't care what is passed through a channel.
- We care when and if it is passed.

We can block and wait until something is sent on a channel using the following syntax

```go
<-ch
```

This will block until it pops a single item off the channel, then continue, discarding the item.

# Assignment

Our Mailio server isn't able to boot up until it receives the signal that its databases are all online,

- and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code.

It never exits! The channel passed to waitForDBs stays blocked.

Fix the `waitForDBs` function.

- It should pause execution until it receives a token for every database from the `dbChan` channel.
- Each time `waitForDBs` reads a token, the `getDBsChannel` goroutine will print a message to the console for you.
- The succinctly named `numDBs` input is the total number of databases.
- Look at the test code to see how these functions are used so you can understand the control flow.
