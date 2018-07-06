package src

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

func NewCalculationMethod(name string, params map[TIMES]float64, config map[string]string) *CalculationMethod {

	c := new(CalculationMethod)
	c.Name = name
	c.params = params
	c.config = config
	return c
}

type Settings struct {
	imsakAngleAdj   float64 // Angle in minutes
	fajrAngleAdj    float64
	dhuhrAngleAdj   float64 // Angle in minutes
	asrFactor       ASR_FACTOR
	maghribAngleAdj float64
	ishaAngleAdj    float64
}

type Location struct {
	Lat      float64
	Lng      float64
	Elv      float64
	mth      CalculationMethod
	settings Settings
}

func NewLocation(lat float64, lng float64, elv float64, mth *CalculationMethod) *Location {
	l := new(Location)
	l.Lat = lat
	l.Lng = lng
	l.Elv = elv
	l.mth = *mth
	l.settings = Settings{
		imsakAngleAdj: 10,
		fajrAngleAdj:  l.mth.params[FAJR],
		dhuhrAngleAdj: 0,
		asrFactor:     STANDARD,
		ishaAngleAdj:  l.mth.params[ISHA],
	}

	return l
}

func (l Location) param(param TIMES) float64 {
	return l.mth.params[param]
}

func formatTime(t float64) string {

	// Minize rounding errors by explicitly casting to float64.
	var fixedTime float64 = fixhour(t + 0.5/60.0)
	var hours float64 = math.Floor(fixedTime)

	minutes := math.Floor((fixedTime - hours) * 60.0)

	formattedTime := fmt.Sprintf("%02.0f:%02.0f", hours, minutes)

	return formattedTime
}

func (l Location) ComputePrayerTimes(datetime time.Time, format string) map[TIMES]string {

	times := l.computePrayerTimes(datetime)

	// Format Times to string
	var fmtTimes = make(map[TIMES]string)
	for i, t := range times.salat {
		fmtTimes[i] = formatTime(t)
	}

	return fmtTimes

}

func (l Location) computePrayerTimes(datetime time.Time) TimeKeeper {

	jd := JulianFromTime(datetime) - l.Lng/(15.0*24.0)

	times := NewTimeKeeper()

	// Convert hours to day units
	times.toDayUnits()

	times.salat[IMSAK] = SunAngleTime(jd, l.settings.imsakAngleAdj, times.salat[IMSAK], l.Lat, CCW)
	times.salat[FAJR] = SunAngleTime(jd, l.settings.fajrAngleAdj, times.salat[FAJR], l.Lat, CCW)
	times.salat[SUNRISE] = SunAngleTime(jd, RiseSetAngle(l.Elv), times.salat[SUNRISE], l.Lat, CCW)
	times.salat[DHUHR] = MidDay(jd, times.salat[DHUHR])
	times.salat[ASR] = l.asrTime(jd, l.settings.asrFactor, times.salat[ASR])
	times.salat[SUNSET] = SunAngleTime(jd, RiseSetAngle(l.Elv), times.salat[SUNSET], l.Lat, NONE)
	times.salat[MAGHRIB] = SunAngleTime(jd, l.settings.maghribAngleAdj, times.salat[MAGHRIB], l.Lat, NONE)
	times.salat[ISHA] = SunAngleTime(jd, l.settings.ishaAngleAdj, times.salat[ISHA], l.Lat, NONE)

	// Adjust for timezone and daylight savings time
	tzAdjust := l.timezoneAdjustment(datetime)
	for i, t := range times.salat {
		times.salat[i] = t + tzAdjust
	}

	// Adjust Maghrib if configured
	times.salat[MAGHRIB] = l.maghribAdjust(
		times.salat[MAGHRIB],
		times.salat[SUNSET],
		l.settings.maghribAngleAdj)

	return *times
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

// Convert hours to day portion
func dayPortionXXX(times TimeKeeper) *TimeKeeper {
	dayFraction := NewTimeKeeper()
	for i, t := range times.salat {
		dayFraction.salat[i] = t / 24.0
	}
	return dayFraction
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
