package main

import (
	"fmt"
)

type Sieve struct {
	upto  uint64
	sieve []bool
}

func createSieve(upto uint64) *Sieve {
	result := Sieve{upto, make([]bool, upto)}
	result.sieve[0] = false
	result.sieve[1] = false
	for i := uint64(2); i < upto; i++ {
		result.sieve[i] = true
	}
	for i := uint64(2); i < upto; i++ {
		if result.sieve[i] {
			for j := 2 * i; j < upto; j += i {
				result.sieve[j] = false
			}
		}
	}
	return &result
}

func (s *Sieve) isPrime(x uint64) bool {
	if x < s.upto {
		return s.sieve[x]
	}
	if x > (s.upto * s.upto) {
		panic(fmt.Sprintf("%d is too big for this sieve", x))
	}
	for i := uint64(2); (i * i) <= x; i++ {
		if s.sieve[i] && x%i == 0 {
			return false
		}
	}
	return true
}
