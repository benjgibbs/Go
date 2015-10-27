package main

import (
	"testing"
)

func Test_fact(t *testing.T) {
	type TestCase struct {
		input, expect uint64
	}
	tcs := []TestCase{{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}}
	for _, tc := range tcs {
		result := fact(tc.input)
		if result != tc.expect {
			t.Fatalf("Expecting %d but got %d", tc.expect, result)
		}
	}
}

func Test_contains(t *testing.T) {

	type TestCase struct {
		inputArray  []int
		inputSearch int
		resultFound bool
		resultPos   int
	}
	tcs := []TestCase{{[]int{1, 2, 3}, 1, true, 0},
		{[]int{1, 2, 3}, 2, true, 1},
		{[]int{1, 2, 3}, 3, true, 2},
		{[]int{1, 2, 3}, 4, false, -1}}

	for _, tc := range tcs {
		found, pos := containsInt(tc.inputSearch, tc.inputArray)
		if found != tc.resultFound || pos != tc.resultPos {
			t.Fatalf("Expecting %t/%d but got %t/%d", tc.resultFound, tc.resultPos, found, pos)
		}
	}

}
