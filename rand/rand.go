// Package rand provides several utility functions using the math.rand package.
package rand

import (
	"github.com/cyfdecyf/goutil"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Rand returns a pseudo-random number in [l, h).
// It panics if l >= h.
func Rand(l, h int) int {
	if l >= h {
		panic("Rand with l >= h")
	}
	n := h - l
	return l + rand.Intn(n)
}

// GenKRandomLessN generate k unique random integers between [0, n-1]
// Panics if k > n.
func GenKRandomLessN(k, n int) []int {
	if k > n {
		panic("GenKRandomLessN with k > n")
	}
	// If n is much larger than k, returning arr will waste too much memory.
	arr := make([]int, n, n)
	res := make([]int, k, k)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for i := 0; i < k; i++ {
		r := Rand(i, n)
		res[i], arr[r] = arr[r], arr[i]
	}
	return res
}

// ShuffleN randomly shuffles the first n item in the data slice. It uses
// Knuth's shuffle algorithm. Panics if n is larger than data.Len()-1 or is
// negative.
func ShuffleN(data goutil.SwapInterface, n int) {
	if n < 0 {
		panic("ShuffleN with negative n")
	}
	if n > data.Len() {
		panic("ShuffleN n too large")
	}
	for i := 0; i < n; i++ {
		data.Swap(i, Rand(i, n))
	}
}

// Shuffle randomly shuffles the all the item in the data slice.
func Shuffle(data goutil.SwapInterface) {
	ShuffleN(data, data.Len())
}
