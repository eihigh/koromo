package koromo

import "math"

// Wave returns sin(2π*ratio)
func Wave(ratio, offset float64) float64 {
	return math.Sin(math.Pi*2*ratio + math.Pi*offset*2)
}

// UWave returns unsigned Wave, [0.0, 1.0]
func UWave(ratio, offset float64) float64 {
	return 0.5 + 0.5*Wave(ratio, offset)
}
