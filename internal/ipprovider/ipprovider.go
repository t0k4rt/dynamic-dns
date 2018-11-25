package ipprovider

import (
	"net"
)

type IPProvider interface {
	GetIP(version int) (net.IP, error)
	GetIPv4() (net.IP, error)
	GetIPv6() (net.IP, error)
}
