package dnsupdater

import "github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider"

type DNSProvider interface {
	UpdateDNS(ip *ipaddressprovider.ProvidedIP) error
}
