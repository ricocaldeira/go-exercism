package secret

import (
	"math"
	"strconv"
)

var handshakes = map[int]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake takes a secret handshake code and returns all matching handshakes
func Handshake(secret uint) []string {
	if secret == 0 {
		return []string{}
	}

	handShake := []string{}
	binarySecret := strconv.FormatInt(int64(secret), 2)

	for i, val := range binarySecret {
		if val == '1' && len(binarySecret)-i < 5 {
			shake := math.Exp2(float64(len(binarySecret)-i) - 1)
			handShake = append(handShake, handshakes[int(shake)])
		}
	}
	if int(secret) < 16 {
		handShake = reverseInts(handShake)
	}
	return handShake
}

func reverseInts(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
