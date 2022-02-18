package itype

import "math"

// Angle wraps a angle in the form of radian float32.
//
// Its value lies strictly in [0, 2Pi).
type Angle float32

func normalizeRadians(rad float64) float64 {
	if rad < 0 {
		return 2*math.Pi - math.Mod(-rad, 2*math.Pi)
	}
	return math.Mod(rad, 2*math.Pi)
}

// Radians construct an Angle from a radian value.
func Radians(rad float32) Angle {
	return Angle(normalizeRadians(float64(rad)))
}

// Degrees construct an Angle from a degree value.
func Degrees(deg float32) Angle {
	return Radians(deg * math.Pi / 180)
}

// Radians return the value of the Angle in radians.
func (a Angle) Radians() float32 {
	return float32(normalizeRadians(float64(a)))
}

// Degrees return the value of the Angle in degrees.
func (a Angle) Degrees() float32 {
	return float32(normalizeRadians(float64(a)) * 180 / math.Pi)
}
