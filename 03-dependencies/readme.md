# Using External Codes

In Go we can use two types of imports:

1. Packages from our own module
2. Packages from outside of our module

We learned how to do the first one, and now we are learning about using external Go code.

It's similar to:

- JavaScript → `npm install`
- Python → `pip install`
- Go → `go get`

## Getting an External Package

First we need to get the needed package and add it as a dependency requirement inside our `go.mod` file using:

```bash
go get <package-name>
```

For example:

```bash
go get github.com/google/uuid
```

After that, the package is available to import inside our Go packages and files:

```go
import "github.com/google/uuid"
```

## `go.mod`

`go.mod` contains our **dependency requirements**.

For example, after installing `github.com/google/uuid`, Go adds it to `go.mod` with the required version.

So we can think of it as:

```text
go.mod = dependency requirements
```

## `go.sum`

When we use external dependencies, Go also uses a `go.sum` file.

`go.sum` contains **checksums** for module content.

Go uses these checksums to verify dependency integrity.

So:

```text
go.mod = dependency requirements

go.sum = verified dependency checksums
```

We should not handle these manually.

Go manages them for us through commands like:

```bash
go get
```

and:

```bash
go mod tidy
```

## Listing Packages

To see all packages inside the current Go module:

```bash
go list ./...
```

## Listing Modules

To see all modules involved in the current module:

```bash
go list -m all
```
