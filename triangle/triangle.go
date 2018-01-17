// Package triangle implements functions to identify triangle types
package triangle

import (
	"math"
)

// Kind represents triangle types
type Kind string

// Triangle types
const (
	NaT = "Nat" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides takes triangle sides and return triangle kind
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

func isValidTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a > b+c {
		return false
	}
	if b > a+c {
		return false
	}
	if c > a+b {
		return false
	}
	return true
}
