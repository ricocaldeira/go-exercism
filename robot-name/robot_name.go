// Package robotname implements functions to generate robot names
package robotname

import (
	"fmt"
	"math/rand"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Robot defines a robot with name string attribute
type Robot struct {
	name string
}

// generateName generates a random name
// Pattern: AA123, KE318, GL128 ...
func generateName() string {
	serialCode := make([]rune, 2)
	serialNumber := 999

	for i := range serialCode {
		randNumber := rand.Intn(len(letters))
		serialNumber -= randNumber
		serialCode[i] = letters[randNumber]
	}

	return fmt.Sprint(string(serialCode), serialNumber)
}

// Name returns robot name string
func (robot *Robot) Name() string {
	if len(robot.name) == 0 {
		robot.name = generateName()
	}
	return robot.name
}

// Reset wipes robot name
func (robot *Robot) Reset() *Robot {
	robot.name = generateName()
	return robot
}
