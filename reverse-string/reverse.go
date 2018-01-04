// Package reverse implements functions to reverse strings
package reverse

// String takes a string and returns its reverse
func String(sentence string) string {
	reverseSentence := ""
	for _, char := range sentence {
		reverseSentence = string(char) + reverseSentence
	}
	return reverseSentence
}
