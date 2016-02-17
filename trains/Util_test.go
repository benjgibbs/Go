package main

import (
	"testing"
)

func check(t *testing.T, expect, got []byte) {
	if len(expect) != len(got) {
		t.Error("No match:", expect, got)
	}
	for i, c := range expect {
		if c != got[i] {
			t.Error("No match:", expect, got)
		}
	}
}

func TestSlicing(t *testing.T) {
	bytes := []byte{1, 2, 3, 4}
	check(t, bytes, []byte{1, 2, 3, 4})
	check(t, bytes[:len(bytes)-1], []byte{1, 2, 3})
	check(t, bytes[:1], []byte{1})
	check(t, bytes[1:], []byte{2, 3, 4})
}
