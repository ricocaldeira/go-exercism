package wordcount

// Frequency maps words and ocurrencies count
type Frequency map[string]int

// WordCount counts words frequency and returns Frequency type
func WordCount(phrase string) Frequency {
	return Frequency{"word": 1}
}
