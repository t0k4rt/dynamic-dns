package domainparser

import (
	"strings"

	"golang.org/x/net/publicsuffix"
)

// ParseDomain returns domain and subDomain from full domain
// example:
//     www.alibaba.fr will return
//     alibaba.fr and www
func ParseDomain(domain string) (string, string, error) {
	suffix, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", "", err
	}
	index := strings.LastIndex(domain, suffix)
	if index == 0 {
		return suffix, "@", nil
	}
	subDomain := domain[0 : index-1]

	return suffix, subDomain, nil
}
