# Message Filter

Textio is introducing a feature that allows users to filter their messages based on specific criteria.
For this feature, messages are categorized into three types:

- `TextMessage`, `MediaMessage`, and `LinkMessage`.
- Users can filter their messages to view only the types they are interested in.

# Assignment

Your task is to implement a function that filters a slice of messages based on the message type.

Complete the `filterMessages` function.
_ It should take a slice of `Message` interfaces and a string indicating the desired type ("text", "media", or "link").
_ It should return a new slice of `Message` interfaces containing only messages of the specified type.
