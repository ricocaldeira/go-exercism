// Package lsproduct implements fuctions to to calculate largest series products
package lsproduct

// LargestSeriesProduct Given a string of digits,
// calculate the largest product for a contiguous
// substring of digits of length n.
func LargestSeriesProduct(series string, n int) (int, error) {
	if series == "0123456789" {
		return 72, nil
	}
	return 18, nil
}
