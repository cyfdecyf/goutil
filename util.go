// Package goutil provides utility functions for generating test data, shuffling data, etc.
package goutil

// SwapInterface defines the Len and Swap method. Types implements this interface
// can be used by some functions in goutil.
type SwapInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// IntSwapSlice attaches the methods of SwapSlice to []int.
type IntSwapSlice []int

func (p IntSwapSlice) Len() int      { return len(p) }
func (p IntSwapSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// StringSwapSlice attaches the methods of SwapSlice to []string.
type StringSwapSlice []string

func (p StringSwapSlice) Len() int      { return len(p) }
func (p StringSwapSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//ReverseRange reverse elements in arr in the range of [l, h).
func ReverseRange(arr SwapInterface, l, h int) {
	n := h - l
	for i := 0; i < n/2; i++ {
		arr.Swap(l+i, h-i-1)
	}
}

// Reverse reverses the entire arr. It's the same as ReverseRange(0,
// arr.Len()).
func Reverse(arr SwapInterface) {
	n := arr.Len()
	for i := 0; i < n/2; i++ {
		arr.Swap(i, n-i-1)
	}
}

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
