package salat

import (
	"testing"
	"time"
)

// Validate the conversion of the default times.
func TestDayPortion(test *testing.T) {

	times := NewTimeKeeper(42.4223, -71.1328, 0, ISNA)

	assertFloatEquals(test, times.location.salat[FAJR], 0.20833333333333334, "fajr")
	assertFloatEquals(test, times.location.salat[SUNRISE], 0.25, "sunrise")
	assertFloatEquals(test, times.location.salat[DHUHR], 0.5, "dhuhr")
	assertFloatEquals(test, times.location.salat[SUNSET], 0.75, "sunset")
	assertFloatEquals(test, times.location.salat[MAGHRIB], 0.75, "maghrib")
	assertFloatEquals(test, times.location.salat[ISHA], 0.75, "isha")

}

func TestTimeKeeper_GetPrayerTimes_20000101(test *testing.T) {

	timeKeeper := NewTimeKeeper(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2000, time.January, 1, 0, 0, 0, 0, tzLocation)
	times := timeKeeper.GetPrayerTimes(t)

	assertEquals(test, times[FAJR], "06:50", "Fajr")
	assertEquals(test, times[SUNRISE], "08:14", "Sunrise")
	assertEquals(test, times[DHUHR], "12:48", "Dhuhr")
	assertEquals(test, times[ASR], "15:04", "Asr")
	assertEquals(test, times[SUNSET], "17:22", "Sunset")
	assertEquals(test, times[MAGHRIB], "17:22", "Maghrib")
	assertEquals(test, times[ISHA], "18:46", "Isha")
}

func TestTimeKeeper_GetPrayerTimes_20180703(test *testing.T) {

	timeKeeper := NewTimeKeeper(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2018, time.July, 3, 0, 0, 0, 0, tzLocation)
	times := timeKeeper.GetPrayerTimes(t)

	assertEquals(test, times[FAJR], "03:28", "Fajr")
	assertEquals(test, times[SUNRISE], "05:12", "Sunrise")
	assertEquals(test, times[DHUHR], "12:49", "Dhuhr")
	assertEquals(test, times[ASR], "16:52", "Asr")
	assertEquals(test, times[SUNSET], "20:25", "Sunset")
	assertEquals(test, times[MAGHRIB], "20:25", "Maghrib")
	assertEquals(test, times[ISHA], "22:09", "Isha")
}

func TestTimeKeeper_GetPrayerTimes_20180705(test *testing.T) {

	timeKeeper := NewTimeKeeper(42.4223, -71.1328, 0, ISNA)

	t := time.Date(2018, time.July, 5, 0, 0, 0, 0, tzLocation)
	times := timeKeeper.GetPrayerTimes(t)

	assertEquals(test, times[FAJR], "03:30", "Fajr")
	assertEquals(test, times[SUNRISE], "05:14", "Sunrise")
	assertEquals(test, times[DHUHR], "12:49", "Dhuhr")
	assertEquals(test, times[ASR], "16:52", "Asr")
	assertEquals(test, times[SUNSET], "20:25", "Sunset")
	assertEquals(test, times[MAGHRIB], "20:25", "Maghrib")
	assertEquals(test, times[ISHA], "22:08", "Isha")
}
