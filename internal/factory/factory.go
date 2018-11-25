package factory

import (
	"errors"
	"fmt"
	"strings"

	"github.com/t0k4rt/dynamic-dns/internal/dnsprovider"
	"github.com/t0k4rt/dynamic-dns/internal/dnsprovider/gandi"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider/livebox"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider/netlink"
)

func MakeIPProvider(s string) (ipprovider.IPProvider, error) {

	switch {
	case strings.Contains(s, "interface"):
		splitted := strings.Split(s, ":")
		return netlink.NewIPProvider(splitted[1])
	case s == "livebox":
		return livebox.NewIPProvider()
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
