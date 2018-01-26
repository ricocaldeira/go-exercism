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
	// triplets = append(triplets, triplet)
	return triplets
}

func Sum(p int) []Triplet {
	return nil
}
