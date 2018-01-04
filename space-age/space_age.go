package space

// Planet struct defines a planet
type Planet string

const earthYearsInSeconds = 31557600.0
const mercuryYearsInSeconds = earthYearsInSeconds * 0.2408467
const venusYearsInSeconds = earthYearsInSeconds * 0.61519726
const marsYearsInSeconds = earthYearsInSeconds * 1.8808158
const jupiterYearsInSeconds = earthYearsInSeconds * 11.862615
const saturnYearsInSeconds = earthYearsInSeconds * 29.447498
const uranusYearsInSeconds = earthYearsInSeconds * 84.016846
const neptuneYearsInSeconds = earthYearsInSeconds * 164.79132

// Age takes an age in seconds and a Planet
// and returns age in years on that planet
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / earthYearsInSeconds
	case "Mercury":
		return seconds / mercuryYearsInSeconds
	case "Venus":
		return seconds / venusYearsInSeconds
	case "Mars":
		return seconds / marsYearsInSeconds
	case "Jupiter":
		return seconds / jupiterYearsInSeconds
	case "Saturn":
		return seconds / saturnYearsInSeconds
	case "Uranus":
		return seconds / uranusYearsInSeconds
	case "Neptune":
		return seconds / neptuneYearsInSeconds
	}
	return 0
}
