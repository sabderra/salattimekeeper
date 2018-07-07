package salat

import (
	"testing"
	"time"
)

var isnaParams = map[TIMES]float64{FAJR: 15, ISHA: 15}
var isnaConfig = make(map[string]string)
var ISNA = NewCalculationMethod("Islamic Society of North America (ISNA)", isnaParams, isnaConfig)

func TestJulianFromTime_20180704_w_lng(test *testing.T) {
	t1 := time.Date(2018, time.July, 4, 0, 0, 0, 0, time.Local)
	res := JulianFromTime(t1)
	res_lng := res - (-71.1328)/(15*24.0)
	assertFloatEquals(test, res_lng, 2458303.697591111, "20180704 with lng adj")
}

func TestLocation_computePrayerTimes_20000101(test *testing.T) {

	var l = NewLocation(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)
	times := l.computePrayerTimes(t)

	assertClose(test, times.salat[FAJR], 6.8257790973714405, 6, "Fajr")
	assertClose(test, times.salat[SUNRISE], 8.233068183935577, 6, "Sunrise")
	assertClose(test, times.salat[DHUHR], 12.798799310581218, 6, "Dhuhr")
	assertClose(test, times.salat[ASR], 15.06814569864285, 6, "Asr")
	assertClose(test, times.salat[SUNSET], 17.367623852904835, 6, "Sunset")
	assertClose(test, times.salat[MAGHRIB], 17.367623852904835, 6, "Maghrib")
	assertClose(test, times.salat[ISHA], 18.774156492784158, 6, "Isha")
}

func TestLocation_ComputePrayerTimes_20000101(test *testing.T) {

	var l = NewLocation(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)
	times := l.ComputePrayerTimes(t, "")

	assertEquals(test, times[FAJR], "06:50", "Fajr")
	assertEquals(test, times[SUNRISE], "08:14", "Sunrise")
	assertEquals(test, times[DHUHR], "12:48", "Dhuhr")
	assertEquals(test, times[ASR], "15:04", "Asr")
	assertEquals(test, times[SUNSET], "17:22", "Sunset")
	assertEquals(test, times[MAGHRIB], "17:22", "Maghrib")
	assertEquals(test, times[ISHA], "18:46", "Isha")
}

func TestLocation_computePrayerTimes_20180703(test *testing.T) {

	var l = NewLocation(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2018, time.July, 3, 0, 0, 0, 0, time.Local)
	times := l.computePrayerTimes(t)

	assertClose(test, times.salat[FAJR], 3.470203191708414, 6, "Fajr")
	assertClose(test, times.salat[SUNRISE], 5.2068474465620795, 6, "Sunrise")
	assertClose(test, times.salat[DHUHR], 12.813655300116784, 6, "Dhuhr")
	assertClose(test, times.salat[ASR], 16.869467326271497, 6, "Asr")
	assertClose(test, times.salat[SUNSET], 20.4171439960229, 6, "Sunset")
	assertClose(test, times.salat[MAGHRIB], 20.4171439960229, 6, "Maghrib")
	assertClose(test, times.salat[ISHA], 22.15115997592993, 6, "Isha")

}

func TestLocation_ComputePrayerTimes_20180703(test *testing.T) {

	var l = NewLocation(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2018, time.July, 3, 0, 0, 0, 0, time.Local)
	times := l.ComputePrayerTimes(t, "")

	assertEquals(test, times[FAJR], "03:28", "Fajr")
	assertEquals(test, times[SUNRISE], "05:12", "Sunrise")
	assertEquals(test, times[DHUHR], "12:49", "Dhuhr")
	assertEquals(test, times[ASR], "16:52", "Asr")
	assertEquals(test, times[SUNSET], "20:25", "Sunset")
	assertEquals(test, times[MAGHRIB], "20:25", "Maghrib")
	assertEquals(test, times[ISHA], "22:09", "Isha")
}

func TestLocation_computePrayerTimes_20180705(test *testing.T) {

	var l = NewLocation(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2018, time.July, 5, 0, 0, 0, 0, time.Local)
	times := l.computePrayerTimes(t)

	assertClose(test, times.salat[FAJR], 3.498769437304892, 6, "Fajr")
	assertClose(test, times.salat[SUNRISE], 5.226802113465489, 6, "Sunrise")
	assertClose(test, times.salat[DHUHR], 12.819543210494254, 6, "Dhuhr")
	assertClose(test, times.salat[ASR], 16.872044472298832, 6, "Asr")
	assertClose(test, times.salat[SUNSET], 20.408456246958536, 6, "Sunset")
	assertClose(test, times.salat[MAGHRIB], 20.408456246958536, 6, "Maghrib")
	assertClose(test, times.salat[ISHA], 22.133525839438406, 6, "Isha")

}

func TestLocation_ComputePrayerTimes_20180705(test *testing.T) {

	var l = NewLocation(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2018, time.July, 5, 0, 0, 0, 0, time.Local)
	times := l.ComputePrayerTimes(t, "")

	assertEquals(test, times[FAJR], "03:30", "Fajr")
	assertEquals(test, times[SUNRISE], "05:14", "Sunrise")
	assertEquals(test, times[DHUHR], "12:49", "Dhuhr")
	assertEquals(test, times[ASR], "16:52", "Asr")
	assertEquals(test, times[SUNSET], "20:25", "Sunset")
	assertEquals(test, times[MAGHRIB], "20:25", "Maghrib")
	assertEquals(test, times[ISHA], "22:08", "Isha")
}
