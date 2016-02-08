package main

import (
	"testing"
)

func Test_NothingToCollect(t *testing.T) {
	q := make(CountQueue)
	go func() {
		q <- map[string]int{}
		close(q)
	}()
	counts := collectCounts(q)
	if len(counts) != 0 {
		t.Error("counts should be empty")
	}
}

func Test_SingleCollection(t *testing.T) {
	q := make(CountQueue)
	go func() {
		q <- map[string]int{"a": 2}
		close(q)
	}()
	counts := collectCounts(q)
	if len(counts) != 1 {
		t.Error("counts should have one item")
	}
	if counts["a"] != 2 {
		t.Error("count of a should be 2")
	}
}

func Test_Merging(t *testing.T) {
	q := make(CountQueue)
	go func() {
		q <- map[string]int{
			"a": 2,
			"b": 1,
		}
		q <- map[string]int{
			"b": 2,
			"c": 1,
		}
		close(q)
	}()
	counts := collectCounts(q)
	if len(counts) != 3 {
		t.Error("counts should have one item")
	}
	if counts["a"] != 2 {
		t.Error("count of a should be 2")
	}
	if counts["b"] != 3 {
		t.Error("count of b should be 3")
	}
	if counts["c"] != 1 {
		t.Error("count of c should be 1")
	}
}
