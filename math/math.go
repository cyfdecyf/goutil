package math

// gcd expects m <= n.
func gcd(m, n int) int {
	if m == 0 {
		return n
	}
	return gcd(n%m, m)
}

// GCD returns the greatest common divisor of m and n. This function defines
// GCD(0, 0) = 0. (This confirms to J, which uses gcd as the boolean OR.)
func GCD(m, n int) int {
	if m > n {
		m, n = n, m
	}
	if m == 0 && n == 0 {
		return 0
	}
	return gcd(m, n)
}
