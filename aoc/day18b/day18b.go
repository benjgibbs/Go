package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Grid [][]bool

var grid Grid = Grid{}

func main() {
	iters, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = (line[i] == '#')
		}
		grid = append(grid, row)
	}
	printGrid()
	for it := 0; it < iters; it++ {
		grid = step()
		printGrid()
		fmt.Println("Count: ", count())
	}
}

func step() Grid {
	dim := len(grid)
	result := make(Grid, dim)
	for x := 0; x < dim; x++ {
		row := make([]bool, dim)
		for y := 0; y < dim; y++ {
			switch countNbs(x, y) {
			case 3:
				row[y] = true
			case 2:
				if grid[x][y] {
					row[y] = true
				}
			}
		}
		result[x] = row
	}
	result[0][0] = true
	result[0][dim-1] = true
	result[dim-1][0] = true
	result[dim-1][dim-1] = true

	return result
}

func countNbs(x, y int) int {
	nbs := 0

	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i >= 0 && i < len(grid) &&
				j >= 0 && j < len(grid) &&
				!(i == x && j == y) &&
				grid[i][j] {
				nbs++
			}
		}
	}
	return nbs
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			if row[j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func count() int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] {
				res++
			}
		}
	}
	return res
}
