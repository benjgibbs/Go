package main

import (
	"bufio"
	"fmt"
	"log"
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

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	return e[i].val > e[j].val
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Entry struct {
	val, row, col int
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

func p345() {
	file, err := os.Open("p345_given.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bufio.NewReader(file))

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
	fmt.Printf("row=%d, numCols=%d\n", rows, count/rows)
	sort.Sort(entries)
	seenRows := []int{}
	seenCols := []int{}
	sum := 0
	for i := 0; len(seenRows) < rows && len(seenCols) < count/rows; i++ {
		e := entries[i]
		if !cont(e.row, seenRows) && !cont(e.col, seenCols) {
			seenRows = append(seenRows, e.row)
			seenCols = append(seenCols, e.col)
			sum += e.val
			fmt.Printf("Adding %d\n", e)
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
