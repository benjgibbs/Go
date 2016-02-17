package permutations

import (
	"testing"
)

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
