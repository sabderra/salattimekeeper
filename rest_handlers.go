package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/sabderra/salattimekeeper/salat"
	"net/http"
	"strconv"
	"time"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.Path("/salat/").
		HandlerFunc(GetSalatTimes).
		Name("salat")

	router.Path("/salat").
		Queries("lat", "{lat}").
		Queries("lng", "{lng}").
		HandlerFunc(GetSalatTimesLatLng).
		Name("salatlatlng")

	return router
}

func GetSalatTimes(w http.ResponseWriter, req *http.Request) {
	logClientRequest(req)

	timeKeeper := salat.NewTimeKeeper(myLocation.Lat, myLocation.Lng, 0, ISNA)

	t := time.Now()
	times := timeKeeper.GetPrayerTimes(t)

	SendJsonResponse(w, times)
}

func GetSalatTimesLatLng(w http.ResponseWriter, req *http.Request) {
	logClientRequest(req)

	vars := mux.Vars(req)

	var lat, lng float64
	var err error

	lat, err = strconv.ParseFloat(vars["lat"], 64)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest)
	}

	lng, err = strconv.ParseFloat(vars["lng"], 64)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest)
	}

	timeKeeper := salat.NewTimeKeeper(lat, lng, 0, ISNA)
	times := timeKeeper.GetPrayerTimes(time.Now())

	SendJsonResponse(w, times)
}

// timesPayload to give better control over
// what is sent across the wire.
type timesPayload struct {
	Imsak   string `json:"imsak"`
	Fajr    string `json:"fajr"`
	Sunrise string `json:"sunrise"`
	Dhuhr   string `json:"dhuhr"`
	Asr     string `json:"asr"`
	Sunset  string `json:"sunset"`
	Maghrib string `json:"maghrib"`
	Isha    string `json:"isha"`
}

func SendJsonResponse(w http.ResponseWriter, times map[salat.TIMES]string) {

	timesPayload := timesPayload{
		Imsak:   times[salat.IMSAK],
		Fajr:    times[salat.FAJR],
		Sunrise: times[salat.SUNRISE],
		Dhuhr:   times[salat.DHUHR],
		Asr:     times[salat.ASR],
		Sunset:  times[salat.SUNSET],
		Maghrib: times[salat.MAGHRIB],
		Isha:    times[salat.ISHA],
	}

	jData, err := json.Marshal(timesPayload)
	log.Debug(timesPayload)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}
}

func SendErrorResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
