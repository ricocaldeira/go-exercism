// Package isogram implements functions to determine if a sentence is an isogram
package isogram

import (
	"strings"
)

// IsIsogram takes a sentence String and determines if is an isogram
// True: no repeating letters
// False: repeating letters
func IsIsogram(sentence string) bool {
	sentence = strings.Replace(sentence, " ", "", -1)
	sentence = strings.Replace(sentence, "-", "", -1)
	sentence = strings.ToUpper(sentence)
	letters := map[rune]bool{}
	for _, char := range sentence {
		if letters[char] {
			return false
		}
		letters[char] = true
	}
	return true
}
