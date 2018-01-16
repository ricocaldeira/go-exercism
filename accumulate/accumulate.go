// Package accumulate implements functions to accumulate an operation to a collection
package accumulate

type accum func(string) string

// Accumulate takes a collection and applies a function to each collection's element
func Accumulate(input []string, fn accum) (out []string) {
	for _, in := range input {
		out = append(out, fn(in))
	}
	return out
}
