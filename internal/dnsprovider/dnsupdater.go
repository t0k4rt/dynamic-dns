package dnsprovider

import (
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
)

type DNSProvider interface {
	UpdateDNS(domain string, ip *ipprovider.ProvidedIP, ttl int) error
}
