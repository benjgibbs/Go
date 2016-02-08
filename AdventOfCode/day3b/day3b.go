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
	x1, y1, x2, y2 := 0, 0, 0, 0
	houses[Point{0, 0}]++
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			c := line[i]
			if i%2 == 0 {
				switch c {
				case '>':
					x1 += 1
				case '<':
					x1 -= 1
				case '^':
					y1 += 1
				case 'v':
					y1 -= 1
				}
				houses[Point{x1, y1}]++
			} else {
				switch c {
				case '>':
					x2 += 1
				case '<':
					x2 -= 1
				case '^':
					y2 += 1
				case 'v':
					y2 -= 1
				}
				houses[Point{x2, y2}]++
			}
		}
	}
	fmt.Println("Number of houses: ", len(houses))
}
