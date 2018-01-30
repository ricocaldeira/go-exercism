// Package cryptosquare
package cryptosquare

import (
	"strings"
	"unicode"
)

// Encode Implement the classic method for composing secret messages called a square code.
func Encode(message string) string {
	encodedMessage := ""
	for _, l := range strings.ToLower(message) {
		if unicode.IsDigit(l) || unicode.IsLetter(l) {
			encodedMessage = encodedMessage + string(l)
		}
	}
	edge := 1
	for edge*edge < len(encodedMessage) {
		edge++
	}
	square := make([]string, edge)
	for i, c := range encodedMessage {
		square[i%edge] += string(c)
	}
	return strings.Join(square, " ")
}
