/*
Package gigasecond supplies a function to add a Gigasecond to Time
*/
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes Time and adds a Gigasecond to it
// 1 Gigasecond = 1 billion seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1E9)
}
