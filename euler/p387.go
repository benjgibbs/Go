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

func isStrong(x uint64, s *Sieve) bool {
	sum := sumOfDigits(x)
	div := x / sum
	return s.isPrime(uint64(div))
}

func p387() {
	s := createSieve(10000001)
	sum := uint64(0)
	harshads := []uint64{}

	for i := uint64(10); i < 100; i++ {
		if isHarshad(i) {
			harshads = append(harshads, i)
		}
		if s.isPrime(i) && isStrong(i/10, s) {
			sum += i
		}
	}
	for i := 0; i < (14 - 2); i++ {
		nextHarshads := []uint64{}
		for _, x := range harshads {
			for j := uint64(0); j < 10; j++ {
				n := x*10 + j
				if isHarshad(n) {
					nextHarshads = append(nextHarshads, n)
				}
				if s.isPrime(n) && isStrong(n/10, s) {
					sum += n
				}
			}
		}
		harshads = nextHarshads
	}

	fmt.Printf("Total is: %d expected (696067597313468)\n", sum)
}
