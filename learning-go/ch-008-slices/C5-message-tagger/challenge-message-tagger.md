# Message tagger

Textio needs a way to tag messages based on specific criteria

# Assignment

1. Complete the `tagMessages` function.

   - It should take a slice of `sms` messages, and a function (that takes a `sms` as input and returns a slice of strings) as inputs.
   - And it should return a slice of `sms` messages.
   - It should loop through each message and set the `tags` to the result of the passed in function.
   - Be sure to modify the messages of the original slice.

2. Complete the `tagger` function.

   - It should take a sms message and return a slice of strings.
   - For any message that contains "urgent" or "Urgent" in the content, the Urgent tag should be applied first.
   - For any message that contains "sale" or "Sale" in the content, the Promo tag should be applied second.

Example usage:

```go
messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}
taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]
```

# Tip

The go `strings` package, specifically the [contains method](https://pkg.go.dev/strings#Contains) can be very helpful here!
