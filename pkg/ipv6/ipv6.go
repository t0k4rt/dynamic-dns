package ipv6

import (
	"errors"
	"fmt"
	"net"
	"regexp"
)

var cidrRegexp, _ = regexp.Compile("^2[a-z0-9]{3}:.*\\/[0-9]*")
var ipRegexp, _ = regexp.Compile("^2[a-z0-9]{3}:.*")

func ParsePublicIPv6(ipString string) (net.IP, error) {
	if cidrRegexp.MatchString(ipString) {
		ip, mask, err := net.ParseCIDR(ipString)
		if err != nil {
			return nil, err
		}

		if ip.To4() == nil && cidrRegexp.MatchString(mask.String()) {
			return ip, nil
		}
	} else if ipRegexp.MatchString(ipString) {
		ip := net.ParseIP(ipString)
		if ip == nil {
			return nil, errors.New(fmt.Sprintf("could not parse ip v6 %s", ipString))
		}

		if ip.To4() == nil {
			return ip, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("could not parse ip v6 %s", ipString))
}
