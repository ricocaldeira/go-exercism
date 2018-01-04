package space

// Planet struct defines a planet
type Planet string

const earthYearsInSeconds = 31557600.0

// Age takes an age in seconds and a Planet
// and returns age in years on that planet
func Age(seconds float64, planet Planet) float64 {
	planets := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / (planets[planet] * earthYearsInSeconds)
}
