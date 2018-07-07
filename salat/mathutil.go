package salat

import (
	"math"
)

func Radians(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func Degrees(radian float64) float64 {
	return radian * (180 / math.Pi)
}

func Sin(degree float64) float64 {
	return math.Sin(Radians(degree))
}

func Cos(degree float64) float64 {
	return math.Cos(Radians(degree))
}

func Tan(degree float64) float64 {
	return math.Tan(Radians(degree))
}

func ArcSin(radian float64) float64 {
	return Degrees(math.Asin(radian))
}

func ArcCos(radian float64) float64 {
	return Degrees(math.Acos(radian))
}

func ArcTan(radian float64) float64 {
	return Degrees(math.Atan(radian))
}

func ArcCot(radian float64) float64 {
	return Degrees(math.Atan(1.0 / radian))
}

func ArcTan2(y float64, x float64) float64 {
	return Degrees(math.Atan2(y, x))
}

func Sign(n float64) float64 {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

// Trunc removes the fraction of the absolute number but does not round.
// The truncation function < > extracts the integral part of a number. This
// is used in the calculation of the Julian datetime.
func Trunc(n float64) float64 {
	return Sign(n) * math.Floor(math.Abs(n))
}
