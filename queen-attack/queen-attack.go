package queenattack

import (
	"errors"
	"math"
)

// CanQueenAttack Given the position of two queens on a chess board, indicate whether or not they
// are positioned so that they can attack each other.
func CanQueenAttack(bQ string, wQ string) (bool, error) {
	if len(bQ) != 2 || len(wQ) != 2 {
		return false, errors.New("Invalid coordinates")
	}
	if bQ == wQ {
		return false, errors.New("Same Square")
	}
	// same file
	if bQ[0] == wQ[0] {
		return true, nil
	}
	// same rank
	if bQ[1] == wQ[1] {
		return true, nil
	}

	// out of board
	// a = 97 h = 104
	// 1 = 49 8 = 56
	if bQ[0] < 97 || bQ[0] > 104 || wQ[0] < 97 || wQ[0] > 104 ||
		bQ[1] < 49 || bQ[1] > 56 || wQ[1] < 49 || wQ[1] > 56 {
		return false, errors.New("Out of board")
	}

	// same diagonal
	if math.Abs(float64(bQ[0])-float64(wQ[0])) == math.Abs(float64(bQ[1])-float64(wQ[1])) {
		return true, nil
	}
	return false, nil
}
