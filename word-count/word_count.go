package wordcount

import (
	"regexp"
	"strings"
)

// Frequency maps words and ocurrencies count
type Frequency map[string]int

// WordCount counts words frequency and returns Frequency type
func WordCount(phrase string) Frequency {
	frequency := Frequency{}
	phrase = strings.ToLower(phrase)

	regexp := regexp.MustCompile(`[\d\w]+'\w+|[\d\w]+`)
	for _, word := range regexp.FindAllString(phrase, -1) {
		frequency[word]++
	}

	return frequency
}
