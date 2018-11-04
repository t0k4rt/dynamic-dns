package dnsupdater

import "github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider"

type IPProvider interface {
	UpdateDNS(ip *ipaddressprovider.ProvidedIP) error
}
