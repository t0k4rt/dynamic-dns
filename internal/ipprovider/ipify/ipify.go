package ipify

import (
	"net"
	"errors"
	ipify "github.com/rdegges/go-ipify"
)

type ipifyProvider struct {
}

// NewIPProvider creates a new ip provider
func NewIPProvider() (*ipifyProvider, error) {
	return &ipifyProvider{}, nil
}

func (l *ipifyProvider) GetIP(version int) (net.IP, error) {
	var err error
	var ip net.IP
	switch version {
	case 4:
		ip, err = l.GetIPv4()
	case 6:
		ip, err = l.GetIPv6()
	}
	return ip, err
}

func (l *ipifyProvider) GetIPv4() (net.IP, error) {

	ip, err := ipify.GetIp()
	if err != nil {
		return nil, err
	}
	return net.ParseIP(ip), nil

}

func (l *ipifyProvider) GetIPv6() (net.IP, error) {
	return nil, errors.New("Ipify does not support ipv6")
}
