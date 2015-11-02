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

func TestSingleDigitHeapPermutation(t *testing.T) {
	res := Heap([]int{1})
	if len(res) != 1 {
		t.Errorf("Length should be 1 not %d", len(res))
	}
	if len(res[0]) != 1 {
		t.Errorf("Length should be 1 not %d", len(res[0]))
	}
	if res[0][0] != 1 {
		t.Errorf("Should be 1 not %d", res)
	}
}

func TestDoubleDigitHeapPermutation(t *testing.T) {
	res := Heap([]int{1, 2})
	if len(res) != 2 {
		t.Errorf("Length should be 2 not %d", len(res))
	}
	if len(res[0]) != 2 {
		t.Errorf("Length should be 1 not %d", len(res[0]))
	}
	if res[0][0] != 1 {
		t.Errorf("Should be 1 not %d", res)
	}
	if res[0][1] != 2 {
		t.Errorf("Should be 1 not %d", res)
	}
	if res[1][0] != 2 {
		t.Errorf("Should be 1 not %d", res)
	}
	if res[1][1] != 1 {
		t.Errorf("Should be 1 not %d", res)
	}
}

func TestTrebleDigitHeapPermutation(t *testing.T) {
	res := Heap([]int{1, 2, 3})
	if len(res) != 6 {
		t.Errorf("Length should be 6 not %d", len(res))
	}
}
