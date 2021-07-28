package ipify

import (
	"io/ioutil"
	"net"
	"net/http"
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
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return net.ParseIP(ip), nil
}

func (l *ipifyProvider) GetIPv6() (net.IP, error) {
	url := "https://api64.ipify.org?format=text"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return net.ParseIP(ip), nil
}
