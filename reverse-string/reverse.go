// Package reverse implements functions to reverse strings
package reverse

// String takes a string and returns its reverse
func String(sentence string) string {
	runes := []rune(sentence)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
