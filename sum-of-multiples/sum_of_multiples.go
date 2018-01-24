package summultiples

// SumMultiples Given a number, find the sum of all the multiples of particular numbers up to
// but not including that number.
func SumMultiples(number int, divisors ...int) (sum int) {
	for i := 1; i < number; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
