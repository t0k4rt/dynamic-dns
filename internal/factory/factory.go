package factory

import (
	"errors"
	"fmt"

	"github.com/t0k4rt/dynamic-dns/internal/dnsprovider"
	"github.com/t0k4rt/dynamic-dns/internal/dnsprovider/gandi"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider/livebox"
)

func MakeIPProvider(s string) (ipprovider.IPProvider, error) {
	switch s {
	case "livebox":
		return livebox.NewIPProvider(), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown IPProvider: \"%s\"", s))
	}
}

func MakeDNSProvider(s string) (dnsprovider.DNSProvider, error) {
	switch s {
	case "gandi":
		return gandi.NewDNSProvider()
	default:
		return nil, errors.New(fmt.Sprintf("Unknown DNSProvider \"%s\"", s))
	}
}
