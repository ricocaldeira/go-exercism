// Package luhn implements functions to validate luhn numbers
package luhn

import "unicode"

// Valid takes a luhn number String and returns valid or not
func Valid(input string) bool {
	numbers := make([]int, 0, len(input))

	for _, char := range input {
		if unicode.IsSpace(char) {
			continue
		}
		if !unicode.IsDigit(char) {
			return false
		}
		AddNumber(&numbers, char)
	}
	if len(numbers) <= 1 {
		return false
	}

	for i := len(numbers) % 2; i < len(numbers); i += 2 {
		numbers[i] = Double(numbers[i])
	}

	return Sum(numbers)%10 == 0
}

func AddNumber(numbers *[]int, char rune) {
	size := len(*numbers)
	*numbers = (*numbers)[:size+1]
	(*numbers)[size] = int(char - '0')
}

func Double(number int) int {
	number *= 2
	if number > 9 {
		number -= 9
	}
	return number
}

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
