package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const GridSz = 1000

type Grid [][]int

var grid Grid

func main() {
	grid = make(Grid, GridSz)
	for i := 0; i < GridSz; i++ {
		grid[i] = make([]int, GridSz)
	}
	scanner := bufio.NewScanner(os.Stdin)
	pattern := regexp.MustCompile(`(.*) (\d+),(\d+) through (\d+),(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		match := pattern.FindStringSubmatch(line)
		if len(match) > 0 {
			a, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			c, _ := strconv.Atoi(match[4])
			d, _ := strconv.Atoi(match[5])
			switch match[1] {
			case "turn on":
				turnOn(a, b, c, d)
			case "turn off":
				turnOff(a, b, c, d)
			case "toggle":
				toggle(a, b, c, d)
			default:
				fmt.Println("ERROR: ", match[1])
			}
		} else {
			fmt.Println("No match")
		}
	}
	fmt.Println("Total Brightness Is: ", count())
}

func turnOn(a, b, c, d int) {
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			grid[i][j] += 1
		}
	}
}
func turnOff(a, b, c, d int) {
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			if grid[i][j] > 0 {
				grid[i][j] -= 1
			}
		}
	}
}
func toggle(a, b, c, d int) {
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			grid[i][j] += 2
		}
	}
}

func count() int {
	count := 0
	for i := 0; i < GridSz; i++ {
		for j := 0; j < GridSz; j++ {
			count += grid[i][j]
		}
	}
	return count
}
