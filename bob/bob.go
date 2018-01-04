// Package bob anwers questions
package bob

import (
	"regexp"
	"strings"
)

// Hey takes String questions and returns String responses
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if isAsking(remark) && isYellin(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isYellin(remark) {
		return "Whoa, chill out!"
	}
	if isAsking(remark) {
		return "Sure."
	}
	if isSilenced(remark) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isAlpha(remark string) bool {
	match, _ := regexp.MatchString("[a-zA-Z]+", remark)
	return match
}

func isYellin(remark string) bool {
	return remark == strings.ToUpper(remark) && isAlpha(remark)
}

func isAsking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isSilenced(remark string) bool {
	return len(remark) == 0
}
