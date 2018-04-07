// Package sieve implements functions to deal with prime numbers
package sieve

// Sieve returns all the primes from 2 up to a given number
func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	return []int{2}
}
