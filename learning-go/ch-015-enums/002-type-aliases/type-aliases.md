# Type Aliases

For all its faults, [TypeScript](https://www.typescriptlang.org/) does have a pretty incredible type system. Here's one of the things it can do that I often miss in Go:

```go
type sendingChannel = "email" | "sms" | "phone";

function sendNotification(ch: sendingChannel, message: string) {
    // send the message
}
```

This `sendingChannel` type that we've created is a [union type](https://www.typescriptlang.org/docs/handbook/2/everyday-types.html#union-types).

- It can only be one of the three strings that we've defined.

That means when a developer calls `sendNotification()` they can't accidentally pass an invalid `sendingChannel` like "slack" or even a misspelled "emil".

- The TypeScript compiler will catch that mistake at compile time.

In Go, we don't have these nice things.

We embrace the [Grug](https://grugbrain.dev/) mentality.

The closest thing we have to a union type is a [type alias](https://go.dev/ref/spec#Alias_declarations):

```go
type sendingChannel string

const (
    Email sendingChannel = "email"
    SMS   sendingChannel = "sms"
    Phone sendingChannel = "phone"
)

func sendNotification(ch sendingChannel, message string) {
    // send the message
}
```

It's a bit safer than using a plain old `string` in Go, but it's not completely safe. Go will stop us from doing this:

```go
sendingCh := "slack"
sendNotification(sendingCh, "hello") // string is not sendingChannel
```

But it will not stop us from doing this:

```go
// "slack" is automatically implied as a sendingChannel
sendNotification("slack", "hello")
```

And will also not stop us from doing this:

```go
sendingCh := "slack"
convertedSendingCh := sendingChannel(sendingCh)
sendNotification(convertedSendingCh, "hello")
```

The `sendingChannel` type is just an alias for `string`, and because we made some constants of that type,

- most developers will just use those constants: we've made it easy to do the right thing.

That said, Go still doesn't force us to do the right thing like TypeScript does.

# Assignment

Use the following code to answer the questions:

```go
type perm string

const (
    Read  perm = "read"
    Write perm = "write"
    Exec  perm = "execute"
)

var Admin = "admin"
var User = perm("user")

func checkPermission(p perm) {
    // check the permission
}
```
