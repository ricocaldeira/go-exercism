package wordcount

import "strings"

// Frequency maps words and ocurrencies count
type Frequency map[string]int

// WordCount counts words frequency and returns Frequency type
func WordCount(phrase string) Frequency {
	words := strings.Split(phrase, " ")
	frequency := Frequency{}

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}
