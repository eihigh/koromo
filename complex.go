package koromo

import "math/cmplx"

// Vec is alias of complex128
type Vec = complex128

// UnitVector returns rotated unit vector
func UnitVector(angle float64) Vec {
	return cmplx.Pow(complex(1, 0), complex(angle, 0))
}

// Rot returns 360°/n rotated unit vector
func Rot(n float64) Vec {
	return UnitVector(4.0 / n)
}
