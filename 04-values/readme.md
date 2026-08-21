# Variables and Values

Inside Go we can define variables by using:

1. `var`
2. `const`
3. `:=` — only inside functions

## Zero Values

In Go we don't have `undefined` like JavaScript.

Instead, each type has its own **zero value**.

For example:

```text
bool   → false
string → ""
int    → 0
```

And other types have their own zero values too.

## `var`

`var` can be defined at package scope and function scope.

If we define the type but don't pass a value, Go automatically assigns the zero value for that type.

For example:

```go
var name string
```

The value of `name` will be:

```text
""
```

Or we can only pass a value and let Go infer the type:

```go
var name = "hi"
```

Go knows that `name` is a `string`.

The type of the variable cannot later be changed.

## `const`

`const` is used to define constant values.

For example:

```go
const PORT = 3000
const DEBUG = false
```

A constant cannot be reassigned.

## `:=`

`:=` is used to declare a new variable and assign a value to it at the same time.

It can only be used inside functions.

For example:

```go
username := "dnipy"
```

Go infers that `username` is a `string`.

To modify the value later, we use `=`:

```go
username = "dnipy_2"
```

We cannot do:

```go
username := "dnipy"
username := "dnipy_2"
```

because `:=` declares a new variable, and we're trying to redeclare the same variable inside the same scope.

So:

```text
:=  → declare + assign
=   → assign
```
