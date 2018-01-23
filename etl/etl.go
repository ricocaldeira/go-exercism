package etl

import (
	"strings"
)

// Transform formats oldMap data into newMap format like
//
// given{
// 	1: {"A", "E"},
// 	2: {"D", "G"},
// },
//
// expectation{
// 	"a": 1,
// 	"e": 1,
// 	"d": 2,
// 	"g": 2,
// },
//
func Transform(oldMap map[int][]string) map[string]int {
	var newMap = map[string]int{}
	for score, strCollection := range oldMap {
		for _, str := range strCollection {
			lowStr := strings.ToLower(str)
			newMap[lowStr] = score
		}
	}
	return newMap
}
