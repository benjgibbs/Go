package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func removeSpaces(ss []string) []string {
	var result []string
	for _, s := range ss {
		if s != " " && s != "" {
			result = append(result, s)
		}
	}
	return result
}

func p345() {
	file, err := os.Open("p345_matrix.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bufio.NewReader(file))
	orderedRemoval(scanner)
}

type Entry struct {
	val, row, col int
}

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	return e[i].val > e[j].val
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Entries []Entry

func cont(x int, xs []int) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func orderedRemoval(scanner *bufio.Scanner) {
	var entries Entries
	rows := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		parts = removeSpaces(parts)
		for col, str := range parts {
			entry := Entry{atoi(str), rows, col}
			entries = append(entries, entry)
			count++
		}
		rows++
	}
	fmt.Printf("Matrix dimensions: rows=%d, cols=%d\n", rows, count/rows)
	sort.Sort(entries)
	best := 0
	for i := 0; i < 1000; i++ {
		sum := run(entries, rows, count/rows)
		if sum > best {
			fmt.Printf("sum=%d, i=%d\n", sum, i)
			best = sum
		}
	}

}

func run(entries Entries, rows, cols int) int {
	seenRows := []int{}
	seenCols := []int{}
	sum := 0
	for i := 0; len(seenRows) < rows && len(seenCols) < cols; i++ {
		e := entries[i%(rows*cols)]
		if !cont(e.row, seenRows) && !cont(e.col, seenCols) && rand.Intn(3) != 0 {
			seenRows = append(seenRows, e.row)
			seenCols = append(seenCols, e.col)
			sum += e.val
		}
	}
	return sum
}
