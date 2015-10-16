package main

import (
	"fmt"
)

func sumOfDigits(x uint64) uint64 {
	sum := uint64(0)
	for x > 0 {
		sum += x % 10
		x = x / 10
	}
	return sum
}

func isHarshad(x uint64) bool {
	sum := sumOfDigits(x)
	return x%sum == 0
}

func isRTH(x uint64) bool {
	for x > 0 {
		if !isHarshad(x) {
			return false
		}
		x = x / 10
	}
	return true
}

func isStrong(x uint64, s *Sieve) bool {
	sum := sumOfDigits(x)
	div := x / sum
	return s.isPrime(uint64(div))
}

func isRT(x uint64) bool {
	for x > 0 {
		if !isHarshad(x) {
			return false
		}
		x = x / 10
	}
	return true
}

func isSRTHP(x uint64, s *Sieve) bool {
	if !s.isPrime(uint64(x)) {
		return false
	}
	x = x / 10
	if !isStrong(x, s) {
		return false
	}
	return isRT(x)
}

func p387() {
	const limit = uint64(100)
	s := createSieve(limit + 1)
	sum := uint64(0)
	for i := uint64(10); i < limit*limit; i++ {
		if isSRTHP(i, s) {
			sum += i
		}
	}
	fmt.Printf("Total is: %d\n", sum)
}
