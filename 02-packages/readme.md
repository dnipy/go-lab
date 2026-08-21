# Run a package from another one

Inside Go:

```go
package main
```

means **executable**.

So we defined the `main` package and then imported other needed packages.

In this case, we defined the second package named `greeting` and imported it inside the executable file via:

```go
"github.com/dnipy/go-lab/02-packages/greeting"
```

It's exactly our module that we defined inside the root project (`go.mod`).

So there are three layers:

1. **Module** — the whole project
2. **Package** — related Go files grouped together
3. **File** — individual `.go` files inside a package

The packages in Go are basically **namespacing related Go files together** and making them importable.

And in each Go file:

- For private vars and utils, we define the name in **lowercase**:

```go
getTime
extractSome
localVar
```

- For public and exported vars and utils, we define the name in **uppercase**:

```go
Hello
Goodbye
GetUser
```

And to list all available packages inside a Go module, we can easily use:

```bash
go list ./...
```
