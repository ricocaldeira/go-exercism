// Package pascal implements functions to pascal triangle
package pascal

// Triangle returns pascal triangle for a given size
func Triangle(size int) [][]int {
	pascal := [][]int{{1}}
	if size == 1 {
		return pascal
	}

	for i := 1; i < size; i++ {
		pascal = append(pascal, makeRow(pascal[i-1]))
	}
	return pascal
}

func makeRow(row []int) (newRow []int) {
	newRow = append(newRow, 1)
	for i, el := range row {
		if i > 0 {
			newRow = append(newRow, el+row[i-1])
		}
	}
	newRow = append(newRow, 1)
	return newRow
}
