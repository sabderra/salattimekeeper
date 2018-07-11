package main_test

import (
	"github.com/sabderra/salattimekeeper"
	"github.com/sabderra/salattimekeeper/salat"
	"testing"
)

// Ensure the configuration can be parsed.
func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c main.Config
	if err := c.FromToml(`
[server]
bind-address = ":8087"
[location]
enabled = true
lat = 39.9839
lng = -77.8441

[[salat.method]]
name = "Islamic Society of North America (ISNA)"
fajr = 15.0
isha = 15.0

[[salat.method]]
name = "Muslim World League"
fajr = 18.0
isha = 17.0

[salat.settings]
imsakAngleAdj = 10.0
fajrAngleAdj = 0.0
dhuhrAngleAdj = 0.0
asrFactor = 1
ishaAngleAdj = 0.0

`); err != nil {
		t.Fatal(err)
	}

	// Validate configuration.
	if c.Server.BindAddress != ":8087" {
		t.Fatalf("unexpected api bind address: %s", c.Server.BindAddress)
	} else if !c.Location.Enabled {
		t.Fatalf("unexpected subscriber enabled: %v", c.Location.Enabled)
	} else if c.Location.Lat != 39.9839 {
		t.Fatalf("unexpected subscriber enabled: %v", c.Location.Lat)
	} else if c.Location.Lng != -77.8441 {
		t.Fatalf("unexpected subscriber enabled: %v", c.Location.Lng)
	} else if len(c.Salat.Methods) != 2 {
		t.Fatalf("unexpected methods count: %d", len(c.Salat.Methods))
	} else if c.Salat.Methods[0].Name != "Islamic Society of North America (ISNA)" {
		t.Fatalf("unexpected Methods[0] name: %s", c.Salat.Methods[0].Name)
	} else if c.Salat.Methods[0].Fajr != 15. {
		t.Fatalf("unexpected Methods[0] fajr: %f", c.Salat.Methods[0].Fajr)
	} else if c.Salat.Methods[0].Isha != 15. {
		t.Fatalf("unexpected Methods[0] isha: %f", c.Salat.Methods[0].Isha)
	} else if c.Salat.Methods[1].Name != "Muslim World League" {
		t.Fatalf("unexpected Methods[1] name: %s", c.Salat.Methods[0].Name)
	} else if c.Salat.Methods[1].Fajr != 18. {
		t.Fatalf("unexpected Methods[1] fajr: %f", c.Salat.Methods[0].Fajr)
	} else if c.Salat.Methods[1].Isha != 17. {
		t.Fatalf("unexpected Methods[1] isha: %f", c.Salat.Methods[0].Isha)
	} else if c.Salat.Settings.ImsakAngleAdj != 10.0 {
		t.Fatalf("unexpected ImsakAngleAdj: %f", c.Salat.Settings.ImsakAngleAdj)
	} else if c.Salat.Settings.FajrAngleAdj != 0.0 {
		t.Fatalf("unexpected FajrAngleAdj: %f", c.Salat.Settings.FajrAngleAdj)
	} else if c.Salat.Settings.DhuhrAngleAdj != 0.0 {
		t.Fatalf("unexpected DhuhrAngleAdj: %f", c.Salat.Settings.DhuhrAngleAdj)
	} else if c.Salat.Settings.AsrFactor != salat.STANDARD {
		t.Fatalf("unexpected AsrFactor: %d", int(c.Salat.Settings.AsrFactor))
	} else if c.Salat.Settings.IshaAngleAdj != 0.0 {
		t.Fatalf("unexpected IshaAngleAdj: %f", c.Salat.Settings.IshaAngleAdj)
	}

}
