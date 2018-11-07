package dnsupdater

import (
	"errors"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider"
	"github.com/t0k4rt/dynamic-dns/internal/dnsupdater/gandi"
)

type DNSProvider interface {
	UpdateDNS(domain *url.URL, ip *ipaddressprovider.ProvidedIP) error
}

func Make(s string) DNSProvider, err {
	switch s {
	case "gandi":
		return gandi.NewUpdater()
	default:
		return nil, errors.New("Unknown DNSProvider")
	}
}
