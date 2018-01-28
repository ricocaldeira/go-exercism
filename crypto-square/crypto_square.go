// Package cryptosquare
package cryptosquare

import (
	"unicode"
)

func Encode(message string) string {
	encodedMessage := ""
	for _, r := range message {
		if unicode.IsLetter(r) {
			encodedMessage = encodedMessage + string(r)
		}
	}

	// Calculate number of columns
	// c: columns r: rows
	// such that `c >= r` and `c - r <= 1`
	return encodedMessage
}
