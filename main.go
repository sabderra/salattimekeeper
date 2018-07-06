package src

var mwlParams = map[TIMES]float64{FAJR: 18, ISHA: 17}
var mwlConfig = make(map[string]string)
var MWL = NewCalculationMethod("Muslim World League", mwlParams, mwlConfig)

var isnaParams = map[TIMES]float64{FAJR: 15, ISHA: 15}
var isnaConfig = make(map[string]string)
var ISNA = NewCalculationMethod("Islamic Society of North America (ISNA)", isnaParams, isnaConfig)

var eqyptParams = map[TIMES]float64{FAJR: 19.5, ISHA: 17.5}
var eqyptConfig = make(map[string]string)
var EGYPT = NewCalculationMethod("Egyptian General Authority of Survey", eqyptParams, eqyptConfig)

var makkahParams = map[TIMES]float64{FAJR: 18.5, ISHA: 17.5}
var makkahConfig = make(map[string]string)
var MAKKAH = NewCalculationMethod("Umm Al-Qura University, Makkah", makkahParams, makkahConfig)

var karachiParams = map[TIMES]float64{FAJR: 18, ISHA: 18}
var karachiConfig = make(map[string]string)
var KARACHI = NewCalculationMethod("University of Islamic Sciences, Karachi", karachiParams, karachiConfig)

var tehranParams = map[TIMES]float64{FAJR: 17.7, ISHA: 14, MAGHRIB: 4.5}
var tehranConfig = map[string]string{"midnight": "Jafari"}
var TEHRAN = NewCalculationMethod("Institute of Geophysics, University of Tehran", tehranParams, tehranConfig)

var jafariParams = map[TIMES]float64{FAJR: 16, ISHA: 14, MAGHRIB: 4}
var jafariConfig = map[string]string{"midnight": "Jafari"}
var JAFARI = NewCalculationMethod("Shia Ithna-Ashari, Leva Institute, Qum", jafariParams, jafariConfig)
