package bitset

import (
	"testing"
)

func TestBitset32(t *testing.T) {
	bs := NewBitset(32)

	if bs.IsSet(30) {
		t.Error("bitset should be all cleared at initialization")
	}
	bs.Set(2)
	if !bs.IsSet(2) {
		t.Error("Test after set got false")
	}
	bs.Clear(2)
	if bs.IsSet(2) {
		t.Error("Test after clear got true")
	}
}

func TestBitset65(t *testing.T) {
	bs := NewBitset(65)

	if bs.IsSet(32) {
		t.Error("bitset should be all cleared at initialization")
	}
	bs.Set(32)
	if !bs.IsSet(32) {
		t.Error("Test after set got false")
	}

	bs.Set(33)
	if !bs.IsSet(33) {
		t.Error("Test after set got false")
	}

	bs.Set(64)
	if !bs.IsSet(64) {
		t.Error("Test after set 64 got false")
	}
}
