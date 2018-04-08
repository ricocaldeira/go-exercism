// Package sieve implements functions to deal with prime numbers
package sieve

import (
	"math"
)

// Sieve returns all the primes from 2 up to a given number
func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	primeNumbers := []int{2}
	for i := 3; i <= n; i++ {
		for k := 2; k <= i; k++ {
			if k == i {
				primeNumbers = append(primeNumbers, i)
			}
			if math.Mod(float64(i), float64(k)) == 0 {
				break
			}
		}
	}
	return primeNumbers
}
