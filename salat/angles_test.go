package salat

import (
	"testing"
)

func TestSunPosition(test *testing.T) {

	jd20000101 := Julian(2000, 1, 1, 0)

	// Adjust for longitude
	jd := jd20000101 - (-71.1328)/(15*24.0)
	assertFloatEquals(test, jd, 2451544.697591111, "20000101")

	x := jd + 0.20833333333333334
	decl, eqt := SunPosition(x)

	assertClose(test, decl, -23.040933527142226, 16, "decl SunPosition for 20000101")
	assertClose(test, eqt, -0.054309113374497286, 16, "eqt  SunPosition for 20000101")
}

func TestSunPosition2(test *testing.T) {

	// June 28, 2018
	jd := 2458298.447591111
	decl, eqt := SunPosition(jd)

	assertClose(test, decl, 23.241857345760952, 14, "dec;")
	assertClose(test, eqt, -0.056195059047817075, 14, "eqt")
}

func TestSunPosition3(test *testing.T) {

	// July 4, 2018
	jd := 2458302.9059244446
	decl, eqt := SunPosition(jd)

	assertFloatEquals(test, decl, 22.94466842237423, "dec;")

	assertFloatEquals(test, eqt, -0.07058023378383371, "eqt")
}

func TestSunAngleTime(test *testing.T) {
	// June 29, 2018
	jd := 2458298.697591111
	angle := 10.0
	time := 0.20833333333333334
	lat := 42.4223

	sunAngleTime := SunAngleTime(jd, angle, time, lat, CCW)

	assertClose(test, sunAngleTime, 3.3651756027580095, 6, "sunAngleTime")

}

func TestSunAngleTime2(test *testing.T) {
	// June 29, 2018
	jd := 2458298.698658333
	angle := 10.0
	time := 0.20833333333333334
	lat := 42.4917

	sunAngleTime := SunAngleTime(jd, angle, time, lat, CCW)

	assertClose(test, sunAngleTime, 3.3589706301355555, 6, "TestSunAngleTime2")

}

func TestMidDay(test *testing.T) {

	//jd := Julian(2018, 6, 29, 9.5)
	jd := 2458298.697591111
	time := 0.20833333333333334

	noon := MidDay(jd, time)

	assertClose(test, noon, 12.057735824732513, 6, "TestMidDay")

}

func TestRiseSetAngle(test *testing.T) {

	angle := RiseSetAngle(0)
	assertFloatEquals(test, angle, 0.833, "TestRiseSetAngle")
}

func TestFixAngle1(test *testing.T) {

	D := 6757.9059244445525
	g := fixangle(357.529 + 0.98560028*D)

	assertFloatEquals(test, g, 178.12297134620894, "fixangle1")
}

func TestFixAngle2(test *testing.T) {

	D := 6757.9059244445525
	q := fixangle(280.459 + 0.98564736*D)

	assertFloatEquals(test, q, 101.37113355713245, "fixangle2")
}

func TestFixAngle3(test *testing.T) {

	D := 6757.9059244445525
	g := fixangle(357.529 + 0.98560028*D)
	q := fixangle(280.459 + 0.98564736*D)
	L := fixangle(q + 1.915*Sin(g) + 0.020*Sin(2*g))

	assertFloatEquals(test, L, 101.43254889238493, "fixangle_L")
}

func TestFixAngle4(test *testing.T) {

	D := 6757.9059244445525
	g := fixangle(357.529 + 0.98560028*D)
	q := fixangle(280.459 + 0.98564736*D)
	L := fixangle(q + 1.915*Sin(g) + 0.020*Sin(2*g))

	e := 23.439 - 0.00000036*D

	assertFloatEquals(test, e, 23.4365671538672, "fixangle_e")

	RA := ArcTan2(Cos(e)*Sin(L), Cos(L)) / 15.0

	assertFloatEquals(test, RA, 6.828655804259331, "fixangle_RA")
}
