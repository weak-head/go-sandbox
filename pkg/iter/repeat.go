package iter

import (
	str "strings"
)

// Repeat returns the concatened input repeated N times
func Repeat(s string, times int) string {
	// var repeated string
	// for i := 0; i < times; i++ {
	// 	repeated += s
	// }
	// return repeated

	return str.Repeat(s, times)
}
