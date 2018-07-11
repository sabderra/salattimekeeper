package salat

const (
	// DefaultLat is the default fajr adjustment.
	DefaultFajr = 0.

	// DefaultLat is the default isha adjustment.
	DefaultIsha = 0.
)

// Config represents the configuration format for the method settings.
type Config struct {
	Methods  []methodConfig `toml:"method"`
	Settings settingsConfig `toml:"settings"`
}

// Config represents the configuration format for the method settings.
type methodConfig struct {
	Name string  `toml:"name"`
	Fajr float64 `toml:"fajr"`
	Isha float64 `toml:"isha"`

	initialIMSAK   float64 `toml:"initialIMSAK"`
	initialFAJR    float64 `toml:"initialFAJR"`
	initialSUNRISE float64 `toml:"initialSUNRISE"`
	initialDHUHR   float64 `toml:"initialDHUHR"`
	initialASR     float64 `toml:"initialASR"`
	initialSUNSET  float64 `toml:"initialSUNSET"`
	initialMAGHRIB float64 `toml:"initialMAGHRIB"`
	initialISHA    float64 `toml:"initialISHA"`
}

type settingsConfig struct {
	ImsakAngleAdj   float64    `toml:"imsakAngleAdj"` // Angle in minutes
	FajrAngleAdj    float64    `toml:"fajrAngleAdj"`
	DhuhrAngleAdj   float64    `toml:"dhuhrAngleAdj"` // Angle in minutes
	AsrFactor       ASR_FACTOR `toml:"asrFactor"`
	MaghribAngleAdj float64    `toml:"maghribAngleAdj"`
	IshaAngleAdj    float64    `toml:"ishaAngleAdj"`
}

// NewConfig returns an instance of Config with defaults.
func NewConfig() *Config {
	c := &Config{}

	c.Methods = []methodConfig{*NewMethodConfig()}
	c.Settings = *NewSettingsConfig()

	return c
}

// NewMethodConfig returns an instance of Config for Calculation Methods with defaults.
func NewMethodConfig() *methodConfig {
	c := &methodConfig{}

	c.Name = ""
	c.Fajr = DefaultFajr
	c.Isha = DefaultIsha

	c.initialIMSAK = 5
	c.initialFAJR = 5
	c.initialSUNRISE = 6
	c.initialDHUHR = 12
	c.initialASR = 13
	c.initialSUNSET = 18
	c.initialMAGHRIB = 18
	c.initialISHA = 18

	return c
}

// NewSettingConfig returns an instance of settings with defaults.
func NewSettingsConfig() *settingsConfig {
	c := &settingsConfig{}

	c.ImsakAngleAdj = 10
	c.FajrAngleAdj = 0
	c.DhuhrAngleAdj = 0
	c.AsrFactor = STANDARD
	c.IshaAngleAdj = 0

	return c
}
