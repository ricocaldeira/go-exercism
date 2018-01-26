package pythagorean

// Triplet represents a Pythagorean triangle
type Triplet [3]int

// Range returns all Pythagorean triplets between a min and max range
func Range(min, max int) (triplets []Triplet) {
	triplet := Triplet{3, 4, 5}
	for triplet[2] <= max {
		if triplet[0] < min {
			triplet[0] += 3
			triplet[1] += 4
			triplet[2] += 5
			continue
		}
		triplets = append(triplets, triplet)
		triplet[0] += 3
		triplet[1] += 4
		triplet[2] += 5
	}

	return triplets
}

// Sum returns all possible triplets given a sum number
func Sum(sum int) (triplets []Triplet) {
	max := sum / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := sum - a - b; pyth(a, b, c) {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return
}

func pyth(a, b, c int) bool {
	return a*a+b*b == c*c
}
