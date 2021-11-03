package internal

import "reflect"

// InterfacesOf returns values' values as a slice of interface{} values.
// Example:
//     ii := InterfacesOf(
//         reflect.ValuesOf(false),
//     )
//     ii[0] == false
func InterfacesOf(values ...reflect.Value) []interface{} {
	var ii []interface{}
	for i := range values {
		v := values[i]
		ii = append(
			ii,
			v.Interface(),
		)
	}
	return ii
}
