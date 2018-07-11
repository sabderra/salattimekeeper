package location

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const IPLocationURL = "http://ipinfo.io/json"

type MyLocation struct {
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Org      string `json:"org"`
	Lat      float64
	Lng      float64
}

// NewMyLocation returns an instance of MyLocation with the provided Lat and Lng.
func NewMyLocation(c Config) *MyLocation {

	l := &MyLocation{}

	if c.Enabled {
		var err error
		l, err = GetLocationFromIp(l)
		if err != nil {
			log.Warn(err.Error())
		}
	} else {
		l = &MyLocation{}

		l.Lat = c.Lat
		l.Lng = c.Lng
	}

	return l
}

func GetLocationFromIp(myLocation *MyLocation) (*MyLocation, error) {

	var myClient = &http.Client{Timeout: 10 * time.Second}
	var err error

	resp, err := myClient.Get(IPLocationURL)

	if err != nil {
		return new(MyLocation), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return new(MyLocation), err
	}

	err = json.Unmarshal(body, &myLocation)

	loc := strings.Split(myLocation.Loc, ",")

	myLocation.Lat, err = strconv.ParseFloat(loc[0], 64)
	if err != nil {
		return new(MyLocation), err
	}

	myLocation.Lng, err = strconv.ParseFloat(loc[1], 64)
	if err != nil {
		return new(MyLocation), err
	}

	log.Infof("My Location is [%f,%f]: %s, %s, %s",
		myLocation.Lat, myLocation.Lng,
		myLocation.City, myLocation.Region, myLocation.Country)

	return myLocation, err
}
