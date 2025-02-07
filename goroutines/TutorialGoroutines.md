# Tutorial: Concurrent Prime Number Generation Using Goroutines and Channels

## Introduction
The Sieve of Eratosthenes is an ancient algorithm for finding all prime numbers up to a specified integer. Implementing this algorithm using goroutines in Go allows for parallel processing, making the task more efficient. In this tutorial, we'll guide you through implementing a concurrent prime number generator using goroutines and channels.

## Prerequisites
- Familiarity with Go programming syntax.
- Understanding of goroutines and their concurrency model.
- Basic knowledge of the Sieve of Eratosthenes algorithm.

## Step-by-Step Guide

### 1. Define Constants and Channels
- **Constants**: Define constants for the maximum limit (`maxLimit`) and the number of goroutines to spawn (`numGoroutines`).
- **Channels**: Create channels for communication between goroutines:
  - `primeChan`: A channel where primes are reported.
  - `numberChan`: A channel that distributes numbers for checking.

### 2. Generator Function
- **Purpose**: Generate all integers from 2 up to the maximum limit.
- **Implementation**:
```go
func generator(n int64) {
    for i := 2; i <= n; i++ {
        numberChan <- i
    }
}
```

### 3. Prime Checker Function
- **Purpose**: Check if a number is prime and report it through `primeChan`.
- **Implementation**:
```go
func primeChecker(n int64, numGoroutines int) {
    for n := range numberChan {
        go func(p int64) {
            if isPrime(p) {
                primeChan <- p
            }
        }(n)
    }
}
```

### 4. Sieve of Eratosthenes Function
- **Purpose**: Collect all primes up to the maximum limit using goroutines.
- **Implementation**:
```go
func sieve(maxLimit int64, numGoroutines int) []int64 {
    var primes []int64

    for i := 2; i <= maxLimit; i++ {
        go func(p int64) {
            if isPrime(p) {
                primes = append(primes, p)
                primeChan <- p
            }
        }(i)
    }

    return primes
}
```

### 5. Prime Checking Function
- **Purpose**: Determine if a number is prime.
- **Implementation**:
```go
func isPrime(n int64) bool {
    if n < 2 {
        return false
    }
    for i := 2; i <= sqrt(n); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
```

### 6. Coordinating Goroutines with Channels
- **Purpose**: Ensure proper communication between goroutines and avoid race conditions.
- **Implementation**:
```go
func main() {
    maxLimit := 10000
    numGoroutines := 10

    numberChan := make(chan int64, numGoroutines)
    primeChan := make(chan int64, numGoroutines)

    for i := 2; i <= maxLimit; i++ {
        go func(n int64) {
            for ; n <= maxLimit; n++ {
                primeChecker(n, numGoroutines)
            }
        }(i)
    }

    for p := range primeChan {
        primes = append(primes, p)
    }

    fmt.Printf("Primes up to %d: %v\n", maxLimit, primes)
}
```

## Explanation

1. **Generator**: The generator function iterates from 2 to `maxLimit`, sending each number through the `numberChan`.
2. **Prime Checker**: For each number received in `numberChan`, a goroutine is spawned to check if it's prime using the `isPrime` function.
3. **Sieve Function**: The sieve function coordinates all the goroutines, spawning one for each number and collecting primes into the `primeChan`.
4. **Prime Collection**: Each prime found is appended to the primes list and sent through `primeChan`.
5. **Output**: Finally, the collected primes are printed.

This approach efficiently utilizes goroutines to check each number in parallel, significantly reducing the time complexity compared to a single-threaded implementation. The use of channels ensures proper communication between concurrent goroutines, allowing for a scalable and efficient prime generation process.
