# Go Lab

<p align="center">
  <img src="https://go.dev/images/gophers/ladder.svg" alt="Go Logo" width="180">
</p>

<p align="center">
  A hands-on Go learning repository for a TypeScript developer.
</p>

---

## About

This repository is my personal Go lab.

The goal is not to learn programming basics, but to learn **Go from the core** by building small, focused examples and understanding how Go differs from languages I already know, especially TypeScript.

Each section focuses on one Go concept and contains practical code and notes.

The repository will eventually move from language fundamentals into concurrency and systems-oriented Go.

---

# Learning Path

## Go Ecosystem

### 001 — Basics

Learning the Go toolchain and basic commands.

```text
go run
go build
go fmt
```

↓

### 002 — Packages in Go

Understanding:

- Modules
- Packages
- Imports
- Package scope
- Exported / unexported names
- `go list`

↓

### 003 — Dependencies

Understanding:

- External packages
- `go get`
- `go.mod`
- `go.sum`
- Dependency checksums
- `go mod tidy`
- `go list -m all`

---

# Go Core

### 004 — Values & Variables

Understanding:

- `var`
- `const`
- `:=`
- Type inference
- Reassignment
- Zero values
- Package scope vs function scope

↓

### 005 — Functions

Understanding:

- Function syntax
- Parameters
- Return values
- Multiple return values
- Blank identifier `_`
- Basic error-return pattern

↓

### 006 — Structs

Understanding:

- Struct types
- Struct literals
- Struct fields
- Zero-value structs
- Nested structs

↓

### 007 — Methods

Understanding:

- Methods
- Receivers
- Value receivers
- Pointer receivers
- Exported / unexported methods

↓

### 008 — Pointers

Understanding:

- Addresses
- `&`
- `*`
- Dereferencing
- Value semantics
- Pointer semantics
- Pointer receivers

↓

### 009 — Slices & Maps

Understanding:

- Arrays
- Slices
- Slice length and capacity
- `append`
- Slice copying
- Maps
- Map lookup
- `nil` maps

↓

### 010 — Interfaces

Understanding:

- Interfaces
- Implicit interface implementation
- Interface values
- Type assertions
- Type switches
- `any`

↓

### 011 — Errors

Understanding:

- `error`
- `errors.New`
- `fmt.Errorf`
- Error wrapping
- `errors.Is`
- `errors.As`
- Custom errors

↓

### 012 — Generics

Understanding:

- Type parameters
- Constraints
- Generic functions
- Generic types
- When to use generics
- When **not** to use generics

↓

### 013 — Defer / Panic / Recover

Understanding:

- `defer`
- Execution order
- `panic`
- `recover`
- When these mechanisms should and shouldn't be used

↓

### 014 — Files & I/O

Understanding:

- Files
- Readers
- Writers
- `io`
- `os`
- Buffers
- Streaming data

↓

### 015 — JSON

Understanding:

- JSON encoding
- JSON decoding
- Struct tags
- `encoding/json`
- Custom JSON behavior

↓

### 016 — Testing

Understanding:

- `testing`
- Unit tests
- Table-driven tests
- Test helpers
- Benchmarks
- Examples

---

# Concurrency

### 017 — Goroutines

Understanding:

- Goroutines
- The Go scheduler
- Lightweight concurrent execution

↓

### 018 — Channels

Understanding:

- Channels
- Sending / receiving
- Blocking
- Buffered channels
- Closing channels
- `select`

↓

### 019 — Context

Understanding:

- `context.Context`
- Cancellation
- Deadlines
- Timeouts
- Request-scoped values

↓

### 020 — Synchronization

Understanding:

- `sync.Mutex`
- `sync.RWMutex`
- `sync.WaitGroup`
- `sync.Once`
- Atomic operations
- Race conditions
- Race detector

---

# Goal

After completing this repository, the goal is to be comfortable enough with Go's language, tooling, standard library, error handling, and concurrency model to start building real backend services.

The next repository will focus on **building actual services with Go** rather than isolated language exercises.

---

## Philosophy

> Learn the standard library first.

Prefer Go's standard library whenever it provides a reasonable solution.

External dependencies should be introduced when they provide meaningful value rather than simply because a package exists for something.

---

## Progress

- [x] Go ecosystem
- [x] Basics
- [x] Packages
- [x] Dependencies
- [x] Values & variables
- [x] Functions
- [x] Structs
- [x] Methods
- [ ] Pointers
- [ ] Slices & Maps
- [ ] Interfaces
- [ ] Errors
- [ ] Generics
- [ ] Defer / Panic / Recover
- [ ] Files & I/O
- [ ] JSON
- [ ] Testing
- [ ] Goroutines
- [ ] Channels
- [ ] Context
- [ ] Synchronization
