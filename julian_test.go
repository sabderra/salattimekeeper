package src

import (
	"fmt"
	"testing"
	"time"
)

func TestJulian(test *testing.T) {
	var res float64
	var see string

	res = Julian(1877, 8, 11, 7.5)
	see = fmt.Sprintf("%.2f", res)
	if res != 2406842.8125 {
		test.Error("1877/8/11 7.5 Expected 2406842.8125 got ", see)
	}

	res = Julian(2018, 6, 21, 0)
	see = fmt.Sprintf("%.2f", res)
	if res != 2458290.5 {
		test.Error("2018/6/21 0 Expected 2458290.5 got ", see)
	}

	res = Julian(2016, 1, 13, 0)
	see = fmt.Sprintf("%.2f", res)
	if res != 2457400.5 {
		test.Error("2016/1/13 Expected 2457400.5 got ", see)
	}

	res = Julian(1978, 1, 1, 0)
	see = fmt.Sprintf("%.4f", res)
	if res != 2443509.5 {
		test.Error("1978/1/1 0h Expected 2443509.5 got ", see)
	}

	res = Julian(1978, 7, 21, 15)
	see = fmt.Sprintf("%.4f", res)
	if res != 2443711.125 {
		test.Error("1978/7/21 15h Expected 2443711.125 got ", see)
	}

	res = Julian(2018, 12, 28, 0)
	see = fmt.Sprintf("%.2f", res)
	if res != 2458480.5 {
		test.Error("2018/12/28 Expected 2458480.5 got ", see)
	}

	res = Julian(2000, 1, 1.5, 0)
	see = fmt.Sprintf("%.2f", res)
	if res != 2451545.0 {
		test.Error("2000/1/1.5 Expected 2451545.0 got ", see)
	}
}

func TestJulianFromTime(test *testing.T) {
	var res float64
	var see string

	t1 := time.Date(1877, time.August, 11, 7, 30, 0, 0, time.Local)
	res = JulianFromTime(t1)
	see = fmt.Sprintf("%.2f", res)
	if res != 2406842.8125 {
		test.Error("1877/8/11 7.5 Expected 2406842.8125 got ", see)
	}

	t2 := time.Date(2018, time.June, 21, 0, 0, 0, 0, time.Local)
	res = JulianFromTime(t2)
	see = fmt.Sprintf("%.2f", res)
	if res != 2458290.5 {
		test.Error("2018/6/21 0 Expected 2458290.5 got ", see)
	}

	t3 := time.Date(2016, time.January, 13, 0, 0, 0, 0, time.Local)
	res = JulianFromTime(t3)
	see = fmt.Sprintf("%.2f", res)
	if res != 2457400.5 {
		test.Error("2016/1/13 Expected 2457400.5 got ", see)
	}

	t4 := time.Date(1978, time.January, 1, 0, 0, 0, 0, time.Local)
	res = JulianFromTime(t4)
	see = fmt.Sprintf("%.4f", res)
	if res != 2443509.5 {
		test.Error("1978/1/1 0h Expected 2443509.5 got ", see)
	}

	t5 := time.Date(1978, time.July, 21, 15, 0, 0, 0, time.Local)
	res = JulianFromTime(t5)
	see = fmt.Sprintf("%.4f", res)
	if res != 2443711.125 {
		test.Error("1978/7/21 15h Expected 2443711.125 got ", see)
	}

	t6 := time.Date(2018, time.December, 28, 0, 0, 0, 0, time.Local)
	res = JulianFromTime(t6)
	see = fmt.Sprintf("%.2f", res)
	if res != 2458480.5 {
		test.Error("2018/12/28 Expected 2458480.5 got ", see)
	}

	t7 := time.Date(2000, time.January, 1, 12, 0, 0, 0, time.Local)
	res = JulianFromTime(t7)
	see = fmt.Sprintf("%.2f", res)
	if res != 2451545.0 {
		test.Error("2000/1/1.5 Expected 2451545.0 got ", see)
	}
}

func TestJulianFromTime_20000101(test *testing.T) {
	t1 := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)
	res := JulianFromTime(t1)
	assertFloatEquals(test, res, 2451544.5, "20000101")
}

func TestJulianFromTime_20180704(test *testing.T) {
	t1 := time.Date(2018, time.July, 4, 0, 0, 0, 0, time.Local)
	res := JulianFromTime(t1)
	assertFloatEquals(test, res, 2458303.5, "20180704")
}

func TestJulianFromTime_20180705(test *testing.T) {
	t1 := time.Date(2018, time.July, 5, 0, 0, 0, 0, time.Local)
	res := JulianFromTime(t1)
	assertFloatEquals(test, res, 2458304.5, "20180705")
}
