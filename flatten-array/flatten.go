// Package flatten implements functions to flatten arrays
package flatten

// Flatten takes an interface and returns a flatten array
func Flatten(input interface{}) []interface{} {
	output := []interface{}{}
	for _, val := range input.([]interface{}) {
		switch val.(type) {
		case []interface{}:
			tmp := Flatten(val)
			output = append(output, tmp...)
		case int:
			output = append(output, val)
		}
	}
	return output
}
