package netlink

import (
	"errors"
	"net"

	"github.com/t0k4rt/dynamic-dns/pkg/ipv6"
)

type netlinkProvider struct {
	netInterface *net.Interface
}

func NewIPProvider(netInterface string) (*netlinkProvider, error) {
	byNameInterface, err := net.InterfaceByName(netInterface)
	if err != nil {
		return nil, err
	}
	return &netlinkProvider{
		netInterface: byNameInterface,
	}, nil
}

func (n *netlinkProvider) GetIP(version int) (net.IP, error) {
	var err error
	var ip net.IP
	switch version {
	case 4:
		ip, err = n.GetIPv4()
	case 6:
		ip, err = n.GetIPv6()
	}
	return ip, err
}

func (n *netlinkProvider) GetIPv4() (net.IP, error) {

	addresses, err := n.netInterface.Addrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addresses {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			return nil, err
		}
		if ip.To4() != nil {
			return ip, nil
		}
	}
	return nil, errors.New("Could not get ipv4, no ipv4 associated to interface")
}

func (n *netlinkProvider) GetIPv6() (net.IP, error) {
	addresses, err := n.netInterface.Addrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addresses {
		ip, err := ipv6.ParsePublicIPv6(addr.String())

		if err != nil {
			continue
		}
		return ip, nil
	}
	return nil, errors.New("Could not find ipv6, no ipv6 associated to interface")
}
