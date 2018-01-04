// Package raindrops implements functions to convert numbers into raindrops
package raindrops

import "strconv"

// Convert take a number to a string.
// Contents of which depend on the number's factors.
func Convert(number int) string {
	rain := ""
	if number%3 == 0 {
		rain += "Pling"
	}
	if number%5 == 0 {
		rain += "Plang"
	}
	if number%7 == 0 {
		rain += "Plong"
	}
	if rain == "" {
		rain = strconv.Itoa(number)
	}
	return rain
}
