package goutil

import (
	"reflect"
	"testing"
)

func TestReverseRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr_rev := []int{1, 5, 4, 3, 2, 6, 7, 8}

	ReverseRange(IntSwapSlice(arr_rev), 1, 5)
	if !reflect.DeepEqual(arr, arr_rev) {
		t.Error("ReverseRange incorrect, reversed array:", arr_rev)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arr_rev := []int{5, 4, 3, 2, 1}

	Reverse(IntSwapSlice(arr_rev))
	if !reflect.DeepEqual(arr, arr_rev) {
		t.Error("Reverse incorrect, reversed array:", arr_rev)
	}
}

func TestSequence(t *testing.T) {
	s := Sequence(1, 2)
	if len(s) != 1 {
		t.Error("Sequence should return a length 1 slice if l+1 = h")
	}

	const l = 10
	const h = 100
	s = Sequence(l, h)
	for i := 0; i < h-l; i++ {
		if s[i] != l+i {
			t.Fatalf("Sequence(%d, %d) item %d got %d, should be %d\n",
				l, h, i, s[i], l+i)
		}
	}
}
