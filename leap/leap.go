// Package implements functions that Given a year, report if it is a leap year.
// ```plain
// on every year that is evenly divisible by 4
//   except every year that is evenly divisible by 100
//     unless the year is also evenly divisible by 400
// ```
package leap

// IsLeapYear takes a year, report if it is a leap year.
func IsLeapYear(year int) bool {
	leapYear := false
	if year%4 == 0 {
		leapYear = true
		if year%100 == 0 {
			leapYear = false
			if year%400 == 0 {
				leapYear = true
			}
		}
	}
	return leapYear
}
