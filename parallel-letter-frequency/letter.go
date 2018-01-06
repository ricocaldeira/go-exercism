// Package letter implements functions to count letter frequency
package letter

// ConcurrentFrequency takes an array of string sentences and returns a FreqMap
func ConcurrentFrequency(inputs []string) FreqMap {
	c := make(chan FreqMap, len(inputs))
	defer close(c)

	for _, s := range inputs {

		go func(in string) {
			c <- Frequency(in)
		}(s)
	}

	m := FreqMap{}

	for range inputs {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
