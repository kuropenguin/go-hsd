// This is package hsd, which is a simple package for calculating the Hamming
package hsd

// StringDistance calculates the Hamming distance between two strings.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// Distance calculates the Hamming distance between two slices of runes.
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
