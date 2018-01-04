// Package bob anwers questions
package bob

import (
	"regexp"
	"strings"
)

// Hey takes String questions and returns String responses
func Hey(remark string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9?]+")
	remark = reg.ReplaceAllString(remark, "")

	selectedOptions := 0

	if isYellin(remark) {
		selectedOptions++
	}
	if isAsking(remark) {
		selectedOptions += 2
	}
	if isSilenced(remark) {
		selectedOptions += 4
	}

	switch selectedOptions {
	case 1:
		return "Whoa, chill out!"
	case 2:
		return "Sure."
	case 3:
		return "Calm down, I know what I'm doing!"
	case 4:
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isSilenced(remark string) bool {
	return len(remark) == 0
}

func isAlpha(remark string) bool {
	match, _ := regexp.MatchString("[a-zA-Z]+", remark)
	return match
}

func isYellin(remark string) bool {
	return remark == strings.ToUpper(remark) && len(remark) > 1 && isAlpha(remark)
}

func isAsking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
