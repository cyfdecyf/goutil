// Package goutil provides utility functions for generating test data, shuffling data, etc.
package goutil

// Sequence generates an ordered slices of int in the range [l, h). Panics if
// l >= h.
func Sequence(l, h int) []int {
	if h <= l {
		panic("Sequence with h <= l")
	}
	s := make([]int, h-l, h-l)
	for i := 0; i < h-l; i++ {
		s[i] = l + i
	}
	return s
}
