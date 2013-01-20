package goutil

import (
	"testing"
)

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
