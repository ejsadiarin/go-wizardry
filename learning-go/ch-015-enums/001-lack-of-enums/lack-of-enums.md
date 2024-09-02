# Lack of Enums

My least favorite part of Go? Glad you asked.

It's Go's lack of [enums](https://en.wikipedia.org/wiki/Enumerated_type), [sum types](https://en.wikipedia.org/wiki/Algebraic_data_type), [tagged unions](https://en.wikipedia.org/wiki/Tagged_union), etc.

Compared to other statically typed languages like:

- [Rust](https://www.rust-lang.org/)
- [TypeScript](https://www.typescriptlang.org/)
- [OCaml](https://ocaml.org/)

Go's type system just isn't as powerful.

- It's more similar to C's type system than it is to Rust's.
- It's more concerned with simplicity than it is with expressiveness.

## Error Handling

In Rust, like Go, errors are just values.

In Go, we write something like this:

```go
user, err := getUser()
if err != nil {
    return fmt.Errorf("failed to get user: %w", err)
}
// do something with user
```

In Rust, we can do something like this:

```rust
let user_result = get_user();
let user = match user_result {
    Ok(user) => user,
    Err(error) => return Err(format!("failed to get user: {}", error)),
};
```

In Rust, the `get_user` function returns a Result type:

- a type that is either an `Ok` or an `Err`.
- The compiler forces the developer to handle the error case
  - before they can continue with the happy path (using the user data).

In Go, the developer can choose to happily ignore the error value if they choose and use the user data,

- even if it's invalid (probably `nil` or an empty struct).

The support for enums in Rust makes it easier to write bug-free code.

## Assignment

A lazy Go programmer wrote the `handleEmailBounce` function...

- just because the compiler doesn't force us to check errors doesn't mean we shouldn't!

Take a look at the `updateStatus` and `track` methods in the `user.go` file. Handle their errors properly, and use the [`fmt.Errorf` function](https://pkg.go.dev/fmt#Errorf) and the `%w` formatting verb to add useful context to the errors.

- If `updateStatus` fails, return an error saying "error updating user status: ERR"
- If `track` fails, return an error saying "error tracking user bounce: ERR"

Where `ERR` is the error returned by the method.
