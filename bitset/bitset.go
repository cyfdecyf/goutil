// Package bitset implementes bit set operations.
package bitset

// Bitset represents a fixed-size sequence of bits.
type Bitset struct {
	bit []int32
	n   int
}

const shift = 5
const mask = (1 << shift) - 1

// NewBitset creates a Bitset representing n bits.
// Panics if n <= 0.
func NewBitset(n int) *Bitset {
	if n <= 0 {
		panic("NewBitset with size <= 0")
	}
	nn := (n-1)/32 + 1
	return &Bitset{bit: make([]int32, nn, nn), n: n}
}

// Set set the i'th bit in the bit set.
// Panics if i < 0 or i >= bit set size.
func (bs *Bitset) Set(i int) {
	if i < 0 || i >= bs.n {
		panic("Bitset.Set with invalid bit index")
	}
	bs.bit[i>>shift] |= 1 << (uint32(i) & mask)
}

// Clear clear the i'th bit in the bit set.
// Panics if i < 0 or i >= bit set size.
func (bs *Bitset) Clear(i int) {
	if i < 0 || i >= bs.n {
		panic("Bitset.Set with invalid bit index")
	}
	bs.bit[i>>shift] &^= 1 << (uint32(i) & mask) // using the bit clear operator
}

// IsSet returns true if the i'th bit is set.
// Panics if i < 0 or i >= bit set size.
func (bs *Bitset) IsSet(i int) bool {
	if i < 0 || i >= bs.n {
		panic("Bitset.Set with invalid bit index")
	}
	return (bs.bit[i>>shift] & (1 << (uint32(i) & mask))) != 0
}
