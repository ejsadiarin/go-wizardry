# Remote Packages

Let's learn how to use an open-source package that's available online.

## A note on "replace"

Be aware that using the "replace" keyword like we did in the last assignment is not advised, but can be useful to get up and running quickly.

- The proper way to create and depend on modules is to publish them to a remote repository.

When you do that, the "replace" keyword can be dropped from the `go.mod`:

This only works for local-only development:

```go
replace github.com/wagslane/mystrings v0.0.0 => ../mystrings
```

If we want the import to work for everyone, we need to make sure the dependency (`mystrings` in this case) actually exists on `https://github.com/wagslane/mystrings.`

[Watch this video](https://youtu.be/jEVVEI2UIts)

# Assignment

1. Create a new directory in the same parent directory as `hellogo` and `mystrings` called `datetest`.
2. Create `main.go` in `datetest` and add the following code:

```go
package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1585750374)
	tt = tt.Add(time.Hour * 48)
	fmt.Println("1585750374 converted to a tinytime is:", tt)
}
```

3. Initialize a module:

```bash
go mod init {REMOTE}/{USERNAME}/datetest
```

4. Download and install the remote go-tinydate package using `go get`:

```bash
go get github.com/wagslane/go-tinytime
```

5. Print the contents of your go.mod file to see the changes:

```bash
cat go.mod
```

6. Compile and run your program:

```go
go build
./datetest
```

7. Run and submit the CLI tests from the root of the `datetest` package.

## Run the CLI commands to test your solution.

1. ./datetest
   - Expecting exit code: 0
   - Expecting stdout to contain all of:
     - 1585750374 converted to a tinytime is: 2020-04-03T14:12:54Z
