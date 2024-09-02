# First Local Program

1. Once inside your personal workspace, create a new directory and enter it:

```go
mkdir hellogo
cd hellogo
```

2. Inside the directory declare your module's name:

```go
go mod init {REMOTE}/{USERNAME}/hellogo
```

Where `{REMOTE}` is your preferred remote source provider (i.e. `github.com`) and `{USERNAME}` is your Git username. If you don't use a remote provider yet, just use `example.com/username/hellogo`

3. Print your `go.mod` file:

```bash
cat go.mod
```

4. Run and submit the CLI tests from the root of the hellogo package.

## Run the CLI commands to test your solution.

1. cat go.mod
   - Expecting exit code: 0
   - Expecting stdout to contain all of:
     - hellogo
