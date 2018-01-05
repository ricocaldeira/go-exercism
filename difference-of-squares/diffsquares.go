// Package diffsquares implements functions to
// calculate difference between the square of
// the sum and the sum of the squares
package diffsquares

// SquareOfSums takes a number and returns square of the sums
func SquareOfSums(n int) int {
	return (n * (1 + n) / 2) * (n * (1 + n) / 2)
}

// SumOfSquares takes a number and returns sum of the squares
func SumOfSquares(n int) int {
	sumOfSquares := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

// Difference takes a number and the difference between the square of the sum and sum of the squares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
