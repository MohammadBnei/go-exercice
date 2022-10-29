package prime

import (
	"math"
)

func generator(limit int64, ch chan<- int64) {
	var i int64 = 2
	for ; i <= limit; i++ {
		ch <- i
	}

	close(ch)
}

func filter(src <-chan int64, dst chan<- int64, prime int64) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}

	close(dst)
}

func Sieve(limit int64) []int64 {
	var primes []int64
	ch := make(chan int64)

	go generator(limit, ch)

	for {
		prime, ok := <-ch
		if !ok {
			break
		}

		ch1 := make(chan int64)
		go filter(ch, ch1, prime)

		ch = ch1

		primes = append(primes, prime)
	}

	return primes
}

func PrimeNumbers(max int64) []int64 {
	var primes []int64
	var i int64 = 2

	for ; i < max; i++ {
		isPrime := true
		var j int64 = 2
		for ; j <= int64(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
