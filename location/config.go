package location

const (
	// DefaultLat is the default latitude to use for the servers location.
	DefaultLat = 42.374146

	// DefaultLng is the default longitude to use for the servers location.
	DefaultLng = -71.110279
)

// Config represents the configuration format for the Location lookup.
type Config struct {
	Enabled bool    `toml:"enabled"`
	Lat     float64 `toml:"lat"`
	Lng     float64 `toml:"lng"`
}

// NewConfig returns an instance of Config with defaults.
func NewConfig() *Config {
	c := &Config{}

	c.Enabled = false
	c.Lat = DefaultLat
	c.Lng = DefaultLng

	return c
}
