package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (chars []string) {
	seriesSize := len(s) - n + 1

	for i := 0; i < seriesSize; i++ {
		chars = append(chars, s[i:n+i])
	}
	return chars
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
