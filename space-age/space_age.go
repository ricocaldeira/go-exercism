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
	planets := map[Planet]float64{
		"Earth":   earthYearsInSeconds,
		"Mercury": mercuryYearsInSeconds,
		"Venus":   venusYearsInSeconds,
		"Mars":    marsYearsInSeconds,
		"Jupiter": jupiterYearsInSeconds,
		"Saturn":  saturnYearsInSeconds,
		"Uranus":  uranusYearsInSeconds,
		"Neptune": neptuneYearsInSeconds,
	}
	return seconds / planets[planet]
}
