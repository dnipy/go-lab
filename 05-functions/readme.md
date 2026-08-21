# Functions

In Go, functions are a bit different from other languages such as TS and Python.

Here we can have a function with multiple return values.

> Not an array with multiple values. The function can really return multiple values.

And it's super common to see and handle multiple return values from one function.

In fact, we should handle all returned values.

Or we can pass `_` for the values we don't want to access, to tell the compiler that we don't need them here.

## Multiple Return Values

For example:

```go
func getUser() (string, int) {
	return "Daniel", 23
}
```

We can receive both values:

```go
name, age := getUser()
```

Or ignore one of them:

```go
name, _ := getUser()
```

## Error Handling

As functions can return multiple values, one of those values could be an `error`.

In Go there is no `try/catch` scope for normal error handling.

The error of a function can simply be one of the return values of that function, and we should handle it gracefully.

For example:

```go
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
```

And usage:

```go
result, err := divide(10, 2)

if err != nil {
	fmt.Print("error while dividing: ", err)
	return
}
```

If there is no error:

```go
return result, nil
```

`nil` means that there is no error.

## Common Go Pattern

This pattern is extremely common in Go:

```go
result, err := doSomething()

if err != nil {
	return err
}
```

We will see this pattern everywhere when working with Go backend code.
