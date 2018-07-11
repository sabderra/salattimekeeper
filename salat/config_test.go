package salat_test

import (
	"github.com/BurntSushi/toml"
	"github.com/sabderra/salattimekeeper/salat"
	"testing"
)

func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c salat.Config
	if _, err := toml.Decode(`
[[method]]
name = "Islamic Society of North America (ISNA)"
fajr = 15.0
isha = 15.0

[[method]]
name = "Muslim World League"
fajr = 18.0
isha = 17.0

[settings]
imsakAngleAdj = 10.0
fajrAngleAdj = 0.0
dhuhrAngleAdj = 0.0
asrFactor = 1
ishaAngleAdj = 0.0

`, &c); err != nil {
		t.Fatal(err)
	}

	// Validate configuration.
	if len(c.Methods) != 2 {
		t.Fatalf("unexpected methods count: %d", len(c.Methods))
	} else if c.Methods[0].Name != "Islamic Society of North America (ISNA)" {
		t.Fatalf("unexpected Methods[0] name: %s", c.Methods[0].Name)
	} else if c.Methods[0].Fajr != 15. {
		t.Fatalf("unexpected Methods[0] fajr: %f", c.Methods[0].Fajr)
	} else if c.Methods[0].Isha != 15. {
		t.Fatalf("unexpected Methods[0] fajr: %f", c.Methods[0].Isha)
	} else if c.Methods[1].Name != "Muslim World League" {
		t.Fatalf("unexpected Methods[1] name: %s", c.Methods[0].Name)
	} else if c.Methods[1].Fajr != 18. {
		t.Fatalf("unexpected Methods[1] fajr: %f", c.Methods[0].Fajr)
	} else if c.Methods[1].Isha != 17. {
		t.Fatalf("unexpected Methods[1] isha: %f", c.Methods[0].Isha)
	} else if c.Settings.ImsakAngleAdj != 10.0 {
		t.Fatalf("unexpected ImsakAngleAdj: %f", c.Settings.ImsakAngleAdj)
	} else if c.Settings.FajrAngleAdj != 0.0 {
		t.Fatalf("unexpected FajrAngleAdj: %f", c.Settings.FajrAngleAdj)
	} else if c.Settings.DhuhrAngleAdj != 0.0 {
		t.Fatalf("unexpected DhuhrAngleAdj: %f", c.Settings.DhuhrAngleAdj)
	} else if c.Settings.AsrFactor != salat.STANDARD {
		t.Fatalf("unexpected AsrFactor: %d", int(c.Settings.AsrFactor))
	} else if c.Settings.IshaAngleAdj != 0.0 {
		t.Fatalf("unexpected IshaAngleAdj: %f", c.Settings.IshaAngleAdj)
	}
}
