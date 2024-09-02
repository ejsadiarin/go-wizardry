# Arrays in Go

Arrays are fixed-size groups of variables of the same type.

The type `[n]T` is an array of n values of type `T`

To declare an array of 10 integers:

```go
var myInts `[10]int`
```

or to declare an initialized literal:

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
```

# Assignment

When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

`getMessageWithRetries` returns:

- An array of 3 strings
- An array of 3 integers

The strings are just the original messages structured as an array.
The first is the primary message, the second is the first reminder, and the third is the last reminder.
The integers represent the cost of sending each message.

The cost of a message is equal to the length of the message, plus the cost of any previous messages.

For example:

1. "hello" costs 5
2. "world" costs 5, adding "hello" makes total cost 10 (5 + 5)
3. "!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)
