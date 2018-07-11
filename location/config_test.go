package location_test

import (
	"github.com/BurntSushi/toml"
	"github.com/sabderra/salattimekeeper/location"
	"testing"
)

func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c location.Config
	if _, err := toml.Decode(`
enabled = true
lat = 38.983939
lng = -76.844116
`, &c); err != nil {
		t.Fatal(err)
	}

	// Validate configuration.
	if !c.Enabled {
		t.Fatalf("unexpected enabled: %v", c.Enabled)
	} else if c.Lat != 38.983939 {
		t.Fatalf("unexpected bind address: %f", c.Lat)
	} else if c.Lng != -76.844116 {
		t.Fatalf("unexpected auth enabled: %f", c.Lng)
	}
}
