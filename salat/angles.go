package salat

import (
	"math"
)

type Direction int

const (
	CCW Direction = 1 + iota
	NONE
)

func fix(a float64, mode float64) float64 {
	if math.IsNaN(a) {
		return a
	}

	a = a - mode*(math.Floor(a/mode))
	if a < 0 {
		return a + mode
	} else {
		return a
	}
}

func fixangle(angle float64) float64 {
	return fix(angle, 360.0)
}

func fixhour(hour float64) float64 {
	return fix(hour, 24.0)
}

// SunPosition computes the declination angle of the sun and
// equation of time.
//
// The declination is the angle between the rays
// of the sun and the plane of the earth equator. The earth orbits the
// sun at a tilt, as such the declination will change throughout
// the year.
//
// The equation of time is the difference between time as read from a
// sundial and a clock. It results from an apparent irregular movement
// of the Sun caused by a combination of the obliquity of the Earth's
// rotation axis and the eccentricity of its orbit. The sundial can be
// ahead (fast) by as much as 16 min 33 s (around November 3) or fall
// behind by as much as 14 min 6 s (around February 12).
//
// References:
// http://praytimes.org/calculation
// http://aa.usno.navy.mil/faq/docs/SunApprox.php
func SunPosition(jd float64) (decl float64, eqt float64) {

	// Compute the number of days and fraction (+ or –) from
	// the epoch referred to as "J2000.0", which is 2000
	// January 1.5, Julian date 2451545.0
	D := jd - 2451545.0

	// Mean anomaly of the Sun
	g := fixangle(357.529 + 0.98560028*D)

	// Mean longitude of the Sun
	q := fixangle(280.459 + 0.98564736*D)

	// Geocentric apparent ecliptic longitude
	// of the Sun (adjusted for aberration)
	L := fixangle(q + 1.915*Sin(g) + 0.020*Sin(2*g))

	// The Sun's ecliptic latitude, b, can be approximated by b=0.
	// The distance of the Sun from the Earth, R, in astronomical
	// units (AU), can be approximated by
	// R := 1.00014 - 0.01671*Cos(g) - 0.00014*Cos(2*g)

	// The mean obliquity of the ecliptic, in degrees
	e := 23.439 - 0.00000036*D

	// The Sun's right ascension
	RA := ArcTan2(Cos(e)*Sin(L), Cos(L)) / 15.0

	//  The Equation of Time apparent solar time minus mean solar time
	eqt = q/15.0 - fixhour(RA)
	decl = ArcSin(Sin(e) * Sin(L))

	return decl, eqt
}

// MidDay computes the mid-day time
func MidDay(jDate float64, time float64) float64 {
	_, eqt := SunPosition(jDate + time)
	return fixhour(12 - eqt)
}

// SunAngleTime computes the time at which sun reaches a specific angle below horizon
func SunAngleTime(jDate float64, angle float64, time float64, lat float64, direction Direction) float64 {

	decl, _ := SunPosition(jDate + time)
	noon := MidDay(jDate, time)
	t := 1.0 / 15.0 * ArcCos((-Sin(angle)-Sin(decl)*Sin(lat))/(Cos(decl)*Cos(lat)))

	var sunAngleTime float64
	if direction == CCW {
		sunAngleTime = noon - t
	} else {
		sunAngleTime = noon + t
	}

	// XXX - for debugging remove
	//fmt.Printf("-self.sin(angle) = %f\n", -Sin(angle))
	//fmt.Printf("self.sin(decl) = %f\n", Sin(decl))
	//fmt.Printf("self.sin(lat) = %f\n", Sin(lat))
	//fmt.Printf( "t = %f\n", t)
	//fmt.Printf( "noon = %f\n", noon)
	//fmt.Printf( "sunAngleTime = %f\n", sunAngleTime)

	return sunAngleTime
}

// RiseSetAngle returns an approximation of the sun angle for sunset/sunrise
func RiseSetAngle(elevation float64) float64 {
	return 0.833 + 0.0347*math.Sqrt(elevation)
}
