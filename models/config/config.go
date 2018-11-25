package config

import (
	"time"

	"github.com/t0k4rt/dynamic-dns/internal/dnsprovider"
	"github.com/t0k4rt/dynamic-dns/internal/factory"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
)

type TomlConfig struct {
	General General
	Domain  []Domain
}

func NewTomlConfig() TomlConfig {
	return TomlConfig{
		General: newGeneral(),
		Domain:  []Domain{},
	}
}

// [general]
// refresh rate, default 1h
// default_refresh_delay = "1h"
// default_ttl = 300

type General struct {
	DefaultRefreshDelay duration `toml:"default_refresh_delay"`
	DefaultTTL          int      `toml:"default_ttl"`
	LogPath             string   `toml:"log_path"`
	LogLevel            string   `toml:"log_level"`
}

func newGeneral() General {
	d, _ := time.ParseDuration("1h")
	return General{
		DefaultRefreshDelay: duration{d},
		DefaultTTL:          300,
		LogPath:             "/var/log/dynamicdns.log",
		LogLevel:            "info",
	}
}

// # [[domain]]
// # name="subdomain.domain2.com"
// # ip_provider = "livebox"
// # dns_updater = "gandi"
// # ttl=3000
// # refresh_delay = 1h
// # ip_version  = 4
type Domain struct {
	Name         string
	IPProvider   *cIPProvider  `toml:"ip_provider"`
	DNSProvider  *cDNSProvider `toml:"dns_updater"`
	IPVersion    int           `toml:"ip_version"`
	TTL          int
	RefreshDelay duration `toml:"refresh_delay"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type cIPProvider struct {
	ipprovider.IPProvider
}

func (i *cIPProvider) UnmarshalText(text []byte) error {
	var err error
	i.IPProvider, err = factory.MakeIPProvider(string(text))
	return err
}

type cDNSProvider struct {
	dnsprovider.DNSProvider
}

func (d *cDNSProvider) UnmarshalText(text []byte) error {
	var err error
	d.DNSProvider, err = factory.MakeDNSProvider(string(text))
	return err
}
