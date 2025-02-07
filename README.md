# Go Programming Exercises Course

Welcome to the Go Programming Exercises course! This repository contains hands-on exercises to help you master key Go language concepts through practical implementation.

## Course Objectives
- Learn Go's concurrency model with goroutines and channels
- Understand memory management with pointers and data structures
- Implement common patterns like worker pools and broadcasters
- Practice test-driven development (TDD) in Go
- Compare different approaches (sync vs async, sequential vs concurrent)

## Key Topics Covered
🧩 **Concurrency Patterns**
- Goroutines (see `goroutines/example/`)
- WaitGroups (`waitgroup.go`)
- Prime number sieve (`goroutines/prime/`)
- Channel-based messaging (`channels/`)

💡 **Memory Management**
- Linked list implementation (`pointers/`)
- Node structure and interfaces
- Pointer manipulation techniques

📡 **Channel Patterns**
- Broadcaster implementation (`channels/broadcast.go`)
- Simple message passing (`channels/simple.go`)
- Channel-based command line tool (`channels/cmd/main.go`)

## Getting Started
```bash
# Run all tests
go test ./...

# Benchmark prime number implementations
go test -bench=. ./goroutines/prime/

# Run concurrency examples
go run ./goroutines/example/*.go
```

## Course Structure
1. Start with `goroutines/example/` to learn basic concurrency
2. Move to `channels/` for communication patterns
3. Explore `pointers/` for memory management
4. Review tests in `*_test.go` files for validation approaches

> 💡 Pro Tip: Use `go test -v` to see verbose test output and better understand the validation checks!

This course is designed to be hands-on - clone the repo and modify the code as you follow along with the exercises. Each directory contains progressively more challenging implementations to help you level up your Go skills!
