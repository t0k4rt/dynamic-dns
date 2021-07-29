package dns

import (
	"bytes"
	"net"
	"os/exec"
	"strings"
)

type dns struct {
}

func NewIPProvider() (*dns, error) {
	_, err := exec.LookPath("dig")
	if err != nil {
		return nil, err
	}
	return &dns{}, nil
}

func (l *dns) GetIP(version int) (net.IP, error) {
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

func (l *dns) GetIPv4() (net.IP, error) {
	cmd := exec.Command("dig", "+short", "myip.opendns.com", "@resolver1.opendns.com", "-4")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return net.ParseIP(strings.TrimSpace(out.String())), nil
}

func (l *dns) GetIPv6() (net.IP, error) {
	cmd := exec.Command("dig", "+short", "AAAA", "myip.opendns.com", "@resolver1.opendns.com")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return net.ParseIP(strings.TrimSpace(out.String())), nil
}
