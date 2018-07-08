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

var time_string = map[TIMES]string{
	IMSAK:   "imsak",
	FAJR:    "fajr",
	SUNRISE: "sunrise",
	DHUHR:   "dhuhr",
	ASR:     "asr",
	SUNSET:  "sunset",
	MAGHRIB: "maghrib",
	ISHA:    "isha",
}

type TimeKeeper struct {
	location Location
	cDate    time.Time
}

func NewTimeKeeper(lat float64, lng float64, elv float64, mth *CalculationMethod) *TimeKeeper {

	l := NewLocation(lat, lng, 0, mth)

	t := new(TimeKeeper)
	t.location = *l

	return t
}

// Convert hours to day units
func (timeKeeper TimeKeeper) toDayUnits() {
	for index, t := range timeKeeper.location.salat {
		timeKeeper.location.salat[index] = t / 24.0
	}
}

// Set date to use for salat calculations
func (timeKeeper TimeKeeper) SetDate(year int, month int, day int) {
	timeKeeper.cDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func (timeKeeper TimeKeeper) SetLocationLatLng(lat float64, lng float64, elv float64, mth *CalculationMethod) *Location {
	var l = NewLocation(lat, lng, 0, mth)
	timeKeeper.location = *l
	return l
}

func (timeKeeper TimeKeeper) SetCalculationMethod(mth *CalculationMethod) {
	timeKeeper.location.mth = *mth
}

func (timeKeeper TimeKeeper) GetPrayerTimes(datetime time.Time) map[TIMES]string {

	times := timeKeeper.location.computePrayerTimes(datetime)

	// Return format times as string
	var fmtTimes = make(map[TIMES]string)

	for i, t := range times {
		fmtTimes[i] = formatTime(t)
	}

	return fmtTimes
}
