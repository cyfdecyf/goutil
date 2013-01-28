package math

import (
	"testing"
)

func TestGCD(t *testing.T) {
	if GCD(0, 0) != 0 {
		t.Error("GCD(0, 0) should return 0")
	}
	if GCD(1, 0) != 1 {
		t.Error("GCD(1, 0) should return 1")
	}
	if GCD(8, 12) != 4 {
		t.Error("GCD(8, 12) should return 4")
	}
	if GCD(36, 24) != 12 {
		t.Error("GCD(36, 24) should return 12")
	}
}
