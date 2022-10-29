package prime_test

import (
	"testing"

	"github.com/MohammadBnei/go-exercice/goroutines/prime"
)

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime.Sieve(int64(10000))
	}
	b.ReportAllocs()
}

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime.PrimeNumbers(int64(10000))
	}
	b.ReportAllocs()
}
