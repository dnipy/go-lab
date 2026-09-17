# Arrays, Slices, Maps

## Arrays

Let's get started with Array.

In data structures, an array is a way to store multiple values in memory and it has a fixed length.

And we define it in Go exactly the way the data structure really is:

```go
numbers := [3]int{10, 20, 30}
```

We can modify and assign values to existing indexes:

```go
numbers[0] = 99
```

But we can't access an index that doesn't exist:

```go
numbers[3] = 40
```

because the array has a fixed length of `3`.

But in real-world apps we technically don't know the array length.

For example:

- users
- messages
- orders

What should we do?

That's when slices shine.

---

## Slices

Slices on first sight are arrays without a fixed size.

Like JS/TS/Python and etc:

```go
numbers := []int{10, 20, 30}
```

The syntax is similar to an array, but under the hood they are different.

A slice is a data structure that contains information about:

- a pointer to an underlying array
- length
- capacity

So we can think about it like:

```text
slice
 ├── pointer → underlying array
 ├── len
 └── cap
```

When we append new values and the current underlying array doesn't have enough capacity, Go creates a new bigger underlying array and the slice points to it.

So basically, the slice is an abstraction for handling an underlying array dynamically.

And it's fun.

Because if we copy a slice:

```go
numbers := []int{10, 20, 30}
other := numbers
```

both slices can point to the same underlying array.

So if we change an existing index:

```go
other[0] = 99
```

the other slice can see the change too:

```text
numbers → [99 20 30]
other   → [99 20 30]
```

But if we append to a slice, it may need a new underlying array.

If that happens, the two slices become disconnected from each other.

So we should not simply think:

> slice = reference

A better mental model is:

> A slice is a small value that points to an underlying array and contains its length and capacity.

---

## Slice Operations

For slices we have:

- `len`
- `cap`
- `append`

### `len`

`len` takes an array or slice and returns its current length.

```go
len(numbers)
```

### `cap`

`cap` returns the capacity of the slice.

```go
cap(numbers)
```

The basic difference is:

```text
len → how many elements currently exist

cap → how much capacity the slice currently has
```

### `append`

`append` adds elements to a slice:

```go
numbers = append(numbers, 40)
```

`append` may reuse the existing underlying array if there is enough capacity.

If there isn't enough capacity, Go creates a new underlying array.

---

## Nil Slices

Slices can also be `nil`:

```go
var numbers []int
```

A nil slice has:

```text
len → 0
```

and:

```go
len(numbers)
cap(numbers)
```

are safe.

We can also append to a nil slice:

```go
numbers = append(numbers, 10)
```

But we can't access an element that doesn't exist:

```go
numbers[0]
```

because the slice is empty.

---

# Maps

And the Maps?

It's a key-value data structure.

It's similar to JS objects and other key-value stores, but Go maps have a defined key type and value type.

We can define one like:

```go
users := map[string]string{
	"1": "Daniel",
	"2": "Alex",
}
```

And for access:

```go
users["3"]
```

For add and modify:

```go
users["1"] = "Danial"
```

And delete:

```go
delete(users, "2")
```

But be aware:

```go
name := users["999"]
```

might not exist.

Go doesn't throw an error for this.

Based on the map's value type, it returns the zero value.

For example, because our map is:

```go
map[string]string
```

the zero value is:

```text
""
```

But for cases where we need to know exactly whether the data exists or not, we can use:

```go
name, ok := users["999"]
```

Now:

```text
name → zero value if it doesn't exist

ok → true/false depending on whether the key exists
```

---

## Read and Write from a Nil Map

We can define a nil map:

```go
var users map[string]string
```

It's okay to read from a nil map:

```go
name := users["1"]
```

We get the zero value.

But for writing:

```go
users["1"] = "Daniel"
```

we get a panic.

For writing, we need to initialize the map first:

```go
users := make(map[string]string)
```

or:

```go
users := map[string]string{}
```

Then:

```go
users["1"] = "Daniel"
```

works.

---

## `make`

And `make` is not only for maps.

It can be used for:

- slices
- maps
- channels

For example:

```go
users := make(map[string]string)
```

or:

```go
numbers := make([]int, 10)
```

The second one creates a slice with a length of `10`.
