# Install

For the majority of this chapter, coding in the browser won't be useful because you need to learn how Go code is organized in a local development environment.

For the majority of this chapter, you will use our [official CLI tool](https://github.com/bootdotdev/bootdev) (written in Go, ofc) to pass off lessons.

# Assignment

All instructions on Boot.dev assume a Unix-like environment.
If you're on Windows, I recommend using [WSL 2 (Windows Subsystem for Linux)](https://docs.microsoft.com/en-us/windows/wsl/install) to get access to an Ubuntu system.

1. Install Go 1.23+ locally.

   - The easiest way (in my opinion) is to use [Webi](https://webinstall.dev/golang/) to install Go.
   - Otherwise, the [official download](https://golang.org/doc/install) page is also a great option.

2. Install and login to the bootdev CLI tool. Instructions are on the GitHub page here.
3. With the CLI working, and Go installed, run and submit the tests.

## Run the CLI commands to test your solution.

1. `go version`
   - Expecting exit code: 0
   - Expecting stdout to contain all of:
     - go version
