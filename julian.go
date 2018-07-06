package src

import "time"

// Julian converts golang Time gregorian date to a Julian date.
func JulianFromTime(datetime time.Time) float64 {

	// Convert remaining time to fractions of hours.
	diffTime := datetime.Sub(time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, time.Local))
	h := diffTime.Minutes() / 60.0

	return Julian(datetime.Year(), int(datetime.Month()), float64(datetime.Day()), h)
}

// Julian converts gregorian date to a Julian date.
// Reference http://aa.usno.navy.mil/faq/docs/JD_Formula.php for
// details on the algorithm.
func Julian(year int, month int, day float64, hour float64) float64 {

	k := float64(year)
	m := float64(month)
	i := day

	// For readability, the conversion is split into multiple parts

	p1 := 367.0*k - Trunc(7*(k+Trunc((m+9)/12))/4)
	p2 := Trunc(275*m/9) + i
	p3 := 1721013.5 + hour/24.0
	p4 := -0.5*Sign(100*k+m-190002.5) + 0.5

	jd := p1 + p2 + p3 + p4
	return jd
}
