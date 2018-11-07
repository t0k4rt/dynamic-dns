package models

import (
	"net/url"
	"time"
)

type TomlConfig struct {
	General general
	Domain  []domain
}

func NewTomlConfig() TomlConfig {
	return TomlConfig{
		General: newGeneral(),
		Domain:  []domain{},
	}
}

// [general]
// refresh rate, default 1h
// default_refresh_delay = "1h"
// default_ttl = 300
// restrict_ip_versions = ["v4", "v6"]

type general struct {
	DefaultRefreshDelay duration `toml:"default_refresh_delay"`
	DefaultTTL          int      `toml:"default_ttl"`
	IPVersions          []string `toml:"ip_versions"`
}

func newGeneral() general {
	d, _ := time.ParseDuration("1h")
	return general{
		DefaultRefreshDelay: duration{d},
		DefaultTTL:          300,
		IPVersions:          []string{"v4", "v6"},
	}
}

// [[domain]]
// name="domain2.fr"
// ip_provider = "ip provider"
// dns_updater = "domain2 dns provider"
// # ttl =
// # refresh_delay =
type domain struct {
	Name         urll
	IPProvider   string `toml:"ip_provider"`
	DNSUpdater   string `toml:"dns_updater"`
	TTL          int
	RefreshDelay duration `toml:"refresh_delay"`
}

type duration struct {
	time.Duration
}

type urll struct {
	*url.URL
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (u *urll) UnmarshalText(text []byte) error {
	var err error
	u.URL, err = url.Parse(string(text))
	return err
}
