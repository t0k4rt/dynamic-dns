package ipaddressprovider

import "net"

type IPProvider interface {
	GetIP() (*ProvidedIP, error)
}

type ProvidedIP struct {
	ipv4 net.IP
	ipv6 net.IP
}

func NewFromString(ipv4 string, ipv6 string) *ProvidedIP {
	return &ProvidedIP{
		ipv4: net.ParseIP(ipv4),
		ipv6: net.ParseIP(ipv6),
	}
}

func NewFromIPv4String(ipv4 string) *ProvidedIP {
	return &ProvidedIP{
		ipv4: net.ParseIP(ipv4),
		ipv6: nil,
	}
}

func NewFromIPv6String(ipv6 string) *ProvidedIP {
	return &ProvidedIP{
		ipv4: nil,
		ipv6: net.ParseIP(ipv6),
	}
}

func (ip *ProvidedIP) GetIPV4() net.IP {
	return ip.ipv4
}

func (ip *ProvidedIP) GetIPV6() net.IP {
	return ip.ipv6
}

func (ip *ProvidedIP) GetIPV4String() string {
	if ip.ipv4 == nil {
		return ""
	}
	return ip.ipv4.String()
}

func (ip *ProvidedIP) GetIPV6String() string {
	if ip.ipv6 == nil {
		return ""
	}
	return ip.ipv6.String()
}
