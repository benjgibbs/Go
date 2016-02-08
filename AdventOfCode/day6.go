package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const GridSz = 1000

type Grid [][]bool

func main() {
	var grid Grid
	grid = make(Grid, GridSz)
	for i := 0; i < GridSz; i++ {
		grid[i] = make([]bool, GridSz)
	}
	scanner := bufio.NewScanner(os.Stdin)
	pattern := regexp.MustCompile(`(.*) (\d+),(\d+) through (\d+),(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		match := pattern.FindStringSubmatch(line)
		if len(match) > 0 {
			//fmt.Printf("%s %s %s\n", match[1], match[2], match[3])
			a, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			c, _ := strconv.Atoi(match[4])
			d, _ := strconv.Atoi(match[5])
			switch match[1] {
			case "turn on":
				turnOn(grid, a, b, c, d)
			case "turn off":
				turnOff(grid, a, b, c, d)
			case "toggle":
				toggle(grid, a, b, c, d)
			default:
				fmt.Println("ERROR: ", match[1])
			}
		} else {
			fmt.Println("No match")
		}
	}
	fmt.Println("Num On Is: ", count(grid))
}

func turnOn(g Grid, a, b, c, d int) {
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			g[i][j] = true
		}
	}
}
func turnOff(g Grid, a, b, c, d int) {
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			g[i][j] = false
		}
	}
}
func toggle(g Grid, a, b, c, d int) {
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			g[i][j] = !g[i][j]
		}
	}
}

func count(grid Grid) int {
	count := 0
	for i := 0; i < GridSz; i++ {
		for j := 0; j < GridSz; j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}
