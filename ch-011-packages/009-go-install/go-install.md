# Go Install

The go install command compiles and installs a package or packages on your local machine for your personal usage.

- It installs the package's compiled binary in the GOBIN directory. (We installed the bootdev cli with it, after all)

# Assignment

1. Ensure you are in your hellogo repo, then run:

```bash
go install
```

2. Navigate out of your project directory:

```bash
cd ../
```

3. If all went well, Go compiled and installed the hellogo program globally on your machine. Run it with:

```bash
hellogo
```

4. Run and submit the CLI tests from anywhere on your machine.

# Tips

If you get an error regarding "hellogo not found" it means you probably don't have your Go environment setup properly.

Specifically, `go install` is adding your binary to your `GOBIN` directory, but that may not be in your [PATH](https://en.wikipedia.org/wiki/PATH_%28variable%29).

You can read more about that here in the [go install docs](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies).

## Run the CLI commands to test your solution.

1. `hellogo`
   - Expecting exit code: `0`
   - Expecting stdout to contain all of:
     - `hello world`
