package pythagorean

// Triplet represents a Pythagorean triangle
type Triplet [3]int

// Range returns all Pythagorean triplets between a min and max range
func Range(min, max int) (triplets []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if isPythagorian(a, b, c) {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum returns all possible triplets given a sum number
func Sum(sum int) (triplets []Triplet) {
	max := sum / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := sum - a - b; isPythagorian(a, b, c) {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return
}

func isPythagorian(a, b, c int) bool {
	return a*a+b*b == c*c
}
