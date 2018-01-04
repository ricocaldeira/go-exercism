// Package acronym supplies functions to reduce sentences to acronyms
package acronym

import (
	"strings"
)

// Abbreviate reduces String sentences to String acronyms
func Abbreviate(s string) string {
	words := sentenceToWords(s)
	return acronym(words)
}

func sentenceToWords(sentence string) []string {
	sentence = strings.Replace(sentence, "-", " ", -1)
	return strings.Split(sentence, " ")
}

func acronym(words []string) string {
	var acronym string

	for _, word := range words {
		acronym += string(word[0])
	}

	return strings.ToUpper(acronym)
}
