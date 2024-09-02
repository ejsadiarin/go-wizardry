# Omitting conditions from a for loop in Go

Loops in Go can omit sections of a for loop.

- For example, the `CONDITION` (middle part) can be omitted which causes the loop to run forever.

```go
for INITIAL; ; AFTER {
  // do something forever
}
```

# Assignment

Complete the `maxMessages` function.

Given a cost threshold, it should calculate the maximum number of messages that can be sent.

Each message costs `100` pennies, plus an additional fee. The fee structure is:

- 1st message: `100 + 0`
- 2nd message: `100 + 1`
- 3rd message: `100 + 2`
- 4th message: `100 + 3`

# Browser freeze

If you lock up your browser by creating an infinite loop that isn't breaking, just click the `cancel` button.
