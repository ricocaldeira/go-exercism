// Package sublist implements functions to deal with lists
package sublist

import (
	"fmt"
	"strings"
)

// Relation returns relation type between lists
// equal, sublist, superlist, unequal
type Relation string

// Sublist takes listOne, listTwo and returns Relation type
// A = [1, 2, 3], B = [1, 2, 3, 4, 5] sublist
// A = [1, 2, 3], B = [1, 2, 3]       equal
// A = [1, 2, 3, 4, 5], B = [1, 2, 3] superlist
// A = [4, 5, 6], B = [1, 2, 3]       unequal
func Sublist(listOne []int, listTwo []int) Relation {
	lenOne := len(listOne)
	lenTwo := len(listTwo)
	stringOne := strings.Trim(fmt.Sprint(listOne), "[]")
	stringTwo := strings.Trim(fmt.Sprint(listTwo), "[]")

	if lenOne == 0 || lenTwo == 0 {
		listOne = append(listOne, 0)
		listTwo = append(listTwo, 0)
	}
	if stringOne == stringTwo {
		return "equal"
	}
	if lenOne > lenTwo && strings.Contains(stringOne, stringTwo) {
		return "superlist"
	}
	if lenOne < lenTwo && strings.Contains(stringTwo, stringOne) {
		return "sublist"
	}
	return "unequal"
}
