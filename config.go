package main

import (
	"github.com/BurntSushi/toml"
	"github.com/sabderra/salattimekeeper/location"
	"github.com/sabderra/salattimekeeper/salat"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

const (
	// DefaultBindAddress is the default address for various RPC services.
	DefaultBindAddress  = "127.0.0.1:8999"
	DefaultWriteTimeout = 15
	DefaultReadTimeout  = 15
	DefaultIdleTimeout  = 60
)

// Config represents the configuration format for the Salat TimeKeeper.
type Config struct {
	Server   serverConfig    `toml:"server"`
	Location location.Config `toml:"location"`
	Salat    salat.Config    `toml:"salat"`
}

type serverConfig struct {
	// BindAddress is the address that the REST services will be listening on
	BindAddress  string        `toml:"bind-address"`
	WriteTimeout time.Duration `toml:"write-timeout"`
	ReadTimeout  time.Duration `toml:"read-timeout"`
	IdleTimeout  time.Duration `toml:"idle-timeout"`
}

// NewConfig returns an instance of Config with defaults.
func NewConfig() *Config {
	c := &Config{}
	c.Server.BindAddress = DefaultBindAddress
	c.Server.WriteTimeout = DefaultWriteTimeout
	c.Server.ReadTimeout = DefaultReadTimeout
	c.Server.IdleTimeout = DefaultIdleTimeout

	c.Location = *location.NewConfig()

	c.Salat = *salat.NewConfig()

	return c
}

// FromTomlFile loads the config from a TOML file.
func (c *Config) FromTomlFile(fpath string) error {
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	// Handle any potential Byte-Order-Marks that may be in the config file.
	// This is for Windows compatibility only.
	// See https://github.com/influxdata/telegraf/issues/1378 and
	// https://github.com/influxdata/influxdb/issues/8965.
	bom := unicode.BOMOverride(transform.Nop)
	bs, _, err = transform.Bytes(bom, bs)
	if err != nil {
		return err
	}
	return c.FromToml(string(bs))
}

// FromToml loads the config from TOML.
func (c *Config) FromToml(input string) error {
	// Replace deprecated [cluster] with [coordinator]
	re := regexp.MustCompile(`(?m)^\s*\[cluster\]`)
	input = re.ReplaceAllStringFunc(input, func(in string) string {
		in = strings.TrimSpace(in)
		out := "[coordinator]"
		log.Printf("deprecated config option %s replaced with %s; %s will not be supported in a future release\n", in, out, in)
		return out
	})

	_, err := toml.Decode(input, c)
	return err
}
