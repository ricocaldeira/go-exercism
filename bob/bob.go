// Package bob anwers questions
package bob

import (
	"regexp"
	"strconv"
	"strings"
)

// Hey takes String questions and returns String responses
func Hey(remark string) string {
	lastRemarkChar := remark[len(remark)-1 : len(remark)]
	remark = remark[0 : len(remark)-1]
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	remark = reg.ReplaceAllString(remark, "")

	selectedOptions := 0

	if isYellin(remark) {
		selectedOptions++
	}
	if isAsking(lastRemarkChar) {
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

func isNumber(remark string) bool {
	if _, err := strconv.Atoi(remark); err == nil {
		return true
	}
	return false
}

func isYellin(remark string) bool {
	if remark == strings.ToUpper(remark) && len(remark) > 0 && !isNumber(remark) {
		return true
	}
	return false
}

func isAsking(lastRemarkChar string) bool {
	if lastRemarkChar == "?" {
		return true
	}
	return false
}
