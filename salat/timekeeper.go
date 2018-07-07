package salat

import (
	"time"
)

type TIMES int

const (
	IMSAK TIMES = 1 + iota
	FAJR
	SUNRISE
	DHUHR
	ASR
	SUNSET
	MAGHRIB
	ISHA
)

type TimeKeeper struct {
	salat map[TIMES]float64
	jDate float64
	cDate time.Time
}

func NewTimeKeeper() *TimeKeeper {
	t := new(TimeKeeper)
	t.salat = make(map[TIMES]float64)

	t.salat[IMSAK] = 5
	t.salat[FAJR] = 5
	t.salat[SUNRISE] = 6
	t.salat[DHUHR] = 12
	t.salat[ASR] = 13
	t.salat[SUNSET] = 18
	t.salat[MAGHRIB] = 18
	t.salat[ISHA] = 18

	return t
}

// Convert hours to day units
func (timeKeeper TimeKeeper) toDayUnits() {
	for index, t := range timeKeeper.salat {
		timeKeeper.salat[index] = t / 24.0
	}
}

// Set date to use for salat calculations
func (timeKeeper TimeKeeper) SetDate(year int, month int, day int) {
	timeKeeper.cDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	timeKeeper.jDate = JulianFromTime(timeKeeper.cDate)
}
