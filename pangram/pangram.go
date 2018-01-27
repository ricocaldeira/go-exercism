// Package pangram implements IsPangram function to identify a pangram sentence
package pangram

import (
	"strings"
)

const pangram = "abcdefghijklmnopqrstuvwxyz"

// IsPangram determines if a sentence is a pangram
// A pangram (Greek: παν γράμμα, pan gramma, "every letter")
// is a sentence using every letter of the alphabet at least once.
func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	for _, letter := range pangram {
		if !strings.ContainsRune(sentence, letter) {
			return false
		}
	}
	return true
}
