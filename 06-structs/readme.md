# Struct

Struct is the Go answer to Object.

It's somewhat similar to a JS object or a Python dictionary, but a Go struct is a **typed value with a fixed set of fields**.

It defines the shape and we can define new types with structs.

```go
type User struct {
	id       int
	name     string
	age      int
	isActive bool
}
```

Then:

```go
user := User{
	id:       1,
	name:     "Daniel",
	age:      23,
	isActive: true,
}
```

And by the way, structs also have zero values too.

Like:

```go
var user User
```

Then:

```text
id       → 0
name     → ""
age      → 0
isActive → false
```

So:

```go
fmt.Printf("%+v\n", user)
```

returns:

```text
{id:0 name: age:0 isActive:false}
```

## Nested Structs

Structs can contain other structs.

For example:

```go
type Address struct {
	City    string
	Country string
}

type User struct {
	ID      string
	Name    string
	Address Address
}
```

Then:

```go
user := User{
	ID:   "123",
	Name: "Daniel",
	Address: Address{
		City:    "Paris",
		Country: "France",
	},
}
```

Access:

```go
fmt.Println(user.Address.City)
```

## Methods

In Go, methods are functions that are attached to a type.

For example:

```go
type User struct {
	name string
}

func (u User) getName() string {
	return u.name
}
```

Here:

```go
func (u User) getName()
```

is a method of the `User` type.

The `u User` part is called the **receiver**.

We can call the method using:

```go
user.getName()
```

The receiver is similar to `this` in TypeScript, but Go handles it differently.

## Value Receiver

A method can have a value receiver:

```go
func (u User) getName() string {
	return u.name
}
```

Here `u` is a copy of the `User`.

So if we modify the receiver:

```go
func (u User) changeName(name string) {
	u.name = name
}
```

the original `User` will not be changed because `u` is a copy.

## Pointer Receiver

We can instead use a pointer receiver:

```go
func (u *User) changeName(name string) {
	u.name = name
}
```

Now `u` is a pointer to the original `User`.

So modifying:

```go
u.name = name
```

modifies the original struct too.

This is why we commonly use pointer receivers when a method needs to modify the original value.

For example:

```go
func (u *User) modifyActivate(active bool) bool {
	u.isActive = active
	return active
}
```

Then:

```go
user.modifyActivate(false)
```

changes the original `user`.

## Exported Methods

Just like fields, methods can be exported or unexported based on their first letter.

Lowercase:

```go
func (u *User) getName() string
```

is **unexported**.

Uppercase:

```go
func (u *User) GetName() string
```

is **exported**.

This means another package can access `GetName`, but cannot access `getName`.
