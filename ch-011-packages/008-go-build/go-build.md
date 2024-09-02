# Go Build

The [`go build` command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies) compiles go code into a single, statically linked executable program.

One of the beauties of Go is that you always go build for production, and because the output is a statically compiled binary:

- **you can ship it to production or end users without them needing the Go toolchain installed.**

Some new Go devs use go run on a server in production, which is a huge mistake.

# Assignment

1. Ensure you are in your hellogo repo, then run:

```bash
go build
```

2. Run the new program:

```bash
./hellogo
```

3. Run and submit the CLI tests from the root of the repo.

## Run the CLI commands to test your solution.

1. ./hellogo
   - Expecting exit code: 0
   - Expecting stdout to contain all of:
     - hello world
