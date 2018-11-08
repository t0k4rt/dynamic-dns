package dnsprovider

import (
	"net/url"

	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
)

type DNSProvider interface {
	UpdateDNS(domain *url.URL, ip *ipprovider.ProvidedIP) error
}
