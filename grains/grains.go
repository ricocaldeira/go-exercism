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
	return intExp2(squareNumber), nil
}

// Total returns total uint64 number of grains
func Total() uint64 {
	var totalGrains uint64
	for i := 1; i <= 64; i++ {
		totalGrains += intExp2(i)
	}
	return totalGrains
}

func intExp2(number int) uint64 {
	return uint64(math.Exp2(float64(number - 1)))
}
