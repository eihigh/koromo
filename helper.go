package koromo

import "math"

// Wave returns sin(2π*ratio)
func Wave(ratio float64) float64 {
	return math.Sin(math.Pi * 2 * ratio)
}
