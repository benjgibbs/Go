package main

import (
	"testing"
)

func TestSievePrimes(t *testing.T) {
	s := createSieve(100)
	primes := map[int]struct{}{2: {}, 3: {}, 5: {}, 7: {}, 11: {}, 13: {},
		17: {}, 19: {}, 23: {}, 29: {}, 31: {}, 37: {}, 41: {}, 43: {}, 47: {}}
	for i := 0; i < 50; i++ {
		_, expect := primes[i]
		result := s.isPrime(uint64(i))
		if expect != result {
			t.Errorf("%d wanted %t got %t", i, expect, result)
		}
	}
}

func TestSieveHigherPrimes(t *testing.T) {
	s := createSieve(10)
	primes := map[int]struct{}{2: {}, 3: {}, 5: {}, 7: {}, 11: {}, 13: {},
		17: {}, 19: {}, 23: {}, 29: {}, 31: {}, 37: {}, 41: {}, 43: {}, 47: {}}
	for i := 0; i < 50; i++ {
		_, expect := primes[i]
		result := s.isPrime(uint64(i))
		if expect != result {
			t.Errorf("%d wanted %t got %t", i, expect, result)
		}
	}
}
