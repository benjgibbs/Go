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
	x = x / 10
	sum := sumOfDigits(x)
	div := x / sum
	return s.isPrime(div)
}

func p387() {
	s := createSieve(10000000)
	sum := uint64(0)
	harshads := []uint64{}

	for i := uint64(10); i < 100; i++ {
		if isHarshad(i) {
			harshads = append(harshads, i)
		}
		if isStrong(i, s) && s.isPrime(i) {
			sum += i
		}
	}
	for i := 0; i < (14 - 2); i++ {
		fmt.Printf("Searching %d harshads. i=%d, sum=%d \n", len(harshads), i, sum)
		nextHarshads := []uint64{}
		for _, x := range harshads {
			for j := uint64(0); j < 10; j++ {
				n := x*10 + j
				if isHarshad(n) {
					nextHarshads = append(nextHarshads, n)
				}
				if isStrong(n, s) && s.isPrime(n) {
					sum += n
				}
			}
		}
		harshads = nextHarshads
	}
	fmt.Printf("Total is: %d\n", sum)
}
