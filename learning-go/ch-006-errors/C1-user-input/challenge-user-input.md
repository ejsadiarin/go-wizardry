# User Input

In Textio, users can set their profile status to communicate their current activity to those that choose to read it...

- However, there are some restrictions on what these statuses can contain.

Your task is to implement a function that validates a user's status update.

- The status update cannot be empty and must not exceed 140 characters.

# Assignment

Complete the `validateStatus` function. It should return an error when the status update violates any of the rules:

- If the status is empty, return an error that reads `status cannot be empty`.
- If the status exceeds 140 characters, return an error that says `status exceeds 140 characters`.
