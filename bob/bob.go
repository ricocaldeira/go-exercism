// Package bob anwers questions
package bob

import (
	"regexp"
	"strings"
)

// Hey takes String questions and returns String responses
func Hey(remark string) string {
	reg, _ := regexp.Compile("[^a-zA-Z?]+")
	remark = reg.ReplaceAllString(remark, "")

	selectedOptions := 0

	if isYellin(remark) {
		selectedOptions++
	}
	if isAsking(remark) {
		selectedOptions += 2
	}

	switch selectedOptions {
	case 1:
		return "Whoa, chill out!"
	case 2:
		return "Sure."
	case 3:
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}

func isYellin(remark string) bool {
	return remark == strings.ToUpper(remark) && len(remark) > 1
}

func isAsking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
