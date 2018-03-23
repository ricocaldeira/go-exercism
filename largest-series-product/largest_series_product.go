// Package lsproduct implements fuctions to to calculate largest series products
package lsproduct

import "strconv"

// LargestSeriesProduct Given a string of digits,
// calculate the largest product for a contiguous
// substring of digits of length n.
func LargestSeriesProduct(series string, n int) (int, error) {
	largestProduct := 0
	for i := 0; i < len(series); i = i + n {
		numbers := series[i : i+n]
		mul := 1
		for j := 0; j < len(numbers); j++ {
			intNum, _ := strconv.Atoi(numbers[j : j+1])
			mul = mul * intNum
		}
		if mul > largestProduct {
			largestProduct = mul
		}
		if i+n+1 == len(series) {
			break
		}
	}
	return largestProduct, nil
}
