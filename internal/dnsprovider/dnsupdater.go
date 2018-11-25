package dnsprovider

import "net"

type DNSProvider interface {
	UpdateDNS(fullDomain string, ip net.IP, ttl int, version int) error
}
