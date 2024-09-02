# Process Notifications

Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts.

Each notification type has a unique way of calculating its importance score based on user interaction and content.

# Assignment

1. Implement the `importance` methods for each message type. They should return the importance score for each message type.

- For a `directMessage` the importance score is based on if the message `isUrgent` or not.
  - If it is urgent, the importance score is `50` otherwise the importance score is equal to the DM's `priorityLevel`.
- For a `groupMessage` the importance score is based on the messages `priorityLevel`
- All `systemAlert` messages should return a `100` importance score.

2. Complete the processNotification function. It should identify the type and return different values for each type

- For a `directMessage`, return the sender's username and importance score.
- For a `groupMessage`, return the group's name and the importance score.
- For an `systemAlert`, return the alert code and the importance score.
- If the notification does not match any known type, return an empty string and a score of 0.
