package main

import (
	"fmt"
	"os"
)

const pLen = 8

func toString(num uint64) string {
	output := ""
	for len(output) < pLen {
		c := num % 26
		output = string('a'+c) + output
		num = num / uint64(26)
	}
	return output
}

func hasIncreasingStraight(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			return true
		}
	}
	return false
}

func contains(s string, banned []rune) bool {
	for _, c := range s {
		for _, b := range banned {
			if c == b {
				return true
			}
		}
	}
	return false
}

func fromString(s string) uint64 {
	var res uint64 = 0
	var base uint64 = 1
	for i := len(s) - 1; i >= 0; i-- {
		d := uint64(s[i] - 'a')
		res += d * base
		base *= uint64(26)
	}
	return res
}

func countPairs(s string) int {
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
			i = i + 2
		}
	}
	return count
}

func main() {
	fmt.Println(toString(fromString("abcdefgh")))
	passStr := os.Args[1]

	for {
		if hasIncreasingStraight(passStr) && !contains(passStr, []rune{'i', 'o', 'l'}) &&
			countPairs(passStr) > 1 {
			fmt.Println("New Password: ", passStr)
			break
		}
		passStr = toString(fromString(passStr) + 1)
	}

}
