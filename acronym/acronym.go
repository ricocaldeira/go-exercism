// Package acronym supplies functions to reduce sentences to acronyms
package acronym

import (
	"strings"
)

// Abbreviate reduces String sentences to String acronyms
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	words := strings.Split(s, " ")
	var acronym string

	for index := range words {
		acronym += string(words[index][0])
	}

	acronym = strings.ToUpper(acronym)
	return acronym
}
