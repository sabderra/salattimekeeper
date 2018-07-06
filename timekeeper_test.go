package src

import "testing"

// Validate the conversion of the default times.
func TestDayPortion(test *testing.T) {

	times := NewTimeKeeper()

	times.toDayUnits()

	assertFloatEquals(test, times.salat[FAJR], 0.20833333333333334, "fajr")
	assertFloatEquals(test, times.salat[SUNRISE], 0.25, "sunrise")
	assertFloatEquals(test, times.salat[DHUHR], 0.5, "dhuhr")
	assertFloatEquals(test, times.salat[SUNSET], 0.75, "sunset")
	assertFloatEquals(test, times.salat[MAGHRIB], 0.75, "maghrib")
	assertFloatEquals(test, times.salat[ISHA], 0.75, "isha")

}
