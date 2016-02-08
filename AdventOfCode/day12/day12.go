package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	numberRe := regexp.MustCompile(`(-?\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := numberRe.FindAllString(line, -1)
		if numbers != nil {
			for _, nstr := range numbers {
				val, err := strconv.Atoi(nstr)
				if err != nil {
					panic("Bad number, reisie your regex: " + nstr)
				}
				sum += val
			}
		}
	}
	fmt.Println("Sum is: ", sum)
}
