package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	houses := make(map[Point]int)
	x, y := 0, 0
	houses[Point{x, y}]++
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			switch c {
			case '>':
				x += 1
			case '<':
				x -= 1
			case '^':
				y += 1
			case 'v':
				y -= 1
			}
			houses[Point{x, y}]++
		}
	}
	fmt.Println("Number of houses: ", len(houses))
}
