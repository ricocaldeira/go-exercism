// Package tournament implements functions to return results of competitions
package tournament

import (
	"bufio"
	"fmt"
	"io"
)

// Team struct to store team's performance at tournament
type Team struct {
	MatchesPlayed int
	Win           int
	Drawn         int
	Lost          int
	Points        int
}

var teams = map[string]Team{}

// Tally receives an input and writes results on buffer
func Tally(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	// Not actually needed since it’s a default split function.
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println()
	}

	return nil
}
