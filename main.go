package main

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/sabderra/salattimekeeper/location"
	"github.com/sabderra/salattimekeeper/salat"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var config Config

var myLocation location.MyLocation

// shutdownChan for stopping go routines
var shutdownChan = make(chan struct{})

var mwlParams = map[salat.TIMES]float64{salat.FAJR: 18, salat.ISHA: 17}
var mwlConfig = make(map[string]string)
var MWL = salat.NewCalculationMethod("Muslim World League", mwlParams, mwlConfig)

var isnaParams = map[salat.TIMES]float64{salat.FAJR: 15, salat.ISHA: 15}
var isnaConfig = make(map[string]string)
var ISNA = salat.NewCalculationMethod("Islamic Society of North America (ISNA)", isnaParams, isnaConfig)

var eqyptParams = map[salat.TIMES]float64{salat.FAJR: 19.5, salat.ISHA: 17.5}
var eqyptConfig = make(map[string]string)
var EGYPT = salat.NewCalculationMethod("Egyptian General Authority of Survey", eqyptParams, eqyptConfig)

var makkahParams = map[salat.TIMES]float64{salat.FAJR: 18.5, salat.ISHA: 17.5}
var makkahConfig = make(map[string]string)
var MAKKAH = salat.NewCalculationMethod("Umm Al-Qura University, Makkah", makkahParams, makkahConfig)

var karachiParams = map[salat.TIMES]float64{salat.FAJR: 18, salat.ISHA: 18}
var karachiConfig = make(map[string]string)
var KARACHI = salat.NewCalculationMethod("University of Islamic Sciences, Karachi", karachiParams, karachiConfig)

var tehranParams = map[salat.TIMES]float64{salat.FAJR: 17.7, salat.ISHA: 14, salat.MAGHRIB: 4.5}
var tehranConfig = map[string]string{"midnight": "Jafari"}
var TEHRAN = salat.NewCalculationMethod("Institute of Geophysics, University of Tehran", tehranParams, tehranConfig)

var jafariParams = map[salat.TIMES]float64{salat.FAJR: 16, salat.ISHA: 14, salat.MAGHRIB: 4}
var jafariConfig = map[string]string{"midnight": "Jafari"}
var JAFARI = salat.NewCalculationMethod("Shia Ithna-Ashari, Leva Institute, Qum", jafariParams, jafariConfig)

func init() {

	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		fmt.Println(err)
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Set severity and only log above.
	log.SetLevel(log.DebugLevel)
}

func main() {

	log.Info("SalatTimeKeeper starting")

	// Set up  OS signal handlers
	registerSignalTermHandler()
	registerSignalHupHandler()

	myLocation = *location.NewMyLocation(config.Location)

	start(config, CreateRouter())
	log.Info("SalatTimeKeeper stopping")

}

func logClientRequest(req *http.Request) {
	forwardedFor := req.Header.Get("X-Forwarded-For")
	log.Debugf("request  %s from %s, %s", req.RequestURI, req.RemoteAddr, forwardedFor)
}

func start(config Config, handler *mux.Router) {
	var wait time.Duration

	srv := &http.Server{
		Addr:         config.Server.BindAddress,
		WriteTimeout: time.Second * config.Server.WriteTimeout,
		ReadTimeout:  time.Second * config.Server.ReadTimeout,
		IdleTimeout:  time.Second * config.Server.IdleTimeout,
		Handler:      handler,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Infof("Listening on %s", config.Server.BindAddress)
	<-shutdownChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
}

// RegisterSignalTermHandler catches os signals to properly shutdown process.
func registerSignalTermHandler() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Debug("Shutdown Listener for SIGINT, SIGTERM.")
		sig := <-sigs
		log.Debugf("Caught shutdown signal: %d", sig)
		log.Debug("Sending a shutdown message.")
		shutdown()
	}()

}

// RegisterSignalHupHandler catches os signals to reload configuration.
func registerSignalHupHandler() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)

	go func() {
		log.Debug("Reload Listener for SIGHUP.")
		sig := <-sigs
		log.Debugf("Caught reload signal: %d", sig)
		log.Info("Reloading configuration.")
		// XXX loadConfig()
	}()

}

// Shutdown stops the worker to stop listening for work requests.
//
func shutdown() {
	close(shutdownChan)
}
