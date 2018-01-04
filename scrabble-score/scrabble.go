// Package scrabble implements functions to return word scrabble scores
package scrabble

import (
	"strings"
)

var letterScorings = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score takes a word String and returns its score int
func Score(word string) int {
	word = strings.ToUpper(word)
	scrabbleScore := 0
	for _, char := range word {
		for letter, value := range letterScorings {
			if strings.ContainsRune(letter, char) {
				scrabbleScore += value
			}
		}
	}
	return scrabbleScore
}
