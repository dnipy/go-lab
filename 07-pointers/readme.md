# Pointer

In Go variables have two related concepts:

- values
- addresses

Now what's that mean?

The important thing is that when we pass a value to a function, Go passes it by value.

That means the function receives a copy of that value.

But if we want the function to work with the original value, we can pass a pointer to it.

Let's get deeper.

Assume we define a user:

```go
user := User{
	Name: "Danial",
	Age:  22,
}
```

This `user` is a value stored somewhere in memory.

Now assume we have a function and we want to modify the user.

If we pass it without a pointer:

```go
func changeName(user User, name string) {
	...
}
```

the function receives a copy of the original `User`.

We can do whatever we want to that copy, but the original `User` outside the function will not be modified.

So:

```text
original User
     ↓
   copy
     ↓
function
```

If we really want to modify the original value, we can pass a pointer:

```go
func changeName(user *User, name string) {
	...
}
```

Now the function receives a pointer to the original `User`.

But how do we create that pointer?

By using `&`.

For example:

```go
changeName(&user, "Alex")
```

`&user` means:

> Give me the address of `user`.

So:

```go
changeName(user, "Alex")
```

passes a copy of the `User`.

While:

```go
changeName(&user, "Alex")
```

passes a pointer to the original `User`.

## `*`

When we define:

```go
func changeName(user *User, name string) {
	...
}
```

the `*` means that `user` is a pointer to a `User`.

A pointer contains the address of a value.

We can use `*` to dereference a pointer and access the value it points to.

For example:

```go
user := User{
	Name: "Danial",
	Age:  22,
}

pointer := &user

fmt.Println(*pointer)
```

Here:

```text
&user
```

means:

> Get the address of `user`.

And:

```text
*pointer
```

means:

> Get the value that `pointer` points to.

Go also automatically handles pointer dereferencing when accessing struct fields.

So:

```go
pointer.Name
```

works even though `pointer` is a `*User`.

## Value vs Pointer

The basic idea is:

```text
User
 ↓
actual value

*User
 ↓
pointer to a User
```

When we pass:

```go
changeName(user, "Alex")
```

the function gets a copy.

When we pass:

```go
changeName(&user, "Alex")
```

the function gets a pointer to the original value.

## Use Cases

Sometimes we want a function to work with its own copy of a value.

Sometimes we want a function to modify the original value.

For example:

```go
func changeName(user User, name string) {
	user.Name = name
}
```

works with a copy.

While:

```go
func changeName(user *User, name string) {
	user.Name = name
}
```

works with the original value through a pointer.

This is the main reason we use pointers.
