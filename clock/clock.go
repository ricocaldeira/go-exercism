// Package clock implements a clock without dates
package clock

import (
	"fmt"
	"time"
)

// Clock has hour int and minute int
type Clock struct {
	hour   int
	minute int
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add takes an integer interval, add it to clock and return new clock time
func (clock Clock) Add(interval int) Clock {
	date := time.Date(2000, time.January, 1, clock.hour, clock.minute+interval, 0, 0, time.UTC)
	clock.hour, clock.minute = date.Hour(), date.Minute()
	return clock
}

// New takes integers hours and minutes and returns a clock time
func New(hour int, minute int) Clock {
	date := time.Date(2000, time.January, 1, hour, minute, 0, 0, time.UTC)
	return Clock{date.Hour(), date.Minute()}
}
