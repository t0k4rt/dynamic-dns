package config

import (
	"net/url"
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
// restrict_ip_versions = ["v4", "v6"]

type General struct {
	DefaultRefreshDelay duration `toml:"default_refresh_delay"`
	DefaultTTL          int      `toml:"default_ttl"`
	IPVersions          []string `toml:"ip_versions"`
}

func newGeneral() General {
	d, _ := time.ParseDuration("1h")
	return General{
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
type Domain struct {
	Name         curl
	IPProvider   cIPProvider  `toml:"ip_provider"`
	DNSProvider  cDNSProvider `toml:"dns_updater"`
	TTL          int
	RefreshDelay duration `toml:"refresh_delay"`
}

// func newDomain() Domain {
// 	d, _ := time.ParseDuration("1h")
// 	return Domain{
// 		Name:       nil,
// 		IPProvider: nil,
// 		DNSProvider: nil,
// 		TTL:	300,
// 		IPVersions: []string{"v4", "v6"},
// 	}
// }

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type curl struct {
	*url.URL
}

func (u *curl) UnmarshalText(text []byte) error {
	var err error
	u.URL, err = url.Parse(string(text))
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
