package rand

import (
	"github.com/cyfdecyf/goutil"
	"sort"
	"testing"
)

func TestRand(t *testing.T) {
	test := func(l, h int) {
		const n = 20
		for i := 0; i < n; i++ {
			r := Rand(l, h)
			if r < l || r >= h {
				t.Errorf("Rand generates integer outside range: [%d, %d)\n", l, h)
			}
		}
	}
	rg := [][2]int{
		{10, 11},
		{1, 100},
	}
	for _, r := range rg {
		test(r[0], r[1])
	}
}

func TestGenKRandomLessN(t *testing.T) {
	const k = 30
	const n = 100
	a := GenKRandomLessN(k, n)
	// Only check if there's duplicates
	// TODO: How to check randomness?
	sort.Ints(a)
	for i := 1; i < k; i++ {
		if a[i] == a[i-1] {
			t.Error("Duplicates found:", a)
		}
		if a[i] >= n {
			t.Error("Got integer larger than n:", a)
		}
	}
}

func TestShuffleSlice(t *testing.T) {
	const n = 100
	seq := goutil.Sequence(1, n+1)

	shf1 := make([]int, n, n)
	shf2 := make([]int, n, n)
	copy(shf1, seq)
	copy(shf2, seq)

	Shuffle(IntShuffleSlice(shf1))
	Shuffle(IntShuffleSlice(shf2))

	diff := false
	for i := 0; i < n; i++ {
		if shf1[i] != shf2[i] {
			diff = true
			break
		}
	}
	if !diff {
		t.Error("Shuffle 100 element slice twice got same result, probably error in implementation.")
	}
}
