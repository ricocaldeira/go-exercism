// Package hamming implements functions to calculate DNA hamming
package hamming

import "errors"

// Distance takes two DNA strings to calculate hamming distance
// if have same length returns integer hamming distance and nil errors
// else returns hamming -1 and error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNAs need to have same length")
	}

	distance := 0

	for index := range a {
		if a[index] != b[index] {
			distance++
		}
	}

	return distance, nil
}
