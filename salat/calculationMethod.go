package salat

import (
	"fmt"
	"math"
	"time"
)

type ASR_FACTOR int

const (
	STANDARD ASR_FACTOR = 1
	HANAFI   ASR_FACTOR = 2
)

type CalculationMethod struct {
	Name   string
	params map[TIMES]float64
	config map[string]string
}

// NewMyLocation returns an instance of MyLocation with the provided Lat and Lng.
func NewMethod(c methodConfig) *CalculationMethod {

	m := &CalculationMethod{}

	m.Name = c.Name

	return m
}

func NewCalculationMethod(name string, params map[TIMES]float64, config map[string]string) *CalculationMethod {

	c := CalculationMethod{name, params, config}
	return &c
}

type Settings struct {
	imsakAngleAdj   float64 // Angle in minutes
	fajrAngleAdj    float64
	dhuhrAngleAdj   float64 // Angle in minutes
	asrFactor       ASR_FACTOR
	maghribAngleAdj float64
	ishaAngleAdj    float64
}

type Times struct {
	salat map[TIMES]float64
}

func NewTimes() *Times {
	t := new(Times)
	return t
}

type Location struct {
	Lat      float64
	Lng      float64
	Elv      float64
	mth      CalculationMethod
	salat    map[TIMES]float64
	settings Settings
}

type FmtSalat struct {
	salat string
	time  string
}

func NewLocation(lat float64, lng float64, elv float64, mth *CalculationMethod) *Location {

	settings := Settings{
		imsakAngleAdj: 10,
		fajrAngleAdj:  0,
		dhuhrAngleAdj: 0,
		asrFactor:     STANDARD,
		ishaAngleAdj:  0,
	}

	l := Location{lat, lng, elv, *mth, make(map[TIMES]float64), settings}

	l.settings.fajrAngleAdj = l.mth.params[FAJR]
	l.settings.ishaAngleAdj = l.mth.params[ISHA]

	l.salat[IMSAK] = 5
	l.salat[FAJR] = 5
	l.salat[SUNRISE] = 6
	l.salat[DHUHR] = 12
	l.salat[ASR] = 13
	l.salat[SUNSET] = 18
	l.salat[MAGHRIB] = 18
	l.salat[ISHA] = 18

	for index, t := range l.salat {
		l.salat[index] = t / 24.0
	}

	return &l
}

func (l Location) param(param TIMES) float64 {
	return l.mth.params[param]
}

func formatTime(t float64) string {

	// Minimize rounding errors by explicitly casting to float64.
	var fixedTime float64 = fixhour(t + 0.5/60.0)
	var hours float64 = math.Floor(fixedTime)

	minutes := math.Floor((fixedTime - hours) * 60.0)

	formattedTime := fmt.Sprintf("%02.0f:%02.0f", hours, minutes)

	return formattedTime
}

//func (l Location) ComputePrayerTimes(datetime time.Time, format string) map[TIMES]string {
//
//	times := l.computePrayerTimes(datetime)
//
//	// Return an ordered array with format times as string
//	// orderedTimes := make([]FmtSalat, 0, len(times.salat))
//	var orderedTimes []FmtSalat
//
//	var fmtTimes = make(map[TIMES]string)
//	for i, t := range times.salat {
//		fmtTimes[i] = formatTime(t)
//		// println(int(i), int(i)-1, time_string[i], formatTime(t) )
//		e := FmtSalat{salat:time_string[i], time:formatTime(t)}
//		orderedTimes = append( orderedTimes, e)
//		//println(orderedTimes)
//	}
//
//	return fmtTimes
//
//}

//func (l Location) computePrayerTimes(datetime time.Time) TimeKeeper {
//
//	jd := JulianFromTime(datetime) - l.Lng/(15.0*24.0)
//
//	times := NewTimeKeeper()
//
//	times.salat[IMSAK] = SunAngleTime(jd, l.settings.imsakAngleAdj, times.salat[IMSAK], l.Lat, CCW)
//	times.salat[FAJR] = SunAngleTime(jd, l.settings.fajrAngleAdj, times.salat[FAJR], l.Lat, CCW)
//	times.salat[SUNRISE] = SunAngleTime(jd, RiseSetAngle(l.Elv), times.salat[SUNRISE], l.Lat, CCW)
//	times.salat[DHUHR] = MidDay(jd, times.salat[DHUHR])
//	times.salat[ASR] = l.asrTime(jd, l.settings.asrFactor, times.salat[ASR])
//	times.salat[SUNSET] = SunAngleTime(jd, RiseSetAngle(l.Elv), times.salat[SUNSET], l.Lat, NONE)
//	times.salat[MAGHRIB] = SunAngleTime(jd, l.settings.maghribAngleAdj, times.salat[MAGHRIB], l.Lat, NONE)
//	times.salat[ISHA] = SunAngleTime(jd, l.settings.ishaAngleAdj, times.salat[ISHA], l.Lat, NONE)
//
//	// Adjust for timezone and daylight savings time
//	tzAdjust := l.timezoneAdjustment(datetime)
//	for i, t := range times.salat {
//		times.salat[i] = t + tzAdjust
//	}
//
//	// Adjust Maghrib if configured
//	times.salat[MAGHRIB] = l.maghribAdjust(
//		times.salat[MAGHRIB],
//		times.salat[SUNSET],
//		l.settings.maghribAngleAdj)
//
//	return *times
//}

func (l Location) computePrayerTimes(datetime time.Time) map[TIMES]float64 {

	jd := JulianFromTime(datetime) - l.Lng/(15.0*24.0)

	l.salat[IMSAK] = SunAngleTime(jd, l.settings.imsakAngleAdj, l.salat[IMSAK], l.Lat, CCW)
	l.salat[FAJR] = SunAngleTime(jd, l.settings.fajrAngleAdj, l.salat[FAJR], l.Lat, CCW)
	l.salat[SUNRISE] = SunAngleTime(jd, RiseSetAngle(l.Elv), l.salat[SUNRISE], l.Lat, CCW)
	l.salat[DHUHR] = MidDay(jd, l.salat[DHUHR])
	l.salat[ASR] = l.asrTime(jd, l.settings.asrFactor, l.salat[ASR])
	l.salat[SUNSET] = SunAngleTime(jd, RiseSetAngle(l.Elv), l.salat[SUNSET], l.Lat, NONE)
	l.salat[MAGHRIB] = SunAngleTime(jd, l.settings.maghribAngleAdj, l.salat[MAGHRIB], l.Lat, NONE)
	l.salat[ISHA] = SunAngleTime(jd, l.settings.ishaAngleAdj, l.salat[ISHA], l.Lat, NONE)

	// Adjust for timezone and daylight savings time
	tzAdjust := l.timezoneAdjustment(datetime)
	for i, t := range l.salat {
		l.salat[i] = t + tzAdjust
	}

	// Adjust Maghrib if configured
	l.salat[MAGHRIB] = l.maghribAdjust(
		l.salat[MAGHRIB],
		l.salat[SUNSET],
		l.settings.maghribAngleAdj)

	return l.salat
}

// Note the timezone is relative to UTC. On the US east coast the timezone offset is -5.
// During daylight saving time, roughly early March to early November, the offset is -4.
// Adjust with a +1 when not in daylight savings.
func (l Location) timezoneAdjustment(datetime time.Time) (tzAdjust float64) {

	var dst float64 = 1.0

	_, offset := datetime.Zone()
	_, winterOffset := time.Date(datetime.Year(), time.January, 1, 0, 0, 0, 0, time.Local).Zone()
	_, summerOffset := time.Date(datetime.Year(), time.July, 1, 0, 0, 0, 0, time.Local).Zone()

	if winterOffset != summerOffset { // the location has daylight saving
		if offset == summerOffset {
			dst = 0.0
		}
	}

	tzoneOffset := float64(offset/60.0/60.0) + dst
	tzAdjust = tzoneOffset - l.Lng/15.0

	return tzAdjust
}

// compute asr time
func (l Location) asrTime(jDate float64, factor ASR_FACTOR, time float64) float64 {

	decl, _ := SunPosition(jDate + time)
	angle := -ArcCot(float64(factor) + Tan(math.Abs(l.Lat-decl)))

	return SunAngleTime(jDate, angle, time, l.Lat, NONE)
}

// Adjust maghrib to be an offset from sunset.
func (l Location) maghribAdjust(maghribTime float64, sunsetTime float64, min float64) float64 {
	return sunsetTime - min/60.0
}
