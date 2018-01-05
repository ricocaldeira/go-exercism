// Package grains implements functions to calculate number of grains of wheat on a chessboard
package grains

import (
	"errors"
	"math"
)

// Square takes number of grains and returs square
// Returns an error for square greater then 64 or lesser then 1
func Square(squareNumber int) (uint64, error) {
	if squareNumber < 1 || squareNumber > 64 {
		return 0, errors.New("Invalid square number")
	}
	var squares uint64
	squares = uint64(math.Pow(2, float64(squareNumber-1)))
	return squares, nil
}

// Total returns total uint64 number of grains
func Total() uint64 {
	var totalGrains uint64
	for i := 1; i <= 64; i++ {
		totalGrains += uint64(math.Pow(2, float64(i-1)))
	}
	return totalGrains
}
